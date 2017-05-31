package main

import (
	"os/exec"
	"strings"
)

type Tmux struct {
	Path    string
	Options []string
}

func NewTmux() (tmux *Tmux, err error) {
	path, err := exec.LookPath("tmux")
	if err != nil {
		return
	}
	return &Tmux{
		Path:    path,
		Options: []string{},
	}, err
}

func (t *Tmux) Run(args []string) error {
	return exec.Command(
		t.Path, "split-window", strings.Join(args, " "),
	).Run()
}
