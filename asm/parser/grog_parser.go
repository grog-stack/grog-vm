// Code generated from Grog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grog

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 164,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 46,
	10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3, 5, 3, 5,
	5, 5, 58, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 96, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 101, 10,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 107, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 124, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 139, 10, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 151, 10, 13,
	3, 13, 5, 13, 154, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 160, 10,
	13, 3, 14, 3, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 2, 2, 2, 192, 2, 29, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 6, 47,
	3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10, 95, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2,
	14, 108, 3, 2, 2, 2, 16, 112, 3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 127,
	3, 2, 2, 2, 22, 132, 3, 2, 2, 2, 24, 153, 3, 2, 2, 2, 26, 161, 3, 2, 2,
	2, 28, 30, 5, 4, 3, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 2, 2, 3,
	34, 3, 3, 2, 2, 2, 35, 46, 5, 14, 8, 2, 36, 46, 5, 16, 9, 2, 37, 46, 5,
	18, 10, 2, 38, 46, 5, 20, 11, 2, 39, 46, 5, 22, 12, 2, 40, 46, 5, 10, 6,
	2, 41, 46, 5, 6, 4, 2, 42, 46, 5, 8, 5, 2, 43, 46, 5, 24, 13, 2, 44, 46,
	5, 26, 14, 2, 45, 35, 3, 2, 2, 2, 45, 36, 3, 2, 2, 2, 45, 37, 3, 2, 2,
	2, 45, 38, 3, 2, 2, 2, 45, 39, 3, 2, 2, 2, 45, 40, 3, 2, 2, 2, 45, 41,
	3, 2, 2, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2,
	46, 5, 3, 2, 2, 2, 47, 48, 7, 32, 2, 2, 48, 53, 7, 8, 2, 2, 49, 54, 7,
	30, 2, 2, 50, 54, 7, 33, 2, 2, 51, 54, 7, 34, 2, 2, 52, 54, 7, 35, 2, 2,
	53, 49, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3,
	2, 2, 2, 54, 7, 3, 2, 2, 2, 55, 58, 7, 32, 2, 2, 56, 58, 7, 30, 2, 2, 57,
	55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 63, 7, 9, 2,
	2, 60, 64, 7, 33, 2, 2, 61, 64, 7, 34, 2, 2, 62, 64, 7, 35, 2, 2, 63, 60,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 9, 3, 2, 2, 2,
	65, 66, 7, 32, 2, 2, 66, 67, 7, 9, 2, 2, 67, 96, 7, 32, 2, 2, 68, 69, 7,
	33, 2, 2, 69, 70, 7, 9, 2, 2, 70, 96, 7, 33, 2, 2, 71, 72, 7, 33, 2, 2,
	72, 73, 7, 9, 2, 2, 73, 96, 7, 34, 2, 2, 74, 75, 7, 33, 2, 2, 75, 76, 7,
	9, 2, 2, 76, 96, 7, 35, 2, 2, 77, 78, 7, 34, 2, 2, 78, 79, 7, 9, 2, 2,
	79, 96, 7, 33, 2, 2, 80, 81, 7, 34, 2, 2, 81, 82, 7, 9, 2, 2, 82, 96, 7,
	34, 2, 2, 83, 84, 7, 34, 2, 2, 84, 85, 7, 9, 2, 2, 85, 96, 7, 35, 2, 2,
	86, 87, 7, 35, 2, 2, 87, 88, 7, 9, 2, 2, 88, 96, 7, 33, 2, 2, 89, 90, 7,
	35, 2, 2, 90, 91, 7, 9, 2, 2, 91, 96, 7, 34, 2, 2, 92, 93, 7, 35, 2, 2,
	93, 94, 7, 9, 2, 2, 94, 96, 7, 35, 2, 2, 95, 65, 3, 2, 2, 2, 95, 68, 3,
	2, 2, 2, 95, 71, 3, 2, 2, 2, 95, 74, 3, 2, 2, 2, 95, 77, 3, 2, 2, 2, 95,
	80, 3, 2, 2, 2, 95, 83, 3, 2, 2, 2, 95, 86, 3, 2, 2, 2, 95, 89, 3, 2, 2,
	2, 95, 92, 3, 2, 2, 2, 96, 11, 3, 2, 2, 2, 97, 101, 7, 33, 2, 2, 98, 101,
	7, 34, 2, 2, 99, 101, 7, 35, 2, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 106, 7, 3, 2, 2,
	103, 107, 7, 33, 2, 2, 104, 107, 7, 34, 2, 2, 105, 107, 7, 35, 2, 2, 106,
	103, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 13, 3,
	2, 2, 2, 108, 109, 7, 32, 2, 2, 109, 110, 7, 8, 2, 2, 110, 111, 7, 10,
	2, 2, 111, 15, 3, 2, 2, 2, 112, 113, 7, 32, 2, 2, 113, 114, 7, 8, 2, 2,
	114, 115, 7, 11, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117, 7, 32, 2, 2, 117,
	118, 7, 8, 2, 2, 118, 123, 7, 32, 2, 2, 119, 124, 7, 12, 2, 2, 120, 124,
	7, 13, 2, 2, 121, 124, 7, 15, 2, 2, 122, 124, 7, 14, 2, 2, 123, 119, 3,
	2, 2, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2, 2,
	2, 124, 125, 3, 2, 2, 2, 125, 126, 7, 32, 2, 2, 126, 19, 3, 2, 2, 2, 127,
	128, 7, 32, 2, 2, 128, 129, 7, 8, 2, 2, 129, 130, 7, 22, 2, 2, 130, 131,
	7, 32, 2, 2, 131, 21, 3, 2, 2, 2, 132, 133, 7, 32, 2, 2, 133, 134, 7, 8,
	2, 2, 134, 138, 7, 32, 2, 2, 135, 139, 7, 23, 2, 2, 136, 139, 7, 25, 2,
	2, 137, 139, 7, 24, 2, 2, 138, 135, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138,
	137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 7, 32, 2, 2, 141, 23,
	3, 2, 2, 2, 142, 143, 7, 28, 2, 2, 143, 150, 7, 32, 2, 2, 144, 151, 7,
	16, 2, 2, 145, 151, 7, 21, 2, 2, 146, 151, 7, 17, 2, 2, 147, 151, 7, 18,
	2, 2, 148, 151, 7, 19, 2, 2, 149, 151, 7, 20, 2, 2, 150, 144, 3, 2, 2,
	2, 150, 145, 3, 2, 2, 2, 150, 146, 3, 2, 2, 2, 150, 147, 3, 2, 2, 2, 150,
	148, 3, 2, 2, 2, 150, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 154,
	7, 32, 2, 2, 153, 142, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 159, 7, 27, 2, 2, 156, 160, 7, 33, 2, 2, 157, 160, 7, 34, 2,
	2, 158, 160, 7, 35, 2, 2, 159, 156, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159,
	158, 3, 2, 2, 2, 160, 25, 3, 2, 2, 2, 161, 162, 7, 26, 2, 2, 162, 27, 3,
	2, 2, 2, 15, 31, 45, 53, 57, 63, 95, 100, 106, 123, 138, 150, 153, 159,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<->'", "", "", "", "", "'<-'", "'->'", "'++'", "'--'", "'+'", "'-'",
	"'/'", "'*'", "'='", "'>'", "'>='", "'<'", "'<='", "'!='", "'NOT'", "'AND'",
	"'XOR'", "'OR'", "'STOP'", "'JUMP'", "'IF'",
}
var symbolicNames = []string{
	"", "", "WHITESPACE", "WS", "COMMENT", "LINE_COMMENT", "LOAD", "STORE",
	"INCREMENT", "DECREMENT", "ADD", "SUBTRACT", "DIVIDE", "MULTIPLY", "EQUAL",
	"GREATER", "GREATER_OR_EQUAL", "LESS", "LESS_OR_EQUAL", "NOT_EQUAL", "NOT",
	"AND", "XOR", "OR", "STOP", "JUMP", "IF", "HEX_DIGIT", "HEXA_BYTE", "WORD",
	"REGISTER", "ABSOLUTE_ADDRESS", "OFFSET_ADDRESS", "POINTER_ADDRESS",
}

