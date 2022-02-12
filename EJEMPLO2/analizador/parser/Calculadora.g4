parser grammar Calculadora;

options {
    tokenVocab = CalLexer;
}


/*
   expr
        : expr  '*' expr
        | expr  '+' expr
        | subexpr
    ;

    subexpr
        : expr produccion1  expr
    ;

 */

inicio returns [int salida]
@init{
    $salida = 10
}
    : expresion   
;

expresion returns [ int salida]
    : opIz = expresion op = (ADD | SUB ) opDe = expresion 
    | opIz = expresion op = (MUL | DIV ) opDe = expresion
    | primitivo
;

primitivo returns [int salida]
    : NUMBER
;