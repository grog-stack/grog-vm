package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinstraus/grog/vm"
)

type Parameters struct {
	file    string
	display bool
}

func main() {
	parameters, err := parameters(os.Args)
	if err != nil {
		panic(err)
	}
	memory := readOrPanic(parameters.file)
	machine := vm.NewMachine("Grog", len(memory), parameters.display)
	machine.Explain()
	machine.Load(memory)
	machine.Run()
	machine.DumpStatus()
}

func parameters(args []string) (Parameters, error) {
	parameters := Parameters{display: false}
	length := len(args)
	if length == 0 {
		return parameters, errors.New("file is mandatory")
	}
	parameters.file = args[length-1]
	if length > 0 {
		for _, value := range args {
			switch value {
			case "--enable_display":
				parameters.display = true
			}
		}
	}
	return parameters, nil
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
