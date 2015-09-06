package main

import (
	"os"
)

func initFiles(name string) {
	if !exists(gfDir) {
		if err := os.Mkdir(gfDir, 0755); err != nil {
			//Log error
		}
	}
	if !exists(name) {
		if err := os.Mkdir(name, 0755); err != nil {
			//Log error
		}
	}
	if !exists(gfDir + "/.default") {
		if err := os.Symlink(name, gfDir+"/.default"); err != nil {
			//Log error
		}
	}
}
