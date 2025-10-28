package protocol

import (
    "io"
    
    "github.com/lucas-clemente/quic-go"
)

func HandleQUICTunnel(stream quic.Stream, dnsHandler func([]byte) ([]byte, error)) {
    defer stream.Close()
    
    buf := make([]byte, 4096)
    for {
        n, err := stream.Read(buf)
        if err != nil {
            if err != io.EOF {
                // Handle error
            }
            return
        }
        
        // Process DNS query through existing logic
        response, err := dnsHandler(buf[:n])
        if err != nil {
            // Handle error
            continue
        }
        
        // Send response back
        if _, err := stream.Write(response); err != nil {
            // Handle write error
            return
        }
    }
}