var ruleNames = []string{
	"program", "instruction", "load", "store", "copyValue", "copyRightToLeft",
	"increment", "decrement", "arithmeticOperation", "unaryBooleanOperation",
	"binaryBooleanOperation", "jump", "stop",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrogParser struct {
	*antlr.BaseParser
}

func NewGrogParser(input antlr.TokenStream) *GrogParser {
	this := new(GrogParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Grog.g4"

	return this
}

// GrogParser tokens.
const (
	GrogParserEOF              = antlr.TokenEOF
	GrogParserT__0             = 1
	GrogParserWHITESPACE       = 2
	GrogParserWS               = 3
	GrogParserCOMMENT          = 4
	GrogParserLINE_COMMENT     = 5
	GrogParserLOAD             = 6
	GrogParserSTORE            = 7
	GrogParserINCREMENT        = 8
	GrogParserDECREMENT        = 9
	GrogParserADD              = 10
	GrogParserSUBTRACT         = 11
	GrogParserDIVIDE           = 12
	GrogParserMULTIPLY         = 13
	GrogParserEQUAL            = 14
	GrogParserGREATER          = 15
	GrogParserGREATER_OR_EQUAL = 16
	GrogParserLESS             = 17
	GrogParserLESS_OR_EQUAL    = 18
	GrogParserNOT_EQUAL        = 19
	GrogParserNOT              = 20
	GrogParserAND              = 21
	GrogParserXOR              = 22
	GrogParserOR               = 23
	GrogParserSTOP             = 24
	GrogParserJUMP             = 25
	GrogParserIF               = 26
	GrogParserHEX_DIGIT        = 27
	GrogParserHEXA_BYTE        = 28
	GrogParserWORD             = 29
	GrogParserREGISTER         = 30
	GrogParserABSOLUTE_ADDRESS = 31
	GrogParserOFFSET_ADDRESS   = 32
	GrogParserPOINTER_ADDRESS  = 33
)

// GrogParser rules.
const (
	GrogParserRULE_program                = 0
	GrogParserRULE_instruction            = 1
	GrogParserRULE_load                   = 2
	GrogParserRULE_store                  = 3
	GrogParserRULE_copyValue              = 4
	GrogParserRULE_copyRightToLeft        = 5
	GrogParserRULE_increment              = 6
	GrogParserRULE_decrement              = 7
	GrogParserRULE_arithmeticOperation    = 8
	GrogParserRULE_unaryBooleanOperation  = 9
	GrogParserRULE_binaryBooleanOperation = 10
	GrogParserRULE_jump                   = 11
	GrogParserRULE_stop                   = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrogParserEOF, 0)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GrogParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrogParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(GrogParserSTOP-24))|(1<<(GrogParserJUMP-24))|(1<<(GrogParserIF-24))|(1<<(GrogParserHEXA_BYTE-24))|(1<<(GrogParserREGISTER-24))|(1<<(GrogParserABSOLUTE_ADDRESS-24))|(1<<(GrogParserOFFSET_ADDRESS-24))|(1<<(GrogParserPOINTER_ADDRESS-24)))) != 0) {
		{
			p.SetState(26)
			p.Instruction()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(GrogParserEOF)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Increment() IIncrementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncrementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncrementContext)
}

func (s *InstructionContext) Decrement() IDecrementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecrementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecrementContext)
}

func (s *InstructionContext) ArithmeticOperation() IArithmeticOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticOperationContext)
}

func (s *InstructionContext) UnaryBooleanOperation() IUnaryBooleanOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryBooleanOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryBooleanOperationContext)
}

