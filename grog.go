package main

import "fmt"
import "io/ioutil"
import "os"

type Machine struct {
	Name           string
	Registers      [16]Register
	Memory         []byte
	ProgramCounter int32
}

func (m *Machine) Explain() {
	fmt.Printf("Machine: %s\n", m.Name)
	fmt.Printf("Memory size: %d bytes\n", len(m.Memory))
	fmt.Printf("Program counter: %d\n", m.ProgramCounter)
}

func (m *Machine) Run() {
	fmt.Println("Running...")
	fmt.Println("Finished!")
}

func (m *Machine) Load(memory []byte) {

}

type Register struct {
	Name  string
	Value int16
}

func main() {
	machine := Machine{Name: "Grog", Memory: make([]byte, 32*1024)}
	memory := readOrPanic(os.Args[0])
	machine.Explain()
	machine.Load(memory)
	machine.Run()
}

func readOrPanic(filename string) (data []byte) {
	data, err := ioutil.ReadFile(os.Args[0])
	check(err)
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
