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
