/*
Glosary:
- MSN: most significant nibble. In 0xf0, MSN is f.
- LSN: less significant nibble. In 0x0f, LSN is f.
- MSB: most significant byte. In 0xff00, MSB is ff.
- MSB: less significant byte. In 0x00ff, LSN is ff.
- Word: Two bytes, little endian. In 0xffaa, MSB is 0xff and LSB is 0xaa.
*/

package vm

import (
	"fmt"
	"sync"
)

type Machine struct {
	Name           string
	Registers      [16]Register
	Memory         []byte
	ProgramCounter uint16
	Flags          Flags
	running        bool
	Devices        [255]Device
	mutex          *sync.Mutex
}

type Flags struct {
	Zero bool
}

type Device interface {
	Read() byte
	Write(value byte)
}

/* Instruction codes. Each code is one byte. Instructions from 0x01 to 0x7F are for transferring
information to and from registers, memory, and devices. Instructions fro 0x80 to 0xFF are for
arithmetic, jumps, function calls.
*/

/* Miscellaneous. */

const (
	STOP = 0x00 // Halts the machine. No further instructions will be executed.
	WAIT = 0x01 // Wait for an interruption. Effectively pauses the machine.
)

/*
Transfer instructions have a source and a destination. In the following table, each row is a source, each
column is a destination.

When the source and target of the operation are the same type (for instance, register and register),
we call it COPY. When the target is a register, we call it LOAD. When the target is a memory location,
 we call it STORE.

The following table shows all allowed combinations. (Rows are destinations, columns are sources.)

         | byte  | register | memory |
register | LOAD  | COPY     | LOAD   |
memory   | STORE | STORE    | COPY   |

*/

const (
	LOAD_BYTE              = 0x11
	LOAD_ADDRESS           = 0x12
	LOAD_OFFSET            = 0x13
	LOAD_POINTER           = 0x14
	STORE_BYTE_ADDRESS     = 0x15
	STORE_BYTE_OFFSET      = 0x16
	STORE_BYTE_POINTER     = 0x17
	STORE_REGISTER_ADDRESS = 0x18
	STORE_REGISTER_OFFSET  = 0x19
	STORE_REGISTER_POINTER = 0x1A
	COPY_REGISTER          = 0x1B
	COPY_ADDRESS_ADDRESS   = 0x1C
	COPY_ADDRESS_OFFSET    = 0x1D
	COPY_ADDRESS_POINTER   = 0x1E
	COPY_OFFSET_ADDRESS    = 0x1F
	COPY_OFFSET_OFFSET     = 0x20
	COPY_OFFSET_POINTER    = 0x21
	COPY_POINTER_ADDRESS   = 0x22
	COPY_POINTER_OFFSET    = 0x23
	COPY_POINTER_POINTER   = 0x24
)

/* Single operand arithmentic register operations. */

const (
	INCREMENT byte = 0x80 // Increment register
	DECREMENT byte = 0x81 // Decrement register
)

/*
   Two-operand register arithmetic operations. Each operation has a left and right operand.
   All operations store the result in the right operand.
*/

const (
	ADD      byte = 0x82
	SUBTRACT byte = 0x83
	MULTIPLY byte = 0x84
	DIVIDE   byte = 0x85
)

/* Boolean operations. All operations store the result in the right operand.*/

const (
	AND byte = 0x86
	OR  byte = 0x87
	XOR byte = 0x88
	NOT byte = 0x89
)

/* Branching instructions */
const (
	JUMP_ADDRESS               byte = 0x90
	JUMP_OFFSET                byte = 0x91
	JUMP_POINTER               byte = 0x92
	JUMP_EQUAL_ADDRESS         byte = 0x93
	JUMP_EQUAL_OFFSET          byte = 0x94
	JUMP_EQUAL_POINTER         byte = 0x95
	JUMP_NOT_EQUAL_ADDRESS     byte = 0x96
	JUMP_NOT_EQUAL_OFFSET      byte = 0x97
	JUMP_NOT_EQUAL_POINTER     byte = 0x98
	JUMP_GREATER_ADDRESS       byte = 0x99
	JUMP_GREATER_OFFSET        byte = 0x9A
	JUMP_GREATER_POINTER       byte = 0x9B
	JUMP_GREATER_EQUAL_ADDRESS byte = 0x9C
	JUMP_GREATER_EQUAL_OFFSET  byte = 0x9D
	JUMP_GREATER_EQUAL_POINTER byte = 0x9E
	JUMP_LESS_ADDRESS          byte = 0x9F
	JUMP_LESS_OFFSET           byte = 0xA0
	JUMP_LESS_POINTER          byte = 0xA1
	JUMP_LESS_EQUAL_ADDRESS    byte = 0xA2
	JUMP_LESS_EQUAL_OFFSET     byte = 0xA3
	JUMP_LESS_EQUAL_POINTER    byte = 0xA4
)

