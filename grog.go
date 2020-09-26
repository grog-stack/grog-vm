/*
Glosary:
- MSN: most significant nibble. In 0xf0, MSN is f.
- LSN: less significant nibble. In 0x0f, LSN is f.
- MSB: most significant byte. In 0xff00, MSB is ff.
- MSB: less significant byte. In 0x00ff, LSN is ff.
- Word: Two bytes, little endian. In 0xffaa, MSB is 0xff and LSB is 0xaa.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Machine struct {
	Name           string
	Registers      [16]Register
	Memory         []byte
	ProgramCounter uint16
	Flags          Flags
	running        bool
}

type Flags struct {
	Zero bool
}

/* Instruction codes. Each code is one byte. Instructions from 0x01 to 0x7F are for transferring
information to and from registers, memory, and devices. Instructions fro 0x80 to 0xFF are for
arithmetic, jumps, function calls.
*/

/* Miscellaneous. */

const (
	STOP = 0x00 // Halts the machine. No further instructions will be executed.
)

/*
Transfer instructions have a source and a destination. In the following table, each row is a source, each
column is a destination.

When the source and target of the operation are the same type (for instance, register and register),
we call it COPY. When the target is a register, we call it LOAD. When the target is a memory location,
 we call it STORE.

The following table shows all allowed combinations. (Rows are sources, columns are destinations.)

         | byte | register | address | pointer |
byte     |    - | LOAD     | STORE   | STORE   |
register |    - | COPY     | STORE   | STORE   |
address  |    - | LOAD     | COPY    | COPY    |
pointer  |    - | LOAD     | COPY    | COPY    |
*/

const (
	LOAD_BYTE              = 0x01
	LOAD_ADDRESS           = 0x02
	LOAD_OFFSET            = 0x03
	LOAD_POINTER           = 0x04
	STORE_BYTE_ADDRESS     = 0x05
	STORE_BYTE_OFFSET      = 0x06
	STORE_BYTE_POINTER     = 0x07
	STORE_REGISTER_ADDRESS = 0x08
	STORE_REGISTER_OFFSET  = 0x09
	STORE_REGISTER_POINTER = 0x0A
	COPY_REGISTER          = 0x0B
	COPY_ADDRESS_ADDRESS   = 0x0C
	COPY_ADDRESS_OFFSET    = 0x0D
	COPY_ADDRESS_POINTER   = 0x0E
	COPY_OFFSET_ADDRESS    = 0x0F
	COPY_OFFSET_OFFSET     = 0x10
	COPY_OFFSET_POINTER    = 0x11
	COPY_POINTER_ADRESS    = 0x12
	COPY_POINTER_OFFSET    = 0x13
	COPY_POINTER_POINTER   = 0x14
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
	ADD      byte = 0x83
	SUBTRACT byte = 0x84
	MULTIPLY byte = 0x85
	DIVIDE   byte = 0x86
)

/* Boolean operations. All operations store the result in the right operand.*/

const (
	AND byte = 0x87
	OR  byte = 0x88
	XOR byte = 0x89
	NOT byte = 0x8A
)

/* Branching instructions */
const (
	JUMP_ADDRESS                 byte = 0x90
	JUMP_OFFSET                  byte = 0x91
	JUMP_POINTER                 byte = 0x92
	JUMP_POINTER_OFFSET          byte = 0x93
	JUMP_ZERO_ADDRESS            byte = 0x94
	JUMP_ZERO_OFFSET             byte = 0x95
	JUMP_ZERO_POINTER            byte = 0x96
	JUMP_ZERO_POINTER_OFFSET     byte = 0x97
	JUMP_NOT_ZERO_ADDRESS        byte = 0x98
	JUMP_NOT_ZERO_OFFSET         byte = 0x99
	JUMP_NOT_ZERO_POINTER        byte = 0x9A
	JUMP_NOT_ZERO_POINTER_OFFSET byte = 0x9B
)

type Instruction struct {
	address uint16
	code    byte
}

func (instruction *Instruction) matches(code byte) bool {
	return instruction.code&0x0f == code
}

// Returns program counter intecrement
func (instruction *Instruction) execute(machine *Machine) int {
	fmt.Printf("Executing instruction %X\n", instruction.code)
	if instruction.code == STOP {
		machine.Stop()
		return 0
	} else if instruction.code == LOAD_BYTE {
		value := machine.ReadByteOffset(1)
		register := machine.ReadByteOffset(2)
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
		machine.writeAddress(machine.ReadAddressOffset(2), machine.ReadNextByte())
		return 4
	} else if instruction.code == STORE_BYTE_OFFSET {
		machine.writeOffset(machine.ReadAddressOffset(2), machine.ReadNextByte())
		return 4
	} else if instruction.code == STORE_BYTE_POINTER {
		pointer := machine.ReadAddressOffset(2)
		address := machine.ReadAddress(pointer)
		machine.writeOffset(address, machine.ReadNextByte())
		return 4
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

func (m *Machine) Run() {
	fmt.Println("Running...")
	m.running = true
	for m.running {
		m.execute(m.currentInstruction())
	}
	fmt.Println("Finished!")
}

func (m *Machine) currentInstruction() Instruction {
	currentByte := m.Memory[m.ProgramCounter]
	return Instruction{m.ProgramCounter, currentByte}
}

func (m *Machine) execute(instruction Instruction) {
	m.ProgramCounter += uint16(instruction.execute(m))
}

func (m *Machine) load(memory []byte) {
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

func (m *Machine) writeOffset(offset uint16, value byte) {
	m.writeAddress(m.ProgramCounter+offset, value)
}

func (m *Machine) writeAddress(address uint16, value byte) {
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

func newMachine(name string, memorySize int) Machine {
	return Machine{
		Name:      "Grog",
		Registers: registers(),
		Memory:    make([]byte, memorySize),
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

type Register struct {
	Name  string
	Value byte
}

func main() {
	memory := readOrPanic(os.Args[1])
	machine := newMachine("Grog", len(memory))
	machine.Explain()
	machine.load(memory)
	machine.Run()
	machine.DumpStatus()
}

func readOrPanic(filename string) (data []byte) {
	fmt.Printf("Reading memory from file %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	check(err)
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
