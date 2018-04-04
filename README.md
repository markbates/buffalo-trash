# Buffalo Trash

This plugin basically runs the following commands:

```bash
$ buffalo new <app-name> -f; cd <app-name>; buffalo db drop -d; buffalo db create -d;
```

## Installation

```bash
$ go get -u -v github.com/markbates/buffalo-trash
```

## Usage

```bash
$ buffalo trash <app-name>
```

**NOTE**: You must **NOT** be in your project directory when you run this. You should be directly above it.

```bash
$ pwd
$GOPATH/src/github.com/markbates

$ ls -la | rg <app-name>
drwxr-xr-x   26 markbates  staff       832 Apr  4 14:57 <app-name>
```

### Options

You can pass in all of the same options you can pass to `buffalo new`.

```bash
$ buffalo trash <app-name> --db-type=mysql --api
```

