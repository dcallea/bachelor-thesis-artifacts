/* 
   Reads packets.pcap file and outputs contents to terminal.

   Program structure adapted from code examples in gopacket documentation:
   https://pkg.go.dev/github.com/google/gopacket
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open the pcap file for reading
	pcapFile, err := os.Open("packets.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer pcapFile.Close()

	// Create a pcap reader
	handle, err := pcap.OpenOffline("packets.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create a packet source from the handle
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Iterate over packets and display information
	for packet := range packetSource.Packets() {
		// Display packet timestamp
		fmt.Printf("Timestamp: %s\n", packet.Metadata().Timestamp)

		// Display packet length
		fmt.Printf("Length: %d bytes\n", packet.Metadata().Length)

		// Display packet application layer data
		if appLayer := packet.ApplicationLayer(); appLayer != nil {
			fmt.Printf("Data: %s\n", appLayer.Payload())
		} else {
			fmt.Println("No application layer data")
		}

		// Display packet details using gopacket's string representation
		fmt.Printf("Packet Details:\n%s\n", packet)

		// Add a separator between packets
		fmt.Println("-------------------------------------------------")
	}
}
