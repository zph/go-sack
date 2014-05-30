search:
	go run sack.go -s ruby ${HOME}/.zsh.d/

edit:
	go run sack.go -e 0

print:
	go run sack.go -p

lint:
	./bin/go-lint

build:
	go build sack.go
