package main

import (
	"flag"
	"k8s.io/klog"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "show your name", "help name")
	flag.StringVar(&name, "n", "show your name n", "help n")
	flag.Parse()

	klog.Infof("name: %s", name)
}
