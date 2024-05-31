DEPTH=3

build:
	go build -o baka cmd/main.go

fmt:
	go fmt ./...

perft:
	cd board && go test -v -run 'Perft' -perft true

play:
	go run cmd/main.go -selfPlay true

profile:
	go run cmd/main.go -cpuprofile cpu.prof -depth $(DEPTH)

profile-web: profile
	go tool pprof -http=":8000" cpu.prof

test:
	go test -v ./...