func (s *InstructionContext) BinaryBooleanOperation() IBinaryBooleanOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryBooleanOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryBooleanOperationContext)
}

func (s *InstructionContext) CopyValue() ICopyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyValueContext)
}

func (s *InstructionContext) Load() ILoadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoadContext)
}

func (s *InstructionContext) Store() IStoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStoreContext)
}

func (s *InstructionContext) Jump() IJumpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJumpContext)
}

func (s *InstructionContext) Stop() IStopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStopContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *GrogParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrogParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Increment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Decrement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.ArithmeticOperation()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.UnaryBooleanOperation()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(37)
			p.BinaryBooleanOperation()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(38)
			p.CopyValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(39)
			p.Load()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(40)
			p.Store()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(41)
			p.Jump()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(42)
			p.Stop()
		}

	}

	return localctx
}

// ILoadContext is an interface to support dynamic dispatch.
type ILoadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// GetValue returns the Value token.
	GetValue() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// SetValue sets the Value token.
	SetValue(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsLoadContext differentiates from other interfaces.
	IsLoadContext()
}

type LoadContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
	Value    antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyLoadContext() *LoadContext {
	var p = new(LoadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_load
	return p
}

func (*LoadContext) IsLoadContext() {}

func NewLoadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadContext {
	var p = new(LoadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_load

	return p
}

func (s *LoadContext) GetParser() antlr.Parser { return s.parser }

func (s *LoadContext) GetRegister() antlr.Token { return s.Register }

func (s *LoadContext) GetValue() antlr.Token { return s.Value }

func (s *LoadContext) GetAddress() antlr.Token { return s.Address }

func (s *LoadContext) GetOffset() antlr.Token { return s.Offset }

func (s *LoadContext) GetPointer() antlr.Token { return s.Pointer }

func (s *LoadContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *LoadContext) SetValue(v antlr.Token) { s.Value = v }

func (s *LoadContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *LoadContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *LoadContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *LoadContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *LoadContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *LoadContext) HEXA_BYTE() antlr.TerminalNode {
	return s.GetToken(GrogParserHEXA_BYTE, 0)
}

func (s *LoadContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *LoadContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *LoadContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *LoadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterLoad(s)
	}
}

func (s *LoadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitLoad(s)
	}
}

func (p *GrogParser) Load() (localctx ILoadContext) {
	localctx = NewLoadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrogParserRULE_load)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*LoadContext).Register = _m
	}
	{
		p.SetState(46)
		p.Match(GrogParserLOAD)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserHEXA_BYTE:
		{
			p.SetState(47)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*LoadContext).Value = _m
		}

	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(48)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*LoadContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(49)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*LoadContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(50)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*LoadContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStoreContext is an interface to support dynamic dispatch.
type IStoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// GetValue returns the Value token.
	GetValue() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// SetValue sets the Value token.
	SetValue(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsStoreContext differentiates from other interfaces.
	IsStoreContext()
}

type StoreContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
	Value    antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyStoreContext() *StoreContext {
	var p = new(StoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_store
	return p
}

func (*StoreContext) IsStoreContext() {}

func NewStoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreContext {
	var p = new(StoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_store

	return p
}

func (s *StoreContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreContext) GetRegister() antlr.Token { return s.Register }

func (s *StoreContext) GetValue() antlr.Token { return s.Value }

func (s *StoreContext) GetAddress() antlr.Token { return s.Address }

func (s *StoreContext) GetOffset() antlr.Token { return s.Offset }

func (s *StoreContext) GetPointer() antlr.Token { return s.Pointer }

func (s *StoreContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *StoreContext) SetValue(v antlr.Token) { s.Value = v }

func (s *StoreContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *StoreContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *StoreContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *StoreContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *StoreContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *StoreContext) HEXA_BYTE() antlr.TerminalNode {
	return s.GetToken(GrogParserHEXA_BYTE, 0)
}

func (s *StoreContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *StoreContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *StoreContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *StoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterStore(s)
	}
}

func (s *StoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitStore(s)
	}
}

func (p *GrogParser) Store() (localctx IStoreContext) {
	localctx = NewStoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrogParserRULE_store)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserREGISTER:
		{
			p.SetState(53)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*StoreContext).Register = _m
		}

	case GrogParserHEXA_BYTE:
		{
			p.SetState(54)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*StoreContext).Value = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(57)
		p.Match(GrogParserSTORE)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(58)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*StoreContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(59)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*StoreContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(60)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*StoreContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICopyValueContext is an interface to support dynamic dispatch.
type ICopyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopyValueContext differentiates from other interfaces.
	IsCopyValueContext()
}

type CopyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopyValueContext() *CopyValueContext {
	var p = new(CopyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyValue
	return p
}

func (*CopyValueContext) IsCopyValueContext() {}

func NewCopyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyValueContext {
	var p = new(CopyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyValue

	return p
}

func (s *CopyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyValueContext) CopyFrom(ctx *CopyValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CopyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CopyRegisterContext struct {
	*CopyValueContext
	SourceRegister      antlr.Token
	DestinationRegister antlr.Token
}

func NewCopyRegisterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyRegisterContext {
	var p = new(CopyRegisterContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyRegisterContext) GetSourceRegister() antlr.Token { return s.SourceRegister }

func (s *CopyRegisterContext) GetDestinationRegister() antlr.Token { return s.DestinationRegister }

func (s *CopyRegisterContext) SetSourceRegister(v antlr.Token) { s.SourceRegister = v }

func (s *CopyRegisterContext) SetDestinationRegister(v antlr.Token) { s.DestinationRegister = v }

func (s *CopyRegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyRegisterContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyRegisterContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *CopyRegisterContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *CopyRegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyRegister(s)
	}
}

func (s *CopyRegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyRegister(s)
	}
}

type CopyAbsoluteToPointerContext struct {
	*CopyValueContext
	SourceAddress      antlr.Token
	DestinationPointer antlr.Token
}

func NewCopyAbsoluteToPointerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyAbsoluteToPointerContext {
	var p = new(CopyAbsoluteToPointerContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyAbsoluteToPointerContext) GetSourceAddress() antlr.Token { return s.SourceAddress }

func (s *CopyAbsoluteToPointerContext) GetDestinationPointer() antlr.Token {
	return s.DestinationPointer
}

func (s *CopyAbsoluteToPointerContext) SetSourceAddress(v antlr.Token) { s.SourceAddress = v }

func (s *CopyAbsoluteToPointerContext) SetDestinationPointer(v antlr.Token) { s.DestinationPointer = v }

func (s *CopyAbsoluteToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToPointerContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyAbsoluteToPointerContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyAbsoluteToPointerContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyAbsoluteToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToPointer(s)
	}
}

func (s *CopyAbsoluteToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToPointer(s)
	}
}

type CopyOffsetToAbsoluteContext struct {
	*CopyValueContext
	SourceOffset       antlr.Token
	DestinationAddress antlr.Token
}

func NewCopyOffsetToAbsoluteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyOffsetToAbsoluteContext {
	var p = new(CopyOffsetToAbsoluteContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyOffsetToAbsoluteContext) GetSourceOffset() antlr.Token { return s.SourceOffset }

func (s *CopyOffsetToAbsoluteContext) GetDestinationAddress() antlr.Token {
	return s.DestinationAddress
}

func (s *CopyOffsetToAbsoluteContext) SetSourceOffset(v antlr.Token) { s.SourceOffset = v }

func (s *CopyOffsetToAbsoluteContext) SetDestinationAddress(v antlr.Token) { s.DestinationAddress = v }

func (s *CopyOffsetToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToAbsoluteContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyOffsetToAbsoluteContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyOffsetToAbsoluteContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyOffsetToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToAbsolute(s)
	}
}

func (s *CopyOffsetToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToAbsolute(s)
	}
}

type CopyOffsetToOffsetContext struct {
	*CopyValueContext
	SourceOffset      antlr.Token
	DestinationOffset antlr.Token
}

func NewCopyOffsetToOffsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyOffsetToOffsetContext {
	var p = new(CopyOffsetToOffsetContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyOffsetToOffsetContext) GetSourceOffset() antlr.Token { return s.SourceOffset }

func (s *CopyOffsetToOffsetContext) GetDestinationOffset() antlr.Token { return s.DestinationOffset }

func (s *CopyOffsetToOffsetContext) SetSourceOffset(v antlr.Token) { s.SourceOffset = v }

func (s *CopyOffsetToOffsetContext) SetDestinationOffset(v antlr.Token) { s.DestinationOffset = v }

func (s *CopyOffsetToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToOffsetContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyOffsetToOffsetContext) AllOFFSET_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserOFFSET_ADDRESS)
}

func (s *CopyOffsetToOffsetContext) OFFSET_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, i)
}

func (s *CopyOffsetToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToOffset(s)
	}
}

func (s *CopyOffsetToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToOffset(s)
	}
}

type CopyPointerToOffsetContext struct {
	*CopyValueContext
	SourcePointer     antlr.Token
	DestinationOffset antlr.Token
}

func NewCopyPointerToOffsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyPointerToOffsetContext {
	var p = new(CopyPointerToOffsetContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyPointerToOffsetContext) GetSourcePointer() antlr.Token { return s.SourcePointer }

func (s *CopyPointerToOffsetContext) GetDestinationOffset() antlr.Token { return s.DestinationOffset }

func (s *CopyPointerToOffsetContext) SetSourcePointer(v antlr.Token) { s.SourcePointer = v }

func (s *CopyPointerToOffsetContext) SetDestinationOffset(v antlr.Token) { s.DestinationOffset = v }

func (s *CopyPointerToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToOffsetContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyPointerToOffsetContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyPointerToOffsetContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyPointerToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToOffset(s)
	}
}

func (s *CopyPointerToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToOffset(s)
	}
}

type CopyAbsoluteToOffsetContext struct {
	*CopyValueContext
	SourceAddress     antlr.Token
	DestinationOffset antlr.Token
}

func NewCopyAbsoluteToOffsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyAbsoluteToOffsetContext {
	var p = new(CopyAbsoluteToOffsetContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyAbsoluteToOffsetContext) GetSourceAddress() antlr.Token { return s.SourceAddress }

func (s *CopyAbsoluteToOffsetContext) GetDestinationOffset() antlr.Token { return s.DestinationOffset }

func (s *CopyAbsoluteToOffsetContext) SetSourceAddress(v antlr.Token) { s.SourceAddress = v }

func (s *CopyAbsoluteToOffsetContext) SetDestinationOffset(v antlr.Token) { s.DestinationOffset = v }

func (s *CopyAbsoluteToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToOffsetContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyAbsoluteToOffsetContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyAbsoluteToOffsetContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyAbsoluteToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToOffset(s)
	}
}

func (s *CopyAbsoluteToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToOffset(s)
	}
}

type CopyOffsetToPointerContext struct {
	*CopyValueContext
	SourceOffset       antlr.Token
	DestinationPointer antlr.Token
}

func NewCopyOffsetToPointerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyOffsetToPointerContext {
	var p = new(CopyOffsetToPointerContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyOffsetToPointerContext) GetSourceOffset() antlr.Token { return s.SourceOffset }

func (s *CopyOffsetToPointerContext) GetDestinationPointer() antlr.Token { return s.DestinationPointer }

func (s *CopyOffsetToPointerContext) SetSourceOffset(v antlr.Token) { s.SourceOffset = v }

func (s *CopyOffsetToPointerContext) SetDestinationPointer(v antlr.Token) { s.DestinationPointer = v }

func (s *CopyOffsetToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToPointerContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyOffsetToPointerContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyOffsetToPointerContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyOffsetToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToPointer(s)
	}
}

