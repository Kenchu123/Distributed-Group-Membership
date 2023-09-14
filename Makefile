all: build

build:
	mkdir -p bin
	go build -o bin/membership-server cmd/main.go

run:
	go run cmd/main.go $(ARGS)

clean:
	rm -rf bin
