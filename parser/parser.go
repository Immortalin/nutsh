
//line parser.l:2

package parser
import __yyfmt__ "fmt"
//line parser.l:3
		
import (
	"strings"
)

var lesson Node

type Node struct {
	typ string
	children []Node
}

func node(typ string, n ...Node) Node {
	return Node{typ: typ, children: n}
}

func (n Node) String() string {
	r := n.typ
	r += ":"
	for _, child := range(n.children) {
		lines := strings.Split(child.String(), "\n")
		for _, line := range(lines) {
			r += "\n    "+line
		}
	}
	return r
}

/*
type Node interface {
	Execute()
}

type Block struct {
	lines []Line
}

func (b Block) Execute() {
	for _, line := range(b.lines) {
		line.Execute()
	}
}

type Line struct {
	message string	
}

func (l Line) Execute() {
	println(l.message)
}

type IfNode struct {
	condition BoolExpressionNode
	block BlockNode
}

func (n IfNode) Execute() {
	if (n.condition.Execute()) {
		block.Execute()
	}
}

type BoolExpressionNode struct {
	typ: string
	child1, child2 BoolExpressioonNode
}

func (n BoolExpressionNode) Execute() bool {
	switch n.typ {
		case "not":
			return ! n.child1.Execute()
		case "and":
			return n.child1.Execute() && n.child2.Execute()
		case "or":
			return n.child1.Execute() || n.child2.Execute()
		case "=~":
			return n.child1.Execute() // ... match
	}
}
*/


//line parser.l:87
type NutshSymType struct {
	yys int
    val string
	node Node
}

const DEF = 57346
const IDENTIFIER = 57347
const LINE = 57348
const LINESEP = 57349
const IF = 57350
const PROMPT = 57351
const STRING = 57352
const AND = 57353
const OR = 57354
const NOT = 57355
const MATCH = 57356

var NutshToknames = []string{
	"DEF",
	"IDENTIFIER",
	"LINE",
	"LINESEP",
	"IF",
	"PROMPT",
	"STRING",
	"AND",
	"OR",
	"NOT",
	"MATCH",
}
var NutshStatenames = []string{}

const NutshEofCode = 1
const NutshErrCode = 2
const NutshMaxDepth = 200

//line parser.l:150


func Parse(l lexer) {
	pos = 0
	first = 0
	NutshParse(l)
}

//line yacctab:1
var NutshExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	17, 22,
	19, 22,
	-2, 15,
	-1, 17,
	14, 29,
	20, 29,
	-2, 31,
}

const NutshNprod = 37
const NutshPrivate = 57344

var NutshTokenNames []string
var NutshStates []string

const NutshLast = 74

var NutshAct = []int{

	17, 5, 18, 10, 22, 33, 32, 23, 28, 14,
	16, 32, 11, 12, 53, 56, 51, 54, 24, 52,
	26, 29, 49, 23, 43, 25, 37, 40, 14, 39,
	34, 35, 41, 40, 40, 46, 47, 10, 50, 30,
	31, 44, 45, 30, 31, 23, 14, 14, 48, 30,
	31, 21, 21, 40, 20, 55, 19, 6, 14, 15,
	13, 11, 12, 9, 8, 7, 36, 42, 27, 4,
	38, 3, 2, 1,
}
var NutshPact = []int{

	-1000, -1000, 53, -1000, -1000, -1000, 54, -1000, -1000, -1000,
	-1000, 41, -10, 6, 5, -7, 28, -1000, -9, 41,
	41, -1000, -1000, -1000, -1000, 23, 42, -10, 19, -1000,
	41, 41, 42, 42, 32, 38, 4, -1000, 0, -14,
	-1000, -1000, -2, -1000, 38, 38, -14, -14, -1000, -1000,
	-1000, -1000, 42, -1000, 10, -14, -1000,
}
var NutshPgo = []int{

	0, 73, 72, 71, 1, 0, 70, 2, 69, 68,
	67, 4, 66, 65, 64, 63, 10, 60,
}
var NutshR1 = []int{

	0, 1, 8, 9, 9, 11, 2, 2, 3, 3,
	12, 12, 4, 4, 4, 4, 13, 14, 15, 10,
	10, 10, 17, 17, 5, 5, 6, 6, 7, 7,
	7, 16, 16, 16, 16, 16, 16,
}
var NutshR2 = []int{

	0, 1, 4, 0, 3, 3, 0, 2, 1, 1,
	0, 2, 1, 1, 1, 1, 3, 2, 2, 0,
	1, 3, 1, 3, 1, 4, 1, 3, 1, 1,
	3, 1, 3, 3, 2, 3, 3,
}
var NutshChk = []int{

	-1000, -1, -2, -3, -8, -4, 4, -13, -14, -15,
	-5, 8, 9, -17, 5, 5, -16, -5, -7, 15,
	13, 10, -11, 17, -11, 19, 15, -9, 15, -11,
	11, 12, 20, 14, -16, -16, -12, -5, -6, -7,
	-5, -11, -10, 5, -16, -16, -7, -7, 16, 18,
	-4, 16, 19, 16, 19, -7, 5,
}
var NutshDef = []int{

	6, -2, 1, 7, 8, 9, 0, 12, 13, 14,
	-2, 0, 0, 0, 24, 3, 0, -2, 0, 0,
	0, 28, 17, 10, 18, 0, 0, 0, 19, 16,
	0, 0, 0, 0, 0, 34, 0, 23, 0, 26,
	29, 2, 0, 20, 35, 36, 30, 32, 33, 5,
	11, 25, 0, 4, 0, 27, 21,
}
var NutshTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 20, 19, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 17, 3, 18,
}
var NutshTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var NutshTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var NutshDebug = 0

