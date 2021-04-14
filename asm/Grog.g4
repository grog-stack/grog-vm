grammar Grog;

// Rules
program 
    : instruction+ EOF 
    ; 

instruction
    : increment | decrement
    | arithmeticOperation
    | unaryBooleanOperation | binaryBooleanOperation
    | load
    | store
    | copyValue
    | jump
    | input | output
    | stop
    | wait;

load
    : LOAD Register=REGISTER (
        Value=HEXA_BYTE | 
        Address=ABSOLUTE_ADDRESS | 
        Offset=OFFSET_ADDRESS | 
        Pointer=POINTER_ADDRESS
    ) 
    ;

store
    : (Register=REGISTER | Value=HEXA_BYTE) 
      '->' 
      (Address=ABSOLUTE_ADDRESS | Offset=OFFSET_ADDRESS | Pointer=POINTER_ADDRESS) 
    ;

copyValue
    : copyRegister
    | copyAbsoluteToAbsolute | copyOffsetToAbsolute | copyPointerToAbsolute
    | copyAbsoluteToOffset | copyOffsetToOffset | copyPointerToOffset
    | copyAbsoluteToPointer | copyOffsetToPointer | copyPointerToPointer
    ;

copyRegister
    : COPY DestinationRegister=REGISTER SourceRegister=REGISTER
    ;

copyAbsoluteToAbsolute
    : COPY Destination=ABSOLUTE_ADDRESS Source=ABSOLUTE_ADDRESS
    ;

copyOffsetToAbsolute
    : COPY Destination=ABSOLUTE_ADDRESS Source=OFFSET_ADDRESS
    ;

copyPointerToAbsolute
    : COPY Destination=ABSOLUTE_ADDRESS Source=POINTER_ADDRESS
    ;

copyAbsoluteToOffset
    : COPY Destination=OFFSET_ADDRESS Source=ABSOLUTE_ADDRESS
    ;

copyOffsetToOffset
    : COPY Destination=OFFSET_ADDRESS Source=OFFSET_ADDRESS
    ;

copyPointerToOffset
    : COPY Destination=OFFSET_ADDRESS Source=POINTER_ADDRESS
    ;

copyAbsoluteToPointer
    : COPY Destination=POINTER_ADDRESS Source=ABSOLUTE_ADDRESS
    ;

copyOffsetToPointer
    : COPY Destination=POINTER_ADDRESS Source=OFFSET_ADDRESS
    ;

copyPointerToPointer
    : COPY Destination=POINTER_ADDRESS Source=POINTER_ADDRESS
    ;


increment
    : INCREMENT Register=REGISTER;

decrement
    : DECREMENT Register=REGISTER;

arithmeticOperation
    : (Operator=ADD| Operator=SUBTRACT| Operator=MULTIPLY | Operator=DIVIDE)
      Destination=REGISTER Source=REGISTER
    ;

unaryBooleanOperation
    : NOT Destination=REGISTER
    ;

binaryBooleanOperation
    : (Operator=AND| Operator=OR| Operator=XOR) Destination=REGISTER Source=REGISTER 
    ;

jump
    :(Operator=JUMP | Operator=JUMP_IF_EQUAL | Operator=JUMP_IF_NOT_EQUAL
     | Operator=JUMP_IF_GREATER | Operator=JUMP_IF_NOT_GREATER
     | Operator=JUMP_IF_GREATER_OR_EQUAL | Operator=JUMP_IF_NOT_GREATER_OR_EQUAL
     | Operator=JUMP_IF_LESS | Operator=JUMP_IF_NOT_LESS
     | Operator=JUMP_IF_LESS_OR_EQUAL | Operator=JUMP_IF_NOT_LESS_OR_EQUAL)
      (Address=ABSOLUTE_ADDRESS | Offset=OFFSET_ADDRESS | Pointer=POINTER_ADDRESS)
    ;

input: Destination=REGISTER '<-' Source=DEVICE;

output: Source=REGISTER '->' Destination=DEVICE;

stop
    : STOP
    ;

wait
    : WAIT
    ;

WHITESPACE: [ \r\n\t]+ -> skip;

// Tokens
WS:  [ \t\r\n\u000C]+ -> skip;
COMMENT:   '/*' .*? '*/' -> skip;
LINE_COMMENT :   '//' ~[\r\n]* -> skip;
LOAD: 'load';
MOVE: 'move';
COPY: 'copy';
INCREMENT: 'increment';
DECREMENT: 'decrement';
ADD: 'add';
SUBTRACT: 'subtract';
DIVIDE: 'divide';
MULTIPLY: 'multiply';
JUMP: 'jump';
JUMP_IF_EQUAL: 'je';
JUMP_IF_NOT_EQUAL: 'jne';
JUMP_IF_GREATER: 'jg';
JUMP_IF_NOT_GREATER: 'jng';
JUMP_IF_GREATER_OR_EQUAL: 'jge';
JUMP_IF_NOT_GREATER_OR_EQUAL: 'jnge';
JUMP_IF_LESS: 'jl';
JUMP_IF_NOT_LESS: 'jnl';
JUMP_IF_LESS_OR_EQUAL: 'jle';
JUMP_IF_NOT_LESS_OR_EQUAL: 'jnle';
EQUAL: '=';
GREATER: '>';
GREATER_OR_EQUAL: '>=';
LESS: '<';
LESS_OR_EQUAL: '<=';
NOT_EQUAL: '!=';
NOT: 'not';
AND: 'and';
OR: 'or';
XOR: 'xor';
STOP: 'STOP';
WAIT: 'WAIT';
IF: 'IF';
HEX_DIGIT: [0-9a-fA-F];
HEXA_BYTE: HEX_DIGIT HEX_DIGIT; 
WORD: HEXA_BYTE HEXA_BYTE; 
REGISTER: 'R' HEX_DIGIT;
DEVICE: 'D' HEXA_BYTE;
ABSOLUTE_ADDRESS: '@'WORD;
OFFSET_ADDRESS: '#'WORD;
POINTER_ADDRESS: '*'WORD;

