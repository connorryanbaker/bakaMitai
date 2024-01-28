fmt:
	go fmt ./...

perft:
	cd board && go test -v -run 'Perft' -perft true

profile:
	go run cmd/main.go -cpuprofile cpu.prof

profile-web: profile
	go tool pprof -http=":8000" cpu

test:
	go test -v ./...

