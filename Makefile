cfg = ./config/config.toml

run: building
	./build/remember

building:
	go build -o ./build/remember ./cmd/remember/main.go 

clean:
	rm ./build/remember