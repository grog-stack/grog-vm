
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

The Grog Virtual Machine has 64Kb of total memory. Thus, memory locations are 16 bits long. Each memory location stores 1 byte.

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

# Instruction set

First, some terminology.

* Nibble: a set of four bits.
* LSN: Least Significant Nibble.
* MSN: Most Significant Nibble.
* LSB: Least Significant Byte.
* MSB: Most Significant Byte.

All memory addresses and values are [big-endian](https://en.wikipedia.org/wiki/Endianness).

## STOP (0x00)

Stops the machine and exits.

## LOAD_M_R (0x1?)

Load next byte in Memory into Register specified by the LSN.

Examples: 

* `0x10` means Load next byte in memory into register 0. 
* `0x11` means Load next byte in memory into register 1. 
* ...
* `0x1F` means Load next byte in memory into register F. 

## STORE_M_R (0x2?)

Store register specified by the LSN into memory address specified next.

Examples: 

* `0x20 0x0001`: Store register 0 in memory location `0x0001`. 
* `0x21 0x0001`: Store register 1 in memory location `0x0001`.  
* ...
* `0x2F 0x0001`: Store register F in memory location `0x0001`.  

## INC_R (0x3?)

Increment register specified by the LSN.

Examples:

* `0x30`: Increment register 0. 
* `0x31`: Increment register 1. 
* ...
* `0x3F`: Increment register F. 

## DEC_R (0x4?)

Decrement the register specified by the LSN.

Examples:

* `0x40`: Decrement register 0. 
* `0x41`: Decrement register 1. 
* ...
* `0x4F`: Decrement register F. 

## ADD_M_R (0x5?)

Add the byte in thennext memory location to the register specified by the LSN.

Examples:

* `0x50 0x05`: Add `0x05` to register 0. 
* `0x51 0x05`: Add `0x05` to register 1. 
* ...
* `0x5F 0x05`: Add `0x05` to register F. 

## SUB_M_R (0x6?)

Subtract the value in the next memory location to the register specified by the LSN.

Examples:

* `0x60 0x05`: Subtract `0x05` to register 0. 
* `0x61 0x05`: Subtract `0x05` to register 1. 
* ...
* `0x6F 0x05`: Subtract `0x05` to register F. 

# JUMP (`0xF1`)

Unconditionally jump to the absolute address specified in the next two bytes.

Example:

* `0xF1 0x0F 0x00`: Jumps to the address `0x000F`.

# JUMP_Z (`0xF2`)

If the `Zero` flag is set, jump to the absolute address specified in the next two bytes.

Example:

* `0xF2 0x0F 0x00`: Jumps to the address `0x000F` if `Zero` is set.


# JUMP_N_Z (`0xF3`)

If the `Zero` flag is _not_ set, jump to the absolute address specified in the next two bytes.

Example:

* `0xF3 0x0F 0x00`: Jumps to the address `0x000F` if `Zero` is not set.