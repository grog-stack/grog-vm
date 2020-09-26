
# Architecture

A Grog Virtual Machine has these componentes:

* 16 registers.
* A 16 bits memory address space.
* A program counter.
* Flags for signaling several conditions.

All instructions consist of an operation code (opcode). Most instructions expect parameters. Opcodes are 1 byte long, so the Grog has a maximum of 256 opcodes.

## Registers

There are 16 registers available. Each registers is identified by an hexadecimal symbol, 
from `0` to `F`. Each register stores 1 byte.

## Memory

The Grog Virtual Machine has 64Kb of total memory. Thus, memory locations are 16 bits long. Each memory location stores 1 byte. It should be enough...

## Program Counter

The program counter always points to the current instruction. It cannot be modified directly. 
It's modified automatically when each instruction completes. It's always initialized with 
`0x0000`.

## Flags

* `ZERO`: This flag is set when the last operation on a register resulted in a result of `0x00`.

# Execution model

Really simple:

* Read the instruction from memory location indicated by the program counter.
* Execute the instruction.
* Increment the program counter.
* Continue.

The machine stops with the `STOP` instruction or when an error occurs.
