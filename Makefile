.PHONY: run build clean

# run the app
run:
	go run cmd/chatnode/main.go

build:
	go build -o build/chatnode cmd/chatnode/main.go

clean:
	rm build/*
