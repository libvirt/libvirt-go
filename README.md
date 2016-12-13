# libvirt-go [![Build Status](https://travis-ci.org/libvirt/libvirt-go.svg?branch=master)](https://travis-ci.org/libvirt/libvirt-go) [![GoDoc](https://godoc.org/github.com/libvirt/libvirt-go?status.svg)](https://godoc.org/github.com/libvirt/libvirt-go)

Go bindings for libvirt.

Make sure to have `libvirt-dev` package (or the development files otherwise somewhere in your include path)

## Version Support

The libvirt go package provides API coverage for libvirt versions
from 1.2.0 onwards, through conditional compilation of newer APIs.

## Documentation

* [api documentation for the bindings](https://godoc.org/github.com/libvirt/libvirt-go)
* [api documentation for libvirt](http://libvirt.org/html/libvirt-libvirt.html)

## Contributing

Please fork and write tests.

Integration tests are available where functionality isn't provided by the test driver, see `integration_test.go`.

A `Vagrantfile` is included to run the integration tests:

* `cd ./vagrant`
* `vagrant up` to provision the virtual machine
* `vagrant ssh` to login to the virtual machine

Once inside, `sudo su -` and `go test -tags integration libvirt`.
