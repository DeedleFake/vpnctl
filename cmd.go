package main

import (
	"fmt"
	"sort"
)

type Cmd interface {
	Name() string
	Desc() string

	Run(Config) error
}

type Config struct {
	Config string
}

type cmdList struct {
	cmds []Cmd
}

var (
	cmds cmdList
)

func RegisterCmd(cmd Cmd) {
	cmds.cmds = append(cmds.cmds, cmd)
	sort.Sort(&cmds)
}

func GetCmd(name string) Cmd {
	i := sort.Search(len(cmds.cmds), func(i int) bool {
		return cmds.cmds[i].Name() >= name
	})

	if (i < len(cmds.cmds)) && (cmds.cmds[i].Name() == name) {
		return cmds.cmds[i]
	}

	return nil
}

func GetCmds() []Cmd {
	return cmds.cmds
}

func (c *cmdList) Len() int {
	return len(c.cmds)
}

func (c *cmdList) Swap(i1, i2 int) {
	c.cmds[i1], c.cmds[i2] = c.cmds[i2], c.cmds[i1]
}

func (c *cmdList) Less(i1, i2 int) bool {
	return c.cmds[i1].Name() < c.cmds[i2].Name()
}

type NeedRootError struct {
	Cmd string
}

func (err NeedRootError) Error() string {
	return fmt.Sprintf("Command %q needs root privileges.", err.Cmd)
}