func (s *CopyOffsetToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToPointer(s)
	}
}

type CopyPointerToAbsoluteContext struct {
	*CopyValueContext
	SourcePointer      antlr.Token
	DestinationAddress antlr.Token
}

func NewCopyPointerToAbsoluteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyPointerToAbsoluteContext {
	var p = new(CopyPointerToAbsoluteContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyPointerToAbsoluteContext) GetSourcePointer() antlr.Token { return s.SourcePointer }

func (s *CopyPointerToAbsoluteContext) GetDestinationAddress() antlr.Token {
	return s.DestinationAddress
}

func (s *CopyPointerToAbsoluteContext) SetSourcePointer(v antlr.Token) { s.SourcePointer = v }

func (s *CopyPointerToAbsoluteContext) SetDestinationAddress(v antlr.Token) { s.DestinationAddress = v }

func (s *CopyPointerToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToAbsoluteContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyPointerToAbsoluteContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyPointerToAbsoluteContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyPointerToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToAbsolute(s)
	}
}

func (s *CopyPointerToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToAbsolute(s)
	}
}

type CopyPointerToPointerContext struct {
	*CopyValueContext
	SourcePointer      antlr.Token
	DestinationPointer antlr.Token
}

func NewCopyPointerToPointerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyPointerToPointerContext {
	var p = new(CopyPointerToPointerContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyPointerToPointerContext) GetSourcePointer() antlr.Token { return s.SourcePointer }

func (s *CopyPointerToPointerContext) GetDestinationPointer() antlr.Token {
	return s.DestinationPointer
}

func (s *CopyPointerToPointerContext) SetSourcePointer(v antlr.Token) { s.SourcePointer = v }

func (s *CopyPointerToPointerContext) SetDestinationPointer(v antlr.Token) { s.DestinationPointer = v }

func (s *CopyPointerToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToPointerContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyPointerToPointerContext) AllPOINTER_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserPOINTER_ADDRESS)
}

func (s *CopyPointerToPointerContext) POINTER_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, i)
}

func (s *CopyPointerToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToPointer(s)
	}
}

func (s *CopyPointerToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToPointer(s)
	}
}

type CopyAbsoluteToAbsoluteContext struct {
	*CopyValueContext
	SourceAddress      antlr.Token
	DestinationAddress antlr.Token
}

func NewCopyAbsoluteToAbsoluteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyAbsoluteToAbsoluteContext {
	var p = new(CopyAbsoluteToAbsoluteContext)

	p.CopyValueContext = NewEmptyCopyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CopyValueContext))

	return p
}

func (s *CopyAbsoluteToAbsoluteContext) GetSourceAddress() antlr.Token { return s.SourceAddress }

func (s *CopyAbsoluteToAbsoluteContext) GetDestinationAddress() antlr.Token {
	return s.DestinationAddress
}

func (s *CopyAbsoluteToAbsoluteContext) SetSourceAddress(v antlr.Token) { s.SourceAddress = v }

func (s *CopyAbsoluteToAbsoluteContext) SetDestinationAddress(v antlr.Token) {
	s.DestinationAddress = v
}

func (s *CopyAbsoluteToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToAbsoluteContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *CopyAbsoluteToAbsoluteContext) AllABSOLUTE_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserABSOLUTE_ADDRESS)
}

func (s *CopyAbsoluteToAbsoluteContext) ABSOLUTE_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, i)
}

func (s *CopyAbsoluteToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToAbsolute(s)
	}
}

func (s *CopyAbsoluteToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToAbsolute(s)
	}
}

func (p *GrogParser) CopyValue() (localctx ICopyValueContext) {
	localctx = NewCopyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrogParserRULE_copyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCopyRegisterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*CopyRegisterContext).SourceRegister = _m
		}
		{
			p.SetState(64)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(65)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*CopyRegisterContext).DestinationRegister = _m
		}

	case 2:
		localctx = NewCopyAbsoluteToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToAbsoluteContext).SourceAddress = _m
		}
		{
			p.SetState(67)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(68)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToAbsoluteContext).DestinationAddress = _m
		}

	case 3:
		localctx = NewCopyAbsoluteToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToOffsetContext).SourceAddress = _m
		}
		{
			p.SetState(70)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(71)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyAbsoluteToOffsetContext).DestinationOffset = _m
		}

	case 4:
		localctx = NewCopyAbsoluteToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToPointerContext).SourceAddress = _m
		}
		{
			p.SetState(73)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(74)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyAbsoluteToPointerContext).DestinationPointer = _m
		}

	case 5:
		localctx = NewCopyOffsetToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToAbsoluteContext).SourceOffset = _m
		}
		{
			p.SetState(76)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(77)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyOffsetToAbsoluteContext).DestinationAddress = _m
		}

	case 6:
		localctx = NewCopyOffsetToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(78)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToOffsetContext).SourceOffset = _m
		}
		{
			p.SetState(79)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(80)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToOffsetContext).DestinationOffset = _m
		}

	case 7:
		localctx = NewCopyOffsetToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToPointerContext).SourceOffset = _m
		}
		{
			p.SetState(82)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(83)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyOffsetToPointerContext).DestinationPointer = _m
		}

	case 8:
		localctx = NewCopyPointerToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(84)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToAbsoluteContext).SourcePointer = _m
		}
		{
			p.SetState(85)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(86)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyPointerToAbsoluteContext).DestinationAddress = _m
		}

	case 9:
		localctx = NewCopyPointerToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(87)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToOffsetContext).SourcePointer = _m
		}
		{
			p.SetState(88)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(89)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyPointerToOffsetContext).DestinationOffset = _m
		}

	case 10:
		localctx = NewCopyPointerToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(90)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToPointerContext).SourcePointer = _m
		}
		{
			p.SetState(91)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(92)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToPointerContext).DestinationPointer = _m
		}

	}

	return localctx
}

