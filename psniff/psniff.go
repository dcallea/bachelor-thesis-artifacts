/*
   This program captures tcp packets from device set by dev at set port and
   writes it to a packets.pcap file.
*/

package main

import (
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers" // for writing .pcap files
	"github.com/google/gopacket/pcap"   // for using pcap
	"github.com/google/gopacket/pcapgo" // for writing .pcap files
)

const dev = "lo" // loopback device on linux

func main() {
	snaplen := int32(1024)
	promisc := false // we only capture from localhost (non promiscuous)
	timeout := pcap.BlockForever

	// open network device for capturing
	handle, err := pcap.OpenLive(dev, snaplen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	// clean up when we're done
	defer handle.Close()

	// create pcap file
	pcapFile, err := os.Create("packets.pcap")
	if err != nil {
		log.Fatal(err)
	}
	// clean up when we're done
	defer pcapFile.Close()

	// create a pcapgo writer
	pcapWriter := pcapgo.NewWriter(pcapFile)

	// write the pcap file header
	err = pcapWriter.WriteFileHeader(65535, layers.LinkTypeEthernet)
	if err != nil {
		log.Fatal(err)
	}

	// set berkeley packet filter to only capture tcp packets on port 5050
	if err := handle.SetBPFFilter("tcp and port 5050"); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// capture packets and write to pcap file
	for packet := range packetSource.Packets() {

		// create CaptureInfo from packet metadata
		captureInfo := gopacket.CaptureInfo{
			Timestamp:      packet.Metadata().Timestamp,
			CaptureLength:  packet.Metadata().Length,
			Length:         len(packet.Data()),
			InterfaceIndex: 0,
		}

		err := pcapWriter.WritePacket(captureInfo, packet.Data())
		if err != nil {
			log.Fatal(err)
		}
	}
}
