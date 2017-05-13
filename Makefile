all: build

build: *.go cmd/mangaeden/*.go
	mkdir -p bin
	cd cmd/mangaeden/ && go build
	mv cmd/mangaeden/mangaeden bin/mangaeden

clean:
	rm -r bin

.PHONY = clean
