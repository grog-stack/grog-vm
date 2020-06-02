package main

import "fmt"
import "loader"

type Machine struct {
	Name      string
	Registers [16]Register
	Memory    []int8
}

func (m *Machine) Explain() {
	fmt.Printf("Machine: %s\n", m.Name)
	fmt.Printf("Memory size: %d bytes\n", len(m.Memory))
}

type Register struct {
	Name  string
	Value int16
}

func main() {
	machine := Machine{Name: "Grog", Memory: make([]int8, 32*1024)}
	machine.Explain()
}
