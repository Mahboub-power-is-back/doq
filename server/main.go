package main

import (
    "flag"
    "log"
    
    "github.com/Mygod/dnstt/server"
)

func main() {
    listenAddr := flag.String("listen", ":784", "QUIC listen address")
    flag.Parse()
    
    quicServer, err := server.NewQUICServer(*listenAddr, handleQUICStream)
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
    defer quicServer.Close()
    
    log.Println("QUIC server listening on", *listenAddr)
    if err := quicServer.Serve(); err != nil {
        log.Fatal("Server error:", err)
    }
}

func handleQUICStream(stream quic.Stream) {
    // Handle DNS tunnel logic over QUIC stream
    // ... existing DNS processing logic
}
