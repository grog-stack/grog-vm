package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinstraus/grog/vm"
)

func main() {
	memory := readOrPanic(os.Args[1])
	machine := vm.NewMachine("Grog", len(memory))
	machine.Explain()
	machine.Load(memory)
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
