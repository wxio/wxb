package adlproc

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	parser "github.com/wxio/wxb/internal/grammars/adl"
	walker "github.com/wxio/wxb/internal/grammars/adlwalker"
)

type walkADL struct {
	File string `type:"arg" help:"proto file" predict:"files"`
}

type TTType struct{}

func (*TTType) Eof() int  { return walker.ADLWalkerEOF }
func (*TTType) Down() int { return walker.ADLWalkerDOWN }
func (*TTType) Up() int   { return walker.ADLWalkerUP }

func (et *walkADL) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}

	var tr ctree.Tree
	{
		errListener := &errorListener{}
		r := &ADLNode{MyToken: MyToken{Token: nil, TType: parser.ADLParserADL}}
		tbl := &ADLBuildListener{
			debug:   true,
			Builder: ctree.NewWalkableBuild("TREE", r),
		}
		is := antlr.NewInputStream(string(by))
		lexer := parser.NewadlLexer(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewADLParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.AddErrorListener(tbl)
		p.BuildParseTrees = true
		ctx := p.Adl()
		fmt.Println("--------")
		antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
		tr = tbl.Builder.Build()
	}

	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewADLWalker(tts)
	i := 1
	for {
		to := tts.Get(i)
		if -1 == to.GetTokenType() {
			break
		}
		fmt.Printf("%d ", to.GetTokenType())
		fmt.Printf("%d : %v\n", i,
			p.BaseRecognizer.SymbolicNames[to.GetTokenType()],
			// to.GetText(),
		)
		i++
	}
	// return nil

	p.SetTokenStream(tts)
	// p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tbl := &WalkListener{}
	tbl.debug = true
	// p.AddErrorListener(tbl)
	// p.BuildParseTrees = false
	ret := p.Adl()
	// fmt.Printf("ret %#+v\n", tbl)
	antlr.ParseTreeWalkerDefault.Walk(tbl, ret)
	return tbl.err
}

type WalkListener struct {
	*walker.BaseADLWalkerListener
	//
	protoDS *descriptor.FileDescriptorSet
	indent  string
	debug   bool
	warning string
	err     error
}

// EnterTld is called when production tld is entered.
func (tr *WalkListener) EnterTld(ctx *walker.TldContext) {
	// fmt.Printf("tld %v\n", ctx)
}

func (tr *WalkListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%s  '%T'\n", tr.indent, node.GetPayload())
}
func (tr *WalkListener) VisitErrorNode(node antlr.ErrorNode) {
	tid := node.GetSymbol().GetTokenType()
	fmt.Printf("2.ERROR %d %+v\n", tid, node)
}
func (tr *WalkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%s>>%T\n", tr.indent, ctx)
	tr.indent += "  "
}
func (tr *WalkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("\n")
	tr.indent = tr.indent[2:]
	// fmt.Printf("\n%s<<%T", tr.indent, ctx)
}

func (tbl *WalkListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		tbl.warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
	} else {
		tbl.err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
	}
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *WalkListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *WalkListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *WalkListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("ReportContextSensitivity")
}
