upaste
======
upaste is a dead-simple pastebin with zero dependencies, designed to be used from the terminal.

```
upaste(1)                          UPASTE                          upaste(1)
NAME
    upaste: simple command line pastebin.

SYNOPSIS
    <command> | curl -F 'upaste=<-' <server url>

DESCRIPTION
    upaste is a dead-simple pastebin with zero dependencies,
    meant to be used from the terminal.

    To upload from a browser, go here:
    <server url>/upload

    This server is configured to delete pastes after 100 days.

EXAMPLES
    $ echo "Hello, upaste!" | curl -F 'upaste=<-' <server url>
        Upload the text "Hello, upaste!" to this server

    $ curl <server url> | less
        Display this page in a terminal

```

Getting Started
---------------
```
$ git clone https://github.com/shawnanastasio/upaste && cd upaste
$ go build
$ cp config.sample.json config.json
$ $EDITOR config.json
$ ./upaste
```

License
-------
upaste is licensed under the GNU GPL v3. For the full license text, see LICENSE.md.
Contributions are welcome.
