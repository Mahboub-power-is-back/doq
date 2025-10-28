package server

import (
    "context"
    "crypto/rand"
    "crypto/rsa"
    "crypto/tls"
    "crypto/x509"
    "encoding/pem"
    "math/big"
    "net"

    "github.com/lucas-clemente/quic-go"
)

type QUICServer struct {
    listener quic.Listener
    handler  func(quic.Stream)
}

func NewQUICServer(addr string, handler func(quic.Stream)) (*QUICServer, error) {
    listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
    if err != nil {
        return nil, err
    }
    
    return &QUICServer{
        listener: listener,
        handler:  handler,
    }, nil
}

func (s *QUICServer) Serve() error {
    for {
        session, err := s.listener.Accept(context.Background())
        if err != nil {
            return err
        }
        
        go s.handleSession(session)
    }
}

func (s *QUICServer) handleSession(session quic.Session) {
    for {
        stream, err := session.AcceptStream(context.Background())
        if err != nil {
            return
        }
        
        go s.handler(stream)
    }
}

func (s *QUICServer) Close() error {
    return s.listener.Close()
}

func generateTLSConfig() *tls.Config {
    key, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }
    
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        NotBefore:    time.Now(),
        NotAfter:     time.Now().Add(365 * 24 * time.Hour),
    }
    
    certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
    if err != nil {
        panic(err)
    }
    
    keyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(key),
    })
    certPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "CERTIFICATE",
        Bytes: certDER,
    })
    
    tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
    if err != nil {
        panic(err)
    }
    
    return &tls.Config{
        Certificates: []tls.Certificate{tlsCert},
        NextProtos:   []string{"dnstt-quic"},
    }
}
