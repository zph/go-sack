search: pkg/sack
	./pkg/sack -s ruby ${HOME}/.zsh.d/

edit: pkg/sack
	./pkg/sack -e 0

print: pkg/sack
	./pkg/sack -p

lint:
	./bin/go-lint

build: clean
	go build -o pkg/sack main.go

clean:
	rm -f pkg/*

install: clean build
	cp pkg/sack ~/bin/sack

readme: clean build
	ruby -rerb -e "puts ERB.new(File.read('src/README.md.erb')).result" > README.md

