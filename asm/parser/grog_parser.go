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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 181,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 6,
	2, 36, 10, 2, 13, 2, 14, 2, 37, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 55, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3, 5, 5, 5, 67,
	10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 105, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 110, 10, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 116, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 131, 10, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 146, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 158, 10, 13, 3, 13, 5, 13, 161,
	10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 167, 10, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2,
	2, 2, 209, 2, 35, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8,
	66, 3, 2, 2, 2, 10, 104, 3, 2, 2, 2, 12, 109, 3, 2, 2, 2, 14, 117, 3, 2,
	2, 2, 16, 120, 3, 2, 2, 2, 18, 123, 3, 2, 2, 2, 20, 134, 3, 2, 2, 2, 22,
	139, 3, 2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 168, 3, 2, 2, 2, 28, 172, 3,
	2, 2, 2, 30, 176, 3, 2, 2, 2, 32, 178, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2,
	35, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3,
	2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 7, 2, 2, 3, 40, 3, 3, 2, 2, 2, 41,
	55, 5, 14, 8, 2, 42, 55, 5, 16, 9, 2, 43, 55, 5, 18, 10, 2, 44, 55, 5,
	20, 11, 2, 45, 55, 5, 22, 12, 2, 46, 55, 5, 10, 6, 2, 47, 55, 5, 6, 4,
	2, 48, 55, 5, 8, 5, 2, 49, 55, 5, 24, 13, 2, 50, 55, 5, 26, 14, 2, 51,
	55, 5, 28, 15, 2, 52, 55, 5, 30, 16, 2, 53, 55, 5, 32, 17, 2, 54, 41, 3,
	2, 2, 2, 54, 42, 3, 2, 2, 2, 54, 43, 3, 2, 2, 2, 54, 44, 3, 2, 2, 2, 54,
	45, 3, 2, 2, 2, 54, 46, 3, 2, 2, 2, 54, 47, 3, 2, 2, 2, 54, 48, 3, 2, 2,
	2, 54, 49, 3, 2, 2, 2, 54, 50, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 5, 3, 2, 2, 2, 56, 57, 7, 9, 2, 2,
	57, 62, 7, 34, 2, 2, 58, 63, 7, 32, 2, 2, 59, 63, 7, 36, 2, 2, 60, 63,
	7, 37, 2, 2, 61, 63, 7, 38, 2, 2, 62, 58, 3, 2, 2, 2, 62, 59, 3, 2, 2,
	2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 7, 3, 2, 2, 2, 64, 67, 7,
	34, 2, 2, 65, 67, 7, 32, 2, 2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 72, 7, 10, 2, 2, 69, 73, 7, 36, 2, 2, 70, 73, 7,
	37, 2, 2, 71, 73, 7, 38, 2, 2, 72, 69, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2,
	72, 71, 3, 2, 2, 2, 73, 9, 3, 2, 2, 2, 74, 75, 7, 34, 2, 2, 75, 76, 7,
	10, 2, 2, 76, 105, 7, 34, 2, 2, 77, 78, 7, 36, 2, 2, 78, 79, 7, 10, 2,
	2, 79, 105, 7, 36, 2, 2, 80, 81, 7, 36, 2, 2, 81, 82, 7, 10, 2, 2, 82,
	105, 7, 37, 2, 2, 83, 84, 7, 36, 2, 2, 84, 85, 7, 10, 2, 2, 85, 105, 7,
	38, 2, 2, 86, 87, 7, 37, 2, 2, 87, 88, 7, 10, 2, 2, 88, 105, 7, 36, 2,
	2, 89, 90, 7, 37, 2, 2, 90, 91, 7, 10, 2, 2, 91, 105, 7, 37, 2, 2, 92,
	93, 7, 37, 2, 2, 93, 94, 7, 10, 2, 2, 94, 105, 7, 38, 2, 2, 95, 96, 7,
	38, 2, 2, 96, 97, 7, 10, 2, 2, 97, 105, 7, 36, 2, 2, 98, 99, 7, 38, 2,
	2, 99, 100, 7, 10, 2, 2, 100, 105, 7, 37, 2, 2, 101, 102, 7, 38, 2, 2,
	102, 103, 7, 10, 2, 2, 103, 105, 7, 38, 2, 2, 104, 74, 3, 2, 2, 2, 104,
	77, 3, 2, 2, 2, 104, 80, 3, 2, 2, 2, 104, 83, 3, 2, 2, 2, 104, 86, 3, 2,
	2, 2, 104, 89, 3, 2, 2, 2, 104, 92, 3, 2, 2, 2, 104, 95, 3, 2, 2, 2, 104,
	98, 3, 2, 2, 2, 104, 101, 3, 2, 2, 2, 105, 11, 3, 2, 2, 2, 106, 110, 7,
	36, 2, 2, 107, 110, 7, 37, 2, 2, 108, 110, 7, 38, 2, 2, 109, 106, 3, 2,
	2, 2, 109, 107, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2,
	111, 115, 7, 3, 2, 2, 112, 116, 7, 36, 2, 2, 113, 116, 7, 37, 2, 2, 114,
	116, 7, 38, 2, 2, 115, 112, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 114,
	3, 2, 2, 2, 116, 13, 3, 2, 2, 2, 117, 118, 7, 11, 2, 2, 118, 119, 7, 34,
	2, 2, 119, 15, 3, 2, 2, 2, 120, 121, 7, 12, 2, 2, 121, 122, 7, 34, 2, 2,
	122, 17, 3, 2, 2, 2, 123, 124, 7, 34, 2, 2, 124, 125, 7, 4, 2, 2, 125,
	130, 7, 34, 2, 2, 126, 131, 7, 13, 2, 2, 127, 131, 7, 14, 2, 2, 128, 131,
	7, 16, 2, 2, 129, 131, 7, 15, 2, 2, 130, 126, 3, 2, 2, 2, 130, 127, 3,
	2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2,
	2, 132, 133, 7, 34, 2, 2, 133, 19, 3, 2, 2, 2, 134, 135, 7, 34, 2, 2, 135,
	136, 7, 4, 2, 2, 136, 137, 7, 23, 2, 2, 137, 138, 7, 34, 2, 2, 138, 21,
	3, 2, 2, 2, 139, 140, 7, 34, 2, 2, 140, 141, 7, 4, 2, 2, 141, 145, 7, 34,
	2, 2, 142, 146, 7, 24, 2, 2, 143, 146, 7, 26, 2, 2, 144, 146, 7, 25, 2,
	2, 145, 142, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 148, 7, 34, 2, 2, 148, 23, 3, 2, 2, 2, 149, 150,
	7, 30, 2, 2, 150, 157, 7, 34, 2, 2, 151, 158, 7, 17, 2, 2, 152, 158, 7,
	22, 2, 2, 153, 158, 7, 18, 2, 2, 154, 158, 7, 19, 2, 2, 155, 158, 7, 20,
	2, 2, 156, 158, 7, 21, 2, 2, 157, 151, 3, 2, 2, 2, 157, 152, 3, 2, 2, 2,
	157, 153, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157,
	156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 7, 34, 2, 2, 160, 149,
	3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 166, 7, 29,
	2, 2, 163, 167, 7, 36, 2, 2, 164, 167, 7, 37, 2, 2, 165, 167, 7, 38, 2,
	2, 166, 163, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167,
	25, 3, 2, 2, 2, 168, 169, 7, 34, 2, 2, 169, 170, 7, 4, 2, 2, 170, 171,
	7, 35, 2, 2, 171, 27, 3, 2, 2, 2, 172, 173, 7, 34, 2, 2, 173, 174, 7, 10,
	2, 2, 174, 175, 7, 35, 2, 2, 175, 29, 3, 2, 2, 2, 176, 177, 7, 27, 2, 2,
	177, 31, 3, 2, 2, 2, 178, 179, 7, 28, 2, 2, 179, 33, 3, 2, 2, 2, 15, 37,
	54, 62, 66, 72, 104, 109, 115, 130, 145, 157, 160, 166,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<->'", "'<-'", "", "", "", "", "'load'", "'->'", "'increment'", "'decrement'",
	"'+'", "'-'", "'/'", "'*'", "'='", "'>'", "'>='", "'<'", "'<='", "'!='",
	"'NOT'", "'AND'", "'XOR'", "'OR'", "'STOP'", "'WAIT'", "'JUMP'", "'IF'",
}
var symbolicNames = []string{
	"", "", "", "WHITESPACE", "WS", "COMMENT", "LINE_COMMENT", "LOAD", "STORE",
	"INCREMENT", "DECREMENT", "ADD", "SUBTRACT", "DIVIDE", "MULTIPLY", "EQUAL",
	"GREATER", "GREATER_OR_EQUAL", "LESS", "LESS_OR_EQUAL", "NOT_EQUAL", "NOT",
	"AND", "XOR", "OR", "STOP", "WAIT", "JUMP", "IF", "HEX_DIGIT", "HEXA_BYTE",
	"WORD", "REGISTER", "DEVICE", "ABSOLUTE_ADDRESS", "OFFSET_ADDRESS", "POINTER_ADDRESS",
}

