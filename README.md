# libvirt-go [![Build Status](https://travis-ci.org/rgbkrk/libvirt-go.svg?branch=master)](https://travis-ci.org/rgbkrk/libvirt-go) [![GoDoc](https://godoc.org/gopkg.in/alexzorin/libvirt-go.v2?status.svg)](http://godoc.org/gopkg.in/alexzorin/libvirt-go.v2)

Go bindings for libvirt.

Make sure to have `libvirt-dev` package (or the development files otherwise somewhere in your include path)

## Version Support

The libvirt go package provides API coverage for libvirt versions
from 1.3.1 onwards, through conditional compilation of newer APIs.

## Documentation

* [api documentation for the bindings](http://godoc.org/github.com/rgbkrk/libvirt-go)
* [api documentation for libvirt](http://libvirt.org/html/libvirt-libvirt.html)

## Contributing

Please fork and write tests.

Integration tests are available where functionality isn't provided by the test driver, see `integration_test.go`.

A `Vagrantfile` is included to run the integration tests:

* `cd ./vagrant`
* `vagrant up` to provision the virtual machine
* `vagrant ssh` to login to the virtual machine

Once inside, `sudo su -` and `go test -tags integration libvirt`.
