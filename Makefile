search: dist/sack
	go run main.go -s ruby ${HOME}/.zsh.d/

edit: dist/sack
	go run main.go -e 0

print: dist/sack
	go run main.go -p

lint:
	./bin/go-lint

hooks:
	cp -f hooks/* .git/hooks/

build: clean
	go build -o dist/sack main.go; \
		GOARCH=amd64 GOOS=linux go build -o dist/sack.linux_amd64 main.go; \
		GOARCH=amd64 GOOS=freebsd go build -o dist/sack.freebsd_amd64 main.go; \
		go build -gcflags '-N' -o dist/sack.debug main.go;

clean:
	rm -f dist/*

install: clean build
	cp dist/sack ~/bin/sack

readme: clean build
	ruby -rerb -e "puts ERB.new(File.read('src/README.md.erb')).result" > README.md \
		&& cat README.md

