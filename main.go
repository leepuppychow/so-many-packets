package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"

	"github.com/google/gopacket/pcap"
)

var (
	device       string
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Description:", device.Description)
		fmt.Println("Addresses:", device.Addresses)
	}

	fmt.Println("Capturing packets for device en0")
	CapturePackets("en0")
}

func CapturePackets(device string) {
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
