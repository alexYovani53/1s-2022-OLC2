// Code generated from Calculadora.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calculadora

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 27, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23, 11, 3, 3,
	4, 3, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6, 2, 4, 3, 2, 28, 29, 3, 2, 26, 27, 2,
	25, 2, 8, 3, 2, 2, 2, 4, 10, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 9, 5, 4,
	3, 2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1, 2, 11, 12, 5, 6, 4, 2, 12, 21,
	3, 2, 2, 2, 13, 14, 12, 5, 2, 2, 14, 15, 9, 2, 2, 2, 15, 20, 5, 4, 3, 6,
	16, 17, 12, 4, 2, 2, 17, 18, 9, 3, 2, 2, 18, 20, 5, 4, 3, 5, 19, 13, 3,
	2, 2, 2, 19, 16, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21,
	22, 3, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 25, 7, 30, 2,
	2, 25, 7, 3, 2, 2, 2, 4, 19, 21,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'system'", "'out'", "'println'", "'int'",
	"'float'", "'string'", "'boolean'", "'.'", "','", "';'", "'&&'", "'||'",
	"'!'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'",
	"'-'", "", "", "", "", "'true'",
}
var symbolicNames = []string{
	"", "LP", "RP", "L_LLAVE", "R_LLAVE", "SYSTEM", "OUT", "PRINTLN", "INTTYPE",
	"FLOATTYPE", "STRINGTYPE", "BOOLTYPE", "PUNTO", "COMA", "PTCOMA", "AND",
	"OR", "NOT", "IGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR",
	"MENOR", "MUL", "DIV", "ADD", "SUB", "NUMBER", "FLOAT", "STRING", "ID",
	"TRUE", "FALSE", "WHITESPACE", "AB",
}

var ruleNames = []string{
	"inicio", "expresion", "primitivo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Calculadora struct {
	*antlr.BaseParser
}

func NewCalculadora(input antlr.TokenStream) *Calculadora {
	this := new(Calculadora)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calculadora.g4"

	return this
}

// Calculadora tokens.
const (
	CalculadoraEOF        = antlr.TokenEOF
	CalculadoraLP         = 1
	CalculadoraRP         = 2
	CalculadoraL_LLAVE    = 3
	CalculadoraR_LLAVE    = 4
	CalculadoraSYSTEM     = 5
	CalculadoraOUT        = 6
	CalculadoraPRINTLN    = 7
	CalculadoraINTTYPE    = 8
	CalculadoraFLOATTYPE  = 9
	CalculadoraSTRINGTYPE = 10
	CalculadoraBOOLTYPE   = 11
	CalculadoraPUNTO      = 12
	CalculadoraCOMA       = 13
	CalculadoraPTCOMA     = 14
	CalculadoraAND        = 15
	CalculadoraOR         = 16
	CalculadoraNOT        = 17
	CalculadoraIGUAL      = 18
	CalculadoraDIFERENTE  = 19
	CalculadoraMAYORIGUAL = 20
	CalculadoraMENORIGUAL = 21
	CalculadoraMAYOR      = 22
	CalculadoraMENOR      = 23
	CalculadoraMUL        = 24
	CalculadoraDIV        = 25
	CalculadoraADD        = 26
	CalculadoraSUB        = 27
	CalculadoraNUMBER     = 28
	CalculadoraFLOAT      = 29
	CalculadoraSTRING     = 30
	CalculadoraID         = 31
	CalculadoraTRUE       = 32
	CalculadoraFALSE      = 33
	CalculadoraWHITESPACE = 34
	CalculadoraAB         = 35
)

// Calculadora rules.
const (
	CalculadoraRULE_inicio    = 0
	CalculadoraRULE_expresion = 1
	CalculadoraRULE_primitivo = 2
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSalida returns the salida attribute.
	GetSalida() int

	// SetSalida sets the salida attribute.
	SetSalida(int)

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	salida int
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculadoraRULE_inicio
	return p
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculadoraRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) GetSalida() int { return s.salida }

func (s *InicioContext) SetSalida(v int) { s.salida = v }

func (s *InicioContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.ExitInicio(s)
	}
}

func (p *Calculadora) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalculadoraRULE_inicio)

	localctx.(*InicioContext).salida = 10

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
		p.SetState(6)
		p.expresion(0)
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpresionContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpresionContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpresionContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpresionContext)

	// GetSalida returns the salida attribute.
	GetSalida() int

	// SetSalida sets the salida attribute.
	SetSalida(int)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	salida int
	opIz   IExpresionContext
	op     antlr.Token
	opDe   IExpresionContext
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculadoraRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculadoraRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) GetOpIz() IExpresionContext { return s.opIz }

func (s *ExpresionContext) GetOpDe() IExpresionContext { return s.opDe }

func (s *ExpresionContext) SetOpIz(v IExpresionContext) { s.opIz = v }

func (s *ExpresionContext) SetOpDe(v IExpresionContext) { s.opDe = v }

func (s *ExpresionContext) GetSalida() int { return s.salida }

func (s *ExpresionContext) SetSalida(v int) { s.salida = v }

func (s *ExpresionContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalculadoraADD, 0)
}

func (s *ExpresionContext) SUB() antlr.TerminalNode {
	return s.GetToken(CalculadoraSUB, 0)
}

func (s *ExpresionContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalculadoraMUL, 0)
}

func (s *ExpresionContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalculadoraDIV, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *Calculadora) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *Calculadora) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalculadoraRULE_expresion, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(9)
		p.Primitivo()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(17)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculadoraRULE_expresion)
				p.SetState(11)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(12)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculadoraADD || _la == CalculadoraSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(13)

					var _x = p.expresion(4)

					localctx.(*ExpresionContext).opDe = _x
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, CalculadoraRULE_expresion)
				p.SetState(14)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(15)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalculadoraMUL || _la == CalculadoraDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(16)

					var _x = p.expresion(3)

					localctx.(*ExpresionContext).opDe = _x
				}

			}

		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSalida returns the salida attribute.
	GetSalida() int

	// SetSalida sets the salida attribute.
	SetSalida(int)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	salida int
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculadoraRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculadoraRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) GetSalida() int { return s.salida }

func (s *PrimitivoContext) SetSalida(v int) { s.salida = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculadoraNUMBER, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalculadoraListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Calculadora) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalculadoraRULE_primitivo)

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
		p.SetState(22)
		p.Match(CalculadoraNUMBER)
	}

	return localctx
}

func (p *Calculadora) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Calculadora) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
