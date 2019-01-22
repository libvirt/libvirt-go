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
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include <string.h>
#include "typedparams_wrapper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type typedParamsFieldInfo struct {
	set *bool
	i   *int
	ui  *uint
	l   *int64
	ul  *uint64
	b   *bool
	d   *float64
	s   *string
	sl  *[]string
}

func typedParamsUnpackLen(cparams *C.virTypedParameter, nparams int, infomap map[string]typedParamsFieldInfo) (uint, error) {
	count := uint(0)
	for i := 0; i < nparams; i++ {
		var cparam *C.virTypedParameter
		cparam = (*C.virTypedParameter)(unsafe.Pointer(uintptr(unsafe.Pointer(cparams)) + unsafe.Sizeof(*cparam)*uintptr(i)))
		name := C.GoString((*C.char)(unsafe.Pointer(&cparam.field)))
		info, ok := infomap[name]
		if !ok {
			// Ignore unknown keys so that we don't break if
			// run against a newer libvirt that returns more
			// parameters than we currently have code to
			// consume
			continue
		}
		switch cparam._type {
		case C.VIR_TYPED_PARAM_INT:
			if info.i == nil {
				return 0, fmt.Errorf("field %s expects an int", name)
			}
			*info.i = int(*(*C.int)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_UINT:
			if info.ui == nil {
				return 0, fmt.Errorf("field %s expects a uint", name)
			}
			*info.ui = uint(*(*C.uint)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_LLONG:
			if info.l == nil {
				return 0, fmt.Errorf("field %s expects an int64", name)
			}
			*info.l = int64(*(*C.longlong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_ULLONG:
			if info.ul == nil {
				return 0, fmt.Errorf("field %s expects a uint64", name)
			}
			*info.ul = uint64(*(*C.ulonglong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_DOUBLE:
			if info.d == nil {
				return 0, fmt.Errorf("field %s expects a float64", name)
			}
			*info.d = float64(*(*C.double)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if info.b == nil {
				return 0, fmt.Errorf("field %s expects a bool", name)
			}
			*info.b = *(*C.char)(unsafe.Pointer(&cparam.value)) == 1
			*info.set = true
		case C.VIR_TYPED_PARAM_STRING:
			if info.s != nil {
				*info.s = C.GoString(*(**C.char)(unsafe.Pointer(&cparam.value)))
				*info.set = true
			} else if info.sl != nil {
				*info.sl = append(*info.sl, C.GoString(*(**C.char)(unsafe.Pointer(&cparam.value))))
				*info.set = true
			} else {
				return 0, fmt.Errorf("field %s expects a string/string list", name)
			}
		}
		count++
	}

	return count, nil
}

func typedParamsUnpack(cparams []C.virTypedParameter, infomap map[string]typedParamsFieldInfo) (uint, error) {
	return typedParamsUnpackLen(&cparams[0], len(cparams), infomap)
}

func typedParamsPackLen(cparams *C.virTypedParameter, nparams int, infomap map[string]typedParamsFieldInfo) error {
	stringOffsets := make(map[string]uint)

	for i := 0; i < nparams; i++ {
		var cparam *C.virTypedParameter
		cparam = (*C.virTypedParameter)(unsafe.Pointer(uintptr(unsafe.Pointer(cparams)) + unsafe.Sizeof(*cparam)*uintptr(i)))
		name := C.GoString((*C.char)(unsafe.Pointer(&cparam.field)))
		info, ok := infomap[name]
		if !ok {
			// Ignore unknown keys so that we don't break if
			// run against a newer libvirt that returns more
			// parameters than we currently have code to
			// consume
			continue
		}
		if !*info.set {
			continue
		}
		switch cparam._type {
		case C.VIR_TYPED_PARAM_INT:
			if info.i == nil {
				return fmt.Errorf("field %s expects an int", name)
			}
			*(*C.int)(unsafe.Pointer(&cparam.value)) = C.int(*info.i)
		case C.VIR_TYPED_PARAM_UINT:
			if info.ui == nil {
				return fmt.Errorf("field %s expects a uint", name)
			}
			*(*C.uint)(unsafe.Pointer(&cparam.value)) = C.uint(*info.ui)
		case C.VIR_TYPED_PARAM_LLONG:
			if info.l == nil {
				return fmt.Errorf("field %s expects an int64", name)
			}
			*(*C.longlong)(unsafe.Pointer(&cparam.value)) = C.longlong(*info.l)
		case C.VIR_TYPED_PARAM_ULLONG:
			if info.ul == nil {
				return fmt.Errorf("field %s expects a uint64", name)
			}
			*(*C.ulonglong)(unsafe.Pointer(&cparam.value)) = C.ulonglong(*info.ul)
		case C.VIR_TYPED_PARAM_DOUBLE:
			if info.d == nil {
				return fmt.Errorf("field %s expects a float64", name)
			}
			*(*C.double)(unsafe.Pointer(&cparam.value)) = C.double(*info.d)
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if info.b == nil {
				return fmt.Errorf("field %s expects a bool", name)
			}
			if *info.b {
				*(*C.char)(unsafe.Pointer(&cparam.value)) = 1
			} else {
				*(*C.char)(unsafe.Pointer(&cparam.value)) = 0
			}
		case C.VIR_TYPED_PARAM_STRING:
			if info.s != nil {
				*(**C.char)(unsafe.Pointer(&cparam.value)) = C.CString(*info.s)
			} else if info.sl != nil {
				count := stringOffsets[name]
				*(**C.char)(unsafe.Pointer(&cparam.value)) = C.CString((*info.sl)[count])
				stringOffsets[name] = count + 1
			} else {
				return fmt.Errorf("field %s expects a string", name)
			}
		}
	}

	return nil
}

func typedParamsPack(cparams []C.virTypedParameter, infomap map[string]typedParamsFieldInfo) error {
	return typedParamsPackLen(&cparams[0], len(cparams), infomap)
}

func typedParamsPackNew(infomap map[string]typedParamsFieldInfo) (*C.virTypedParameter, C.int, error) {
	var cparams C.virTypedParameterPtr
	var nparams C.int
	var maxparams C.int

	defer C.virTypedParamsFree(cparams, nparams)

	for name, value := range infomap {
		if !*value.set {
			continue
		}

		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
		if value.sl != nil {
			/* We're not actually using virTypedParamsAddStringList, as it is
			 * easier to avoid creating a 'char **' in Go to hold all the strings.
			 * We none the less do a version check, because earlier libvirts
			 * would not expect to see multiple string values in a typed params
			 * list with the same field name
			 */
			if C.LIBVIR_VERSION_NUMBER < 1002017 {
				return nil, 0, makeNotImplementedError("virTypedParamsAddStringList")
			}
			for i := 0; i < len(*value.sl); i++ {
				cvalue := C.CString((*value.sl)[i])
				defer C.free(unsafe.Pointer(cvalue))
				var err C.virError
				ret := C.virTypedParamsAddStringWrapper(&cparams, &nparams, &maxparams, cname, cvalue, &err)
				if ret < 0 {
					return nil, 0, makeError(&err)
				}
			}
		} else {
			var err C.virError
			var ret C.int
			if value.i != nil {
				ret = C.virTypedParamsAddIntWrapper(&cparams, &nparams, &maxparams, cname, C.int(*value.i), &err)
			} else if value.ui != nil {
				ret = C.virTypedParamsAddUIntWrapper(&cparams, &nparams, &maxparams, cname, C.uint(*value.ui), &err)
			} else if value.l != nil {
				ret = C.virTypedParamsAddLLongWrapper(&cparams, &nparams, &maxparams, cname, C.longlong(*value.l), &err)
			} else if value.ul != nil {
				ret = C.virTypedParamsAddULLongWrapper(&cparams, &nparams, &maxparams, cname, C.ulonglong(*value.ul), &err)
			} else if value.b != nil {
				v := 0
				if *value.b {
					v = 1
				}
				ret = C.virTypedParamsAddBooleanWrapper(&cparams, &nparams, &maxparams, cname, C.int(v), &err)
			} else if value.d != nil {
				ret = C.virTypedParamsAddDoubleWrapper(&cparams, &nparams, &maxparams, cname, C.double(*value.i), &err)
			} else if value.s != nil {
				cvalue := C.CString(*value.s)
				defer C.free(unsafe.Pointer(cvalue))
				ret = C.virTypedParamsAddStringWrapper(&cparams, &nparams, &maxparams, cname, cvalue, &err)
			} else {
				return nil, 0, fmt.Errorf("No typed parameter value set for field '%s'", name)
			}
			if ret < 0 {
				return nil, 0, makeError(&err)
			}
		}
	}

	return cparams, nparams, nil
}
