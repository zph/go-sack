sack := bin/sack
binsack := ~/bin/sack

gb := vendor/bin/gb

.PHONY: build test

build: $(gb) clean
	$(gb) build all

test: $(gb)
	$(gb) test all

$(gb):
	bash -c 'GOPATH="$$(pwd)/vendor" go get github.com/constabulary/gb/...' && \
	bash -c 'GOPATH="$$(pwd)/vendor" go install github.com/constabulary/gb/...'

search: $(sack)
	$(sack) -s ruby ${HOME}/.zsh.d/

alias_search:
	$(sack) -s cmd_non-existent ~/bin

edit: $(sack)
	$(sack) -e 0 5

print: $(sack)
	$(sack) -p

lint:
	./bin/go-lint

hooks:
	cp -f hooks/* .git/hooks/

# build_all: clean
# 	go build -o bin/sack main.go; \
# 		GOARCH=amd64 GOOS=linux go build -o bin/sack.linux_amd64 main.go; \
# 		GOARCH=amd64 GOOS=freebsd go build -o bin/sack.freebsd_amd64 main.go; \
# 		go build -gcflags '-N' -o bin/sack.debug main.go;

clean:
	rm -f bin/*

install: clean build
	cp $(sack) $(binsack) && chmod +x $(binsack)

readme: clean build
	ruby -rerb -e "puts ERB.new(File.read('scripts/README.md.erb')).result" > README.md \
		&& cat README.md

