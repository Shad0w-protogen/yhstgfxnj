package main

import (
	"fmt"
	"unsafe"
)

const CPU_TIME_SLICE = 5 // milliseconds per slice in round-robin scheduling

type Task struct {
	ID    int
	State int // 0=ready, 1=running, 2=suspended, 3=zombie
}

var tasks []Task

func NewTask(id int) *Task {
	task := &Task{ID: id, State: 0}
	tasks = append(tasks, *task)
	return task
}

func RunTask(t *Task) {
	t.State = 1 // mark as running
	// Simulate work – in real kernel this would invoke Go code that runs on the CPU slice.
	fmt.Printf("Running task %d\n", t.ID)
}

func Yield() {
	for i := range tasks {
		if tasks[i].State == 1 {
			tasks[i].State = 0 // switch back to ready
			break
		}
	}
}

// main entry point – called from kernel.asm via the stub.
func pizdec() {
	fmt.Printf("memory size: %d\n", GetMemorySize())
	fmt.Printf("allocate 1MiB: %d\n", AllocateMemory(MiB))
	alloc := AllocateMemory(KiB * 256)
	data := make([]byte, alloc)

	printHex(data[:KiB])
	printBytes(data[:KiB])

	freeAddr := uintptr(unsafe.Pointer(&data[0]))
	FreeMemory(freeAddr, KiB*256)
	fmt.Printf("free %d bytes at %x\n", KiB*256, freeAddr)

	fmt.Println("\n--- process scheduling stub ---")
	t1 := NewTask(1)
	t2 := NewTask(2)
	t3 := NewTask(3)

	RunTask(t1)
	Yield()
	RunTask(t2)
	Yield()
	RunTask(t3)
}
