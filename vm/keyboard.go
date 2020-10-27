package vm

import (
	"bytes"
	"os"
)

// Keyboard represents a keyboard device.
type Keyboard struct {
	reader *bytes.Reader
}

func (keyboard Keyboard) Read() byte {
	value, error := keyboard.reader.ReadByte()
	for error != nil {
		value, error = keyboard.reader.ReadByte()
	}
	return value
}

func (keyboard Keyboard) Write(value byte) {

}

// NewKeyboard Creates a new Keyboard device.
func NewKeyboard() Keyboard {
	return Keyboard{reader: bytes.NewReader(os.Stdin)}
}
