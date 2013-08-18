package libvirt

import "C"

import (
	"unsafe"
)

type VirNodeInfo struct {
	ptr _Ctype_virNodeInfo
}

func (ni *VirNodeInfo) GetModel() string {
	model := C.GoString((*C.char)(unsafe.Pointer(&ni.ptr.model)))
	return model
}

func (ni *VirNodeInfo) getMemoryKB() uint64 {
	return uint64(ni.ptr.memory)
}

func (ni *VirNodeInfo) getCPUs() uint32 {
	return uint32(ni.ptr.cpus)
}

func (ni *VirNodeInfo) getMhz() uint32 {
	return uint32(ni.ptr.mhz)
}

func (ni *VirNodeInfo) getNodes() uint32 {
	return uint32(ni.ptr.nodes)
}

func (ni *VirNodeInfo) getSockets() uint32 {
	return uint32(ni.ptr.sockets)
}

func (ni *VirNodeInfo) getCores() uint32 {
	return uint32(ni.ptr.cores)
}

func (ni *VirNodeInfo) getThreads() uint32 {
	return uint32(ni.ptr.threads)
}
