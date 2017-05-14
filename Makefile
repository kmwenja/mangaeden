VERSION := $(shell sh -c 'git describe --always --tags')
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

all: build

build: *.go cmd/mangaeden/*.go clean-build
	cd cmd/mangaeden/ && go build -o ../../bin/mangaeden $(LDFLAGS)

define go_dist
	mkdir -p dist/mangaeden-$(VERSION)_$1_$2/
	cd dist/mangaeden-$(VERSION)_$1_$2/ && GOOS=$1 GOARCH=$2 go build -v $(LDFLAGS) github.com/kmwenja/mangaeden/cmd/mangaeden
	if [ "$1" = "windows" ]; then \
		cd dist/ && zip mangaeden-$(VERSION)_$1_$2.zip mangaeden-$(VERSION)_$1_$2/*; \
	else \
		cd dist/ && tar -cvzf mangaeden-$(VERSION)_$1_$2.tar.gz mangaeden-$(VERSION)_$1_$2/*; \
	fi
endef

dist: clean-dist dist-linux dist-macos dist-windows

dist-macos:
	$(call go_dist,darwin,amd64)

dist-linux:
	$(call go_dist,linux,amd64)
	$(call go_dist,linux,386)

dist-windows:
	go get -v github.com/inconshreveable/mousetrap
	$(call go_dist,windows,amd64)

clean: clean-build clean-dist

clean-build:
	rm -rf ./bin

clean-dist:
	rm -rf ./dist

.PHONY = clean clean-build clean-dist dist dist-linux dist-macos dist-windows