var ruleNames = []string{
	"program", "instruction", "load", "store", "copyValue", "copyRightToLeft",
	"increment", "decrement", "arithmeticOperation", "unaryBooleanOperation",
	"binaryBooleanOperation", "jump", "input", "output", "stop", "wait",
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
	GrogParserT__1             = 2
	GrogParserWHITESPACE       = 3
	GrogParserWS               = 4
	GrogParserCOMMENT          = 5
	GrogParserLINE_COMMENT     = 6
	GrogParserLOAD             = 7
	GrogParserSTORE            = 8
	GrogParserINCREMENT        = 9
	GrogParserDECREMENT        = 10
	GrogParserADD              = 11
	GrogParserSUBTRACT         = 12
	GrogParserDIVIDE           = 13
	GrogParserMULTIPLY         = 14
	GrogParserEQUAL            = 15
	GrogParserGREATER          = 16
	GrogParserGREATER_OR_EQUAL = 17
	GrogParserLESS             = 18
	GrogParserLESS_OR_EQUAL    = 19
	GrogParserNOT_EQUAL        = 20
	GrogParserNOT              = 21
	GrogParserAND              = 22
	GrogParserXOR              = 23
	GrogParserOR               = 24
	GrogParserSTOP             = 25
	GrogParserWAIT             = 26
	GrogParserJUMP             = 27
	GrogParserIF               = 28
	GrogParserHEX_DIGIT        = 29
	GrogParserHEXA_BYTE        = 30
	GrogParserWORD             = 31
	GrogParserREGISTER         = 32
	GrogParserDEVICE           = 33
	GrogParserABSOLUTE_ADDRESS = 34
	GrogParserOFFSET_ADDRESS   = 35
	GrogParserPOINTER_ADDRESS  = 36
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
	GrogParserRULE_input                  = 12
	GrogParserRULE_output                 = 13
	GrogParserRULE_stop                   = 14
	GrogParserRULE_wait                   = 15
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-7)&-(0x1f+1)) == 0 && ((1<<uint((_la-7)))&((1<<(GrogParserLOAD-7))|(1<<(GrogParserINCREMENT-7))|(1<<(GrogParserDECREMENT-7))|(1<<(GrogParserSTOP-7))|(1<<(GrogParserWAIT-7))|(1<<(GrogParserJUMP-7))|(1<<(GrogParserIF-7))|(1<<(GrogParserHEXA_BYTE-7))|(1<<(GrogParserREGISTER-7))|(1<<(GrogParserABSOLUTE_ADDRESS-7))|(1<<(GrogParserOFFSET_ADDRESS-7))|(1<<(GrogParserPOINTER_ADDRESS-7)))) != 0) {
		{
			p.SetState(32)
			p.Instruction()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
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

func (s *InstructionContext) Input() IInputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *InstructionContext) Output() IOutputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutputContext)
}

