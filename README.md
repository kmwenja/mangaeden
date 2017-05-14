Mangaeden
=========

**mangaeden** is a simple lookup and download cli client for the [mangaeden.com API](www.mangaeden.com/api/) in Go.

Make sure to checkout [Mangaeden](www.mangaeden.com). All manga content belongs to them.

Install
-------

**Prebuilt binaries**

See [Releases](https://github.com/kmwenja/mangaeden/releases) for Linux, Windows and Mac OS X binaries.

**From Source**

1. `go get -v github.com/kmwenja/mangaeden/cmd/mangaeden/`
2. `go install github.com/kmwenja/mangaeden/cmd/mangaeden/`

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

Dependencies
------------

- go
- make

Hacking
-------

1. Get the code: `go get -v github.com/kmwenja/mangaeden/cmd/mangaeden/`
2. Cd into the folder: `cd $GOPATH/src/github.com/kmwenja/mangaeden/`
3. Hack hack hack.
4. Build binary: `make`
5. Run the binary: `bin/mangaeden`
6. Make release: `make dist`

`github.com/kmwenja/mangaeden` has the client and supporting structs.
`github.com/kmwenja/mangaeden/cmd/mangaeden/` has the cli code, built with [Cobra](https://github.com/spf13/cobra).

License
-------

MIT
