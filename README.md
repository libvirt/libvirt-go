libvirt-go
============

Go bindings for libvirt.

Versions
--------------
Please use the v1.x branch for libvirt 0.9.8 and below: `gopkg.in/alexzorin/libvirt-go.v1` [(docs)](http://gopkg.in/alexzorin/libvirt-go.v1).

The 2.x branch targets the 1.x version of libvirt: `gopkg.in/alexzorin/libvirt-go.v2` [(docs)http://gopkg.in/alexzorin/libvirt-go.v2].

Make sure to have libvirt-dev package (or the development files otherwise somewhere in your include path)

Documentation
--------------

* [api documentation for the bindings](http://godoc.org/github.com/alexzorin/libvirt-go)
* [api documentation for libvirt](http://libvirt.org/html/libvirt-libvirt.html)

Contributing
-------------

Please fork and write tests.

Integration tests are available where functionality isn't provided by the test driver, see `integration_test.go` for more info.

