
# Architecture

A Grog Virtual Machine has these componentes:

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

LMR means "Load next byte in Memory into Register ?". There're a total of 16 opcodes in this family. Each opcode codifies the target register using the opcode's LSN. 

Examples: 

* `0x10` means _Load next byte in memory into register 0_. 
* `0x11` means _Load next byte in memory into register 1_. 
* ...
* `0x1F` means _Load next byte in memory into register F_. 

## SRM (0x2?)

__S__tore __R__egister __?__ in __M__emory location address in the next two bytes.
There's a total of 16 opcodes in  this family. Each opcode codifies the target register using the opcode's LSN. 

Examples: 

* `0x20 0x0001`: _Store register 0 in memory location `0x0001`_. 
* `0x21 0x0001`: _Store register 1 in memory location `0x0001`_.  
* ...
* `0x2F 0x0001`: _Store register F in memory location `0x0001`_.  

## INC (0x3?)

__INC__crement register __?__. There's a total of 16 opcodes in this family. Each
 opcode codifies the target register using the opcode's LSN. 

Examples:

* `0x30`: _Increment register 0_. 
* `0x31`: _Increment register 1_. 
* ...
* `0x3F`: _Increment register F_. 

## DEC (0x4?)

__DEC__crement register __?__. There's a total of 16 opcodes in this family. Each
 opcode codifies the target register using the opcode's LSN. 

Examples:

* `0x40`: _Decrement register 0_. 
* `0x41`: _Decrement register 1_. 
* ...
* `0x4F`: _Decrement register F_. 

## ADD (0xA?)

__ADD__ value in next memory location to register __?__. There's a total of 16
 opcodes in this family. Each opcode codifies the target register using the opcode's LSN. 

Examples:

* `0xA0 0x05`: _Add `0x05` to register 0_. 
* `0xA1 0x05`: _Add `0x05` to register 1_. 
* ...
* `0xAF 0x05`: _Add `0x05` to register F_. 

## SUB (0xB?)

__SUB__tract value in next memory location to register __?__. There's a total of 16
 opcodes in this family. Each opcode codifies the target register using the opcode's LSN. 

Examples:

* `0xB0 0x05`: _Subtract `0x05` to register 0_. 
* `0xB1 0x05`: _Subtract `0x05` to register 1_. 
* ...
* `0xBF 0x05`: _Subtract `0x05` to register F_. 
