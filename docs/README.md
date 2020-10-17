# grog
The virtual machine you didn't ask for.

## Description

__Grog__ is a very simplistic and naive virtual machine. The machine defines an architecture, an execution model, and an instruction set. Programs are stored in a binary file, representing the memory of the machine.

## How to build

    make

## How to run

Create a memory file with any hexadecimal editor, such as [GHex](http://gitlab.gnome.org/browse/ghex/). Run the Grog Virtual Machine with this command:

    ./grog test-machine-1.grog

The machine will run, trace the instructions, and dump its final state:

    ./grog test-machine-1.grog 
    Reading memory from file test-machine-1.grog
    Machine: Grog
    Memory size: 9 bytes
    Program counter: 0
    Running...
    Executing instruction 10
    Executing instruction A0
    Executing instruction 20
    Executing instruction 0
    Finished!
    Status:
        Registers: 0=2 1=0 2=0 3=0 4=0 5=0 6=0 7=0 8=0 9=0 A=0 B=0 C=0 D=0 E=0 F=0 
        Memory status: 10 1 A0 1 20 8 0 0 2 

## Further reading

* [Grog Virtual Machine reference](./reference.md)
* [Grog Assembler](./assembler.md)
