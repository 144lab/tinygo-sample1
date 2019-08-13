package rtt

/*
#include "sdk_config.h"
#include "SEGGER_RTT.h"
*/
import "C"
import (
	"log"
	"sync"
	"unsafe"
)

func init() {
	C.SEGGER_RTT_Init()
}

// Writer ...
type writer struct {
	sync.Mutex
}

// New ...
func New() *log.Logger {
	return log.New(&writer{}, "", 0)
}

// Write ...
func (w *writer) Write(b []byte) (int, error) {
	w.Lock()
	n := C.SEGGER_RTT_WriteNoLock(0, unsafe.Pointer(&b[0]), C.uint(len(b)))
	w.Unlock()
	return int(n), nil
}

//go:export memcpy
func libc_memcpy(dst, src unsafe.Pointer, size uintptr) {
	for i := uintptr(0); i < size; i++ {
		*(*byte)(unsafe.Pointer(uintptr(dst) + i)) = *(*byte)(unsafe.Pointer(uintptr(src) + i))
	}
}
