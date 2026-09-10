package main

type BS float32

const (
	_      = iota
	KiB BS = 1 << (10 * iota)
	MiB
)

func main() {

}
