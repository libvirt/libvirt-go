# libvirt-go [![Build Status](https://travis-ci.org/rgbkrk/libvirt-go.svg?branch=master)](https://travis-ci.org/rgbkrk/libvirt-go) [![GoDoc](https://godoc.org/gopkg.in/alexzorin/libvirt-go.v2?status.svg)](http://godoc.org/gopkg.in/alexzorin/libvirt-go.v2)

Go bindings for libvirt.

Make sure to have `libvirt-dev` package (or the development files otherwise somewhere in your include path)

## Version Support

The minimum required version of libvirt is **2.4.0**. Due to the
API/ABI compatibility promise of libvirt, more recent versions of
libvirt should work too.

The master branch of libvirt-go will always aim to support all
the APIs present in the master branch of the libvirt core library
API. At time of release, a tag will be created, so applications
that need build compatibility with an older version of libvirt
should checkout the tag corresponding to the version that they
require. This policy is appies from version 2.4.0 onwards. If
compatibility is required against a version of libvirt prior to
2.4.0, it will be necessary to create a branch and strip out
APIs.

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
