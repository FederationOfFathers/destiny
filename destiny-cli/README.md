About
=====

destiny-cli is a command line interface to the [bungie-platform-destiny API](http://destinydevs.github.io/BungieNetPlatform/docs/Getting-Started)

It builds on the Go API client (work in progress) [available on github](https://github.com/FederationOfFathers/destiny/tree/master/destiny-cli)

Status
======

Working but not feature complete.  Documentation of commands and subcommands needs work.

Installation
============

Currently requires a working [Go development setup](https://golang.org/doc/install) I can release binaries if anyone cares enough to need them but is not themselves a developer (which seems unlikely as this tool just spits out gobs of JSON at you, mostly)

```
go get github.com/FederationOfFathers/destiny/destiny-cli
```

You will need [a working API key](https://www.bungie.net/en/User/API) for the program to be of much use

Documentation
=============

[see docs/md/destiny-cli.md](docs/md/destiny-cli.md)
