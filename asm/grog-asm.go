package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/martinstraus/grog/asm/parser"
	"github.com/martinstraus/grog/vm"
)

var arithmeticOperators = map[string]byte{
	"+": vm.ADD,
	"-": vm.SUBTRACT,
	"*": vm.MULTIPLY,
	"/": vm.DIVIDE,
}

var booleanOperators = map[string]byte{
	"AND": vm.AND,
	"OR":  vm.OR,
	"XOR": vm.XOR,
}

type listener struct {
	*parser.BaseGrogListener
	Output *bufio.Writer
}

func (l *listener) ExitLoad(c *parser.LoadContext) {
	if c.Value != nil {
		l.Output.WriteByte(vm.LOAD_BYTE)
		l.Output.WriteByte(valueByte(c.Value.GetText()))
	} else if c.Address != nil {
		l.Output.WriteByte(vm.LOAD_ADDRESS)
		l.Output.Write(absoluteAddressBytes(c.Address.GetText()))
	} else if c.Offset != nil {
		l.Output.WriteByte(vm.LOAD_OFFSET)
		l.Output.Write(offsetAddressBytes(c.Offset.GetText()))
	} else if c.Pointer != nil {
		l.Output.WriteByte(vm.LOAD_POINTER)
		l.Output.Write(pointerAddressBytes(c.Pointer.GetText()))
	}
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitStore(c *parser.StoreContext) {
	if c.Value != nil {
		value := valueByte(c.Value.GetText())
		if c.Address != nil {
			l.Output.WriteByte(vm.STORE_BYTE_ADDRESS)
			l.Output.WriteByte(value)
			l.Output.Write(absoluteAddressBytes(c.Address.GetText()))
		} else if c.Offset != nil {
			l.Output.WriteByte(vm.STORE_BYTE_OFFSET)
			l.Output.WriteByte(value)
			l.Output.Write(offsetAddressBytes(c.Offset.GetText()))
		} else if c.Pointer != nil {
			l.Output.WriteByte(vm.STORE_BYTE_POINTER)
			l.Output.WriteByte(value)
			l.Output.Write(pointerAddressBytes(c.Pointer.GetText()))
		}
	} else if c.Register != nil {
		register := registerByte(c.Register.GetText())
		if c.Address != nil {
			l.Output.WriteByte(vm.STORE_REGISTER_ADDRESS)
			l.Output.WriteByte(register)
			l.Output.Write(absoluteAddressBytes(c.Address.GetText()))
		} else if c.Offset != nil {
			l.Output.WriteByte(vm.STORE_REGISTER_OFFSET)
			l.Output.WriteByte(register)
			l.Output.Write(offsetAddressBytes(c.Offset.GetText()))
		} else if c.Pointer != nil {
			l.Output.WriteByte(vm.STORE_REGISTER_POINTER)
			l.Output.WriteByte(register)
			l.Output.Write(pointerAddressBytes(c.Pointer.GetText()))
		}
	}
}

func (l *listener) ExitCopyRegister(c *parser.CopyRegisterContext) {
	l.Output.WriteByte(vm.COPY_REGISTER)
	l.Output.WriteByte(registerByte(c.SourceRegister.GetText()))
	l.Output.WriteByte(registerByte(c.DestinationRegister.GetText()))
}
func (l *listener) ExitCopyAbsoluteToAbsolute(c *parser.CopyAbsoluteToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_ADDRESS_ADDRESS)
	l.Output.Write(addressBytes(c.SourceAddress.GetText()))
	l.Output.Write(addressBytes(c.DestinationAddress.GetText()))
}
func (l *listener) ExitCopyAbsoluteToOffset(c *parser.CopyAbsoluteToOffsetContext) {
	l.Output.WriteByte(vm.COPY_ADDRESS_OFFSET)
	l.Output.Write(addressBytes(c.SourceAddress.GetText()))
	l.Output.Write(offsetAddressBytes(c.DestinationOffset.GetText()))
}
func (l *listener) ExitCopyAbsoluteToPointer(c *parser.CopyAbsoluteToPointerContext) {
	l.Output.WriteByte(vm.COPY_ADDRESS_POINTER)
	l.Output.Write(addressBytes(c.SourceAddress.GetText()))
	l.Output.Write(pointerAddressBytes(c.DestinationPointer.GetText()))
}
func (l *listener) ExitCopyOffsetToAbsolute(c *parser.CopyOffsetToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_ADDRESS)
	l.Output.Write(offsetAddressBytes(c.SourceOffset.GetText()))
	l.Output.Write(addressBytes(c.DestinationAddress.GetText()))
}
func (l *listener) ExitCopyOffsetToOffset(c *parser.CopyOffsetToOffsetContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_OFFSET)
	l.Output.Write(offsetAddressBytes(c.SourceOffset.GetText()))
	l.Output.Write(offsetAddressBytes(c.DestinationOffset.GetText()))
}
func (l *listener) ExitCopyOffsetToPointer(c *parser.CopyOffsetToPointerContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_POINTER)
	l.Output.Write(offsetAddressBytes(c.SourceOffset.GetText()))
	l.Output.Write(pointerAddressBytes(c.DestinationPointer.GetText()))
}
func (l *listener) ExitCopyPointerToAbsolute(c *parser.CopyPointerToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_POINTER_ADDRESS)
	l.Output.Write(pointerAddressBytes(c.SourcePointer.GetText()))
	l.Output.Write(addressBytes(c.DestinationAddress.GetText()))
}
func (l *listener) ExitCopyPointerToOffset(c *parser.CopyPointerToOffsetContext) {
	l.Output.WriteByte(vm.COPY_POINTER_OFFSET)
	l.Output.Write(pointerAddressBytes(c.SourcePointer.GetText()))
	l.Output.Write(offsetAddressBytes(c.DestinationOffset.GetText()))
}
func (l *listener) ExitCopyPointerToPointer(c *parser.CopyPointerToPointerContext) {
	l.Output.WriteByte(vm.COPY_POINTER_POINTER)
	l.Output.Write(pointerAddressBytes(c.SourcePointer.GetText()))
	l.Output.Write(pointerAddressBytes(c.DestinationPointer.GetText()))
}

