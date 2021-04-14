// Code generated from Grog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grog
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrogListener is a complete listener for a parse tree produced by GrogParser.
type BaseGrogListener struct{}

var _ GrogListener = &BaseGrogListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrogListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrogListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrogListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrogListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGrogListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGrogListener) ExitProgram(ctx *ProgramContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseGrogListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseGrogListener) ExitInstruction(ctx *InstructionContext) {}

// EnterLoad is called when production load is entered.
func (s *BaseGrogListener) EnterLoad(ctx *LoadContext) {}

// ExitLoad is called when production load is exited.
func (s *BaseGrogListener) ExitLoad(ctx *LoadContext) {}

// EnterStore is called when production store is entered.
func (s *BaseGrogListener) EnterStore(ctx *StoreContext) {}

// ExitStore is called when production store is exited.
func (s *BaseGrogListener) ExitStore(ctx *StoreContext) {}

// EnterCopyRegister is called when production CopyRegister is entered.
func (s *BaseGrogListener) EnterCopyRegister(ctx *CopyRegisterContext) {}

// ExitCopyRegister is called when production CopyRegister is exited.
func (s *BaseGrogListener) ExitCopyRegister(ctx *CopyRegisterContext) {}

// EnterCopyRightToLeft is called when production copyRightToLeft is entered.
func (s *BaseGrogListener) EnterCopyRightToLeft(ctx *CopyRightToLeftContext) {}

// ExitCopyRightToLeft is called when production copyRightToLeft is exited.
func (s *BaseGrogListener) ExitCopyRightToLeft(ctx *CopyRightToLeftContext) {}

// EnterIncrement is called when production increment is entered.
func (s *BaseGrogListener) EnterIncrement(ctx *IncrementContext) {}

// ExitIncrement is called when production increment is exited.
func (s *BaseGrogListener) ExitIncrement(ctx *IncrementContext) {}

// EnterDecrement is called when production decrement is entered.
func (s *BaseGrogListener) EnterDecrement(ctx *DecrementContext) {}

// ExitDecrement is called when production decrement is exited.
func (s *BaseGrogListener) ExitDecrement(ctx *DecrementContext) {}

// EnterArithmeticOperation is called when production arithmeticOperation is entered.
func (s *BaseGrogListener) EnterArithmeticOperation(ctx *ArithmeticOperationContext) {}

// ExitArithmeticOperation is called when production arithmeticOperation is exited.
func (s *BaseGrogListener) ExitArithmeticOperation(ctx *ArithmeticOperationContext) {}

// EnterUnaryBooleanOperation is called when production unaryBooleanOperation is entered.
func (s *BaseGrogListener) EnterUnaryBooleanOperation(ctx *UnaryBooleanOperationContext) {}

// ExitUnaryBooleanOperation is called when production unaryBooleanOperation is exited.
func (s *BaseGrogListener) ExitUnaryBooleanOperation(ctx *UnaryBooleanOperationContext) {}

// EnterBinaryBooleanOperation is called when production binaryBooleanOperation is entered.
func (s *BaseGrogListener) EnterBinaryBooleanOperation(ctx *BinaryBooleanOperationContext) {}

// ExitBinaryBooleanOperation is called when production binaryBooleanOperation is exited.
func (s *BaseGrogListener) ExitBinaryBooleanOperation(ctx *BinaryBooleanOperationContext) {}

// EnterJump is called when production jump is entered.
func (s *BaseGrogListener) EnterJump(ctx *JumpContext) {}

// ExitJump is called when production jump is exited.
func (s *BaseGrogListener) ExitJump(ctx *JumpContext) {}

// EnterInput is called when production input is entered.
func (s *BaseGrogListener) EnterInput(ctx *InputContext) {}

// ExitInput is called when production input is exited.
func (s *BaseGrogListener) ExitInput(ctx *InputContext) {}

// EnterOutput is called when production output is entered.
func (s *BaseGrogListener) EnterOutput(ctx *OutputContext) {}

// ExitOutput is called when production output is exited.
func (s *BaseGrogListener) ExitOutput(ctx *OutputContext) {}

// EnterStop is called when production stop is entered.
func (s *BaseGrogListener) EnterStop(ctx *StopContext) {}

// ExitStop is called when production stop is exited.
func (s *BaseGrogListener) ExitStop(ctx *StopContext) {}

// EnterWait is called when production wait is entered.
func (s *BaseGrogListener) EnterWait(ctx *WaitContext) {}

// ExitWait is called when production wait is exited.
func (s *BaseGrogListener) ExitWait(ctx *WaitContext) {}
