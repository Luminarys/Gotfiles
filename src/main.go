package main

import (
	"os"
	"os/exec"
)

var homeDir string = os.Getenv("HOME")
var gfDir string = homeDir + "/gotfiles"

func main() {
	args := os.Args[1:]
	argNum := len(args)
	cmd := args[0]

	if cmd == "init" {
		if argNum == 2 {
			initFiles(gfDir + "/" + args[1])
		} else {
			// Log error
		}
	}

	if cmd == "add" {
		if argNum >= 2 {
			if exists(gfDir + "/" + args[1]) {
				for _, p := range args[2:] {
					addPath(gfDir+"/"+args[1], p)
				}
			} else {
				if exists(gfDir + "/.default") {
					for _, p := range args[1:] {
						addPath(gfDir+"/.default", p)
					}
				} else {
					// Log error
				}
			}
		} else {
			// Log error
		}
	}

	if cmd == "sync" {
		// rsync this because I'm too lazy to think of anything better
		if argNum == 2 {
			if exists(gfDir + "/" + args[1]) {
				cmd := "rsync"
				args := []string{"-r", "--existing", homeDir + "/", gfDir + "/" + args[1]}
				if err := exec.Command(cmd, args...).Run(); err != nil {
					os.Exit(1)
				}
				if err := sync(gfDir + "/.default"); err != nil {
					// Log error
				}
			} else {
				// Log error
			}
		} else {
			if exists(gfDir + "/.default") {
				if err := sync(gfDir + "/.default"); err != nil {
					// Log error
				}
			} else {
				// Log error
			}
		}
	}
}
