package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	RegisterCmd(&upCmd{})
}

type upCmd struct {
}

func (c upCmd) Name() string {
	return "up"
}

func (c upCmd) Desc() string {
	return "Brings OpenVPN up."
}

func (c upCmd) Run(config Config) error {
	if !chkroot() {
		return &NeedRootError{
			Cmd: c.Name(),
		}
	}

	openvpn := exec.Command("openvpn", "--config", config.Config)
	openvpn.Stdout = os.Stdout
	openvpn.Stderr = os.Stderr

	err := openvpn.Run()
	if err != nil {
		return fmt.Errorf("OpenVPN failed: %v", err)
	}

	return nil
}
