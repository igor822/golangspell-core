# golangspell-core
Core Plugin with the main Golang Spell commands

## Golang Spell
The Core project contains the core commands (and the respective templates) of the platform [Golang Spell](https://github.com/danilovalente/golangspell).

## Environment
# In order to properly build the project export the current environment variable before building:
export GO111MODULE=off
# or on Windows
set GO111MODULE=off

![alt text](https://golangspell.com/golangspell/blob/master/img/gopher_spell.png?raw=true)

## Install
To install the golangspell-core spell use the command
golangspell addspell github.com/danilovalente/golangspell-core golangspell-core

## Update
To update the golangspell-core version use the command
golangspell updatespell github.com/danilovalente/golangspell-core golangspell-core

## Version history
* 0.4.2 - Removed unused imports from template
* 0.4.1 - Fixed directory and naming issues
* 0.4.0 - Updated echo to the v4, added Prometheus to init template, created feature addspellcommand
