package vm

import "fmt"

// First-input first-output device
type Fifo struct {
	buffer []byte
	read   int
	write  int
}

func (fifo Fifo) Read() byte {
	value := fifo.buffer[fifo.read]
	fmt.Printf("Read %X from %d\n", value, fifo.read)
	if fifo.read != fifo.write {
		if fifo.read < fifo.maxIndex() {
			fifo.read++
		} else {
			fifo.read = 0
		}
	}
	return value
}

func (fifo Fifo) Write(value byte) {
	if fifo.write == fifo.read {
		fmt.Printf("Wrote %d into %X\n", value, fifo.write)
		fifo.buffer[fifo.write] = value
		if fifo.write == fifo.maxIndex() {
			fifo.write = 0
		} else {
			fifo.write++
		}
	}
}

func (fifo Fifo) maxIndex() int {
	return len(fifo.buffer) - 1
}

func NewFifo(size int) *Fifo {
	return &Fifo{buffer: make([]byte, size), read: 0, write: 0}
}
