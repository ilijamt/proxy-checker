Proxy Checker
=============

Check if you proxy is working or not very easily

## Installation

If you don't want to compile your own version, you can use the following repository to install it 

### Debian

```bash
echo "deb http://packages.matoski.com/ debian main" | sudo tee /etc/apt/sources.list.d/packages-matoski-com.list
curl -s http://packages.matoski.com/keyring.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install proxy-checker
```

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
in $GOPATH/bin directory with the name proxy-checker and add the bash autocomplete files.

```bash
make install
```

### CSV file

The csv file structure is like this, you can add as many as you want

```csv
<schema>://<host>[:<port>],<username>,<password>
```

* **\<schema>** is http, or https
* **\<host>** is the actually host of the proxy, or it can even be an IP address
* **\<port>** is the port of the proxy, if not supplied it defaults to 80
* **\<username>** is the username for the proxy
* **\<password>** is the password for the proxy

## Run

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

### Single proxy

```bash
proxy-checker check http://example.com:60000 username password
http://username:password@example.com:60000 is valid
```

### Checking multiple proxies

Where **proxies.csv** contains

```csv
http://proxy01.com:60000,username1,password1
http://proxy02.com:9000,username2,password2
http://proxy03.com:14000,username3,password3
http://proxy04.com:60000,username4,password4
http://proxy05.com,username5,password5
```

And you can easily check all of them with this command

```bash
proxy-checker csv-file proxies.csv
http://username1:password1@proxy01.com:60000 is valid
http://username2:password2@proxy02.com:9000 is invalid
http://username3:password3@proxy03.com:14000 is invalid
http://username4:password4@proxy04.com:60000 is valid
http://username5:password5@proxy05.com is valid
```

## License

Apache License
Version 2.0, January 2004

See [License](LICENSE) file
