package main

import (
	"fmt"
	"os"

	"github.com/jpillora/opts"
	"github.com/wxio/wxb/internal/wxb"
)

var (
	Version string
	Data    string
	Commit  string
)

type root struct {
}

func main() {
	r := root{}
	ro := opts.New(&r).
		Complete("wxb").
		Version(Version)

	wxb.Register(ro)

	err := ro.Parse().Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run error %v\n", err)
		os.Exit(2)
	}
}
