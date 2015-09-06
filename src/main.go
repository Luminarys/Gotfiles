package main

import (
	"os"
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

	if cmd == "default" {
		if argNum == 2 {
			if exists(gfDir + "/" + args[1]) {
				if exists(gfDir + "/.default") {
					os.Remove(gfDir + "/.default")
				}
				if err := os.Symlink(gfDir+"/"+args[1], gfDir+"/.default"); err != nil {
					//Log error
				}
			}
		}
	}

	if cmd == "sync" {
		// rsync this because I'm too lazy to think of anything better
		if argNum == 2 {
			if exists(gfDir + "/" + args[1]) {
				if err := sync(gfDir + "/" + args[1]); err != nil {
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

	if cmd == "deploy" {
		force := false
		if exists, ind := stringInSliceInd("-f", args); exists {
			args = append(args[:ind], args[ind+1:]...)
			argNum--
			force = true
		}
		// Deploys the entire default repo
		if argNum == 1 {
			if exists(gfDir + "/.default") {
				deploy(gfDir+"/.default/", force)
			} else {
				// Log error
			}
		} else if argNum == 2 {
			// If a repo or a specific file/dir has been specified
			if exists(gfDir + "/" + args[1]) {
				deploy(gfDir+"/"+args[1]+"/", force)
			} else {
				if exists(gfDir + "/.default/" + args[1]) {
					deploy(gfDir+"/.default/"+args[1], force)
				} else {
					// Log error
				}
			}
		} else {
			if exists(gfDir+"/"+args[1]) && exists(gfDir+"/"+args[1]+"/"+args[2]) {
				deploy(gfDir+"/"+args[1]+"/"+args[2], force)
			} else {
				// Log error
			}
		}
	}
}