type NutshLexer interface {
	Lex(lval *NutshSymType) int
	Error(s string)
}

const NutshFlag = -1000

func NutshTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(NutshToknames) {
		if NutshToknames[c-4] != "" {
			return NutshToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func NutshStatname(s int) string {
	if s >= 0 && s < len(NutshStatenames) {
		if NutshStatenames[s] != "" {
			return NutshStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Nutshlex1(lex NutshLexer, lval *NutshSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = NutshTok1[0]
		goto out
	}
	if char < len(NutshTok1) {
		c = NutshTok1[char]
		goto out
	}
	if char >= NutshPrivate {
		if char < NutshPrivate+len(NutshTok2) {
			c = NutshTok2[char-NutshPrivate]
			goto out
		}
	}
	for i := 0; i < len(NutshTok3); i += 2 {
		c = NutshTok3[i+0]
		if c == char {
			c = NutshTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = NutshTok2[1] /* unknown char */
	}
	if NutshDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), NutshTokname(c))
	}
	return c
}

func NutshParse(Nutshlex NutshLexer) int {
	var Nutshn int
	var Nutshlval NutshSymType
	var NutshVAL NutshSymType
	NutshS := make([]NutshSymType, NutshMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Nutshstate := 0
	Nutshchar := -1
	Nutshp := -1
	goto Nutshstack

ret0:
	return 0

ret1:
	return 1

Nutshstack:
	/* put a state and value onto the stack */
	if NutshDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", NutshTokname(Nutshchar), NutshStatname(Nutshstate))
	}

	Nutshp++
	if Nutshp >= len(NutshS) {
		nyys := make([]NutshSymType, len(NutshS)*2)
		copy(nyys, NutshS)
		NutshS = nyys
	}
	NutshS[Nutshp] = NutshVAL
	NutshS[Nutshp].yys = Nutshstate

Nutshnewstate:
	Nutshn = NutshPact[Nutshstate]
	if Nutshn <= NutshFlag {
		goto Nutshdefault /* simple state */
	}
	if Nutshchar < 0 {
		Nutshchar = Nutshlex1(Nutshlex, &Nutshlval)
	}
	Nutshn += Nutshchar
	if Nutshn < 0 || Nutshn >= NutshLast {
		goto Nutshdefault
	}
	Nutshn = NutshAct[Nutshn]
	if NutshChk[Nutshn] == Nutshchar { /* valid shift */
		Nutshchar = -1
		NutshVAL = Nutshlval
		Nutshstate = Nutshn
		if Errflag > 0 {
			Errflag--
		}
		goto Nutshstack
	}

Nutshdefault:
	/* default state action */
	Nutshn = NutshDef[Nutshstate]
	if Nutshn == -2 {
		if Nutshchar < 0 {
			Nutshchar = Nutshlex1(Nutshlex, &Nutshlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if NutshExca[xi+0] == -1 && NutshExca[xi+1] == Nutshstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Nutshn = NutshExca[xi+0]
			if Nutshn < 0 || Nutshn == Nutshchar {
				break
			}
		}
		Nutshn = NutshExca[xi+1]
		if Nutshn < 0 {
			goto ret0
		}
	}
	if Nutshn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Nutshlex.Error("syntax error")
			Nerrs++
			if NutshDebug >= 1 {
				__yyfmt__.Printf("%s", NutshStatname(Nutshstate))
				__yyfmt__.Printf("saw %s\n", NutshTokname(Nutshchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Nutshp >= 0 {
				Nutshn = NutshPact[NutshS[Nutshp].yys] + NutshErrCode
				if Nutshn >= 0 && Nutshn < NutshLast {
					Nutshstate = NutshAct[Nutshn] /* simulate a shift of "error" */
					if NutshChk[Nutshstate] == NutshErrCode {
						goto Nutshstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if NutshDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", NutshS[Nutshp].yys)
				}
				Nutshp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if NutshDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", NutshTokname(Nutshchar))
			}
			if Nutshchar == NutshEofCode {
				goto ret1
			}
			Nutshchar = -1
			goto Nutshnewstate /* try again in the same state */
		}
	}

	/* reduction by production Nutshn */
	if NutshDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Nutshn, NutshStatname(Nutshstate))
	}

	Nutshnt := Nutshn
	Nutshpt := Nutshp
	_ = Nutshpt // guard against "declared and not used"

	Nutshp -= NutshR2[Nutshn]
	NutshVAL = NutshS[Nutshp+1]

	/* consult goto table to find next state */
	Nutshn = NutshR1[Nutshn]
	Nutshg := NutshPgo[Nutshn]
	Nutshj := Nutshg + NutshS[Nutshp].yys + 1

	if Nutshj >= NutshLast {
		Nutshstate = NutshAct[Nutshg]
	} else {
		Nutshstate = NutshAct[Nutshj]
		if NutshChk[Nutshstate] != -Nutshn {
			Nutshstate = NutshAct[Nutshg]
		}
	}
	// dummy call; replaced with literal code
	switch Nutshnt {

	case 1:
		//line parser.l:97
		{ NutshVAL.node = NutshS[Nutshpt-0].node; lesson = NutshVAL.node }
	case 2:
		//line parser.l:99
		{ NutshVAL.node = node("def", node(NutshS[Nutshpt-2].val), NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node) }
	case 3:
		//line parser.l:101
		{ NutshVAL.node = node("arguments") }
	case 4:
		//line parser.l:102
		{ NutshVAL.node = node("arguments", NutshS[Nutshpt-1].node.children...) }
	case 5:
		//line parser.l:104
		{ NutshVAL.node = NutshS[Nutshpt-1].node }
	case 6:
		//line parser.l:106
		{ NutshVAL.node = node("block") }
	case 7:
		//line parser.l:107
		{ NutshVAL.node = node("block", append(NutshS[Nutshpt-1].node.children, NutshS[Nutshpt-0].node)...) }
	case 8:
		//line parser.l:109
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 9:
		//line parser.l:110
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 10:
		//line parser.l:112
		{ NutshVAL.node = node("block") }
	case 11:
		//line parser.l:113
		{ NutshVAL.node = node("block", append(NutshS[Nutshpt-1].node.children, NutshS[Nutshpt-0].node)...) }
	case 12:
		//line parser.l:115
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 13:
		//line parser.l:116
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 14:
		//line parser.l:117
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 15:
		//line parser.l:118
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 16:
		//line parser.l:120
		{ NutshVAL.node = node("if", NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node) }
	case 17:
		//line parser.l:122
		{ NutshVAL.node = node("prompt", NutshS[Nutshpt-0].node) }
	case 18:
		//line parser.l:124
		{ NutshVAL.node = node("state", NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node) }
	case 19:
		//line parser.l:126
		{ NutshVAL.node = node("identifiers") }
	case 20:
		//line parser.l:127
		{ NutshVAL.node = node("identifiers", node("id", node(NutshS[Nutshpt-0].val))) }
	case 21:
		//line parser.l:128
		{ NutshVAL.node = node("identifiers", append(NutshS[Nutshpt-2].node.children, node(NutshS[Nutshpt-0].val))...) }
	case 22:
		//line parser.l:130
		{ node("calls", NutshS[Nutshpt-0].node) }
	case 23:
		//line parser.l:131
		{ node("calls", append(NutshS[Nutshpt-2].node.children, NutshS[Nutshpt-0].node)...) }
	case 24:
		//line parser.l:133
		{ NutshVAL.node = node("call", node(NutshS[Nutshpt-0].val)) }
	case 25:
		//line parser.l:134
		{ NutshVAL.node = node("call", node(NutshS[Nutshpt-3].val), NutshS[Nutshpt-1].node) }
	case 26:
		//line parser.l:136
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 27:
		//line parser.l:137
		{ NutshVAL.node = node("stringexpressions", append(NutshS[Nutshpt-2].node.children, NutshS[Nutshpt-0].node)...) }
	case 28:
		//line parser.l:139
		{ NutshVAL.node = node("string", node(NutshS[Nutshpt-0].val)) }
	case 29:
		//line parser.l:140
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 30:
		//line parser.l:141
		{ NutshVAL.node = node("+", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 31:
		//line parser.l:143
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 32:
		//line parser.l:144
		{ NutshVAL.node = node("match", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 33:
		//line parser.l:145
		{ NutshVAL.node = NutshS[Nutshpt-1].node }
	case 34:
		//line parser.l:146
		{ NutshVAL.node = node("not", NutshS[Nutshpt-0].node) }
	case 35:
		//line parser.l:147
		{ NutshVAL.node = node("and", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 36:
		//line parser.l:148
		{ NutshVAL.node = node("or", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	}
	goto Nutshstack /* stack new state and value */
}