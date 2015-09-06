# Gotfiles
Dotfiles made (more) sane.

# Requirements
* Go

# Setup
Run `make && sudo make install` and you should be good to go.

# Usage
* `gf init [name]` creates a new dotfiles repo with a given name. This dir will be located in `~/gotfiles/[name]`
* `gf default [name]` sets the dotfiles repo with the specified name as the default one.
* `gf add ?[name] [file1/dir1] [file2/dir2]...` adds a file or directory to the dotfile repo with specified name. [name] is an optional argument, and if not given the default repo will be used.
* `gf sync ?[name]` synchronizes the dotfiles in your home directory with the repo with [name]. [name] is an optional argument, and if not given the default repo will be used
* `gf deploy -f ?[name] ?[file/dir]` deploys files from the repo to the home dir. If [name] is given, that repo is used, otherwise the default is used. 
If ?[file/dir] is given the specified file or dir will be deployed. If -f is used, files will be overwritten in the home dir, otherwise they will not be.
