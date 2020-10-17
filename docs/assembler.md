# grog-asm
Assembler por Grog Virtual Machine

## How to compile

    make

## How to use

Write your program using any text editor.

    R1 <- ff
    R2 <- @ff00
    STOP

Compile it

    ./grog-asm my-program.grog-asm

Run it

    ./grog my-program.grog

## The language

### Value types

Constant byte hexdecimal values are written directly:

    00
    0f
    ba

Absolute memory addresses are prefixed with `@`:

    @0001
    @0f00

Offset memory addresses are prefixed with `#`. The base of the offset is always the Program Counter:

    #0001

Pointers to memory addresses are prefixed with `*`:

    *a2f1

### Load values into registers

Load values into registers with the `<-` operation. Source values are: constant byte values,
values in absolute memory locations, values in offset memory locations, and values in memory
locations pointed to in other memory locations.

Examples:

    R0 <- fa // Load the value fa into the register 0
    R0 <- @00fa // Load the value stored in the memory location 00fa into the register 0
    R0 <- #00fa // Load the value stored in the memory location PC+00fa into the register 0
    R0 <- *00fa // Load into the register 0 the value stored in the memory location 
                // which is stored in the memory location 00fa

### Store values in memory locations

Store values in memory locations with the `->` operation. Sources can be constant values,
or values in registers.

Examples:

    fa -> @00fa // Store the value fa into the memory location 00fa
    R1 -> #00fa // Store the value in register R1 into the memory location PC+00fa
    Rf -> *00fa // Store the value in register Rf into the memory location 
                // which is stored in the memory location 00fa

### Copy values

Copy values between registers and memory locations using the `->` operation.

    R0 -> R1 // Copy the value in R0 into R1
    @00fa -> *00fa // Copy the value in memory location 00fa in the memory location 
                    // which is stored in the memory location 00fa

## Increment and decrement register

    R1 <- ++ // Increment
    R1 <- -- // Decrement

## Arithmetic operations

    R1 <- R2 + R3 
    R1 <- R1 - R2
    R1 <- R4 * R2
    R1 <- R4 / R2

## Boolean operations

    R1 <- R2 AND R3 
    R1 <- R1 OR R2
    R1 <- R4 XOR R2
    R1 <- NOT R4

## Branching

    JUMP @0af1 // Unconditional jump to address 0af1
    JUMP #0af1 // Unconditional jump to offset 0af1
    JUMP *0af1 // Unconditional jump to address located at 0af1
    IF R1 = R2 JUMP @0af1 // Conditional jump to absolute address. 
    IF R1 > R2 JUMP #0af1 // Conditional jump to relative address address.
    IF R1 != R2 JUMP *0af1 // Conditional jump to address located at 0af1

Conditional operators are `=`, `>`,`>=`, `<`,`<=`, and`!=`.

### Comments

    // Single line comment
    R1 -> R2 // Another comment
    /* yet
    another
    commment */