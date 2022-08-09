package main

import (
	"github.com/u-root/u-root/pkg/pci"
)

func main() {
	br, err := pci.NewBusReader()
	if err != nil {
		t.Fatal(err)
	}
	if len(n.(*bus).Devices) == 0 {
		t.Fatal("got 0 devices, want at least 1")
	}

}
