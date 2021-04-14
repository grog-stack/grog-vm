// Code generated from Grog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grog
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrogListener is a complete listener for a parse tree produced by GrogParser.
type GrogListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterLoad is called when entering the load production.
	EnterLoad(c *LoadContext)

	// EnterStore is called when entering the store production.
	EnterStore(c *StoreContext)

	// EnterCopyValue is called when entering the copyValue production.
	EnterCopyValue(c *CopyValueContext)

	// EnterCopyRegister is called when entering the copyRegister production.
	EnterCopyRegister(c *CopyRegisterContext)

	// EnterCopyAbsoluteToAbsolute is called when entering the copyAbsoluteToAbsolute production.
	EnterCopyAbsoluteToAbsolute(c *CopyAbsoluteToAbsoluteContext)

	// EnterCopyOffsetToAbsolute is called when entering the copyOffsetToAbsolute production.
	EnterCopyOffsetToAbsolute(c *CopyOffsetToAbsoluteContext)

	// EnterCopyPointerToAbsolute is called when entering the copyPointerToAbsolute production.
	EnterCopyPointerToAbsolute(c *CopyPointerToAbsoluteContext)

	// EnterCopyAbsoluteToOffset is called when entering the copyAbsoluteToOffset production.
	EnterCopyAbsoluteToOffset(c *CopyAbsoluteToOffsetContext)

	// EnterCopyOffsetToOffset is called when entering the copyOffsetToOffset production.
	EnterCopyOffsetToOffset(c *CopyOffsetToOffsetContext)

	// EnterCopyPointerToOffset is called when entering the copyPointerToOffset production.
	EnterCopyPointerToOffset(c *CopyPointerToOffsetContext)

	// EnterCopyAbsoluteToPointer is called when entering the copyAbsoluteToPointer production.
	EnterCopyAbsoluteToPointer(c *CopyAbsoluteToPointerContext)

	// EnterCopyOffsetToPointer is called when entering the copyOffsetToPointer production.
	EnterCopyOffsetToPointer(c *CopyOffsetToPointerContext)

	// EnterCopyPointerToPointer is called when entering the copyPointerToPointer production.
	EnterCopyPointerToPointer(c *CopyPointerToPointerContext)

	// EnterIncrement is called when entering the increment production.
	EnterIncrement(c *IncrementContext)

	// EnterDecrement is called when entering the decrement production.
	EnterDecrement(c *DecrementContext)

	// EnterArithmeticOperation is called when entering the arithmeticOperation production.
	EnterArithmeticOperation(c *ArithmeticOperationContext)

	// EnterUnaryBooleanOperation is called when entering the unaryBooleanOperation production.
	EnterUnaryBooleanOperation(c *UnaryBooleanOperationContext)

	// EnterBinaryBooleanOperation is called when entering the binaryBooleanOperation production.
	EnterBinaryBooleanOperation(c *BinaryBooleanOperationContext)

	// EnterJump is called when entering the jump production.
	EnterJump(c *JumpContext)

	// EnterIo is called when entering the io production.
	EnterIo(c *IoContext)

	// EnterStop is called when entering the stop production.
	EnterStop(c *StopContext)

	// EnterWait is called when entering the wait production.
	EnterWait(c *WaitContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitLoad is called when exiting the load production.
	ExitLoad(c *LoadContext)

	// ExitStore is called when exiting the store production.
	ExitStore(c *StoreContext)

	// ExitCopyValue is called when exiting the copyValue production.
	ExitCopyValue(c *CopyValueContext)

	// ExitCopyRegister is called when exiting the copyRegister production.
	ExitCopyRegister(c *CopyRegisterContext)

	// ExitCopyAbsoluteToAbsolute is called when exiting the copyAbsoluteToAbsolute production.
	ExitCopyAbsoluteToAbsolute(c *CopyAbsoluteToAbsoluteContext)

	// ExitCopyOffsetToAbsolute is called when exiting the copyOffsetToAbsolute production.
	ExitCopyOffsetToAbsolute(c *CopyOffsetToAbsoluteContext)

	// ExitCopyPointerToAbsolute is called when exiting the copyPointerToAbsolute production.
	ExitCopyPointerToAbsolute(c *CopyPointerToAbsoluteContext)

	// ExitCopyAbsoluteToOffset is called when exiting the copyAbsoluteToOffset production.
	ExitCopyAbsoluteToOffset(c *CopyAbsoluteToOffsetContext)

	// ExitCopyOffsetToOffset is called when exiting the copyOffsetToOffset production.
	ExitCopyOffsetToOffset(c *CopyOffsetToOffsetContext)

	// ExitCopyPointerToOffset is called when exiting the copyPointerToOffset production.
	ExitCopyPointerToOffset(c *CopyPointerToOffsetContext)

	// ExitCopyAbsoluteToPointer is called when exiting the copyAbsoluteToPointer production.
	ExitCopyAbsoluteToPointer(c *CopyAbsoluteToPointerContext)

	// ExitCopyOffsetToPointer is called when exiting the copyOffsetToPointer production.
	ExitCopyOffsetToPointer(c *CopyOffsetToPointerContext)

	// ExitCopyPointerToPointer is called when exiting the copyPointerToPointer production.
	ExitCopyPointerToPointer(c *CopyPointerToPointerContext)

	// ExitIncrement is called when exiting the increment production.
	ExitIncrement(c *IncrementContext)

	// ExitDecrement is called when exiting the decrement production.
	ExitDecrement(c *DecrementContext)

	// ExitArithmeticOperation is called when exiting the arithmeticOperation production.
	ExitArithmeticOperation(c *ArithmeticOperationContext)

	// ExitUnaryBooleanOperation is called when exiting the unaryBooleanOperation production.
	ExitUnaryBooleanOperation(c *UnaryBooleanOperationContext)

	// ExitBinaryBooleanOperation is called when exiting the binaryBooleanOperation production.
	ExitBinaryBooleanOperation(c *BinaryBooleanOperationContext)

	// ExitJump is called when exiting the jump production.
	ExitJump(c *JumpContext)

	// ExitIo is called when exiting the io production.
	ExitIo(c *IoContext)

	// ExitStop is called when exiting the stop production.
	ExitStop(c *StopContext)

	// ExitWait is called when exiting the wait production.
	ExitWait(c *WaitContext)
}