// ICopyRightToLeftContext is an interface to support dynamic dispatch.
type ICopyRightToLeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSourceAddress returns the SourceAddress token.
	GetSourceAddress() antlr.Token

	// GetSourceOffset returns the SourceOffset token.
	GetSourceOffset() antlr.Token

	// GetSourcePointer returns the SourcePointer token.
	GetSourcePointer() antlr.Token

	// SetSourceAddress sets the SourceAddress token.
	SetSourceAddress(antlr.Token)

	// SetSourceOffset sets the SourceOffset token.
	SetSourceOffset(antlr.Token)

	// SetSourcePointer sets the SourcePointer token.
	SetSourcePointer(antlr.Token)

	// IsCopyRightToLeftContext differentiates from other interfaces.
	IsCopyRightToLeftContext()
}

type CopyRightToLeftContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	SourceAddress antlr.Token
	SourceOffset  antlr.Token
	SourcePointer antlr.Token
}

func NewEmptyCopyRightToLeftContext() *CopyRightToLeftContext {
	var p = new(CopyRightToLeftContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyRightToLeft
	return p
}

func (*CopyRightToLeftContext) IsCopyRightToLeftContext() {}

func NewCopyRightToLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyRightToLeftContext {
	var p = new(CopyRightToLeftContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyRightToLeft

	return p
}

func (s *CopyRightToLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyRightToLeftContext) GetSourceAddress() antlr.Token { return s.SourceAddress }

func (s *CopyRightToLeftContext) GetSourceOffset() antlr.Token { return s.SourceOffset }

func (s *CopyRightToLeftContext) GetSourcePointer() antlr.Token { return s.SourcePointer }

func (s *CopyRightToLeftContext) SetSourceAddress(v antlr.Token) { s.SourceAddress = v }

func (s *CopyRightToLeftContext) SetSourceOffset(v antlr.Token) { s.SourceOffset = v }

func (s *CopyRightToLeftContext) SetSourcePointer(v antlr.Token) { s.SourcePointer = v }

func (s *CopyRightToLeftContext) AllABSOLUTE_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserABSOLUTE_ADDRESS)
}

func (s *CopyRightToLeftContext) ABSOLUTE_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, i)
}

func (s *CopyRightToLeftContext) AllOFFSET_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserOFFSET_ADDRESS)
}

func (s *CopyRightToLeftContext) OFFSET_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, i)
}

func (s *CopyRightToLeftContext) AllPOINTER_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserPOINTER_ADDRESS)
}

func (s *CopyRightToLeftContext) POINTER_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, i)
}

func (s *CopyRightToLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyRightToLeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyRightToLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyRightToLeft(s)
	}
}

func (s *CopyRightToLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyRightToLeft(s)
	}
}

func (p *GrogParser) CopyRightToLeft() (localctx ICopyRightToLeftContext) {
	localctx = NewCopyRightToLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrogParserRULE_copyRightToLeft)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(95)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceAddress = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(96)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceOffset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(97)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourcePointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(100)
		p.Match(GrogParserT__0)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(101)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceAddress = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(102)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceOffset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(103)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourcePointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIncrementContext is an interface to support dynamic dispatch.
type IIncrementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// IsIncrementContext differentiates from other interfaces.
	IsIncrementContext()
}

type IncrementContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
}

func NewEmptyIncrementContext() *IncrementContext {
	var p = new(IncrementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_increment
	return p
}

func (*IncrementContext) IsIncrementContext() {}

func NewIncrementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrementContext {
	var p = new(IncrementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_increment

	return p
}

func (s *IncrementContext) GetParser() antlr.Parser { return s.parser }

func (s *IncrementContext) GetRegister() antlr.Token { return s.Register }

func (s *IncrementContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *IncrementContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *IncrementContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(GrogParserINCREMENT, 0)
}

func (s *IncrementContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *IncrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterIncrement(s)
	}
}

func (s *IncrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitIncrement(s)
	}
}

func (p *GrogParser) Increment() (localctx IIncrementContext) {
	localctx = NewIncrementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrogParserRULE_increment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*IncrementContext).Register = _m
	}
	{
		p.SetState(107)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(108)
		p.Match(GrogParserINCREMENT)
	}

	return localctx
}

// IDecrementContext is an interface to support dynamic dispatch.
type IDecrementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// IsDecrementContext differentiates from other interfaces.
	IsDecrementContext()
}

type DecrementContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
}

func NewEmptyDecrementContext() *DecrementContext {
	var p = new(DecrementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_decrement
	return p
}

func (*DecrementContext) IsDecrementContext() {}

func NewDecrementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecrementContext {
	var p = new(DecrementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_decrement

	return p
}

func (s *DecrementContext) GetParser() antlr.Parser { return s.parser }

func (s *DecrementContext) GetRegister() antlr.Token { return s.Register }

func (s *DecrementContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *DecrementContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *DecrementContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(GrogParserDECREMENT, 0)
}

func (s *DecrementContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *DecrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterDecrement(s)
	}
}

func (s *DecrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitDecrement(s)
	}
}

func (p *GrogParser) Decrement() (localctx IDecrementContext) {
	localctx = NewDecrementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrogParserRULE_decrement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*DecrementContext).Register = _m
	}
	{
		p.SetState(111)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(112)
		p.Match(GrogParserDECREMENT)
	}

	return localctx
}

