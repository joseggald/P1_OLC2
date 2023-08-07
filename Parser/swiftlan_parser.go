// Code generated from .\SwiftLan.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftLan

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftLanParser struct {
	*antlr.BaseParser
}

var SwiftLanParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlanParserInit() {
	staticData := &SwiftLanParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'('", "')'", "'?'", "':'", "'['", "']'", "'var'", "'print'",
		"'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "", "", "",
		"", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'^'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'>'", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Declarar", "Print", "IntDecla", "FloatDecla",
		"BoolDecla", "StringDecla", "CharDecla", "BoolVal", "Number", "Id",
		"String", "Or", "And", "Equals", "NotEquals", "MayQEquals", "MinQEquals",
		"Power", "Add", "Minus", "Mult", "Div", "Mod", "MayorQue", "MenorQue",
		"Space", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"inicio", "sentencias", "statement", "asignacion", "llamadaPrint", "expression",
		"indexes",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 96, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 5, 1, 18, 8, 1, 10, 1, 12,
		1, 21, 9, 1, 1, 2, 1, 2, 3, 2, 25, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 3, 4, 35, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 47, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		54, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 83, 8, 5, 10, 5, 12, 5, 86, 9, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 4, 6, 92, 8, 6, 11, 6, 12, 6, 93, 1, 6, 0, 1, 10, 7,
		0, 2, 4, 6, 8, 10, 12, 0, 4, 1, 0, 28, 30, 1, 0, 26, 27, 2, 0, 23, 24,
		31, 32, 1, 0, 21, 22, 107, 0, 14, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 24,
		1, 0, 0, 0, 6, 26, 1, 0, 0, 0, 8, 31, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12,
		91, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 1, 1, 0, 0, 0, 16, 18, 3, 4, 2,
		0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20,
		1, 0, 0, 0, 20, 3, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 25, 3, 6, 3, 0,
		23, 25, 3, 8, 4, 0, 24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 5, 1, 0,
		0, 0, 26, 27, 5, 8, 0, 0, 27, 28, 5, 17, 0, 0, 28, 29, 5, 1, 0, 0, 29,
		30, 3, 10, 5, 0, 30, 7, 1, 0, 0, 0, 31, 32, 5, 9, 0, 0, 32, 34, 5, 2, 0,
		0, 33, 35, 3, 10, 5, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 37, 5, 3, 0, 0, 37, 9, 1, 0, 0, 0, 38, 39, 6, 5, -1, 0,
		39, 40, 5, 27, 0, 0, 40, 54, 3, 10, 5, 15, 41, 54, 5, 16, 0, 0, 42, 54,
		5, 15, 0, 0, 43, 54, 3, 8, 4, 0, 44, 46, 5, 17, 0, 0, 45, 47, 3, 12, 6,
		0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 54, 1, 0, 0, 0, 48, 54,
		5, 18, 0, 0, 49, 50, 5, 2, 0, 0, 50, 51, 3, 10, 5, 0, 51, 52, 5, 3, 0,
		0, 52, 54, 1, 0, 0, 0, 53, 38, 1, 0, 0, 0, 53, 41, 1, 0, 0, 0, 53, 42,
		1, 0, 0, 0, 53, 43, 1, 0, 0, 0, 53, 44, 1, 0, 0, 0, 53, 48, 1, 0, 0, 0,
		53, 49, 1, 0, 0, 0, 54, 84, 1, 0, 0, 0, 55, 56, 10, 14, 0, 0, 56, 57, 5,
		25, 0, 0, 57, 83, 3, 10, 5, 14, 58, 59, 10, 13, 0, 0, 59, 60, 7, 0, 0,
		0, 60, 83, 3, 10, 5, 14, 61, 62, 10, 12, 0, 0, 62, 63, 7, 1, 0, 0, 63,
		83, 3, 10, 5, 13, 64, 65, 10, 11, 0, 0, 65, 66, 7, 2, 0, 0, 66, 83, 3,
		10, 5, 12, 67, 68, 10, 10, 0, 0, 68, 69, 7, 3, 0, 0, 69, 83, 3, 10, 5,
		11, 70, 71, 10, 9, 0, 0, 71, 72, 5, 20, 0, 0, 72, 83, 3, 10, 5, 10, 73,
		74, 10, 8, 0, 0, 74, 75, 5, 19, 0, 0, 75, 83, 3, 10, 5, 9, 76, 77, 10,
		7, 0, 0, 77, 78, 5, 4, 0, 0, 78, 79, 3, 10, 5, 0, 79, 80, 5, 5, 0, 0, 80,
		81, 3, 10, 5, 8, 81, 83, 1, 0, 0, 0, 82, 55, 1, 0, 0, 0, 82, 58, 1, 0,
		0, 0, 82, 61, 1, 0, 0, 0, 82, 64, 1, 0, 0, 0, 82, 67, 1, 0, 0, 0, 82, 70,
		1, 0, 0, 0, 82, 73, 1, 0, 0, 0, 82, 76, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 11, 1, 0, 0, 0, 86, 84, 1,
		0, 0, 0, 87, 88, 5, 6, 0, 0, 88, 89, 3, 10, 5, 0, 89, 90, 5, 7, 0, 0, 90,
		92, 1, 0, 0, 0, 91, 87, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0,
		0, 93, 94, 1, 0, 0, 0, 94, 13, 1, 0, 0, 0, 8, 19, 24, 34, 46, 53, 82, 84,
		93,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftLanParserInit initializes any static state used to implement SwiftLanParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftLanParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLanParserInit() {
	staticData := &SwiftLanParserStaticData
	staticData.once.Do(swiftlanParserInit)
}

// NewSwiftLanParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftLanParser(input antlr.TokenStream) *SwiftLanParser {
	SwiftLanParserInit()
	this := new(SwiftLanParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftLanParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftLan.g4"

	return this
}

// SwiftLanParser tokens.
const (
	SwiftLanParserEOF          = antlr.TokenEOF
	SwiftLanParserT__0         = 1
	SwiftLanParserT__1         = 2
	SwiftLanParserT__2         = 3
	SwiftLanParserT__3         = 4
	SwiftLanParserT__4         = 5
	SwiftLanParserT__5         = 6
	SwiftLanParserT__6         = 7
	SwiftLanParserDeclarar     = 8
	SwiftLanParserPrint        = 9
	SwiftLanParserIntDecla     = 10
	SwiftLanParserFloatDecla   = 11
	SwiftLanParserBoolDecla    = 12
	SwiftLanParserStringDecla  = 13
	SwiftLanParserCharDecla    = 14
	SwiftLanParserBoolVal      = 15
	SwiftLanParserNumber       = 16
	SwiftLanParserId           = 17
	SwiftLanParserString_      = 18
	SwiftLanParserOr           = 19
	SwiftLanParserAnd          = 20
	SwiftLanParserEquals       = 21
	SwiftLanParserNotEquals    = 22
	SwiftLanParserMayQEquals   = 23
	SwiftLanParserMinQEquals   = 24
	SwiftLanParserPower        = 25
	SwiftLanParserAdd          = 26
	SwiftLanParserMinus        = 27
	SwiftLanParserMult         = 28
	SwiftLanParserDiv          = 29
	SwiftLanParserMod          = 30
	SwiftLanParserMayorQue     = 31
	SwiftLanParserMenorQue     = 32
	SwiftLanParserSpace        = 33
	SwiftLanParserCOMMENT      = 34
	SwiftLanParserLINE_COMMENT = 35
)

// SwiftLanParser rules.
const (
	SwiftLanParserRULE_inicio       = 0
	SwiftLanParserRULE_sentencias   = 1
	SwiftLanParserRULE_statement    = 2
	SwiftLanParserRULE_asignacion   = 3
	SwiftLanParserRULE_llamadaPrint = 4
	SwiftLanParserRULE_expression   = 5
	SwiftLanParserRULE_indexes      = 6
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sentencias() ISentenciasContext

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_inicio
	return p
}

func InitEmptyInicioContext(p *InicioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_inicio
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Sentencias() ISentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciasContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitInicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftLanParserRULE_inicio)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Sentencias()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciasContext is an interface to support dynamic dispatch.
type ISentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsSentenciasContext differentiates from other interfaces.
	IsSentenciasContext()
}

type SentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciasContext() *SentenciasContext {
	var p = new(SentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_sentencias
	return p
}

func InitEmptySentenciasContext(p *SentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_sentencias
}

func (*SentenciasContext) IsSentenciasContext() {}

func NewSentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciasContext {
	var p = new(SentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_sentencias

	return p
}

func (s *SentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciasContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SentenciasContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Sentencias() (localctx ISentenciasContext) {
	localctx = NewSentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftLanParserRULE_sentencias)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftLanParserDeclarar || _la == SwiftLanParserPrint {
		{
			p.SetState(16)
			p.Statement()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Asignacion() IAsignacionContext
	LlamadaPrint() ILlamadaPrintContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *StatementContext) LlamadaPrint() ILlamadaPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaPrintContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftLanParserRULE_statement)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftLanParserDeclarar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Asignacion()
		}

	case SwiftLanParserPrint:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.LlamadaPrint()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declarar() antlr.TerminalNode
	Id() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Declarar() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserDeclarar, 0)
}

func (s *AsignacionContext) Id() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserId, 0)
}

func (s *AsignacionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftLanParserRULE_asignacion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(SwiftLanParserDeclarar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(SwiftLanParserId)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(SwiftLanParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaPrintContext is an interface to support dynamic dispatch.
type ILlamadaPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLlamadaPrintContext differentiates from other interfaces.
	IsLlamadaPrintContext()
}

type LlamadaPrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaPrintContext() *LlamadaPrintContext {
	var p = new(LlamadaPrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_llamadaPrint
	return p
}

func InitEmptyLlamadaPrintContext(p *LlamadaPrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_llamadaPrint
}

func (*LlamadaPrintContext) IsLlamadaPrintContext() {}

func NewLlamadaPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaPrintContext {
	var p = new(LlamadaPrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_llamadaPrint

	return p
}

func (s *LlamadaPrintContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaPrintContext) CopyAll(ctx *LlamadaPrintContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LlamadaPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaPrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintLlamadaPrintContext struct {
	LlamadaPrintContext
}

func NewPrintLlamadaPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintLlamadaPrintContext {
	var p = new(PrintLlamadaPrintContext)

	InitEmptyLlamadaPrintContext(&p.LlamadaPrintContext)
	p.parser = parser
	p.CopyAll(ctx.(*LlamadaPrintContext))

	return p
}

func (s *PrintLlamadaPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintLlamadaPrintContext) Print() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserPrint, 0)
}

func (s *PrintLlamadaPrintContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintLlamadaPrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitPrintLlamadaPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) LlamadaPrint() (localctx ILlamadaPrintContext) {
	localctx = NewLlamadaPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftLanParserRULE_llamadaPrint)
	var _la int

	localctx = NewPrintLlamadaPrintContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(SwiftLanParserPrint)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(SwiftLanParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134709764) != 0 {
		{
			p.SetState(33)
			p.expression(0)
		}

	}
	{
		p.SetState(36)
		p.Match(SwiftLanParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LlamadaFuncionExpressionContext struct {
	ExpressionContext
}

func NewLlamadaFuncionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamadaFuncionExpressionContext {
	var p = new(LlamadaFuncionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LlamadaFuncionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionExpressionContext) LlamadaPrint() ILlamadaPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaPrintContext)
}

func (s *LlamadaFuncionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitLlamadaFuncionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExpressionContext struct {
	ExpressionContext
}

func NewBoolExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) BoolVal() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserBoolVal, 0)
}

func (s *BoolExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitBoolExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExpressionContext struct {
	ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserNumber, 0)
}

func (s *NumberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitNumberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpressionContext struct {
	ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserOr, 0)
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExpressionContext struct {
	ExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMinus, 0)
}

func (s *UnaryMinusExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitUnaryMinusExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerExpressionContext struct {
	ExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExpressionContext) Power() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserPower, 0)
}

func (s *PowerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitPowerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewEqExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExpressionContext {
	var p = new(EqExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserEquals, 0)
}

func (s *EqExpressionContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserNotEquals, 0)
}

func (s *EqExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitEqExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpressionContext struct {
	ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserAnd, 0)
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExpressionContext struct {
	ExpressionContext
}

func NewIdExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpressionContext {
	var p = new(IdExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpressionContext) Id() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserId, 0)
}

func (s *IdExpressionContext) Indexes() IIndexesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *IdExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitIdExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExpressionContext struct {
	ExpressionContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) String_() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserString_, 0)
}

