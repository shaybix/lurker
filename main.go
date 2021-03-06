package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	version := pcap.Version()
	fmt.Println(version)

	handle, err := pcap.OpenLive(
		"en0",
		int32(65535),
		true,
		-1*time.Second,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

}
