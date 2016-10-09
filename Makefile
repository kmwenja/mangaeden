all: build

build: manga-list manga-info manga-chapter

manga-list: ./cmd/manga-list/*.go
	go build -o ./bin/manga-list ./cmd/manga-list

manga-info: ./cmd/manga-info/*.go
	go build -o ./bin/manga-info ./cmd/manga-info

manga-chapter: ./cmd/manga-chapter/*.go
	go build -o ./bin/manga-chapter ./cmd/manga-chapter

clean:
	rm ./bin/*

.PHONY = clean

