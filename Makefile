search:
	go run main.go -s ruby ${HOME}/.zsh.d/

edit:
	go run main.go -e 0

print:
	go run main.go -p

lint:
	./bin/go-lint

build:
	go build main.go sack
