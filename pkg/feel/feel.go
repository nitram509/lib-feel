package main

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/repr"
)

// A custom lexer for INI files. This illustrates a relatively complex Regexp lexer, as well
// as use of the Unquote filter, which unquotes string tokens.
//var iniLexer = lexer.Must(lexer.Regexp(
//	`(?m)` +
//		`(\s+)` +
//		`|(^[#;].*$)` +
//		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
//		`|(?P<String>"(?:\\.|[^"])*")` +
//		`|(?P<Float>\d+(?:\.\d+)?)` +
//		`|(?P<Punct>[][=])`,
//))

// Addition
// 16 addition = expression , "+" , expression ;
type Addition struct {
	//LHS *SimpleValue `@@`
	Op  string       `@("+")`
	RHS *SimpleValue `@@ `
}

type Operator int

const (
	OpMul Operator = iota
	OpDiv
	OpAdd
	OpSub
	OpExp
)

var operatorMap = map[string]Operator{"+": OpAdd, "-": OpSub, "*": OpMul, "/": OpDiv, "**": OpExp}

func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

type Value struct {
	Int           *int64      `   @Int`
	Float         *float64    `|  @Float`
	Variable      *string     `| @Ident`
	Subexpression *Expression `| "(" @@ ")"`
}

type Factor struct {
	Base     *Value `@@`
	Exponent *Value `( "^" @@ )?`
}

type OpFactor struct {
	Operator Operator `@("*" | "/")`
	Factor   *Factor  `@@`
}

type Term struct {
	Left  *Factor     `@@`
	Right []*OpFactor `@@*`
}

type OpTerm struct {
	Operator Operator `@("+" | "-")`
	Term     *Term    `@@`
}

// ArithmeticExpression
// 2 arithmetic_expression = addition | subtraction | multiplication | division | exponentiation | arithmetic negation ;
type ArithmeticExpression struct {
	Left  *Term     `@@`
	Right []*OpTerm `@@*`
}

// SimpleValue
// 14 simple_value = qualified name | simple literal ;
type SimpleValue struct {
	SimpleLiteral *SimpleLiteral `  @@`
}

type Comparison struct {
	LHS *Expression `  @@`
	Op  string      `  @( "=" | "!=" | "<" | "<=" | ">" | ">=" )`
	RHS *Expression `  @@`
}

// SimpleExpression
// 3 simple expression = arithmetic expression | simple value | comparison ;
type SimpleExpression struct {
	ArithmeticExpression *ArithmeticExpression `  @@`
	SimpleValue          *SimpleValue          `| @@`
	Comparison           *Comparison           `| @@`
}

type Expression struct {
	SimpleExpression *SimpleExpression `  @@`
}

// SimpleLiteral
//28 simple literal = numeric literal | string literal | boolean literal | date time literal ;
//29 string literal = """, { character – (""" | vertical space) | string escape sequence}, """ ;
//30 boolean literal = "true" | "false" ;
//31 numeric literal=["-"],(digits,[".",digits]|".",digits);
//32 digit = [0-9] ;
//33 digits = digit , {digit} ;
//34 date time literal = ("date" | "time" | "duration" ) , "(" , string literal , ")" ;
type SimpleLiteral struct {
	String   *string  `  @String`
	IntVal   *int64   ` | @Int`
	FloatVal *float64 ` | @Float`
	Boolean  *bool    ` | @("TRUE" | "FALSE" | "true" | "false" )`
}

func main() {
	parser, err := participle.Build(&Expression{},
		//participle.Lexer(iniLexer),
		participle.Unquote("String"),
	)
	if err != nil {
		panic(err)
	}
	ast := &Expression{}
	err = parser.ParseString("", "10 + 10", ast)
	if err != nil {
		panic(err)
	}
	repr.Println(ast, repr.Indent("  "), repr.OmitEmpty(true))
}
