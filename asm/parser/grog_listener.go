// Code generated from Grog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grog
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrogListener is a complete listener for a parse tree produced by GrogParser.
type GrogListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterLoad is called when entering the load production.
	EnterLoad(c *LoadContext)

	// EnterStore is called when entering the store production.
	EnterStore(c *StoreContext)

	// EnterCopyRegister is called when entering the CopyRegister production.
	EnterCopyRegister(c *CopyRegisterContext)

	// EnterCopyAbsoluteToAbsolute is called when entering the CopyAbsoluteToAbsolute production.
	EnterCopyAbsoluteToAbsolute(c *CopyAbsoluteToAbsoluteContext)

	// EnterCopyAbsoluteToOffset is called when entering the CopyAbsoluteToOffset production.
	EnterCopyAbsoluteToOffset(c *CopyAbsoluteToOffsetContext)

	// EnterCopyAbsoluteToPointer is called when entering the CopyAbsoluteToPointer production.
	EnterCopyAbsoluteToPointer(c *CopyAbsoluteToPointerContext)

	// EnterCopyOffsetToAbsolute is called when entering the CopyOffsetToAbsolute production.
	EnterCopyOffsetToAbsolute(c *CopyOffsetToAbsoluteContext)

	// EnterCopyOffsetToOffset is called when entering the CopyOffsetToOffset production.
	EnterCopyOffsetToOffset(c *CopyOffsetToOffsetContext)

	// EnterCopyOffsetToPointer is called when entering the CopyOffsetToPointer production.
	EnterCopyOffsetToPointer(c *CopyOffsetToPointerContext)

	// EnterCopyPointerToAbsolute is called when entering the CopyPointerToAbsolute production.
	EnterCopyPointerToAbsolute(c *CopyPointerToAbsoluteContext)

	// EnterCopyPointerToOffset is called when entering the CopyPointerToOffset production.
	EnterCopyPointerToOffset(c *CopyPointerToOffsetContext)

	// EnterCopyPointerToPointer is called when entering the CopyPointerToPointer production.
	EnterCopyPointerToPointer(c *CopyPointerToPointerContext)

	// EnterCopyRightToLeft is called when entering the copyRightToLeft production.
	EnterCopyRightToLeft(c *CopyRightToLeftContext)

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

	// EnterInput is called when entering the input production.
	EnterInput(c *InputContext)

	// EnterOutput is called when entering the output production.
	EnterOutput(c *OutputContext)

	// EnterStop is called when entering the stop production.
	EnterStop(c *StopContext)

	// EnterWait is called when entering the wait production.
	EnterWait(c *WaitContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitLoad is called when exiting the load production.
	ExitLoad(c *LoadContext)

	// ExitStore is called when exiting the store production.
	ExitStore(c *StoreContext)

	// ExitCopyRegister is called when exiting the CopyRegister production.
	ExitCopyRegister(c *CopyRegisterContext)

	// ExitCopyAbsoluteToAbsolute is called when exiting the CopyAbsoluteToAbsolute production.
	ExitCopyAbsoluteToAbsolute(c *CopyAbsoluteToAbsoluteContext)

	// ExitCopyAbsoluteToOffset is called when exiting the CopyAbsoluteToOffset production.
	ExitCopyAbsoluteToOffset(c *CopyAbsoluteToOffsetContext)

	// ExitCopyAbsoluteToPointer is called when exiting the CopyAbsoluteToPointer production.
	ExitCopyAbsoluteToPointer(c *CopyAbsoluteToPointerContext)

	// ExitCopyOffsetToAbsolute is called when exiting the CopyOffsetToAbsolute production.
	ExitCopyOffsetToAbsolute(c *CopyOffsetToAbsoluteContext)

	// ExitCopyOffsetToOffset is called when exiting the CopyOffsetToOffset production.
	ExitCopyOffsetToOffset(c *CopyOffsetToOffsetContext)

	// ExitCopyOffsetToPointer is called when exiting the CopyOffsetToPointer production.
	ExitCopyOffsetToPointer(c *CopyOffsetToPointerContext)

	// ExitCopyPointerToAbsolute is called when exiting the CopyPointerToAbsolute production.
	ExitCopyPointerToAbsolute(c *CopyPointerToAbsoluteContext)

	// ExitCopyPointerToOffset is called when exiting the CopyPointerToOffset production.
	ExitCopyPointerToOffset(c *CopyPointerToOffsetContext)

	// ExitCopyPointerToPointer is called when exiting the CopyPointerToPointer production.
	ExitCopyPointerToPointer(c *CopyPointerToPointerContext)

	// ExitCopyRightToLeft is called when exiting the copyRightToLeft production.
	ExitCopyRightToLeft(c *CopyRightToLeftContext)

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

	// ExitInput is called when exiting the input production.
	ExitInput(c *InputContext)

	// ExitOutput is called when exiting the output production.
	ExitOutput(c *OutputContext)

	// ExitStop is called when exiting the stop production.
	ExitStop(c *StopContext)

	// ExitWait is called when exiting the wait production.
	ExitWait(c *WaitContext)
}
