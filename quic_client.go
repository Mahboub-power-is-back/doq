package client

import (
    "context"
    "crypto/tls"
    "net"
    "time"

    "github.com/lucas-clemente/quic-go"
)

type QUICClient struct {
    session quic.Session
    stream  quic.Stream
    serverAddr string
}

func NewQUICClient(serverAddr string) *QUICClient {
    return &QUICClient{
        serverAddr: serverAddr,
    }
}

func (c *QUICClient) Connect() error {
    tlsConf := &tls.Config{
        InsecureSkipVerify: true,
        NextProtos:         []string{"dnstt-quic"},
    }
    
    session, err := quic.DialAddr(c.serverAddr, tlsConf, nil)
    if err != nil {
        return err
    }
    
    stream, err := session.OpenStreamSync(context.Background())
    if err != nil {
        return err
    }
    
    c.session = session
    c.stream = stream
    return nil
}

func (c *QUICClient) Read(b []byte) (int, error) {
    return c.stream.Read(b)
}

func (c *QUICClient) Write(b []byte) (int, error) {
    return c.stream.Write(b)
}

func (c *QUICClient) Close() error {
    if c.stream != nil {
        c.stream.Close()
    }
    if c.session != nil {
        return c.session.CloseWithError(0, "")
    }
    return nil
}
