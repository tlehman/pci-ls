# `pci-ls` is a like `lspci` but for go, and not intended to replace it

It's really just a test project to use the u-root pci library.

## how to build

```
go build
```

## how to run

```
./pci-ls
```

Then it will produce output like this:
```
PCI Device: 9b44
        Address: 0000:00:00.0
        VendorId: 32902
        DeviceId: 39748
PCI Device: Xeon E3-1200 v5/E3-1500 v5/6th Gen Core Processor PCIe Controller (x16)
        Address: 0000:00:01.0
        VendorId: 32902
        DeviceId: 6401
PCI Device: 9bc4
        Address: 0000:00:02.0
        VendorId: 32902
        DeviceId: 39876
PCI Device: Xeon E3-1200 v5/E3-1500 v5/6th Gen Core Processor Thermal Subsystem
        Address: 0000:00:04.0
        VendorId: 32902
        DeviceId: 6403
PCI Device: Xeon E3-1200 v5/v6 / E3-1500 v5 / 6th/7th Gen Core Processor Gaussian Mixture Model
        Address: 0000:00:08.0
        VendorId: 32902
        DeviceId: 6417
...
```