func (s *InstructionContext) Stop() IStopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStopContext)
}

func (s *InstructionContext) Wait() IWaitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWaitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWaitContext)
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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Increment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Decrement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.ArithmeticOperation()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(42)
			p.UnaryBooleanOperation()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.BinaryBooleanOperation()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.CopyValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(45)
			p.Load()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(46)
			p.Store()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(47)
			p.Jump()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(48)
			p.Input()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(49)
			p.Output()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(50)
			p.Stop()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(51)
			p.Wait()
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
		p.SetState(54)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(55)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*LoadContext).Register = _m
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserHEXA_BYTE:
		{
			p.SetState(56)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*LoadContext).Value = _m
		}

	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(57)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*LoadContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(58)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*LoadContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(59)

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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserREGISTER:
		{
			p.SetState(62)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*StoreContext).Register = _m
		}

	case GrogParserHEXA_BYTE:
		{
			p.SetState(63)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*StoreContext).Value = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(66)
		p.Match(GrogParserSTORE)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(67)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*StoreContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(68)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*StoreContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(69)

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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCopyRegisterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*CopyRegisterContext).SourceRegister = _m
		}
		{
			p.SetState(73)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(74)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*CopyRegisterContext).DestinationRegister = _m
		}

	case 2:
		localctx = NewCopyAbsoluteToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToAbsoluteContext).SourceAddress = _m
		}
		{
			p.SetState(76)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(77)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToAbsoluteContext).DestinationAddress = _m
		}

	case 3:
		localctx = NewCopyAbsoluteToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToOffsetContext).SourceAddress = _m
		}
		{
			p.SetState(79)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(80)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyAbsoluteToOffsetContext).DestinationOffset = _m
		}

	case 4:
		localctx = NewCopyAbsoluteToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyAbsoluteToPointerContext).SourceAddress = _m
		}
		{
			p.SetState(82)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(83)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyAbsoluteToPointerContext).DestinationPointer = _m
		}

	case 5:
		localctx = NewCopyOffsetToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToAbsoluteContext).SourceOffset = _m
		}
		{
			p.SetState(85)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(86)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyOffsetToAbsoluteContext).DestinationAddress = _m
		}

	case 6:
		localctx = NewCopyOffsetToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(87)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToOffsetContext).SourceOffset = _m
		}
		{
			p.SetState(88)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(89)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToOffsetContext).DestinationOffset = _m
		}

	case 7:
		localctx = NewCopyOffsetToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(90)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyOffsetToPointerContext).SourceOffset = _m
		}
		{
			p.SetState(91)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(92)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyOffsetToPointerContext).DestinationPointer = _m
		}

	case 8:
		localctx = NewCopyPointerToAbsoluteContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(93)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToAbsoluteContext).SourcePointer = _m
		}
		{
			p.SetState(94)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(95)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyPointerToAbsoluteContext).DestinationAddress = _m
		}

	case 9:
		localctx = NewCopyPointerToOffsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(96)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToOffsetContext).SourcePointer = _m
		}
		{
			p.SetState(97)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(98)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyPointerToOffsetContext).DestinationOffset = _m
		}

	case 10:
		localctx = NewCopyPointerToPointerContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(99)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyPointerToPointerContext).SourcePointer = _m
		}
		{
			p.SetState(100)
			p.Match(GrogParserSTORE)
		}
		{
			p.SetState(101)

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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(104)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceAddress = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(105)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceOffset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(106)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourcePointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(109)
		p.Match(GrogParserT__0)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(110)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceAddress = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(111)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*CopyRightToLeftContext).SourceOffset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(112)

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
		p.SetState(115)
		p.Match(GrogParserINCREMENT)
	}
	{
		p.SetState(116)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*IncrementContext).Register = _m
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
		p.SetState(118)
		p.Match(GrogParserDECREMENT)
	}
	{
		p.SetState(119)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*DecrementContext).Register = _m
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
		p.SetState(121)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Destination = _m
	}
	{
		p.SetState(122)
		p.Match(GrogParserT__1)
	}
	{
		p.SetState(123)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Left = _m
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserADD:
		{
			p.SetState(124)

			var _m = p.Match(GrogParserADD)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserSUBTRACT:
		{
			p.SetState(125)

			var _m = p.Match(GrogParserSUBTRACT)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserMULTIPLY:
		{
			p.SetState(126)

			var _m = p.Match(GrogParserMULTIPLY)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserDIVIDE:
		{
			p.SetState(127)

			var _m = p.Match(GrogParserDIVIDE)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(130)

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
		p.SetState(132)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*UnaryBooleanOperationContext).Destination = _m
	}
	{
		p.SetState(133)
		p.Match(GrogParserT__1)
	}
	{
		p.SetState(134)
		p.Match(GrogParserNOT)
	}
	{
		p.SetState(135)

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
		p.SetState(137)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Destination = _m
	}
	{
		p.SetState(138)
		p.Match(GrogParserT__1)
	}
	{
		p.SetState(139)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Left = _m
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserAND:
		{
			p.SetState(140)

			var _m = p.Match(GrogParserAND)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserOR:
		{
			p.SetState(141)

			var _m = p.Match(GrogParserOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserXOR:
		{
			p.SetState(142)

			var _m = p.Match(GrogParserXOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(145)

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
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrogParserIF {
		{
			p.SetState(147)
			p.Match(GrogParserIF)
		}
		{
			p.SetState(148)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*JumpContext).Left = _m
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GrogParserEQUAL:
			{
				p.SetState(149)

				var _m = p.Match(GrogParserEQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserNOT_EQUAL:
			{
				p.SetState(150)

				var _m = p.Match(GrogParserNOT_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserGREATER:
			{
				p.SetState(151)

				var _m = p.Match(GrogParserGREATER)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserGREATER_OR_EQUAL:
			{
				p.SetState(152)

				var _m = p.Match(GrogParserGREATER_OR_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserLESS:
			{
				p.SetState(153)

				var _m = p.Match(GrogParserLESS)

				localctx.(*JumpContext).Operator = _m
			}

		case GrogParserLESS_OR_EQUAL:
			{
				p.SetState(154)

				var _m = p.Match(GrogParserLESS_OR_EQUAL)

				localctx.(*JumpContext).Operator = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(157)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*JumpContext).Right = _m
		}

	}
	{
		p.SetState(160)
		p.Match(GrogParserJUMP)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(161)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*JumpContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(162)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*JumpContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(163)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*JumpContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) GetDestination() antlr.Token { return s.Destination }

func (s *InputContext) GetSource() antlr.Token { return s.Source }

func (s *InputContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *InputContext) SetSource(v antlr.Token) { s.Source = v }

func (s *InputContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *InputContext) DEVICE() antlr.TerminalNode {
	return s.GetToken(GrogParserDEVICE, 0)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitInput(s)
	}
}

func (p *GrogParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrogParserRULE_input)

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
		p.SetState(166)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*InputContext).Destination = _m
	}
	{
		p.SetState(167)
		p.Match(GrogParserT__1)
	}
	{
		p.SetState(168)

		var _m = p.Match(GrogParserDEVICE)

		localctx.(*InputContext).Source = _m
	}

	return localctx
}

// IOutputContext is an interface to support dynamic dispatch.
type IOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// IsOutputContext differentiates from other interfaces.
	IsOutputContext()
}

type OutputContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Source      antlr.Token
	Destination antlr.Token
}

func NewEmptyOutputContext() *OutputContext {
	var p = new(OutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_output
	return p
}

func (*OutputContext) IsOutputContext() {}

func NewOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputContext {
	var p = new(OutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_output

	return p
}

func (s *OutputContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputContext) GetSource() antlr.Token { return s.Source }

func (s *OutputContext) GetDestination() antlr.Token { return s.Destination }

func (s *OutputContext) SetSource(v antlr.Token) { s.Source = v }

func (s *OutputContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *OutputContext) STORE() antlr.TerminalNode {
	return s.GetToken(GrogParserSTORE, 0)
}

func (s *OutputContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *OutputContext) DEVICE() antlr.TerminalNode {
	return s.GetToken(GrogParserDEVICE, 0)
}

func (s *OutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterOutput(s)
	}
}

func (s *OutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitOutput(s)
	}
}

func (p *GrogParser) Output() (localctx IOutputContext) {
	localctx = NewOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrogParserRULE_output)

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
		p.SetState(170)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*OutputContext).Source = _m
	}
	{
		p.SetState(171)
		p.Match(GrogParserSTORE)
	}
	{
		p.SetState(172)

		var _m = p.Match(GrogParserDEVICE)

		localctx.(*OutputContext).Destination = _m
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
	p.EnterRule(localctx, 28, GrogParserRULE_stop)

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
		p.SetState(174)
		p.Match(GrogParserSTOP)
	}

	return localctx
}

// IWaitContext is an interface to support dynamic dispatch.
type IWaitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWaitContext differentiates from other interfaces.
	IsWaitContext()
}

type WaitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWaitContext() *WaitContext {
	var p = new(WaitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_wait
	return p
}

func (*WaitContext) IsWaitContext() {}

func NewWaitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaitContext {
	var p = new(WaitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_wait

	return p
}

func (s *WaitContext) GetParser() antlr.Parser { return s.parser }

func (s *WaitContext) WAIT() antlr.TerminalNode {
	return s.GetToken(GrogParserWAIT, 0)
}

func (s *WaitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WaitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterWait(s)
	}
}

func (s *WaitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitWait(s)
	}
}

func (p *GrogParser) Wait() (localctx IWaitContext) {
	localctx = NewWaitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrogParserRULE_wait)

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
		p.SetState(176)
		p.Match(GrogParserWAIT)
	}

	return localctx
}
