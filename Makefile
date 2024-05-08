build:
	go build -o bin/main -v ./

run:
	nodemon -w . --exec "go run" ./cmd/TownVoice/main.go