// IArithmeticOperationContext is an interface to support dynamic dispatch.
type IArithmeticOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetLeft returns the Left token.
	GetLeft() antlr.Token

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetRight returns the Right token.
	GetRight() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetLeft sets the Left token.
	SetLeft(antlr.Token)

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetRight sets the Right token.
	SetRight(antlr.Token)

	// IsArithmeticOperationContext differentiates from other interfaces.
	IsArithmeticOperationContext()
}

type ArithmeticOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Left        antlr.Token
	Operator    antlr.Token
	Right       antlr.Token
}

func NewEmptyArithmeticOperationContext() *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_arithmeticOperation
	return p
}

func (*ArithmeticOperationContext) IsArithmeticOperationContext() {}

func NewArithmeticOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_arithmeticOperation

	return p
}

func (s *ArithmeticOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *ArithmeticOperationContext) GetLeft() antlr.Token { return s.Left }

func (s *ArithmeticOperationContext) GetOperator() antlr.Token { return s.Operator }

func (s *ArithmeticOperationContext) GetRight() antlr.Token { return s.Right }

func (s *ArithmeticOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *ArithmeticOperationContext) SetLeft(v antlr.Token) { s.Left = v }

func (s *ArithmeticOperationContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *ArithmeticOperationContext) SetRight(v antlr.Token) { s.Right = v }

func (s *ArithmeticOperationContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *ArithmeticOperationContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *ArithmeticOperationContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *ArithmeticOperationContext) ADD() antlr.TerminalNode {
	return s.GetToken(GrogParserADD, 0)
}

func (s *ArithmeticOperationContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(GrogParserSUBTRACT, 0)
}

func (s *ArithmeticOperationContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(GrogParserMULTIPLY, 0)
}

func (s *ArithmeticOperationContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(GrogParserDIVIDE, 0)
}

func (s *ArithmeticOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterArithmeticOperation(s)
	}
}

func (s *ArithmeticOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitArithmeticOperation(s)
	}
}

func (p *GrogParser) ArithmeticOperation() (localctx IArithmeticOperationContext) {
	localctx = NewArithmeticOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrogParserRULE_arithmeticOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Destination = _m
	}
	{
		p.SetState(115)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(116)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Left = _m
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserADD:
		{
			p.SetState(117)

			var _m = p.Match(GrogParserADD)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserSUBTRACT:
		{
			p.SetState(118)

			var _m = p.Match(GrogParserSUBTRACT)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserMULTIPLY:
		{
			p.SetState(119)

			var _m = p.Match(GrogParserMULTIPLY)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserDIVIDE:
		{
			p.SetState(120)

			var _m = p.Match(GrogParserDIVIDE)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(123)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Right = _m
	}

	return localctx
}

// IUnaryBooleanOperationContext is an interface to support dynamic dispatch.
type IUnaryBooleanOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetOperand returns the Operand token.
	GetOperand() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetOperand sets the Operand token.
	SetOperand(antlr.Token)

	// IsUnaryBooleanOperationContext differentiates from other interfaces.
	IsUnaryBooleanOperationContext()
}

type UnaryBooleanOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Operand     antlr.Token
}

func NewEmptyUnaryBooleanOperationContext() *UnaryBooleanOperationContext {
	var p = new(UnaryBooleanOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_unaryBooleanOperation
	return p
}

func (*UnaryBooleanOperationContext) IsUnaryBooleanOperationContext() {}

func NewUnaryBooleanOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryBooleanOperationContext {
	var p = new(UnaryBooleanOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_unaryBooleanOperation

	return p
}

func (s *UnaryBooleanOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryBooleanOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *UnaryBooleanOperationContext) GetOperand() antlr.Token { return s.Operand }

func (s *UnaryBooleanOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *UnaryBooleanOperationContext) SetOperand(v antlr.Token) { s.Operand = v }

func (s *UnaryBooleanOperationContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *UnaryBooleanOperationContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrogParserNOT, 0)
}

func (s *UnaryBooleanOperationContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *UnaryBooleanOperationContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *UnaryBooleanOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryBooleanOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryBooleanOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterUnaryBooleanOperation(s)
	}
}

func (s *UnaryBooleanOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitUnaryBooleanOperation(s)
	}
}

func (p *GrogParser) UnaryBooleanOperation() (localctx IUnaryBooleanOperationContext) {
	localctx = NewUnaryBooleanOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrogParserRULE_unaryBooleanOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*UnaryBooleanOperationContext).Destination = _m
	}
	{
		p.SetState(126)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(127)
		p.Match(GrogParserNOT)
	}
	{
		p.SetState(128)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*UnaryBooleanOperationContext).Operand = _m
	}

	return localctx
}

// IBinaryBooleanOperationContext is an interface to support dynamic dispatch.
type IBinaryBooleanOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetLeft returns the Left token.
	GetLeft() antlr.Token

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetRight returns the Right token.
	GetRight() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetLeft sets the Left token.
	SetLeft(antlr.Token)

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetRight sets the Right token.
	SetRight(antlr.Token)

	// IsBinaryBooleanOperationContext differentiates from other interfaces.
	IsBinaryBooleanOperationContext()
}

type BinaryBooleanOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Left        antlr.Token
	Operator    antlr.Token
	Right       antlr.Token
}

