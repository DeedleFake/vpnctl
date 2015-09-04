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
	"os"
)

func chkroot() bool {
	if os.Geteuid() != 0 {
		return false
	}

	return true
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %v [command]

vpnctl is a wrapper around OpenVPN.

If no command is specified, status is displayed.

Commands:
`, os.Args[0])

		for _, cmd := range GetCmds() {
			fmt.Fprintf(os.Stderr, "\t%v:\t%v\n", cmd.Name(), cmd.Desc())
		}
		fmt.Fprintln(os.Stderr)

		fmt.Fprintf(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	var config Config
	flag.StringVar(&config.Config, "conf", "", "The OpenVPN config to use.")

	flag.Parse()
	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(2)
	}

	cmdname := "stat"
	if flag.NArg() == 1 {
		cmdname = flag.Arg(0)
	}

	cmd := GetCmd(cmdname)
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "Unknown command: %q\n", cmdname)
		os.Exit(2)
	}

	err := cmd.Run(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// vim: noet tw=72 ts=3 sw=3
