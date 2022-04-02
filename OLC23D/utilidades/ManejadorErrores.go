package utilidades

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomSyntaxError struct {
	line, column int
	msg          string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
