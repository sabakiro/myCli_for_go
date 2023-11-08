package main

import (
	"os"
	"slimple_cli/add"
	"slimple_cli/new"
	"slimple_cli/watch"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("no enough args")
	}

	args := make([]string, len(os.Args))
	for i := 1; i < len(os.Args); i++ {
		args[i-1] = os.Args[i]
	}
	switch strings.ToLower(args[0]) {
	case "run": //create watcher
		watch.Instance.Run()
	case "new": //create package
		new.Instance.Run(args[1])

	case "add": //add mod
		add.Instance.Run(args[1])
	}

}
