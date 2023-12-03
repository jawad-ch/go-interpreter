# Introduction

when working with Interpreter we are going to take our `Source Code` and go through the following
transformation, Thins know as lexical analysis

1- Source Code]  => 2- tokens] => 33 Abstract Syntax Tree (AST)

1) `let x = 5 + 5;`
2) To array of tokens
```
[
   LET,
   IDENTIFIER("x"),
   EQUAL_SIGN,
   INTEGER(5),
   PLUS_SIGN,
   INTEGER(5),
   SEMICOLON
]
   ```
3) AST

### Defining Tokens
