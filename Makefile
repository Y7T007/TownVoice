build:
	 go build -o app ./main.go
	chmod +x app


run:
	nodemon -w . --exec "go run" ./cmd/TownVoice/main.go