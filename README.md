Hello World
===========

A simple, from scratch, Go WebApp skeleton.

**Note:** This is a work in progress and should not be taken as a good example
(yet).

## Flags

The application supports the following command line flags:

* `-h`: Print usage instructions
* `-c`: [Configuration file](#configurations)

## Configurations

The application supports a configuration file using YAML serialization standard,
which should be passed via command line using the `-c` flag.

The supported configurations are described bellow:

* `Port`: Port number where HTTP server will listen at (e.g. `80`).

   Remember that using port numbers less than 1024 will make `sudo` require to
   run the application.

## How to build

```
$ mkdir workspace && cd workspace
$ export GOPATH=$(pwd)
$ go get github.com/PauloASilva/hello-world
```

## How to run

```
$ cd $GOPATH
$ sudo bin/hello-world -c src/github.com/PauloASilva/config/config.yml
```

If you want to provide your own configuration file, right now it is as simple
as

```yaml
port: 8080
```

Save it somewhere (e.g. `~/tmp/my.yml`) and than lunch the application

```
$GOPATH/bin/hello-world -c /tmp/my.yml
```

### Available routers

* `/`: Application skeleton home
* `/about`: Application skeleton about
* `/m1`: Module 1 (aka 1st Module) home
* `/m1/about`: About Module 1
