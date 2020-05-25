EXECUTABLE=client

all: client

client:
	go build -o ./build/client ./cmd/client

clean:
	rm -rf ./build/*
