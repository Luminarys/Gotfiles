# Gotfiles
Dotfiles made (more) sane.

# Requirements
* Go

# Setup
Run `make && sudo make install` and you should be good to go.

# Usage
* `gf init [name]` creates a new dotfiles repo with a given name. This dir will be located in `~/gotfiles/[name]`
* `gf add ?[name] [file1/dir1] [file2/dir2]...` adds a file or directory to the dotfile repo with specified name. [name] is an optional argument, and if not given the default repo will be used.
* `gf sync ?[name]` synchronizes the dotfiles in your home directory with the repo with [name]. [name] is an optional argument, and if not given the default repo will be used
