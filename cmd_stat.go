package main

import (
	"fmt"
	"net"
)

func init() {
	RegisterCmd(&statCmd{})
}

type statCmd struct {
}

func (c statCmd) Name() string {
	return "stat"
}

func (c statCmd) Desc() string {
	return "Gives information about the current status of the VPN."
}

func (c statCmd) filterTuns(ifs []net.Interface) []net.Interface {
	tuns := make([]net.Interface, 0, len(ifs))
	for n := 0; n < len(ifs); n++ {
		netif := ifs[n]
		name := netif.Name

		isTun := len(name) > 2 && name[0:3] == "tun" &&
			netif.Flags&net.FlagPointToPoint > 0

		if isTun {
			tuns = append(tuns, netif)
		}
	}
	return tuns
}

func (c statCmd) Run(config Config) error {
	ifs, err := net.Interfaces()
	if err != nil {
		return fmt.Errorf("vpnstat: Failed to get interface info: %v", err)
	}

	tuns := c.filterTuns(ifs)
	fmt.Println(tuns)

	return nil
}
