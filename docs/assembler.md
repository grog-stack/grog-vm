# grog-asm
Assembler por Grog Virtual Machine

## How to compile

    make

## How to use

Write your program using any text editor.

    load R1 01
    load R2 02
    add R1 R2
    store R1 @ff00
    stop

Compile it

    ./grog-asm my-program.grog-asm

Run it

    ./grog my-program.grog

## The language

### Value types

Constant byte hexdecimal values are written directly. For example: `0f`, `0f`, `ba`.

Absolute memory addresses are prefixed with `@`. For example: `@0001`.

Offset memory addresses are prefixed with `#`. For example: `#0001`. The base of the 
offset is always the Program Counter:

Pointers to memory addresses are prefixed with `*`. For example: `*a2f1`. When a pointer
is used, the machine searches for the actual memory location in the memory location specified by the pointer.

### Constants

You can define simple constants at the begining of your program, and then use them 
wherever you need:

    constants (
        a = 01
        b = 02
    )

    load R1 a // Load the value of the 'a' constant into R1

Obviously, constants cannot be redefined.

### Load values into registers

Load values into registers with the `load` operation. 
Source values are: constant byte values, values in absolute memory locations, values in
offset memory locations, and values in memory locations pointed to in other memory 
locations.

Examples:

    load R0 fa // Load the value fa into the register 0
    load R0 @00fa // Load the value stored in the memory location 00fa into the register 0
    load R0 #00fa // Load the value stored in the memory location PC+00fa into the register 0
    load R0 *00fa // Load into the register 0 the value stored in the memory location 
                // which is stored in the memory location 00fa

### Store values in memory locations

Store values in memory locations with the `store` operation. Sources can be constant 
values, or values in registers.

Examples:

    store @00fa aa // Store the value aa into the memory location 00fa
    store #00fa R1 // Store the value in register R1 into the memory location PC+00fa
    store *00fa Rf // Store the value in register Rf into the memory location 
                   // which is stored in the memory location 00fa

### Copy values

Copy values between registers and memory locations using the `copy` operation.

    copy R1 R0       // Copy the value in R0 into R1
    copy *00fa @00fa // Copy the value in memory location 00fa into the memory location 
                     // which is stored in the memory location 00fa

## Increment and decrement register

    increment R1
    decrement R1

## Arithmetic operations

All arithmetic operations have the same format:

* `add` __destination__ __source__
* `subtract` __destination__ __source__
* `multiply` __destination__ __source__
* `divide` __destination__ __source__

__destination__ and __source__ are always a register.

## Comparing numbers

* `compare` __register__ __register__
* `compare` __register__ __value__
* `compare` __register__ __absolute address__
* `compare` __register__ __offset address__
* `compare` __register__ __pointer address__

The `compare` operation affects the `equals`, `greater`, and `less` flags.

## Boolean operations

All boolean operations have the same format:

* `and` __destination__ __source__
* `or` __destination__ __source__
* `xor` __destination__ __source__

__destination__ and __source__ are always a register.

## Branching

### Unconditional jumps

* `jump` __destination__: unconditinal jump

   jump #0001 // Jump to address 0001

### Unconditional jumps

All jumps have the same format:

* `je` __destination__: jump if equals
* `jne` __destination__: jump if not equals
* `jg` __destination__: jump if greater 
* `jge` __destination__: jump if greater or equal
* `jl` __destination__: jump if less 
* `jle` __destination__: jump if less or equal

All jumps check the flags `equals`, `greater`, and `less`, in order to evaluate the 
result of a previous operation.

## Input and output

A GrogVM can receive information from, or send information to, the outside world using
the IN and OUT instructions. There are 256 devices available, each of then addressable with
a unique code.

    input R0 D1 // Reads a byte from device 01 into register R1
    output D2 R1 // Writes a byte from register R1 to device D2

##

### Comments

    // Single line comment
    R1 -> R2 // Another comment
    /* yet
    another
    commment */