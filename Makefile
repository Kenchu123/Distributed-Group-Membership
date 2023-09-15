all: build

build:
	mkdir -p bin
	go build -o bin/msl main.go

run:
	go run main.go $(ARGS)

clean:
	rm -rf bin
