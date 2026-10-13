N ?= 1000000
BALANCE ?= 1000
BET ?= 10
HANDS ?= 10000

.PHONY: run bankroll build test bench fmt vet check clean

run:
	go run ./cmd/poker -n $(N)

bankroll:
	go run ./cmd/bankroll -balance $(BALANCE) -bet $(BET) -hands $(HANDS)

build:
	go build -o bin/poker ./cmd/poker
	go build -o bin/bankroll ./cmd/bankroll

test:
	go test ./...

bench:
	go test -bench . -run '^$$' ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

clean:
	rm -rf bin
