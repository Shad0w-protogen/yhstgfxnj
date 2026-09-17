package main

import (
	"fmt"
	"unsafe"
)

// BS represents a byte size unit; using uint32 for proper memory sizing.
type BS uint32

const (
	_      = iota
	KiB BS = 1 << 10
	MiB BS = 1 << 20
	GiB BS = 1 << 30
)

// Basic memory manager functions.
func GetMemorySize() BS {
	return MiB * 64 // start with a modest heap size
}

func AllocateMemory(size BS) BS {
	// Simple allocation logic placeholder – actual implementation would involve mapping kernel memory.
	if size < KiB {
		panic("allocate: minimum size is 1KiB")
	}
	return size
}

func FreeMemory(addr uintptr, size BS) {
	// Placeholder for freeing; in real code this would unmap or return pages to the system.
}

func printHex(b []byte) {
	for i := range b {
		fmt.Printf("%02x", b[i])
	}
	fmt.Println()
}

func printBytes(b []byte) {
	for _, c := range b {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// main entry point – called from kernel.asm via the stub.
func go_main_stub() {
	memSize := GetMemorySize()
	fmt.Printf("memory size: %d\n", memSize)
	fmt.Printf("allocate 1MiB: %d\n", AllocateMemory(MiB))
	alloc := AllocateMemory(KiB * 256)
	data := make([]byte, alloc)

	printHex(data[:KiB])
	printBytes(data[:KiB])

	freeAddr := uintptr(unsafe.Pointer(&data[0]))
	FreeMemory(freeAddr, KiB*256)
	fmt.Printf("free %d bytes at %x\n", KiB*256, freeAddr)
}

func main() {
	go go_main_stub()
	go pizdec()
}
