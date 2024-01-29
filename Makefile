DEPTH=3

fmt:
	go fmt ./...

perft:
	cd board && go test -v -run 'Perft' -perft true

play:
	go run cmd/main.go -depth $(DEPTH)

profile:
	go run cmd/main.go -cpuprofile cpu.prof -depth $(DEPTH)

profile-web: profile
	go tool pprof -http=":8000" cpu.prof

test:
	go test -v ./...

