// Generated from /home/martinstraus/MEGA/proyectos/propios/grog/asm/Grog.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GrogParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, WHITESPACE=2, WS=3, COMMENT=4, LINE_COMMENT=5, LOAD=6, STORE=7, 
		INCREMENT=8, DECREMENT=9, ADD=10, SUBTRACT=11, DIVIDE=12, MULTIPLY=13, 
		EQUAL=14, GREATER=15, GREATER_OR_EQUAL=16, LESS=17, LESS_OR_EQUAL=18, 
		NOT_EQUAL=19, NOT=20, AND=21, XOR=22, OR=23, STOP=24, JUMP=25, IF=26, 
		HEX_DIGIT=27, HEXA_BYTE=28, WORD=29, REGISTER=30, ABSOLUTE_ADDRESS=31, 
		OFFSET_ADDRESS=32, POINTER_ADDRESS=33;
	public static final int
		RULE_program = 0, RULE_instruction = 1, RULE_load = 2, RULE_store = 3, 
		RULE_copyValue = 4, RULE_copyRightToLeft = 5, RULE_increment = 6, RULE_decrement = 7, 
		RULE_arithmeticOperation = 8, RULE_unaryBooleanOperation = 9, RULE_binaryBooleanOperation = 10, 
		RULE_jump = 11, RULE_stop = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "instruction", "load", "store", "copyValue", "copyRightToLeft", 
			"increment", "decrement", "arithmeticOperation", "unaryBooleanOperation", 
			"binaryBooleanOperation", "jump", "stop"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<->'", null, null, null, null, "'<-'", "'->'", "'++'", "'--'", 
			"'+'", "'-'", "'/'", "'*'", "'='", "'>'", "'>='", "'<'", "'<='", "'!='", 
			"'NOT'", "'AND'", "'XOR'", "'OR'", "'STOP'", "'JUMP'", "'IF'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "WHITESPACE", "WS", "COMMENT", "LINE_COMMENT", "LOAD", "STORE", 
			"INCREMENT", "DECREMENT", "ADD", "SUBTRACT", "DIVIDE", "MULTIPLY", "EQUAL", 
			"GREATER", "GREATER_OR_EQUAL", "LESS", "LESS_OR_EQUAL", "NOT_EQUAL", 
			"NOT", "AND", "XOR", "OR", "STOP", "JUMP", "IF", "HEX_DIGIT", "HEXA_BYTE", 
			"WORD", "REGISTER", "ABSOLUTE_ADDRESS", "OFFSET_ADDRESS", "POINTER_ADDRESS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Grog.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrogParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(GrogParser.EOF, 0); }
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(27); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(26);
				instruction();
				}
				}
				setState(29); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << STOP) | (1L << JUMP) | (1L << IF) | (1L << HEXA_BYTE) | (1L << REGISTER) | (1L << ABSOLUTE_ADDRESS) | (1L << OFFSET_ADDRESS) | (1L << POINTER_ADDRESS))) != 0) );
			setState(31);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstructionContext extends ParserRuleContext {
		public IncrementContext increment() {
			return getRuleContext(IncrementContext.class,0);
		}
		public DecrementContext decrement() {
			return getRuleContext(DecrementContext.class,0);
		}
		public ArithmeticOperationContext arithmeticOperation() {
			return getRuleContext(ArithmeticOperationContext.class,0);
		}
		public UnaryBooleanOperationContext unaryBooleanOperation() {
			return getRuleContext(UnaryBooleanOperationContext.class,0);
		}
		public BinaryBooleanOperationContext binaryBooleanOperation() {
			return getRuleContext(BinaryBooleanOperationContext.class,0);
		}
		public CopyValueContext copyValue() {
			return getRuleContext(CopyValueContext.class,0);
		}
		public LoadContext load() {
			return getRuleContext(LoadContext.class,0);
		}
		public StoreContext store() {
			return getRuleContext(StoreContext.class,0);
		}
		public JumpContext jump() {
			return getRuleContext(JumpContext.class,0);
		}
		public StopContext stop() {
			return getRuleContext(StopContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instruction);
		try {
			setState(43);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				increment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				decrement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				arithmeticOperation();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(36);
				unaryBooleanOperation();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(37);
				binaryBooleanOperation();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(38);
				copyValue();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(39);
				load();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(40);
				store();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(41);
				jump();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(42);
				stop();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LoadContext extends ParserRuleContext {
		public Token Register;
		public Token Value;
		public Token Address;
		public Token Offset;
		public Token Pointer;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public TerminalNode REGISTER() { return getToken(GrogParser.REGISTER, 0); }
		public TerminalNode HEXA_BYTE() { return getToken(GrogParser.HEXA_BYTE, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public LoadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_load; }
	}

	public final LoadContext load() throws RecognitionException {
		LoadContext _localctx = new LoadContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_load);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			((LoadContext)_localctx).Register = match(REGISTER);
			setState(46);
			match(LOAD);
			setState(51);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case HEXA_BYTE:
				{
				setState(47);
				((LoadContext)_localctx).Value = match(HEXA_BYTE);
				}
				break;
			case ABSOLUTE_ADDRESS:
				{
				setState(48);
				((LoadContext)_localctx).Address = match(ABSOLUTE_ADDRESS);
				}
				break;
			case OFFSET_ADDRESS:
				{
				setState(49);
				((LoadContext)_localctx).Offset = match(OFFSET_ADDRESS);
				}
				break;
			case POINTER_ADDRESS:
				{
				setState(50);
				((LoadContext)_localctx).Pointer = match(POINTER_ADDRESS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StoreContext extends ParserRuleContext {
		public Token Register;
		public Token Value;
		public Token Address;
		public Token Offset;
		public Token Pointer;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode REGISTER() { return getToken(GrogParser.REGISTER, 0); }
		public TerminalNode HEXA_BYTE() { return getToken(GrogParser.HEXA_BYTE, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public StoreContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_store; }
	}

	public final StoreContext store() throws RecognitionException {
		StoreContext _localctx = new StoreContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_store);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case REGISTER:
				{
				setState(53);
				((StoreContext)_localctx).Register = match(REGISTER);
				}
				break;
			case HEXA_BYTE:
				{
				setState(54);
				((StoreContext)_localctx).Value = match(HEXA_BYTE);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(57);
			match(STORE);
			setState(61);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ABSOLUTE_ADDRESS:
				{
				setState(58);
				((StoreContext)_localctx).Address = match(ABSOLUTE_ADDRESS);
				}
				break;
			case OFFSET_ADDRESS:
				{
				setState(59);
				((StoreContext)_localctx).Offset = match(OFFSET_ADDRESS);
				}
				break;
			case POINTER_ADDRESS:
				{
				setState(60);
				((StoreContext)_localctx).Pointer = match(POINTER_ADDRESS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CopyValueContext extends ParserRuleContext {
		public CopyValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_copyValue; }
	 
		public CopyValueContext() { }
		public void copyFrom(CopyValueContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class CopyRegisterContext extends CopyValueContext {
		public Token SourceRegister;
		public Token DestinationRegister;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public List<TerminalNode> REGISTER() { return getTokens(GrogParser.REGISTER); }
		public TerminalNode REGISTER(int i) {
			return getToken(GrogParser.REGISTER, i);
		}
		public CopyRegisterContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyAbsoluteToPointerContext extends CopyValueContext {
		public Token SourceAddress;
		public Token DestinationPointer;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public CopyAbsoluteToPointerContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyOffsetToAbsoluteContext extends CopyValueContext {
		public Token SourceOffset;
		public Token DestinationAddress;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public CopyOffsetToAbsoluteContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyOffsetToOffsetContext extends CopyValueContext {
		public Token SourceOffset;
		public Token DestinationOffset;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public List<TerminalNode> OFFSET_ADDRESS() { return getTokens(GrogParser.OFFSET_ADDRESS); }
		public TerminalNode OFFSET_ADDRESS(int i) {
			return getToken(GrogParser.OFFSET_ADDRESS, i);
		}
		public CopyOffsetToOffsetContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyPointerToOffsetContext extends CopyValueContext {
		public Token SourcePointer;
		public Token DestinationOffset;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public CopyPointerToOffsetContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyAbsoluteToOffsetContext extends CopyValueContext {
		public Token SourceAddress;
		public Token DestinationOffset;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public CopyAbsoluteToOffsetContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyOffsetToPointerContext extends CopyValueContext {
		public Token SourceOffset;
		public Token DestinationPointer;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public CopyOffsetToPointerContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyPointerToAbsoluteContext extends CopyValueContext {
		public Token SourcePointer;
		public Token DestinationAddress;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public CopyPointerToAbsoluteContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyPointerToPointerContext extends CopyValueContext {
		public Token SourcePointer;
		public Token DestinationPointer;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public List<TerminalNode> POINTER_ADDRESS() { return getTokens(GrogParser.POINTER_ADDRESS); }
		public TerminalNode POINTER_ADDRESS(int i) {
			return getToken(GrogParser.POINTER_ADDRESS, i);
		}
		public CopyPointerToPointerContext(CopyValueContext ctx) { copyFrom(ctx); }
	}
	public static class CopyAbsoluteToAbsoluteContext extends CopyValueContext {
		public Token SourceAddress;
		public Token DestinationAddress;
		public TerminalNode STORE() { return getToken(GrogParser.STORE, 0); }
		public List<TerminalNode> ABSOLUTE_ADDRESS() { return getTokens(GrogParser.ABSOLUTE_ADDRESS); }
		public TerminalNode ABSOLUTE_ADDRESS(int i) {
			return getToken(GrogParser.ABSOLUTE_ADDRESS, i);
		}
		public CopyAbsoluteToAbsoluteContext(CopyValueContext ctx) { copyFrom(ctx); }
	}

	public final CopyValueContext copyValue() throws RecognitionException {
		CopyValueContext _localctx = new CopyValueContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_copyValue);
		try {
			setState(93);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new CopyRegisterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(63);
				((CopyRegisterContext)_localctx).SourceRegister = match(REGISTER);
				setState(64);
				match(STORE);
				setState(65);
				((CopyRegisterContext)_localctx).DestinationRegister = match(REGISTER);
				}
				}
				break;
			case 2:
				_localctx = new CopyAbsoluteToAbsoluteContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(66);
				((CopyAbsoluteToAbsoluteContext)_localctx).SourceAddress = match(ABSOLUTE_ADDRESS);
				setState(67);
				match(STORE);
				setState(68);
				((CopyAbsoluteToAbsoluteContext)_localctx).DestinationAddress = match(ABSOLUTE_ADDRESS);
				}
				}
				break;
			case 3:
				_localctx = new CopyAbsoluteToOffsetContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(69);
				((CopyAbsoluteToOffsetContext)_localctx).SourceAddress = match(ABSOLUTE_ADDRESS);
				setState(70);
				match(STORE);
				setState(71);
				((CopyAbsoluteToOffsetContext)_localctx).DestinationOffset = match(OFFSET_ADDRESS);
				}
				}
				break;
			case 4:
				_localctx = new CopyAbsoluteToPointerContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				{
				setState(72);
				((CopyAbsoluteToPointerContext)_localctx).SourceAddress = match(ABSOLUTE_ADDRESS);
				setState(73);
				match(STORE);
				setState(74);
				((CopyAbsoluteToPointerContext)_localctx).DestinationPointer = match(POINTER_ADDRESS);
				}
				}
				break;
			case 5:
				_localctx = new CopyOffsetToAbsoluteContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				{
				setState(75);
				((CopyOffsetToAbsoluteContext)_localctx).SourceOffset = match(OFFSET_ADDRESS);
				setState(76);
				match(STORE);
				setState(77);
				((CopyOffsetToAbsoluteContext)_localctx).DestinationAddress = match(ABSOLUTE_ADDRESS);
				}
				}
				break;
			case 6:
				_localctx = new CopyOffsetToOffsetContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				{
				setState(78);
				((CopyOffsetToOffsetContext)_localctx).SourceOffset = match(OFFSET_ADDRESS);
				setState(79);
				match(STORE);
				setState(80);
				((CopyOffsetToOffsetContext)_localctx).DestinationOffset = match(OFFSET_ADDRESS);
				}
				}
				break;
			case 7:
				_localctx = new CopyOffsetToPointerContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				{
				setState(81);
				((CopyOffsetToPointerContext)_localctx).SourceOffset = match(OFFSET_ADDRESS);
				setState(82);
				match(STORE);
				setState(83);
				((CopyOffsetToPointerContext)_localctx).DestinationPointer = match(POINTER_ADDRESS);
				}
				}
				break;
			case 8:
				_localctx = new CopyPointerToAbsoluteContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				{
				setState(84);
				((CopyPointerToAbsoluteContext)_localctx).SourcePointer = match(POINTER_ADDRESS);
				setState(85);
				match(STORE);
				setState(86);
				((CopyPointerToAbsoluteContext)_localctx).DestinationAddress = match(ABSOLUTE_ADDRESS);
				}
				}
				break;
			case 9:
				_localctx = new CopyPointerToOffsetContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				{
				setState(87);
				((CopyPointerToOffsetContext)_localctx).SourcePointer = match(POINTER_ADDRESS);
				setState(88);
				match(STORE);
				setState(89);
				((CopyPointerToOffsetContext)_localctx).DestinationOffset = match(OFFSET_ADDRESS);
				}
				}
				break;
			case 10:
				_localctx = new CopyPointerToPointerContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				{
				setState(90);
				((CopyPointerToPointerContext)_localctx).SourcePointer = match(POINTER_ADDRESS);
				setState(91);
				match(STORE);
				setState(92);
				((CopyPointerToPointerContext)_localctx).DestinationPointer = match(POINTER_ADDRESS);
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CopyRightToLeftContext extends ParserRuleContext {
		public Token SourceAddress;
		public Token SourceOffset;
		public Token SourcePointer;
		public List<TerminalNode> ABSOLUTE_ADDRESS() { return getTokens(GrogParser.ABSOLUTE_ADDRESS); }
		public TerminalNode ABSOLUTE_ADDRESS(int i) {
			return getToken(GrogParser.ABSOLUTE_ADDRESS, i);
		}
		public List<TerminalNode> OFFSET_ADDRESS() { return getTokens(GrogParser.OFFSET_ADDRESS); }
		public TerminalNode OFFSET_ADDRESS(int i) {
			return getToken(GrogParser.OFFSET_ADDRESS, i);
		}
		public List<TerminalNode> POINTER_ADDRESS() { return getTokens(GrogParser.POINTER_ADDRESS); }
		public TerminalNode POINTER_ADDRESS(int i) {
			return getToken(GrogParser.POINTER_ADDRESS, i);
		}
		public CopyRightToLeftContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_copyRightToLeft; }
	}

	public final CopyRightToLeftContext copyRightToLeft() throws RecognitionException {
		CopyRightToLeftContext _localctx = new CopyRightToLeftContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_copyRightToLeft);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ABSOLUTE_ADDRESS:
				{
				setState(95);
				((CopyRightToLeftContext)_localctx).SourceAddress = match(ABSOLUTE_ADDRESS);
				}
				break;
			case OFFSET_ADDRESS:
				{
				setState(96);
				((CopyRightToLeftContext)_localctx).SourceOffset = match(OFFSET_ADDRESS);
				}
				break;
			case POINTER_ADDRESS:
				{
				setState(97);
				((CopyRightToLeftContext)_localctx).SourcePointer = match(POINTER_ADDRESS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(100);
			match(T__0);
			setState(104);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ABSOLUTE_ADDRESS:
				{
				setState(101);
				((CopyRightToLeftContext)_localctx).SourceAddress = match(ABSOLUTE_ADDRESS);
				}
				break;
			case OFFSET_ADDRESS:
				{
				setState(102);
				((CopyRightToLeftContext)_localctx).SourceOffset = match(OFFSET_ADDRESS);
				}
				break;
			case POINTER_ADDRESS:
				{
				setState(103);
				((CopyRightToLeftContext)_localctx).SourcePointer = match(POINTER_ADDRESS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IncrementContext extends ParserRuleContext {
		public Token Register;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public TerminalNode INCREMENT() { return getToken(GrogParser.INCREMENT, 0); }
		public TerminalNode REGISTER() { return getToken(GrogParser.REGISTER, 0); }
		public IncrementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_increment; }
	}

	public final IncrementContext increment() throws RecognitionException {
		IncrementContext _localctx = new IncrementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_increment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			((IncrementContext)_localctx).Register = match(REGISTER);
			setState(107);
			match(LOAD);
			setState(108);
			match(INCREMENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DecrementContext extends ParserRuleContext {
		public Token Register;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public TerminalNode DECREMENT() { return getToken(GrogParser.DECREMENT, 0); }
		public TerminalNode REGISTER() { return getToken(GrogParser.REGISTER, 0); }
		public DecrementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decrement; }
	}

	public final DecrementContext decrement() throws RecognitionException {
		DecrementContext _localctx = new DecrementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_decrement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			((DecrementContext)_localctx).Register = match(REGISTER);
			setState(111);
			match(LOAD);
			setState(112);
			match(DECREMENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArithmeticOperationContext extends ParserRuleContext {
		public Token Destination;
		public Token Left;
		public Token Operator;
		public Token Right;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public List<TerminalNode> REGISTER() { return getTokens(GrogParser.REGISTER); }
		public TerminalNode REGISTER(int i) {
			return getToken(GrogParser.REGISTER, i);
		}
		public TerminalNode ADD() { return getToken(GrogParser.ADD, 0); }
		public TerminalNode SUBTRACT() { return getToken(GrogParser.SUBTRACT, 0); }
		public TerminalNode MULTIPLY() { return getToken(GrogParser.MULTIPLY, 0); }
		public TerminalNode DIVIDE() { return getToken(GrogParser.DIVIDE, 0); }
		public ArithmeticOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arithmeticOperation; }
	}

	public final ArithmeticOperationContext arithmeticOperation() throws RecognitionException {
		ArithmeticOperationContext _localctx = new ArithmeticOperationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_arithmeticOperation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			((ArithmeticOperationContext)_localctx).Destination = match(REGISTER);
			setState(115);
			match(LOAD);
			setState(116);
			((ArithmeticOperationContext)_localctx).Left = match(REGISTER);
			setState(121);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ADD:
				{
				setState(117);
				((ArithmeticOperationContext)_localctx).Operator = match(ADD);
				}
				break;
			case SUBTRACT:
				{
				setState(118);
				((ArithmeticOperationContext)_localctx).Operator = match(SUBTRACT);
				}
				break;
			case MULTIPLY:
				{
				setState(119);
				((ArithmeticOperationContext)_localctx).Operator = match(MULTIPLY);
				}
				break;
			case DIVIDE:
				{
				setState(120);
				((ArithmeticOperationContext)_localctx).Operator = match(DIVIDE);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(123);
			((ArithmeticOperationContext)_localctx).Right = match(REGISTER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnaryBooleanOperationContext extends ParserRuleContext {
		public Token Destination;
		public Token Operand;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public TerminalNode NOT() { return getToken(GrogParser.NOT, 0); }
		public List<TerminalNode> REGISTER() { return getTokens(GrogParser.REGISTER); }
		public TerminalNode REGISTER(int i) {
			return getToken(GrogParser.REGISTER, i);
		}
		public UnaryBooleanOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryBooleanOperation; }
	}

	public final UnaryBooleanOperationContext unaryBooleanOperation() throws RecognitionException {
		UnaryBooleanOperationContext _localctx = new UnaryBooleanOperationContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_unaryBooleanOperation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			((UnaryBooleanOperationContext)_localctx).Destination = match(REGISTER);
			setState(126);
			match(LOAD);
			setState(127);
			match(NOT);
			setState(128);
			((UnaryBooleanOperationContext)_localctx).Operand = match(REGISTER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BinaryBooleanOperationContext extends ParserRuleContext {
		public Token Destination;
		public Token Left;
		public Token Operator;
		public Token Right;
		public TerminalNode LOAD() { return getToken(GrogParser.LOAD, 0); }
		public List<TerminalNode> REGISTER() { return getTokens(GrogParser.REGISTER); }
		public TerminalNode REGISTER(int i) {
			return getToken(GrogParser.REGISTER, i);
		}
		public TerminalNode AND() { return getToken(GrogParser.AND, 0); }
		public TerminalNode OR() { return getToken(GrogParser.OR, 0); }
		public TerminalNode XOR() { return getToken(GrogParser.XOR, 0); }
		public BinaryBooleanOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryBooleanOperation; }
	}

	public final BinaryBooleanOperationContext binaryBooleanOperation() throws RecognitionException {
		BinaryBooleanOperationContext _localctx = new BinaryBooleanOperationContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_binaryBooleanOperation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			((BinaryBooleanOperationContext)_localctx).Destination = match(REGISTER);
			setState(131);
			match(LOAD);
			setState(132);
			((BinaryBooleanOperationContext)_localctx).Left = match(REGISTER);
			setState(136);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case AND:
				{
				setState(133);
				((BinaryBooleanOperationContext)_localctx).Operator = match(AND);
				}
				break;
			case OR:
				{
				setState(134);
				((BinaryBooleanOperationContext)_localctx).Operator = match(OR);
				}
				break;
			case XOR:
				{
				setState(135);
				((BinaryBooleanOperationContext)_localctx).Operator = match(XOR);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(138);
			((BinaryBooleanOperationContext)_localctx).Right = match(REGISTER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JumpContext extends ParserRuleContext {
		public Token Left;
		public Token Operator;
		public Token Right;
		public Token Address;
		public Token Offset;
		public Token Pointer;
		public TerminalNode JUMP() { return getToken(GrogParser.JUMP, 0); }
		public TerminalNode IF() { return getToken(GrogParser.IF, 0); }
		public TerminalNode ABSOLUTE_ADDRESS() { return getToken(GrogParser.ABSOLUTE_ADDRESS, 0); }
		public TerminalNode OFFSET_ADDRESS() { return getToken(GrogParser.OFFSET_ADDRESS, 0); }
		public TerminalNode POINTER_ADDRESS() { return getToken(GrogParser.POINTER_ADDRESS, 0); }
		public List<TerminalNode> REGISTER() { return getTokens(GrogParser.REGISTER); }
		public TerminalNode REGISTER(int i) {
			return getToken(GrogParser.REGISTER, i);
		}
		public TerminalNode EQUAL() { return getToken(GrogParser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(GrogParser.NOT_EQUAL, 0); }
		public TerminalNode GREATER() { return getToken(GrogParser.GREATER, 0); }
		public TerminalNode GREATER_OR_EQUAL() { return getToken(GrogParser.GREATER_OR_EQUAL, 0); }
		public TerminalNode LESS() { return getToken(GrogParser.LESS, 0); }
		public TerminalNode LESS_OR_EQUAL() { return getToken(GrogParser.LESS_OR_EQUAL, 0); }
		public JumpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jump; }
	}

	public final JumpContext jump() throws RecognitionException {
		JumpContext _localctx = new JumpContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_jump);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IF) {
				{
				setState(140);
				match(IF);
				setState(141);
				((JumpContext)_localctx).Left = match(REGISTER);
				setState(148);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case EQUAL:
					{
					setState(142);
					((JumpContext)_localctx).Operator = match(EQUAL);
					}
					break;
				case NOT_EQUAL:
					{
					setState(143);
					((JumpContext)_localctx).Operator = match(NOT_EQUAL);
					}
					break;
				case GREATER:
					{
					setState(144);
					((JumpContext)_localctx).Operator = match(GREATER);
					}
					break;
				case GREATER_OR_EQUAL:
					{
					setState(145);
					((JumpContext)_localctx).Operator = match(GREATER_OR_EQUAL);
					}
					break;
				case LESS:
					{
					setState(146);
					((JumpContext)_localctx).Operator = match(LESS);
					}
					break;
				case LESS_OR_EQUAL:
					{
					setState(147);
					((JumpContext)_localctx).Operator = match(LESS_OR_EQUAL);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(150);
				((JumpContext)_localctx).Right = match(REGISTER);
				}
			}

			setState(153);
			match(JUMP);
			setState(157);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ABSOLUTE_ADDRESS:
				{
				setState(154);
				((JumpContext)_localctx).Address = match(ABSOLUTE_ADDRESS);
				}
				break;
			case OFFSET_ADDRESS:
				{
				setState(155);
				((JumpContext)_localctx).Offset = match(OFFSET_ADDRESS);
				}
				break;
			case POINTER_ADDRESS:
				{
				setState(156);
				((JumpContext)_localctx).Pointer = match(POINTER_ADDRESS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StopContext extends ParserRuleContext {
		public TerminalNode STOP() { return getToken(GrogParser.STOP, 0); }
		public StopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stop; }
	}

	public final StopContext stop() throws RecognitionException {
		StopContext _localctx = new StopContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_stop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(STOP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#\u00a4\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\6\2\36\n\2\r\2\16\2\37\3\2\3\2\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3.\n\3\3\4\3\4\3\4\3\4\3\4\3\4\5"+
		"\4\66\n\4\3\5\3\5\5\5:\n\5\3\5\3\5\3\5\3\5\5\5@\n\5\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6`\n\6\3\7\3\7\3\7\5\7e\n\7\3\7\3\7"+
		"\3\7\3\7\5\7k\n\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\5\n|\n\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\5\f\u008b\n\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u0097\n"+
		"\r\3\r\5\r\u009a\n\r\3\r\3\r\3\r\3\r\5\r\u00a0\n\r\3\16\3\16\3\16\2\2"+
		"\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\2\2\u00c0\2\35\3\2\2\2\4-\3\2\2"+
		"\2\6/\3\2\2\2\b9\3\2\2\2\n_\3\2\2\2\fd\3\2\2\2\16l\3\2\2\2\20p\3\2\2\2"+
		"\22t\3\2\2\2\24\177\3\2\2\2\26\u0084\3\2\2\2\30\u0099\3\2\2\2\32\u00a1"+
		"\3\2\2\2\34\36\5\4\3\2\35\34\3\2\2\2\36\37\3\2\2\2\37\35\3\2\2\2\37 \3"+
		"\2\2\2 !\3\2\2\2!\"\7\2\2\3\"\3\3\2\2\2#.\5\16\b\2$.\5\20\t\2%.\5\22\n"+
		"\2&.\5\24\13\2\'.\5\26\f\2(.\5\n\6\2).\5\6\4\2*.\5\b\5\2+.\5\30\r\2,."+
		"\5\32\16\2-#\3\2\2\2-$\3\2\2\2-%\3\2\2\2-&\3\2\2\2-\'\3\2\2\2-(\3\2\2"+
		"\2-)\3\2\2\2-*\3\2\2\2-+\3\2\2\2-,\3\2\2\2.\5\3\2\2\2/\60\7 \2\2\60\65"+
		"\7\b\2\2\61\66\7\36\2\2\62\66\7!\2\2\63\66\7\"\2\2\64\66\7#\2\2\65\61"+
		"\3\2\2\2\65\62\3\2\2\2\65\63\3\2\2\2\65\64\3\2\2\2\66\7\3\2\2\2\67:\7"+
		" \2\28:\7\36\2\29\67\3\2\2\298\3\2\2\2:;\3\2\2\2;?\7\t\2\2<@\7!\2\2=@"+
		"\7\"\2\2>@\7#\2\2?<\3\2\2\2?=\3\2\2\2?>\3\2\2\2@\t\3\2\2\2AB\7 \2\2BC"+
		"\7\t\2\2C`\7 \2\2DE\7!\2\2EF\7\t\2\2F`\7!\2\2GH\7!\2\2HI\7\t\2\2I`\7\""+
		"\2\2JK\7!\2\2KL\7\t\2\2L`\7#\2\2MN\7\"\2\2NO\7\t\2\2O`\7!\2\2PQ\7\"\2"+
		"\2QR\7\t\2\2R`\7\"\2\2ST\7\"\2\2TU\7\t\2\2U`\7#\2\2VW\7#\2\2WX\7\t\2\2"+
		"X`\7!\2\2YZ\7#\2\2Z[\7\t\2\2[`\7\"\2\2\\]\7#\2\2]^\7\t\2\2^`\7#\2\2_A"+
		"\3\2\2\2_D\3\2\2\2_G\3\2\2\2_J\3\2\2\2_M\3\2\2\2_P\3\2\2\2_S\3\2\2\2_"+
		"V\3\2\2\2_Y\3\2\2\2_\\\3\2\2\2`\13\3\2\2\2ae\7!\2\2be\7\"\2\2ce\7#\2\2"+
		"da\3\2\2\2db\3\2\2\2dc\3\2\2\2ef\3\2\2\2fj\7\3\2\2gk\7!\2\2hk\7\"\2\2"+
		"ik\7#\2\2jg\3\2\2\2jh\3\2\2\2ji\3\2\2\2k\r\3\2\2\2lm\7 \2\2mn\7\b\2\2"+
		"no\7\n\2\2o\17\3\2\2\2pq\7 \2\2qr\7\b\2\2rs\7\13\2\2s\21\3\2\2\2tu\7 "+
		"\2\2uv\7\b\2\2v{\7 \2\2w|\7\f\2\2x|\7\r\2\2y|\7\17\2\2z|\7\16\2\2{w\3"+
		"\2\2\2{x\3\2\2\2{y\3\2\2\2{z\3\2\2\2|}\3\2\2\2}~\7 \2\2~\23\3\2\2\2\177"+
		"\u0080\7 \2\2\u0080\u0081\7\b\2\2\u0081\u0082\7\26\2\2\u0082\u0083\7 "+
		"\2\2\u0083\25\3\2\2\2\u0084\u0085\7 \2\2\u0085\u0086\7\b\2\2\u0086\u008a"+
		"\7 \2\2\u0087\u008b\7\27\2\2\u0088\u008b\7\31\2\2\u0089\u008b\7\30\2\2"+
		"\u008a\u0087\3\2\2\2\u008a\u0088\3\2\2\2\u008a\u0089\3\2\2\2\u008b\u008c"+
		"\3\2\2\2\u008c\u008d\7 \2\2\u008d\27\3\2\2\2\u008e\u008f\7\34\2\2\u008f"+
		"\u0096\7 \2\2\u0090\u0097\7\20\2\2\u0091\u0097\7\25\2\2\u0092\u0097\7"+
		"\21\2\2\u0093\u0097\7\22\2\2\u0094\u0097\7\23\2\2\u0095\u0097\7\24\2\2"+
		"\u0096\u0090\3\2\2\2\u0096\u0091\3\2\2\2\u0096\u0092\3\2\2\2\u0096\u0093"+
		"\3\2\2\2\u0096\u0094\3\2\2\2\u0096\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u009a\7 \2\2\u0099\u008e\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u009b\3\2"+
		"\2\2\u009b\u009f\7\33\2\2\u009c\u00a0\7!\2\2\u009d\u00a0\7\"\2\2\u009e"+
		"\u00a0\7#\2\2\u009f\u009c\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u009e\3\2"+
		"\2\2\u00a0\31\3\2\2\2\u00a1\u00a2\7\32\2\2\u00a2\33\3\2\2\2\17\37-\65"+
		"9?_dj{\u008a\u0096\u0099\u009f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}