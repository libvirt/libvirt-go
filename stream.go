package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"
import (
	"io"
	"unsafe"
)

type StreamFlags int

const (
	STREAM_NONBLOCK = StreamFlags(C.VIR_STREAM_NONBLOCK)
)

type StreamEventType int

const (
	STREAM_EVENT_READABLE = StreamEventType(C.VIR_STREAM_EVENT_READABLE)
	STREAM_EVENT_WRITABLE = StreamEventType(C.VIR_STREAM_EVENT_WRITABLE)
	STREAM_EVENT_ERROR    = StreamEventType(C.VIR_STREAM_EVENT_ERROR)
	STREAM_EVENT_HANGUP   = StreamEventType(C.VIR_STREAM_EVENT_HANGUP)
)

type Stream struct {
	ptr C.virStreamPtr
}

func (v *Stream) Abort() error {
	result := C.virStreamAbort(v.ptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (v *Stream) Close() error {
	result := C.virStreamFinish(v.ptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (v *Stream) Free() error {
	result := C.virStreamFree(v.ptr)
	if result == -1 {
		return GetLastError()
	}
	v.ptr = nil
	return nil
}

func (v *Stream) Read(p []byte) (int, error) {
	n := C.virStreamRecv(v.ptr, (*C.char)(unsafe.Pointer(&p[0])), C.size_t(len(p)))
	if n < 0 {
		return 0, GetLastError()
	}
	if n == 0 {
		return 0, io.EOF
	}

	return int(n), nil
}

func (v *Stream) Write(p []byte) (int, error) {
	n := C.virStreamSend(v.ptr, (*C.char)(unsafe.Pointer(&p[0])), C.size_t(len(p)))
	if n < 0 {
		return 0, GetLastError()
	}
	if n == 0 {
		return 0, io.EOF
	}

	return int(n), nil
}
