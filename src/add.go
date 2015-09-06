package main

import (
	"os"
	"strings"
)

func addPath(confDir, path string) {
	fullPath := strings.Split(path, "/")
	prePath := strings.Join(fullPath[0:len(fullPath)-1], "/")
	if err := os.MkdirAll(confDir+"/"+prePath, 0755); err != nil {
		// Log error
	}
	if info, err := os.Stat(homeDir + "/" + path); err == nil && info.IsDir() {
		err = CopyTree(homeDir+"/"+path, confDir+"/"+path, &CopyTreeOptions{Symlinks: true, IgnoreDanglingSymlinks: true, CopyFunction: Copy, Ignore: nil})
		if err != nil {
			// Log error
		}
	} else if err == nil {
		_, err = Copy(homeDir+"/"+path, confDir+"/"+path, true)
		if err != nil {
			// Log error
		}
	} else {

	}
}
