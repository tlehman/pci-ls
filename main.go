package main

import (
	"fmt"

	"github.com/u-root/u-root/pkg/pci"
)

func main() {
	br, err := pci.NewBusReader()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var devices []*pci.PCI
	devices, err = br.Read()
	if len(devices) == 0 {
		fmt.Println("got 0 devices, want at least 1")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, device := range devices {
		fmt.Printf("PCI Device: %s\n", device.DeviceName)
		fmt.Printf("\tAddress: %s\n", device.Addr)
		fmt.Printf("\tVendorId: %d\n", device.Vendor)
		fmt.Printf("\tDeviceId: %d\n", device.Device)
	}
}
