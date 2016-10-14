all: build

build: manga-list manga-info manga-chapter manga-download

manga-list: ./cmd/manga-list/*.go
	go build -o ./bin/manga-list ./cmd/manga-list

manga-info: ./cmd/manga-info/*.go
	go build -o ./bin/manga-info ./cmd/manga-info

manga-chapter: ./cmd/manga-chapter/*.go
	go build -o ./bin/manga-chapter ./cmd/manga-chapter

manga-download: ./cmd/manga-download/*.go
	go build -o ./bin/manga-download ./cmd/manga-download

clean:
	rm ./bin/*

.PHONY = clean