func (s *StringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionExpressionContext struct {
	ExpressionContext
}

func NewExpressionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionExpressionContext {
	var p = new(ExpressionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitExpressionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) Add() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserAdd, 0)
}

func (s *AddExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMinus, 0)
}

func (s *AddExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitAddExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewCompExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExpressionContext {
	var p = new(CompExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompExpressionContext) GetOp() antlr.Token { return s.op }

func (s *CompExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) MayQEquals() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMayQEquals, 0)
}

func (s *CompExpressionContext) MinQEquals() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMinQEquals, 0)
}

func (s *CompExpressionContext) MayorQue() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMayorQue, 0)
}

func (s *CompExpressionContext) MenorQue() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMenorQue, 0)
}

func (s *CompExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitCompExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMultExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpressionContext {
	var p = new(MultExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) Mult() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMult, 0)
}

func (s *MultExpressionContext) Div() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserDiv, 0)
}

func (s *MultExpressionContext) Mod() antlr.TerminalNode {
	return s.GetToken(SwiftLanParserMod, 0)
}

func (s *MultExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitMultExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExpressionContext struct {
	ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitTernaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SwiftLanParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SwiftLanParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftLanParserMinus:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(39)
			p.Match(SwiftLanParserMinus)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.expression(15)
		}

	case SwiftLanParserNumber:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.Match(SwiftLanParserNumber)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftLanParserBoolVal:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(SwiftLanParserBoolVal)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftLanParserPrint:
		localctx = NewLlamadaFuncionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.LlamadaPrint()
		}

	case SwiftLanParserId:
		localctx = NewIdExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(SwiftLanParserId)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(45)
				p.Indexes()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SwiftLanParserString_:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(SwiftLanParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftLanParserT__1:
		localctx = NewExpressionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(SwiftLanParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.expression(0)
		}
		{
			p.SetState(51)
			p.Match(SwiftLanParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(56)
					p.Match(SwiftLanParserPower)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(57)
					p.expression(14)
				}

			case 2:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(59)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1879048192) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(60)
					p.expression(14)
				}

			case 3:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftLanParserAdd || _la == SwiftLanParserMinus) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(13)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(65)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6467616768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)
					p.expression(12)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(68)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftLanParserEquals || _la == SwiftLanParserNotEquals) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)
					p.expression(11)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(71)
					p.Match(SwiftLanParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(72)
					p.expression(10)
				}

			case 7:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(74)
					p.Match(SwiftLanParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(75)
					p.expression(9)
				}

			case 8:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SwiftLanParserRULE_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(77)
					p.Match(SwiftLanParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(78)
					p.expression(0)
				}
				{
					p.SetState(79)
					p.Match(SwiftLanParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expression(8)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexesContext is an interface to support dynamic dispatch.
type IIndexesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsIndexesContext differentiates from other interfaces.
	IsIndexesContext()
}

type IndexesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexesContext() *IndexesContext {
	var p = new(IndexesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_indexes
	return p
}

func InitEmptyIndexesContext(p *IndexesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftLanParserRULE_indexes
}

func (*IndexesContext) IsIndexesContext() {}

func NewIndexesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexesContext {
	var p = new(IndexesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftLanParserRULE_indexes

	return p
}

func (s *IndexesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexesContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexesContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftLanVisitor:
		return t.VisitIndexes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftLanParser) Indexes() (localctx IIndexesContext) {
	localctx = NewIndexesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftLanParserRULE_indexes)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(87)
				p.Match(SwiftLanParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(88)
				p.expression(0)
			}
			{
				p.SetState(89)
				p.Match(SwiftLanParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftLanParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftLanParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
