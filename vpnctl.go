// Copyright (c) 2015, vixsomnis < vs [at] vczf.io >
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sort"
)

type Command struct {
	Name string
	Desc string

	Run func(Config) error
}

type Config struct {
	Config string
}

type CommandList struct {
	cmds   []*Command
	config Config
}

func (c *CommandList) Add(cmd *Command) {
	c.cmds = append(c.cmds, cmd)
	sort.Sort(c)
}

func (c *CommandList) Run(name string) error {
}

func (c *CommandList) Len() int {
	return len(c.cmds)
}

func (c *CommandList) Swap(i1, i2 int) {
	c.cmds[i1], c.cmds[i2] = c.cmds[i2], c.cmds[i1]
}

func (c *CommandList) Less(i1, i2 int) bool {
	return c.cmds[i1].Name < c.cmds[i2].Name
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %v [command]

vpnctl is a wrapper around OpenVPN.

If no command is specified, status is displayed.

Commands:
	up:	Bring OpenVPN up.
	down:	Bring OpenVPN down.

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}

	var config Config
	flag.StringVar(&config.Config, "conf", "", "The OpenVPN config to use.")

	flag.Parse()

	cmd := flag.Arg(1)

	switch cmd {
	case "up":
		chkroot("up")
		vpnup(*conf)
	case "down":
		chkroot("down")
		//vpndown(*conf)
	default:
		vpnstat()
	}
}

func chkroot(cmd string) {
	if os.Geteuid() != 0 {
		log.Fatalf("%v: Root permissions are required for this command.", cmd)
	}
}

func vpnup(conf string) {
	openvpn := exec.Command("openvpn", "--config", conf)
	openvpn.Stdout = os.Stdout
	openvpn.Stderr = os.Stderr

	err := openvpn.Run()
	if err != nil {
		log.Fatalf("vpnup: openvpn failed: %v", err)
	}
}

func vpnstat() {
	ifs, err := net.Interfaces()
	if err != nil {
		log.Fatalf("vpnstat: Failed to get interface info: %v", err)
	}

	tuns := filterTuns(ifs)
	fmt.Println(tuns)
}

func filterTuns(ifs []net.Interface) []net.Interface {
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

// vim: noet tw=72 ts=3 sw=3