// Input, output

const (
	INPUT_REGISTER  byte = 0xB0
	OUTPUT_REGISTER byte = 0xC0
)

type Instruction struct {
	address uint16
	code    byte
}

func (instruction *Instruction) matches(code byte) bool {
	return instruction.code&0x0f == code
}

// Returns program counter intecrement
func (instruction *Instruction) execute(machine *Machine) uint16 {
	fmt.Printf("Executing instruction %X\n", instruction.code)
	if instruction.code == STOP {
		machine.Stop()
		return 0
	} else if instruction.code == WAIT {
		fmt.Println("Waiting for interrupt...")
		// The instruction locks the machine, and only an interrupt can unlock it.
		machine.lock()
		return 0
	} else if instruction.code == LOAD_BYTE {
		value := machine.ReadByteOffset(1)
		register := machine.ReadByteOffset(2)
		fmt.Printf("Loading %X into register %X\n", value, register)
		machine.register(register).Value = value
		return 3
	} else if instruction.code == LOAD_ADDRESS {
		value := machine.ReadByteAddress(machine.ReadNextAddress())
		register := machine.ReadByteOffset(3)
		machine.register(register).Value = value
		return 4
	} else if instruction.code == LOAD_POINTER {
		value := machine.ReadBytePointer(machine.ReadNextAddress())
		register := machine.ReadByteOffset(3)
		machine.register(register).Value = value
		return 4
	} else if instruction.code == STORE_BYTE_ADDRESS {
		machine.writeInAddress(machine.ReadAddressOffset(2), machine.ReadNextByte())
		return 4
	} else if instruction.code == STORE_BYTE_OFFSET {
		machine.writeInOffset(machine.ReadAddressOffset(2), machine.ReadNextByte())
		return 4
	} else if instruction.code == STORE_BYTE_POINTER {
		pointer := machine.ReadAddressOffset(2)
		address := machine.ReadAddress(pointer)
		machine.writeInOffset(address, machine.ReadNextByte())
		return 4
	} else if instruction.code == STORE_REGISTER_ADDRESS {
		machine.writeInAddress(
			machine.ReadAddressOffset(2),
			machine.Registers[machine.ReadNextByte()].Value,
		)
		return 4
	} else if instruction.code == STORE_REGISTER_OFFSET {
		machine.writeInOffset(
			machine.ReadAddressOffset(2),
			machine.Registers[machine.ReadNextByte()].Value,
		)
		return 4
	} else if instruction.code == STORE_REGISTER_POINTER {
		pointer := machine.ReadAddressOffset(2)
		address := machine.ReadAddress(pointer)
		machine.writeInAddress(
			address,
			machine.Registers[machine.ReadNextByte()].Value,
		)
		return 4
	} else if instruction.code == COPY_REGISTER {
		destination := &machine.Registers[machine.ReadByteOffset(1)]
		source := &machine.Registers[machine.ReadByteOffset(2)]
		destination.Value = source.Value
		return 3
	} else if instruction.code == COPY_ADDRESS_ADDRESS {
		machine.writeInAddress(
			machine.ReadAddressOffset(3),
			machine.Memory[machine.ReadAddressOffset(1)],
		)
		return 5
	} else if instruction.code == COPY_ADDRESS_OFFSET {
		machine.writeInOffset(
			machine.ReadAddressOffset(3),
			machine.Memory[machine.ReadAddressOffset(1)],
		)
		return 5
	} else if instruction.code == COPY_ADDRESS_POINTER {
		machine.writeInOffset(
			machine.ReadAddress(machine.ReadAddressOffset(3)),
			machine.Memory[machine.ReadAddressOffset(1)],
		)
		return 5
	} else if instruction.code == COPY_OFFSET_ADDRESS {
		machine.writeInAddress(
			machine.ReadAddressOffset(3),
			machine.ReadByteOffset(machine.ReadNextAddress()),
		)
		return 5
	} else if instruction.code == COPY_OFFSET_OFFSET {
		machine.writeInOffset(
			machine.ReadAddressOffset(3),
			machine.ReadByteOffset(machine.ReadNextAddress()),
		)
		return 5
	} else if instruction.code == COPY_OFFSET_POINTER {
		machine.writeInOffset(
			machine.ReadAddress(machine.ReadAddressOffset(3)),
			machine.ReadByteOffset(machine.ReadNextAddress()),
		)
		return 5
	} else if instruction.code == COPY_POINTER_ADDRESS {
		machine.writeInAddress(
			machine.ReadAddressOffset(3),
			machine.ReadByteAddress(machine.ReadAddress(machine.ReadAddressOffset(1))),
		)
		return 5
	} else if instruction.code == COPY_POINTER_OFFSET {
		machine.writeInOffset(
			machine.ReadAddressOffset(3),
			machine.ReadByteAddress(machine.ReadAddress(machine.ReadAddressOffset(1))),
		)
		return 5
	} else if instruction.code == COPY_POINTER_POINTER {
		machine.writeInOffset(
			machine.ReadAddress(machine.ReadAddressOffset(3)),
			machine.ReadByteAddress(machine.ReadAddress(machine.ReadAddressOffset(1))),
		)
		return 5
	} else if instruction.code == INCREMENT {
		(&machine.Registers[machine.ReadByteOffset(1)]).Value++
		return 2
	} else if instruction.code == DECREMENT {
		(&machine.Registers[machine.ReadByteOffset(1)]).Value--
		return 2
	} else if instruction.code == ADD {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value + right.Value
		return 4
	} else if instruction.code == SUBTRACT {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value - right.Value
		return 4
	} else if instruction.code == MULTIPLY {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value * right.Value
		return 4
	} else if instruction.code == DIVIDE {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value / right.Value
		return 4
	} else if instruction.code == AND {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value & right.Value
		return 4
	} else if instruction.code == OR {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value | right.Value
		return 4
	} else if instruction.code == XOR {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		destination := &machine.Registers[machine.ReadByteOffset(3)]
		destination.Value = left.Value ^ right.Value
		return 4
	} else if instruction.code == NOT {
		destination := &machine.Registers[machine.ReadByteOffset(2)]
		destination.Value = ^destination.Value
		return 3
	} else if instruction.code == JUMP_ADDRESS {
		machine.Jump(machine.ReadNextAddress())
		return 0
	} else if instruction.code == JUMP_OFFSET {
		machine.Jump(machine.ReadAddressFromNextOffset())
		return 0
	} else if instruction.code == JUMP_POINTER {
		machine.Jump(machine.ReadAddressFromNextPointer())
		return 0
	} else if instruction.code == JUMP_EQUAL_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value == right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_EQUAL_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value == right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_EQUAL_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value == right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_NOT_EQUAL_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value != right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_NOT_EQUAL_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value != right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_NOT_EQUAL_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value != right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value > right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value > right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value > right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_EQUAL_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value >= right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_EQUAL_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value >= right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_GREATER_EQUAL_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value >= right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value < right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value < right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value < right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_EQUAL_ADDRESS {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value <= right.Value {
			machine.Jump(machine.ReadAddressOffset(3))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_EQUAL_OFFSET {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value <= right.Value {
			machine.Jump(machine.ReadAddressOffset(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == JUMP_LESS_EQUAL_POINTER {
		left := &machine.Registers[machine.ReadByteOffset(1)]
		right := &machine.Registers[machine.ReadByteOffset(2)]
		if left.Value <= right.Value {
			machine.Jump(machine.ReadAddress(machine.ReadAddressOffset(3)))
			return 0
		}
		return 5
	} else if instruction.code == INPUT_REGISTER {
		device := machine.ReadNextByte()
		register := machine.ReadByteOffset(2)
		machine.Registers[register].Value = machine.Devices[device].Read()
		return 3
	} else if instruction.code == OUTPUT_REGISTER {
		register := machine.ReadNextByte()
		device := machine.ReadByteOffset(2)
		value := machine.Registers[register].Value
		fmt.Printf("Sending value %X into device %X\n", value, device)
		machine.Devices[device].Write(value)
		return 3
	}

	fmt.Printf("Invalid instruction code: %X. Halting.", instruction.code)
	machine.Stop()
	return 0
}

func (instruction *Instruction) extractRegister() byte {
	return instruction.code & byte(0xF0)
}

func (m *Machine) Explain() {
	fmt.Printf("Machine: %s\n", m.Name)
	fmt.Printf("Memory size: %d bytes\n", len(m.Memory))
	fmt.Printf("Program counter: %d\n", m.ProgramCounter)
}

/* Before executing an instruction, we try to get a lock on the machine, in order to
simulate hardware interrupts. This works for the time being, but we must dive deeper... */
func (m *Machine) Run() {
	fmt.Println("Running...")
	m.running = true
	for m.running {
		m.lock()
		m.execute(m.currentInstruction())
		m.unlock()
	}
	fmt.Println("Finished!")
}

func (m *Machine) lock() {
	m.mutex.Lock()
}

func (m *Machine) unlock() {
	m.mutex.Unlock()
}

func (m *Machine) currentInstruction() Instruction {
	currentByte := m.Memory[m.ProgramCounter]
	return Instruction{m.ProgramCounter, currentByte}
}

func (m *Machine) execute(instruction Instruction) {
	m.ProgramCounter += instruction.execute(m)
}

func (m *Machine) Load(memory []byte) {
	for address, value := range memory {
		m.Memory[address] = value
	}
}

func (m *Machine) memorySize() uint16 {
	return uint16(len(m.Memory))
}

func (m *Machine) register(register byte) *Register {
	return &m.Registers[register]
}

func (m *Machine) ReadNextAddress() uint16 {
	return m.ReadAddressOffset(1)
}

// Assumes that the next memory location is an offset from the current program counter.
func (m *Machine) ReadAddressFromNextOffset() uint16 {
	return m.ProgramCounter + m.ReadNextAddress()
}

// Assumes that the next memory location contains a pointer to a memory location that holds
// an address
func (m *Machine) ReadAddressFromNextPointer() uint16 {
	return m.ReadAddress(m.ReadNextAddress())
}

func (m *Machine) ReadAddressOffset(offset uint16) uint16 {
	return m.ReadAddress(m.ProgramCounter + offset)
}

func (m *Machine) ReadAddress(address uint16) uint16 {
	lsb := uint16(m.Memory[address])
	msb := uint16(m.Memory[address+1])
	return msb<<8 + lsb
}

func (m *Machine) ReadNextByte() byte {
	return m.ReadByteOffset(1)
}

func (m *Machine) ReadByteOffset(offset uint16) byte {
	return m.Memory[m.ProgramCounter+offset]
}

func (m *Machine) ReadByteAddress(address uint16) byte {
	return m.Memory[address]
}

func (m *Machine) ReadBytePointer(pointer uint16) byte {
	return m.ReadByteAddress(m.ReadAddress(pointer))
}

func (m *Machine) Jump(address uint16) {
	m.ProgramCounter = address
}

func (m *Machine) writeInOffset(offset uint16, value byte) {
	m.writeInAddress(m.ProgramCounter+offset, value)
}

func (m *Machine) writeInAddress(address uint16, value byte) {
	m.Memory[address] = value
}

func (m *Machine) Stop() {
	m.running = false
}

func (m *Machine) lastOpWasZero(wasZero bool) {
	m.Flags.Zero = wasZero
}

func (m *Machine) DumpStatus() {
	fmt.Println("Status:")
	m.DumpRegisters()
	m.DumpFlags()
	m.DumpMemory()
}

func (m *Machine) DumpRegisters() {
	fmt.Print("\tRegisters: ")
	for _, value := range m.Registers {
		fmt.Printf("%s=%X ", value.Name, value.Value)
	}
	fmt.Println()
}

func (m *Machine) DumpMemory() {
	fmt.Print("\tMemory status: ")
	for _, value := range m.Memory {
		fmt.Printf("%X ", value)
	}
	fmt.Println()
}

func (m *Machine) DumpFlags() {
	fmt.Printf("\tFlags: Zero=%t\n", m.Flags.Zero)
}

func NewMachine(name string, memorySize int, display bool) Machine {
	return Machine{
		Name:      "Grog",
		Registers: registers(),
		Memory:    make([]byte, memorySize),
		Devices:   makeDevices(display),
		mutex:     &sync.Mutex{},
	}
}

func registers() [16]Register {
	var registerNames = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	registers := [16]Register{}
	for i := 0; i < 16; i++ {
		registers[i] = newRegister(registerNames[i])
	}
	return registers
}

func newRegister(name string) Register {
	return Register{name, 0x00}
}

func makeDevices(display bool) [255]Device {
	devices := [255]Device{}
	if display {
		devices[0] = NewDisplay(120, 90)
	}
	return devices
}

type Register struct {
	Name  string
	Value byte
}
