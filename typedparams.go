package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include <string.h>
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

func typedParamsUnpackLen(cparams *C.virTypedParameter, nparams int, infomap map[string]typedParamsFieldInfo) error {
	for i := 0; i < nparams; i++ {
		var cparam *C.virTypedParameter
		cparam = (*C.virTypedParameter)(unsafe.Pointer(uintptr(unsafe.Pointer(cparams)) + unsafe.Sizeof(*cparam)*uintptr(i)))
		name := C.GoString((*C.char)(unsafe.Pointer(&cparam.field)))
		info, ok := infomap[name]
		if !ok {
			continue
		}
		switch cparam._type {
		case C.VIR_TYPED_PARAM_INT:
			if info.i == nil {
				return fmt.Errorf("field %s expects an int", name)
			}
			*info.i = int(*(*C.int)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_UINT:
			if info.ui == nil {
				return fmt.Errorf("field %s expects a uint", name)
			}
			*info.ui = uint(*(*C.uint)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_LLONG:
			if info.l == nil {
				return fmt.Errorf("field %s expects an int64", name)
			}
			*info.l = int64(*(*C.longlong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_ULLONG:
			if info.ul == nil {
				return fmt.Errorf("field %s expects a uint64", name)
			}
			*info.ul = uint64(*(*C.ulonglong)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_DOUBLE:
			if info.d == nil {
				return fmt.Errorf("field %s expects a float64", name)
			}
			*info.d = float64(*(*C.double)(unsafe.Pointer(&cparam.value)))
			*info.set = true
		case C.VIR_TYPED_PARAM_BOOLEAN:
			if info.b == nil {
				return fmt.Errorf("field %s expects a bool", name)
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
				return fmt.Errorf("field %s expects a string/string list", name)
			}
		}
	}

	return nil
}

func typedParamsUnpack(cparams []C.virTypedParameter, infomap map[string]typedParamsFieldInfo) error {
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

func typedParamsPackNew(infomap map[string]typedParamsFieldInfo) (*[]C.virTypedParameter, error) {
	nparams := 0
	for _, value := range infomap {
		if value.sl != nil {
			nparams += len(*value.sl)
		} else {
			nparams++
		}
	}

	cparams := make([]C.virTypedParameter, nparams)
	nparams = 0
	for key, value := range infomap {
		cfield := C.CString(key)
		defer C.free(cfield)
		clen := len(key) + 1
		if clen > C.VIR_TYPED_PARAM_FIELD_LENGTH {
			clen = C.VIR_TYPED_PARAM_FIELD_LENGTH
		}
		if value.sl != nil {
			for i := 0; i < len(*value.sl); i++ {
				cparam := &cparams[nparams]
				C.memcpy(unsafe.Pointer(&cparam.field[0]), unsafe.Pointer(cfield), C.size_t(clen))
				nparams++
			}
		} else {
			cparam := &cparams[nparams]
			C.memcpy(unsafe.Pointer(&cparam.field[0]), unsafe.Pointer(cfield), C.size_t(clen))
			nparams++
		}
	}

	err := typedParamsPack(cparams, infomap)
	if err != nil {
		C.virTypedParamsClear((*C.virTypedParameter)(unsafe.Pointer(&cparams[0])), C.int(nparams))
		return nil, err
	}
	return &cparams, nil
}
