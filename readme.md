# Grammar

```
PROGRAM               --> block
block                 --> declarations statement_list
function              --> DEFINE ID LPAREN formal_parameters_list? RPAREN LCURLY block RCURLY
formal_parameter_list --> formal_parameters | formal_parameters SEMI_COLON formal_parameter_list
formal_parameters     --> ID (COMMA ID)* COLON type_spec
function_call         --> ID LPAREN (expression (COMMA expression)*)? RPAREN
conditional_statement --> IF logical_statement LCURLY block RCURLY (ELIF logical_statement LCURLY block RCURLY)*
                          (ELSE LCURLY block RCURLY)  {0,1}
loop                  --> LOOP FROM expression TO expression WITH variable LCURLY block RCURLY
declarations          --> LET (variable_declaration SEMI)+ | function* | blank
variable_declaration  --> ID (COMMA ID)* COLON var_type
var_type              --> INTEGER | FLOAT
statement_list        --> statement SEMI_COLON | statement SEMI_COLON statement_list
statement             --> assignment_statement | function_call | conditional_statement | blank
comparison            --> expression comparator expression
assignment_statement  --> variable ASSIGN expression
logical_statement     --> NOT* (comparator ((AND | OR) comparator)*)
variable              --> ID
blank                 -->
expression            --> term ((PLUS | MINUS) term)*
term                  --> factor ((MUL | DIV) factor)*
factor                --> ((PLUS | MINUS) factor) | INTEGER | FLOAT | STRING | LPAREN expression RPAREN | variable
comparator            --> > | < | >= | <= | ==
LPAREN                --> (
RPAREN                --> )
LCURLY                --> {
RCURLY                --> }
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
}`)

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
}`)

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

    loop from 2 to n // 2 using a {
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
}`)

result := interpreter.Interpret()
fmt.Println(result)
```
