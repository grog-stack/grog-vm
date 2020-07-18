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
	ProgramCounter uint32
	running        bool
}

/* Instruction codes. Each code is one byte.
MSB: most significant byte
LSB: less significant byte
*/
const (
	// Stop
	STOP = 0x00
	// Load next byte from memory to register in LSB. LSB is the index of the register.
	LMR byte = 0x10
	// Store register in LSB to memory address in next address
	SRM byte = 0x20
	// Add next byte to register in LSB
	ADD byte = 0xa0
	// Subtract next byte to register in LSB
	SUB byte = 0xb0
)

var registerNames = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

type Instruction struct {
	address uint32
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
	} else if instruction.matches(LMR) {
		registerIndex := instruction.extractRegister()
		machine.Registers[registerIndex].Value = machine.Memory[machine.ProgramCounter+1]
		return 2
	} else if instruction.matches(SRM) {
		registerIndex := instruction.extractRegister()
		address := machine.ReadAddress(machine.ProgramCounter + 1)
		machine.Memory[address] = machine.Registers[registerIndex].Value
		return 3
	} else if instruction.matches(ADD) {
		registerIndex := instruction.extractRegister()
		machine.Registers[registerIndex].Value += machine.Memory[machine.ProgramCounter+1]
		return 2
	} else if instruction.matches(SUB) {
		registerIndex := instruction.extractRegister()
		machine.Registers[registerIndex].Value -= machine.Memory[machine.ProgramCounter+1]
		return 2
	}
	fmt.Printf("Invalid instruction code: %X. Halting.", instruction.code)
	machine.Stop()
	return 0
}

func (instruction *Instruction) extractRegister() byte {
	return instruction.code & 0x0F
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
	m.ProgramCounter += uint32(instruction.execute(m))
}

func (m *Machine) load(memory []byte) {
	for address, value := range memory {
		m.Memory[address] = value
	}
}

func (m *Machine) memorySize() uint32 {
	return uint32(len(m.Memory))
}

func (m *Machine) ReadAddress(address uint32) uint32 {
	lsb := uint32(m.Memory[address])
	msb := uint32(m.Memory[address+1])
	return msb<<8 + lsb
}

func (m *Machine) Stop() {
	m.running = false
}

func (m *Machine) DumpStatus() {
	fmt.Println("Status:")
	m.DumpRegisters()
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

func newMachine(name string, memorySize int) Machine {
	return Machine{
		Name:      "Grog",
		Registers: registers(),
		Memory:    make([]byte, memorySize),
	}
}

func registers() [16]Register {
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
