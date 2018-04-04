# Buffalo Trash

This plugin basically runs the following commands:

```bash
$ buffalo new coke -f; cd coke; buffalo db drop -d; buffalo db create -d;
```

## Installation

```bash
$ go get -u -v github.com/markbates/buffalo-trash
```

## Usage

```bash
$ buffalo trash coke
```

**NOTE**: You must **NOT** be in your project directory when you run this. You should be directly above it.

```bash
$ pwd
$GOPATH/src/github.com/markbates

$ ls -la | rg coke
drwxr-xr-x   26 markbates  staff       832 Apr  4 14:57 coke
```

### Options

You can pass in all of the same options you can pass to `buffalo new`.

```bash
$ buffalo trash coke --db-type=mysql --api
```

