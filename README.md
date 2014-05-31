# go-sack

Search and edit those results instantly.

## Details

Search for specific terms in files using AG (the-silver-searcher) and then jump to instantly edit those lines in VIM.

Sack outputs those shortcuts into `~/.sack_shortcuts` for easy editing across any terminal on the system.

It's a reimplementation of [sack](https://github.com/zph/sack), which was a ruby re-implementation of the original bash script [sack](https://github.com/sampson-chen/sack).

## Installation

`wget -o ~/bin/sack https://github.com/zph/go-sack/raw/master/pkg/sack && chmod +x ~/bin/sack`

For shell integration (sets up two aliases) add the following to your `.zshrc or .bashrc`

`eval "$(sack init)"`

``
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

## Credit

  Original idea & implementation belong to @sampson-chen:
  https://github.com/sampson-chen/sack.
    
  Rewritten for cleanliness in go-lang for speed b/c Ruby standalone scripts with dependencies aren't convenient, & b/c Shell scripts past a certain length are unwieldy.

## License

  2014 (TM) Zander Hill

## Contributing

  Input and pull requests are welcomed!

1. Fork it ( http://github.com/zph/go-sack/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


