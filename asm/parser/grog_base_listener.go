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

// EnterCopyValue is called when production copyValue is entered.
func (s *BaseGrogListener) EnterCopyValue(ctx *CopyValueContext) {}

// ExitCopyValue is called when production copyValue is exited.
func (s *BaseGrogListener) ExitCopyValue(ctx *CopyValueContext) {}

// EnterCopyRegister is called when production copyRegister is entered.
func (s *BaseGrogListener) EnterCopyRegister(ctx *CopyRegisterContext) {}

// ExitCopyRegister is called when production copyRegister is exited.
func (s *BaseGrogListener) ExitCopyRegister(ctx *CopyRegisterContext) {}

// EnterCopyAbsoluteToAbsolute is called when production copyAbsoluteToAbsolute is entered.
func (s *BaseGrogListener) EnterCopyAbsoluteToAbsolute(ctx *CopyAbsoluteToAbsoluteContext) {}

// ExitCopyAbsoluteToAbsolute is called when production copyAbsoluteToAbsolute is exited.
func (s *BaseGrogListener) ExitCopyAbsoluteToAbsolute(ctx *CopyAbsoluteToAbsoluteContext) {}

// EnterCopyOffsetToAbsolute is called when production copyOffsetToAbsolute is entered.
func (s *BaseGrogListener) EnterCopyOffsetToAbsolute(ctx *CopyOffsetToAbsoluteContext) {}

// ExitCopyOffsetToAbsolute is called when production copyOffsetToAbsolute is exited.
func (s *BaseGrogListener) ExitCopyOffsetToAbsolute(ctx *CopyOffsetToAbsoluteContext) {}

// EnterCopyPointerToAbsolute is called when production copyPointerToAbsolute is entered.
func (s *BaseGrogListener) EnterCopyPointerToAbsolute(ctx *CopyPointerToAbsoluteContext) {}

// ExitCopyPointerToAbsolute is called when production copyPointerToAbsolute is exited.
func (s *BaseGrogListener) ExitCopyPointerToAbsolute(ctx *CopyPointerToAbsoluteContext) {}

// EnterCopyAbsoluteToOffset is called when production copyAbsoluteToOffset is entered.
func (s *BaseGrogListener) EnterCopyAbsoluteToOffset(ctx *CopyAbsoluteToOffsetContext) {}

// ExitCopyAbsoluteToOffset is called when production copyAbsoluteToOffset is exited.
func (s *BaseGrogListener) ExitCopyAbsoluteToOffset(ctx *CopyAbsoluteToOffsetContext) {}

// EnterCopyOffsetToOffset is called when production copyOffsetToOffset is entered.
func (s *BaseGrogListener) EnterCopyOffsetToOffset(ctx *CopyOffsetToOffsetContext) {}

// ExitCopyOffsetToOffset is called when production copyOffsetToOffset is exited.
func (s *BaseGrogListener) ExitCopyOffsetToOffset(ctx *CopyOffsetToOffsetContext) {}

// EnterCopyPointerToOffset is called when production copyPointerToOffset is entered.
func (s *BaseGrogListener) EnterCopyPointerToOffset(ctx *CopyPointerToOffsetContext) {}

// ExitCopyPointerToOffset is called when production copyPointerToOffset is exited.
func (s *BaseGrogListener) ExitCopyPointerToOffset(ctx *CopyPointerToOffsetContext) {}

// EnterCopyAbsoluteToPointer is called when production copyAbsoluteToPointer is entered.
func (s *BaseGrogListener) EnterCopyAbsoluteToPointer(ctx *CopyAbsoluteToPointerContext) {}

// ExitCopyAbsoluteToPointer is called when production copyAbsoluteToPointer is exited.
func (s *BaseGrogListener) ExitCopyAbsoluteToPointer(ctx *CopyAbsoluteToPointerContext) {}

// EnterCopyOffsetToPointer is called when production copyOffsetToPointer is entered.
func (s *BaseGrogListener) EnterCopyOffsetToPointer(ctx *CopyOffsetToPointerContext) {}

// ExitCopyOffsetToPointer is called when production copyOffsetToPointer is exited.
func (s *BaseGrogListener) ExitCopyOffsetToPointer(ctx *CopyOffsetToPointerContext) {}

// EnterCopyPointerToPointer is called when production copyPointerToPointer is entered.
func (s *BaseGrogListener) EnterCopyPointerToPointer(ctx *CopyPointerToPointerContext) {}

// ExitCopyPointerToPointer is called when production copyPointerToPointer is exited.
func (s *BaseGrogListener) ExitCopyPointerToPointer(ctx *CopyPointerToPointerContext) {}

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

// EnterIo is called when production io is entered.
func (s *BaseGrogListener) EnterIo(ctx *IoContext) {}

// ExitIo is called when production io is exited.
func (s *BaseGrogListener) ExitIo(ctx *IoContext) {}

// EnterStop is called when production stop is entered.
func (s *BaseGrogListener) EnterStop(ctx *StopContext) {}

// ExitStop is called when production stop is exited.
func (s *BaseGrogListener) ExitStop(ctx *StopContext) {}

// EnterWait is called when production wait is entered.
func (s *BaseGrogListener) EnterWait(ctx *WaitContext) {}

// ExitWait is called when production wait is exited.
func (s *BaseGrogListener) ExitWait(ctx *WaitContext) {}
