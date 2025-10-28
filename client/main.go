package main

import (
    "flag"
    "log"
    
    "github.com/Mygod/dnstt/client"
)

func main() {
    serverAddr := flag.String("server", "localhost:784", "QUIC server address")
    flag.Parse()
    
    quicClient := client.NewQUICClient(*serverAddr)
    if err := quicClient.Connect(); err != nil {
        log.Fatal("Failed to connect:", err)
    }
    defer quicClient.Close()
    
    // Existing DNS tunnel logic adapted for QUIC streams
    // ... rest of tunnel implementation
}
