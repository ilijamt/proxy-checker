Proxy Checker
=============

Check if you proxy is working or not very easily

## Getting Started with Proxy Checker

### Requirements

* [Golang](https://golang.org/dl/) >= 1.7
* [Glide](https://github.com/Masterminds/glide) >= 0.12.3

### Autocomplete

You know what to do with this

* [Bash](contrib/proxy-checker.bash)
* [Zsh](contrib/proxy-checker.zsh)

### Dependencies 

This project uses glide to manage dependencies so download them before trying to build/install by running 

```bash
glide install
```

### Build

To build the binary for Proxy Checker run the command below. This will generate a binary
in the bin directory with the name proxy-checker.

```bash
make build
```

### Install

To install the binary for Proxy Checker run the command below. This will generate a binary
in $GOPATH/bin directory with the name proxy-checker.

```bash
make install
```

### Run

### Help
```bash
$ proxy-checker --help
usage: proxy-checker [<flags>] <command> [<args> ...]

Checks if an http proxy with basic auth works by querying https://api.ipify.org/

Flags:
  -h, --help         Show context-sensitive help (also try --help-long and --help-man).
  -v, --version      Show version and terminate
      --queue=25     How many request to process at one time
      --failed-only  Show only failed proxies

Commands:
  help [<command>...]
    Show help.

  check <host-port> [<username>] [<password>]
    Check the single proxy

  csv-file <name>
    Check all the proxies in the file specified
```

