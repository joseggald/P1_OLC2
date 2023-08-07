// Code generated from .\SwiftLan.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftLan

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftLanParser.
type SwiftLanVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftLanParser#inicio.
	VisitInicio(ctx *InicioContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#sentencias.
	VisitSentencias(ctx *SentenciasContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#printLlamadaPrint.
	VisitPrintLlamadaPrint(ctx *PrintLlamadaPrintContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#llamadaFuncionExpression.
	VisitLlamadaFuncionExpression(ctx *LlamadaFuncionExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#boolExpression.
	VisitBoolExpression(ctx *BoolExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#numberExpression.
	VisitNumberExpression(ctx *NumberExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#unaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#powerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#eqExpression.
	VisitEqExpression(ctx *EqExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#idExpression.
	VisitIdExpression(ctx *IdExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#stringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#expressionExpression.
	VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#compExpression.
	VisitCompExpression(ctx *CompExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#multExpression.
	VisitMultExpression(ctx *MultExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#ternaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by SwiftLanParser#indexes.
	VisitIndexes(ctx *IndexesContext) interface{}
}
