// Code generated from .\SwiftLan.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftLan

import "github.com/antlr4-go/antlr/v4"

type BaseSwiftLanVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftLanVisitor) VisitInicio(ctx *InicioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitSentencias(ctx *SentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitPrintLlamadaPrint(ctx *PrintLlamadaPrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitLlamadaFuncionExpression(ctx *LlamadaFuncionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitBoolExpression(ctx *BoolExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitNumberExpression(ctx *NumberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitEqExpression(ctx *EqExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitIdExpression(ctx *IdExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitStringExpression(ctx *StringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitExpressionExpression(ctx *ExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitCompExpression(ctx *CompExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitMultExpression(ctx *MultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftLanVisitor) VisitIndexes(ctx *IndexesContext) interface{} {
	return v.VisitChildren(ctx)
}
