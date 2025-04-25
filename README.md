Local Development
=================

To get started, run the following commands:

```bash
$ git clone https://github.com/gregdayne/auth1.git
$ cd auth1
$ make run
```

The `make run` command will build and run the application.

The command `make dev` will run the application in development mode, which will
use a constant port.

[Root](http://0.0.0.0:4000/)
[OpenAPI](http://0.0.0.0:4000/docs)


## Generate TLS files for local dev

_Path for the generate_cert.go module will be dependent on your install_

```sh
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

or using asdf

```sh
go run ~/.asdf/installs/golang/1.24.0/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```


Testing
=======

To run the tests, run the following command:

```bash
$ make audit
```

This will run the tests and perform various static analysis checks, such as
linting and verifying the module dependencies.


Building
========

To build the application, run the following command:

```bash
$ make build
```

This will build the application and output the binary to the current directory.


Tidying
=======

To tidy the application, run the following command:

```bash
$ make tidy
```

This will format the `.go` files, tidy the module dependencies, and vendor the
dependencies.

