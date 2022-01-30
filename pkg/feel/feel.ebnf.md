section 9.2 S-FEEL syntax
=====================================

Just copied from the specification
  DMN-1.3-formal-21-01-01.pdf

Extended Backus–Naur form
--------------------------------------------

```
1 expression = simple_expression ;
2 arithmetic_expression = addition | subtraction | multiplication | division | exponentiation | arithmetic negation ;
3 simple_expression = arithmetic_expression | simple_value | comparison ;
4 simple_expressions = simple_expression , { "," , simple_expression } ;
5 simple_positive_unary_test = ["<"|"<="|">"|">="] , endpoint | interval ;
6 interval = ( open interval start | closed interval start ) , endpoint , ".." , endpoint , ( open interval end | closed interval end ) ;
7 open interval start = "(" | "]" ;
8 closed interval start = "[" ;
9 open interval end = ")" | "[" ;
10 closed interval end = "]" ;
11 simple_positive_unary_tests = simple_positive_unary_test , { "," , simple_positive_unary_test } ;
12 simple unary tests = simple_positive_unary_tests | "not", "(", simple_positive_unary_tests, ")" | "-";
13 endpoint = simple_value ;
14 simple_value = qualified name | simple literal ;
15 qualified name = name , { "." , name } ;
16 addition = expression , "+" , expression ;
17 subtraction = expression , "-" , expression ;
18 multiplication = expression , "*" , expression ;
19 division = expression , "/" , expression ;
20 exponentiation = expression, "**", expression ;
21 arithmetic negation = "-" , expression ;
22 name = name start , { name part | additional name symbols } ;
23 name start = name start char, { name part char } ;
24 name part = name part char , { name part char } ;
25 name start char = "?" | [A-Z] | "_" | [a-z] | [\uC0-\uD6] | [\uD8-\uF6] | [\uF8-\u2FF] | [\u370-\u37D] | [\u37F- \u1FFF] | [\u200C-\u200D] | [\u2070-\u218F] | [\u2C00-\u2FEF] | [\u3001-\uD7FF] | [\uF900-\uFDCF] | [\uFDF0-\uFFFD] | [\u1 0000-\uEFFFF] ;
26 name part char = name start char | digit | \uB7 | [\u0300-\u036F] | [\u203F-\u2040]
27 additional name symbols = "." | "/" | "-" | "’" | "+" | "*" ;
28 simple literal = numeric literal | string literal | boolean literal | date time literal ;
29 string literal = "\"", { character – ("\"" | vertical space) | string escape sequence}, "\"" ;
30 boolean literal = "true" | "false" ;
31 numeric literal = ["-"],(digits,[".",digits]|".",digits);
32 digit = [0-9] ;
33 digits = digit , {digit} ;
34 date time literal = ("date" | "time" | "duration" ) , "(" , string literal , ")" ;
35 comparison = expression , ( "=" | "!=" | "<" | "<=" | ">" | ">=" ) , expression ;
36 white space = vertical space | \u0009 | \u0020 | \u0085 | \u00A0 | \u1680 | \u180E | [\u2000-\u200B] | \u2028 | \u2029 | \u202F | \u205F | \u3000 | \uFEFF ;
37 vertical space = [\u000A-\u000D] ;
38 string escape sequence = "\\'" | "\\\"" | "\\\\" | "\\n" | "\\r" | "\\t" | "\\u", hex digit, hex digit, hex digit, hex digit ;
```