func (l *listener) ExitIncrement(c *parser.IncrementContext) {
	l.Output.WriteByte(vm.INCREMENT)
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitDecrement(c *parser.DecrementContext) {
	l.Output.WriteByte(vm.DECREMENT)
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitArithmeticOperation(c *parser.ArithmeticOperationContext) {
	l.Output.WriteByte(arithmeticOperators[c.Operator.GetText()])
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
	l.Output.WriteByte(registerByte(c.Source.GetText()))
}

func (l *listener) ExitBinaryBooleanOperation(c *parser.BinaryBooleanOperationContext) {
	l.Output.WriteByte(booleanOperators[c.Operator.GetText()])
	l.Output.WriteByte(registerByte(c.Left.GetText()))
	l.Output.WriteByte(registerByte(c.Right.GetText()))
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
}

func (l *listener) ExitUnaryBooleanOperation(c *parser.UnaryBooleanOperationContext) {
	l.Output.WriteByte(vm.NOT)
	l.Output.WriteByte(registerByte(c.Operand.GetText()))
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
}

func (l *listener) ExitJump(c *parser.JumpContext) {
	if c.Left != nil && c.Right != nil {
		left := registerByte(c.Left.GetText())
		right := registerByte(c.Right.GetText())
		var operation byte = 0
		var destination []byte = nil
		if c.Address != nil {
			destination = absoluteAddressBytes(c.Address.GetText())
		} else if c.Offset != nil {
			destination = offsetAddressBytes(c.Offset.GetText())
		} else if c.Pointer != nil {
			destination = pointerAddressBytes(c.Pointer.GetText())
		}
		switch c.Operator.GetText() {
		case "=":
			if c.Address != nil {
				operation = vm.JUMP_EQUAL_ADDRESS
			} else if c.Offset != nil {
				operation = vm.JUMP_EQUAL_OFFSET
			} else if c.Pointer != nil {
				operation = vm.JUMP_EQUAL_POINTER
			}
		case "!=":
			{
				if c.Address != nil {
					operation = vm.JUMP_NOT_EQUAL_ADDRESS
				} else if c.Offset != nil {
					operation = vm.JUMP_NOT_EQUAL_OFFSET
				} else if c.Pointer != nil {
					operation = vm.JUMP_NOT_EQUAL_POINTER
				}
			}
		case ">":
			{
				if c.Address != nil {
					operation = vm.JUMP_GREATER_ADDRESS
				} else if c.Offset != nil {
					operation = vm.JUMP_GREATER_OFFSET
				} else if c.Pointer != nil {
					operation = vm.JUMP_GREATER_POINTER
				}
			}
		case ">=":
			{
				if c.Address != nil {
					operation = vm.JUMP_GREATER_EQUAL_ADDRESS
				} else if c.Offset != nil {
					operation = vm.JUMP_GREATER_EQUAL_OFFSET
				} else if c.Pointer != nil {
					operation = vm.JUMP_GREATER_EQUAL_POINTER
				}
			}
		case "<":
			{
				if c.Address != nil {
					operation = vm.JUMP_LESS_ADDRESS
				} else if c.Offset != nil {
					operation = vm.JUMP_LESS_OFFSET
				} else if c.Pointer != nil {
					operation = vm.JUMP_LESS_POINTER
				}
			}
		case "<=":
			{
				if c.Address != nil {
					operation = vm.JUMP_LESS_EQUAL_ADDRESS
				} else if c.Offset != nil {
					operation = vm.JUMP_LESS_EQUAL_OFFSET
				} else if c.Pointer != nil {
					operation = vm.JUMP_LESS_EQUAL_POINTER
				}
			}
		}
		l.Output.WriteByte(operation)
		l.Output.WriteByte(left)
		l.Output.WriteByte(right)
		l.Output.Write(destination)
	} else {
		if c.Address != nil {
			l.Output.WriteByte(vm.JUMP_ADDRESS)
			l.Output.Write(addressBytes(c.Address.GetText()))
		} else if c.Offset != nil {
			l.Output.WriteByte(vm.JUMP_OFFSET)
			l.Output.Write(addressBytes(c.Offset.GetText()))
		} else if c.Pointer != nil {
			l.Output.WriteByte(vm.JUMP_POINTER)
			l.Output.Write(addressBytes(c.Pointer.GetText()))
		} else {
			panic("Unsupported jump.")
		}
	}
}

func (l *listener) ExitInput(c *parser.InputContext) {
	register := registerByte(c.Destination.GetText())
	device := deviceByte(c.Source.GetText())
	l.Output.Write([]byte{vm.INPUT_REGISTER, device, register})
}

func (l *listener) ExitOutput(c *parser.OutputContext) {
	register := registerByte(c.Source.GetText())
	device := deviceByte(c.Destination.GetText())
	l.Output.Write([]byte{vm.OUTPUT_REGISTER, register, device})
}

func (l *listener) ExitStop(c *parser.StopContext) {
	l.Output.WriteByte(vm.STOP)
}

func (l *listener) ExitWait(c *parser.WaitContext) {
	l.Output.WriteByte(vm.WAIT)
}

func absoluteAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "@")
}

func offsetAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "#")
}

func pointerAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "*")
}

func prefixedAddressBytes(value string, prefix string) []byte {
	return addressBytes(strings.Split(value, prefix)[1])
}

func writeValue(l *listener, value string) {
	l.Output.WriteByte(valueByte(value))
}

func writeAddress(l *listener, address string) {
	l.Output.Write(addressBytes(address))
}

func registerOpCode(opcodeByte byte, registerValue byte) byte {
	return opcodeByte | registerValue
}

func registerByte(registerName string) byte {
	register := strings.TrimLeft(registerName, "R")
	registerBytes, _ := hex.DecodeString("0" + register) // We need to prepend the "0" to complete the byte.
	return registerBytes[0]
}

func deviceByte(deviceName string) byte {
	value := strings.TrimLeft(deviceName, "D")
	bytes, _ := hex.DecodeString(value)
	return bytes[0]
}

func valueByte(value string) byte {
	valueBytes, _ := hex.DecodeString(value)
	return valueBytes[0]
}

func addressBytes(address string) []byte {
	addressBytes, _ := hex.DecodeString(address)
	return addressBytes
}

func main() {
	fileName := os.Args[1]
	fmt.Printf("Reading file %s\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes\n", len(data))
	input := antlr.NewInputStream(string(data))
	lexer := parser.NewGrogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewGrogParser(stream)
	outputFile, err := os.Create(fileName + ".grog")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	listener := listener{Output: bufio.NewWriter(outputFile)}
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Program())
	listener.Output.Flush()
}
