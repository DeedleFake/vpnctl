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
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"io"
	//"time"
)

// Need to check for current tun devices, then assign the next highest
// one to the current attempt. Using same, openvpn initialization fails.

func main() {
	if os.Geteuid() != 0 {
		log.Fatalln( "Root permissions required." )
	}

	fmt.Println( net.Interfaces() )

	openvpn := exec.Command( "openvpn",
		"--config","/etc/openvpn/west.conf" )
	stdout,_ := openvpn.StdoutPipe()
	stderr,_ := openvpn.StderrPipe()
	go io.Copy( os.Stdout,stdout )
	go io.Copy( os.Stderr,stderr )
	err := openvpn.Start()
	if err != nil {
		log.Fatal( err )
	}

	//fmt.Print( "Opening VPN connection." )
	//// TODO timeout should not be hard coded
	//for t := 0; t < 5; t++ {
	//	if verifyTunnel() {
	//		// why on earth does this not work?
	//		break
	//	}
	//	time.Sleep( 1 )
	//	fmt.Print( "." )
	//}

	for {
	}
}

// vim: noet tw=72 ts=3 sw=3
