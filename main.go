package main

import (
	p "github.com/leepuppychow/so-many-packets/packets"
)

func main() {
	p.ShowDevices()
	p.CapturePackets("en0")
}
