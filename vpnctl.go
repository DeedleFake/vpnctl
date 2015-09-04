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
	//"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"io"
	"time"
)

// Need to check for current tun devices, then assign the next highest
// one to the current attempt. Using same, openvpn initialization fails.

func main() {

	conf := os.Args[ 1 ]
	cmd := os.Args[ 2 ]

	switch cmd {
	case "up":
		chkroot()
		vpnup( conf )
	case "down":
		chkroot()
		//vpndown( conf )
	default:
		vpnstat()
	}

	for {
		time.Sleep(10000)
	}
}

func chkroot() {
	if os.Geteuid() != 0 {
		log.Fatalln( "Root permissions are required for this command." )
	}
}

func vpnup( conf string ) {

	openvpn := exec.Command( "openvpn","--config",conf )

	stdout,_ := openvpn.StdoutPipe()
	stderr,_ := openvpn.StderrPipe()

	go io.Copy( os.Stdout,stdout )
	go io.Copy( os.Stderr,stderr )

	err := openvpn.Start()
	chk( err )
}

func vpnstat() {
	ifs,err := net.Interfaces();
	chk( err )
	parseInterfaces( ifs )
}

func parseInterfaces( ifs []net.Interface ){
	// tun, up, point-to-point -- what need
	for n := 0; n < len( ifs ); n++ {
		fmt.Println( ifs[ n ].Name )
	}
}

func chk( err error ) {
	if err != nil {
		log.Fatal( err )
	}
}

// vim: noet tw=72 ts=3 sw=3
