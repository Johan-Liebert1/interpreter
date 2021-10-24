# Overview

A small programming language and interpreter. Use it in shell mode or pass a file to interpret.

# Grammar

```
PROGRAM               --> block
block                 --> declarations statement_list
function              --> DEFINE ID LPAREN formal_parameters_list? RPAREN LCURLY block (RETURN expression)? RCURLY
formal_parameter_list --> formal_parameters | formal_parameters SEMI_COLON formal_parameter_list
formal_parameters     --> ID (COMMA ID)* COLON type_spec
function_call         --> ID LPAREN (expression (COMMA expression)*)? RPAREN
conditional_statement --> IF logical_statement LCURLY block RCURLY (ELIF logical_statement LCURLY block RCURLY)*
                          (ELSE LCURLY block RCURLY)  {0,1}
loop                  --> LOOP FROM expression TO expression WITH variable LCURLY block RCURLY
declarations          --> LET (variable_declaration SEMI)+ | function* | blank
variable_declaration  --> ID (COMMA ID)* COLON var_type
var_type              --> INTEGER | FLOAT | STRING
statement_list        --> statement SEMI_COLON | statement SEMI_COLON statement_list
statement             --> assignment_statement | function_call | conditional_statement | blank
comparison            --> expression comparator expression
assignment_statement  --> variable ASSIGN expression
logical_statement     --> NOT* (comparator ((AND | OR) comparator)*)
variable              --> ID
blank                 -->
comment               --> HASH (UNICODE_CHARACTER)* \n
expression            --> term ((PLUS | MINUS) term)*
term                  --> factor ((MUL | DIV | EXPONENT) factor)*
factor                --> ((PLUS | MINUS) factor) | INTEGER | FLOAT | STRING | BOOLEAN | LPAREN expression RPAREN | variable
comparator            --> > | < | >= | <= | == | !=
LPAREN                --> (
RPAREN                --> )
LCURLY                --> {
RCURLY                --> }
HASH                  --> #
```

# Syntax

### Variable Declaration

```
let varName1: int;
let varName2: float;
let varName3: bool;
let varName4: str;

let varName1, varName2, varName3: int;
```

### Variable Definition

```
varName1 := 20;
varName2 := 34.85;
varName3 := "This is a string";
```

### Comment

```
# this is a comment
```

### Arithmetic Operations on Variables

```
Addition : a + b;
Subtraction : a - b;
Multiplication : a * b;
Float Division : a / b;
Integer Division : a // b;
Modulo : a % b;
Exponent : a ^ b;
```

### Logical Operations on Variables

```
Greater Than, Greater than equal to : a > b, a >= b;
Less Than, Less than equal to : a <> b, a <>= b;
Equal to, Not Equal to : ==, !=
Logical AND, OR and NOT : and, or, not

Ex: not(a > b and (b != c or d <= 3))
```

### Printing to stdout

```
let variable: int;
variable := 32;

output("This will print to stdout");
output(1 + 34);
output(variable);
```

### Loop

```
loop from 1 to 10 using iterator {
    # do stuff 10 times
}
```

### Conditionals

```
let var : int;
var := 6;

if var > 0 and var <= 5 {
    output("Between 0 and 5");
} elif var > 5 and var <= 10 {
    output("Between 5 and 10");
} else {
    output("Greater than 10")
}

```

### Functions

```
define add(a, b : int) {
    return a + b;
}

var c : int;

# Calling
c := add(1, 2);
```

# FizzBuzz

```golang
interpreter := Interpreter{}
interpreter.InitConcrete()

interpreter.Init(`
loop from 1 to 50 using i {
    if i % 3 == 0 and i % 5 == 0 {
        output("FizzBuzz")
    } elif i % 3 == 0 {
        output("Fizz")
    } elif i % 5 == 0 {
        output("Buzz")
    } else {
        output(i)
    }
}`, false)

result := interpreter.Interpret()
fmt.Println(result)
```

# Fibonacci

```go
interpreter := Interpreter{}
interpreter.InitConcrete()

interpreter.Init(`
let first, second, third : int;

first := 1;
second := 1;

third := first + second;

output(first);
output(second);

loop from 1 to 20 using i {
    output(third);
    first := second;
    second := third;
    third := first + second;
}`, false)

result := interpreter.Interpret()
fmt.Println(result)
```

# Prime Numbers up to a range

```go
interpreter := Interpreter{}
interpreter.InitConcrete()

interpreter.Init(`
let p : bool;

define isPrime(n : int) {
    let value : int;
    value := true;

    loop from 2 to n ^ 0.5 using a {
        if n % a == 0 {
            value := false;
        }
    }

    return value;
}

loop from 1 to 50 using i {
    p := isPrime(i);

    if p == true {
        output(i, " is prime = ", p);
    }
}`, false)

result := interpreter.Interpret()
fmt.Println(result)
```
