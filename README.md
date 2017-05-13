Mangaeden
=========

**mangaeden** is a simple lookup and download cli client for mangaeden.com API in Go.

*This is a project to practice writing Go.*

Install
-------

`go install github.com/kmwenja/mangaeden/cmd/`

Usage
-----

**List all the manga**

`mangaeden list > mangalist`

**Find the manga you're looking for**

`grep "manga name" mangalist`

**See info about that manga**

`mangaeden info [manga id]`

**Download the manga**

`mangaeden download [manga id]`

Hacking
-------

1. Get the code: `go get -v github.com/kmwenja/mangaeden`
2. Cd into the folder: `cd $GOPATH/src/github.com/kmwenja/mangaeden/`
3. Hack hack hack.
4. Build binary: `cd cmd && go build -o mangaeden`

`github.com/kmwenja/mangaeden` has the client and supporting structs.

`github.com/kmwenja/mangaeden/cmd` has the cli code, built with [Cobra](https://github.com/spf13/cobra).
