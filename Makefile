search: dist/sack
	go run main.go -s ruby ${HOME}/.zsh.d/

edit: dist/sack
	go run main.go -e 0

print: dist/sack
	go run main.go -p

lint:
	./bin/go-lint

build: clean
	go build -o dist/sack main.go

clean:
	rm -f dist/*

install: clean build
	cp dist/sack ~/bin/sack

readme: clean build
	ruby -rerb -e "puts ERB.new(File.read('src/README.md.erb')).result" > README.md \
		&& cat README.md

