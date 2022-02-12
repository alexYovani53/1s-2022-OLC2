lexer grammar CalLexer;


// Tokens



LP       : '(';
RP       : ')';
L_LLAVE  : '{';
R_LLAVE  : '}';


SYSTEM:  'system';
OUT:      'out';
PRINTLN:   'println';
INTTYPE:    'int';
FLOATTYPE:  'float';
STRINGTYPE: 'string';
BOOLTYPE:   'boolean';


PUNTO       : '.';
COMA        : ',';
PTCOMA      : ';';

AND:        '&&';
OR:         '||';
NOT:         '!' ;
IGUAL:     '=';
DIFERENTE: '!=';
MAYORIGUAL: '>=';
MENORIGUAL: '<=';
MAYOR: '>';
MENOR: '<';
MUL: '*' ;
DIV: '/' ;
ADD: '+' ;
SUB: '-' ;


NUMBER: [0-9]+;
FLOAT: [0-9]+[.][0-9]+;
STRING: '"'~["]*'"';
ID: [a-zA-Z_] [a-zA-Z0-9_]*;
TRUE: 'true';
FALSE: 'false' ESC_SEQ;


WHITESPACE: [ \r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;


AB: A B;

fragment
    A: 'a' | 'A';

fragment
    B: 'a' | 'A';

    