func NewEmptyBinaryBooleanOperationContext() *BinaryBooleanOperationContext {
	var p = new(BinaryBooleanOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_binaryBooleanOperation
	return p
}

func (*BinaryBooleanOperationContext) IsBinaryBooleanOperationContext() {}

func NewBinaryBooleanOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryBooleanOperationContext {
	var p = new(BinaryBooleanOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_binaryBooleanOperation

	return p
}

func (s *BinaryBooleanOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryBooleanOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *BinaryBooleanOperationContext) GetLeft() antlr.Token { return s.Left }

func (s *BinaryBooleanOperationContext) GetOperator() antlr.Token { return s.Operator }

func (s *BinaryBooleanOperationContext) GetRight() antlr.Token { return s.Right }

func (s *BinaryBooleanOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *BinaryBooleanOperationContext) SetLeft(v antlr.Token) { s.Left = v }

func (s *BinaryBooleanOperationContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *BinaryBooleanOperationContext) SetRight(v antlr.Token) { s.Right = v }

func (s *BinaryBooleanOperationContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *BinaryBooleanOperationContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *BinaryBooleanOperationContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *BinaryBooleanOperationContext) AND() antlr.TerminalNode {
	return s.GetToken(GrogParserAND, 0)
}

func (s *BinaryBooleanOperationContext) OR() antlr.TerminalNode {
	return s.GetToken(GrogParserOR, 0)
}

func (s *BinaryBooleanOperationContext) XOR() antlr.TerminalNode {
	return s.GetToken(GrogParserXOR, 0)
}

func (s *BinaryBooleanOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryBooleanOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryBooleanOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterBinaryBooleanOperation(s)
	}
}

func (s *BinaryBooleanOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitBinaryBooleanOperation(s)
	}
}

func (p *GrogParser) BinaryBooleanOperation() (localctx IBinaryBooleanOperationContext) {
	localctx = NewBinaryBooleanOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrogParserRULE_binaryBooleanOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Destination = _m
	}
	{
		p.SetState(131)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(132)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Left = _m
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserAND:
		{
			p.SetState(133)

			var _m = p.Match(GrogParserAND)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserOR:
		{
			p.SetState(134)

			var _m = p.Match(GrogParserOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserXOR:
		{
			p.SetState(135)

			var _m = p.Match(GrogParserXOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(138)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Right = _m
	}

	return localctx
}

// IJumpContext is an interface to support dynamic dispatch.
type IJumpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the Left token.
	GetLeft() antlr.Token

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetRight returns the Right token.
	GetRight() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetLeft sets the Left token.
	SetLeft(antlr.Token)

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetRight sets the Right token.
	SetRight(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsJumpContext differentiates from other interfaces.
	IsJumpContext()
}

type JumpContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Left     antlr.Token
	Operator antlr.Token
	Right    antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyJumpContext() *JumpContext {
	var p = new(JumpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_jump
	return p
}

func (*JumpContext) IsJumpContext() {}

func NewJumpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpContext {
	var p = new(JumpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_jump

	return p
}

func (s *JumpContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpContext) GetLeft() antlr.Token { return s.Left }

func (s *JumpContext) GetOperator() antlr.Token { return s.Operator }

func (s *JumpContext) GetRight() antlr.Token { return s.Right }

func (s *JumpContext) GetAddress() antlr.Token { return s.Address }

func (s *JumpContext) GetOffset() antlr.Token { return s.Offset }

func (s *JumpContext) GetPointer() antlr.Token { return s.Pointer }

func (s *JumpContext) SetLeft(v antlr.Token) { s.Left = v }

func (s *JumpContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *JumpContext) SetRight(v antlr.Token) { s.Right = v }

func (s *JumpContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *JumpContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *JumpContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *JumpContext) JUMP() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP, 0)
}

func (s *JumpContext) IF() antlr.TerminalNode {
	return s.GetToken(GrogParserIF, 0)
}

func (s *JumpContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *JumpContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *JumpContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *JumpContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *JumpContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *JumpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserEQUAL, 0)
}

func (s *JumpContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserNOT_EQUAL, 0)
}

func (s *JumpContext) GREATER() antlr.TerminalNode {
	return s.GetToken(GrogParserGREATER, 0)
}

func (s *JumpContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserGREATER_OR_EQUAL, 0)
}

func (s *JumpContext) LESS() antlr.TerminalNode {
	return s.GetToken(GrogParserLESS, 0)
}

func (s *JumpContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserLESS_OR_EQUAL, 0)
}

func (s *JumpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterJump(s)
	}
}

func (s *JumpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitJump(s)
	}
}

func (p *GrogParser) Jump() (localctx IJumpContext) {
	localctx = NewJumpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrogParserRULE_jump)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrogParserIF {
		{
			p.SetState(140)
			p.Match(GrogParserIF)
		}
		{
			p.SetState(141)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*JumpContext).Left = _m
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GrogParserEQUAL:
			{
				p.SetState(142)

				var _m = p.Match(GrogParserEQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserNOT_EQUAL:
			{
				p.SetState(143)

				var _m = p.Match(GrogParserNOT_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserGREATER:
			{
				p.SetState(144)

				var _m = p.Match(GrogParserGREATER)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserGREATER_OR_EQUAL:
			{
				p.SetState(145)

				var _m = p.Match(GrogParserGREATER_OR_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserLESS:
			{
				p.SetState(146)

				var _m = p.Match(GrogParserLESS)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserLESS_OR_EQUAL:
			{
				p.SetState(147)

				var _m = p.Match(GrogParserLESS_OR_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(150)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*JumpContext).Right = _m
		}

	}
	{
		p.SetState(153)
		p.Match(GrogParserJUMP)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(154)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*JumpContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(155)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*JumpContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(156)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*JumpContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStopContext is an interface to support dynamic dispatch.
type IStopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStopContext differentiates from other interfaces.
	IsStopContext()
}

type StopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStopContext() *StopContext {
	var p = new(StopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_stop
	return p
}

func (*StopContext) IsStopContext() {}

func NewStopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StopContext {
	var p = new(StopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_stop

	return p
}

func (s *StopContext) GetParser() antlr.Parser { return s.parser }

func (s *StopContext) STOP() antlr.TerminalNode {
	return s.GetToken(GrogParserSTOP, 0)
}

func (s *StopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterStop(s)
	}
}

func (s *StopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitStop(s)
	}
}

func (p *GrogParser) Stop() (localctx IStopContext) {
	localctx = NewStopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrogParserRULE_stop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(GrogParserSTOP)
	}

	return localctx
}
