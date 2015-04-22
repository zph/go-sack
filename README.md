# sack

Search and edit those results instantly. (Named comes from search-ack, now ag).

## Demonstration

  [sack demo](http://showterm.io/8f3421bae1d48c2109e1d#fast)

## Details

Search for specific terms in files using AG (the-silver-searcher) and then jump to instantly edit those lines in VIM.

Sack outputs those shortcuts into `~/.sack_shortcuts` for easy editing/re-use across any terminal on the system.

It's a reimplementation of [sack](https://github.com/zph/sack), which was a ruby re-implementation of the original bash script [sack](https://github.com/sampson-chen/sack).

## Installation (on Mac)

`wget -o ~/bin/sack https://github.com/zph/go-sack/releases/download/v0.3.0/sack && chmod +x ~/bin/sack`

Or for the reckless:

```
curl -sSL https://github.com/zph/go-sack/raw/master/src/install | bash -s
```

## Installation on other platforms

Install script will try to match for the right version of binary (only amd64 arch):

```
curl -sSL https://github.com/zph/go-sack/raw/master/src/install | bash -s
```

Check for your platform's bin on repo's [Release Page](https://github.com/zph/go-sack/releases).

Install that into somewhere in your $PATH and chmod +x that puppy.

OR:

Install from source using usual go build. Clone it && `make build` or `make install`.

For shell integration (sets up two aliases) add the following to your `.zshrc or .bashrc`

```
eval "$(sack init)"
```

## Usage

To search in a directory:

    sack -s SEARCH_TERM [DIRECTORY: defaults to current]

    OR (with shell integration)

    S SEARCH_TERM [DIRECTORY: defaults to current]

To edit one of those results by index number (zero indexed):

    sack -e 1

    OR (with shell integration)

    F 1

To display current shortcuts:

    sack -p

### Options

```
NAME:
   Sack - sack [searchterm] [optional directory]

USAGE:
   Sack [global options] command [command options] [arguments...]

VERSION:
   0.3.1

COMMANDS:
   init		shell init script
   eval		shell eval command to insert into .{zsh,bash}rc
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --edit, -e		edit a given shortcut
   --search, -s		search-ack/ag it
   --print, -p		display existing shortcuts
   --debug, -d		show all the texts
   --flags, -f '-i'	flags to pass to ag
   --version, -v	print the version
   --help, -h		show help
   

```

## Credit

  Original idea & implementation belong to @sampson-chen:
  https://github.com/sampson-chen/sack.

  Rewritten for cleanliness in go-lang for speed b/c Ruby standalone scripts with dependencies aren't convenient, & b/c Shell scripts past a certain length are unwieldy.

## Dependencies

  - ag [the-silver-searcher](https://github.com/ggreer/the_silver_searcher)
    - Falls back to use grep... but don't do that :P.

## License

MIT License (MIT)
Copyright (c) 2014 Zander Hill

## Contributing

  Input and pull requests are welcomed!

1. Fork it ( http://github.com/zph/go-sack/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
