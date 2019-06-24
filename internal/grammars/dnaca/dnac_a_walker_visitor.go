// Generated from DNAC_A_Walker.g4 by ANTLR 4.7.

package dnaca // DNAC_A_Walker
import "github.com/wxio/goantlr"

// Struct of Handlers
type DNAC_A_WalkerHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Adl            func(ctx IAdlContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	NameNode       func(ctx INameNodeContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	NameRule       func(ctx INameRuleContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	TypeNode       func(ctx ITypeNodeContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	ExnotationNode func(ctx IExnotationNodeContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	NameBodyNode   func(ctx INameBodyNodeContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	Annotation     func(ctx IAnnotationContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	TypeExpr_      func(ctx ITypeExpr_Context, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	TypeParams     func(ctx ITypeParamsContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonStr        func(ctx IJsonStrContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonBool       func(ctx IJsonBoolContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonNull       func(ctx IJsonNullContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonInt        func(ctx IJsonIntContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonFloat      func(ctx IJsonFloatContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonArray      func(ctx IJsonArrayContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
	JsonObj        func(ctx IJsonObjContext, this *DNAC_A_WalkerHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by DNAC_A_Walker.
type DNAC_A_WalkerVisitor interface {
	antlr.ParseTreeVisitor
	AdlContextVisitor
	NameNodeContextVisitor
	NameRuleContextVisitor
	TypeNodeContextVisitor
	ExnotationNodeContextVisitor
	NameBodyNodeContextVisitor
	AnnotationContextVisitor
	TypeExpr_ContextVisitor
	TypeParamsContextVisitor
	JsonStrContextVisitor
	JsonBoolContextVisitor
	JsonNullContextVisitor
	JsonIntContextVisitor
	JsonFloatContextVisitor
	JsonArrayContextVisitor
	JsonObjContextVisitor
}

type AdlContextVisitor interface {
	VisitAdl(ctx IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameNodeContextVisitor interface {
	VisitNameNode(ctx INameNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameRuleContextVisitor interface {
	VisitNameRule(ctx INameRuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeNodeContextVisitor interface {
	VisitTypeNode(ctx ITypeNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ExnotationNodeContextVisitor interface {
	VisitExnotationNode(ctx IExnotationNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NameBodyNodeContextVisitor interface {
	VisitNameBodyNode(ctx INameBodyNodeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type AnnotationContextVisitor interface {
	VisitAnnotation(ctx IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExpr_ContextVisitor interface {
	VisitTypeExpr_(ctx ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParamsContextVisitor interface {
	VisitTypeParams(ctx ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonStrContextVisitor interface {
	VisitJsonStr(ctx IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonBoolContextVisitor interface {
	VisitJsonBool(ctx IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonNullContextVisitor interface {
	VisitJsonNull(ctx IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonIntContextVisitor interface {
	VisitJsonInt(ctx IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonFloatContextVisitor interface {
	VisitJsonFloat(ctx IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonArrayContextVisitor interface {
	VisitJsonArray(ctx IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonObjContextVisitor interface {
	VisitJsonObj(ctx IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}