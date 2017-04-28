// +build !without_lxc

/*
 * This file is part of the libvirt-go project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (c) 2013 Alex Zorin
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirt

/*
#cgo pkg-config: libvirt
// Can't rely on pkg-config for libvirt-lxc since it was not
// installed until 2.6.0 onwards
#cgo LDFLAGS: -lvirt-lxc
#include <libvirt/libvirt.h>
#include <libvirt/libvirt-lxc.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

func (d *Domain) LxcOpenNamespace(flags uint32) ([]os.File, error) {
	var cfdlist *C.int

	ret := C.virDomainLxcOpenNamespace(d.ptr, &cfdlist, C.uint(flags))
	if ret == -1 {
		return []os.File{}, GetLastError()
	}
	fdlist := make([]os.File, ret)
	for i := 0; i < int(ret); i++ {
		var cfd C.int
		cfd = *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(cfdlist)) + (unsafe.Sizeof(cfd) * uintptr(i))))
		fdlist[i] = *os.NewFile(uintptr(cfd), "namespace")
	}
	defer C.free(cfdlist)
	return fdlist, nil
}

func (d *Domain) LxcEnterNamespace(fdlist []os.File, flags uint32) ([]os.File, error) {
	var coldfdlist *C.int
	var ncoldfdlist C.uint
	cfdlist := make([]C.int, len(fdlist))
	for i := 0; i < len(fdlist); i++ {
		cfdlist[i] = C.int(fdlist[i].Fd())
	}

	ret := C.virDomainLxcEnterNamespace(d.ptr, C.uint(len(fdlist)), &cfdlist[0], &ncoldfdlist, &coldfdlist, C.uint(flags))
	if ret == -1 {
		return []os.File{}, GetLastError()
	}
	oldfdlist := make([]os.File, ncoldfdlist)
	for i := 0; i < int(ncoldfdlist); i++ {
		var cfd C.int
		cfd = *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(coldfdlist)) + (unsafe.Sizeof(cfd) * uintptr(i))))
		oldfdlist[i] = *os.NewFile(uintptr(cfd), "namespace")
	}
	defer C.free(coldfdlist)
	return oldfdlist, nil
}
