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

/* Instruction codes. Each code is one byte.
MSN: most significant nibble
LSN: less significant nibble
*/
const (
	STOP           = 0x00 // Stop
	LOAD_M_R  byte = 0x10 // Load next byte from memory to register. LSN is the index of the register.
	STORE_R_M byte = 0x20 // Store register in LSB to memory address in next address.
	INC_R     byte = 0x30 // Increment register in LSN.
	DEC_R     byte = 0x40 // Decrement register in LSN.
	ADD_M_R   byte = 0x50 // Add next byte to register in LSN.
	SUB_M_R   byte = 0x60 // Subtract next byte to register in LSN.
	MUL_M_R   byte = 0x70 // Multiply next byte to register in LSN.
	DIV_M_R   byte = 0x80 // Divide next byte to register in LSN.
	AND_M_R   byte = 0x90 // Logical AND between next memory byte and register un LSN. Store result in the same register.
	OR_M_R    byte = 0xA0 // Logical OR between next memory byte and register un LSN. Store result in the same register.
	XOR_M_R   byte = 0xB0 // Logical XOR between next memory byte and register un LSN. Store result in the same register.
	NOT_R     byte = 0xC0 // Logical NOT of register in LSN.
	JUMP      byte = 0xf1 // Unconditional jump
	JUMP_Z    byte = 0xf2 // Jump if Zero is set
	JUMP_N_Z  byte = 0xf3 // Jump if Zero is not set
)

type Instruction struct {
	address uint16
	code    byte
}

func (instruction *Instruction) matches(code byte) bool {
	return instruction.code&0xf0 == code
}

// Returns program counter intecrement
func (instruction *Instruction) execute(machine *Machine) int {
	fmt.Printf("Executing instruction %X\n", instruction.code)
	if instruction.code == STOP {
		machine.Stop()
		return 0
	} else if instruction.code == JUMP {
		return jumpToAbsoluteAddress(machine, instruction)
	} else if instruction.code == JUMP_Z {
		return jumpToAbsoluteAddressIfZero(machine, instruction)
	} else if instruction.code == JUMP_N_Z {
		return jumpToAbsoluteAddressIfNotZero(machine, instruction)
	} else if instruction.matches(LOAD_M_R) {
		return loadMemoryIntoRegister(machine, instruction)
	} else if instruction.matches(STORE_R_M) {
		return storeRegisterIntoMemory(machine, instruction)
	} else if instruction.matches(INC_R) {
		return incrementRegister(machine, instruction)
	} else if instruction.matches(DEC_R) {
		return decrementRegister(machine, instruction)
	} else if instruction.matches(ADD_M_R) {
		return addMemoryToRegister(machine, instruction)
	} else if instruction.matches(SUB_M_R) {
		return subtractMemoryFromRegister(machine, instruction)
	} else if instruction.matches(MUL_M_R) {
		return multiplyMemoryFromRegister(machine, instruction)
	} else if instruction.matches(DIV_M_R) {
		return divideMemoryFromRegister(machine, instruction)
	} else if instruction.matches(AND_M_R) {
		return andBetweenMemoryAndRegister(machine, instruction)
	} else if instruction.matches(OR_M_R) {
		return orBetweenMemoryAndRegister(machine, instruction)
	} else if instruction.matches(XOR_M_R) {
		return xorBetweenMemoryAndRegister(machine, instruction)
	} else if instruction.matches(NOT_R) {
		return notRegister(machine, instruction)
	}
	fmt.Printf("Invalid instruction code: %X. Halting.", instruction.code)
	machine.Stop()
	return 0
}

func jumpToAbsoluteAddress(m *Machine, i *Instruction) int {
	address := m.ReadAddress(m.ProgramCounter + 1)
	m.Jump(address)
	return 0
}

func jumpToAbsoluteAddressIfZero(m *Machine, i *Instruction) int {
	return jumpToAbsoluteAddressIfTrue(m, i, m.Flags.Zero)
}

func jumpToAbsoluteAddressIfNotZero(m *Machine, i *Instruction) int {
	return jumpToAbsoluteAddressIfTrue(m, i, !m.Flags.Zero)
}

func jumpToAbsoluteAddressIfTrue(m *Machine, i *Instruction, test bool) int {
	if test {
		address := m.ReadAddress(m.ProgramCounter + 1)
		m.Jump(address)
		return 0
	}
	return 3
}

func loadMemoryIntoRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	value := m.Memory[m.ProgramCounter+1]
	register.Value = value
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func storeRegisterIntoMemory(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	address := m.ReadAddress(m.ProgramCounter + 1)
	m.Memory[address] = register.Value
	return 3
}

func incrementRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value++
	m.lastOpWasZero(register.Value == 0x00)
	return 1
}

func decrementRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value--
	m.lastOpWasZero(register.Value == 0x00)
	return 1
}

func addMemoryToRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value += m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func subtractMemoryFromRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value -= m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func multiplyMemoryFromRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value *= m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func divideMemoryFromRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value /= m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func andBetweenMemoryAndRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value = register.Value & m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func orBetweenMemoryAndRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value = register.Value | m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func xorBetweenMemoryAndRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value = register.Value ^ m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func notRegister(m *Machine, i *Instruction) int {
	register := &m.Registers[i.extractRegister()]
	register.Value = ^m.Memory[m.ProgramCounter+1]
	m.lastOpWasZero(register.Value == 0x00)
	return 2
}

func (instruction *Instruction) extractRegister() byte {
	return instruction.code & byte(0x0F)
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

func (m *Machine) ReadAddress(address uint16) uint16 {
	lsb := uint16(m.Memory[address])
	msb := uint16(m.Memory[address+1])
	return msb<<8 + lsb
}

func (m *Machine) Jump(address uint16) {
	m.ProgramCounter = address
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
