
# Architecture

A Grog Virtual Machine has theses componentes:

* 16 registers, identified with a hexadecimal symbol, from `0` to `F`. Each register stores 1 byte.
* A 32 bits memory address space. Each memory location stores 1 byte.
* A 32 bits program counter. It points to the next instruction. It's initialialized 
with `0`.

All instructions consist of an operation code (opcode), and optional parameters. Opcodes
are 1 byte long, so the Grog has a maximum of 256 opcodes.

# Execution model

Really simple:

* Read the instruction from memory location indicated by the program counter.
* Execute the instruction.
* Increment the program counter.
* Continue.

The machine stops with the `STOP` instruction or when an error occurs.

# Instruction set

First, some terminology.

* Nibble: a set of four bits.
* LSN: Least Significant Nibble.
* MSN: Most Significant Nibble.
* LSB: Least Significant Byte.
* MSB: Most Significant Byte.

## STOP (0x00)

Stops the machine and exits.

## LMR (0x1?)

LMR means "Load next byte in Memory into Register ?". There're a total of 16 opcodes in this family. Each opcode codifies the target register using the opcode's LSN. Examples: 

* `0x10` means _Load next byte in memory into register 0_. 
* `0x11` means _Load next byte in memory into register 1_. 
* ...
* `0x1F` means _Load next byte in memory into register F_. 
