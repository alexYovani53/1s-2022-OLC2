lexer grammar CalcLexer;


// Tokens

SYSTEM:  'system';
OUT:      'out';
PRINTLN:   'println';

NUMBER: [0-9]+;
STRING: '"'~["]*'"';

PUNTO       : '.';
PTCOMA      : ';';
DIFERENTE: '!';
MAYORIGUAL: '>=';
MENORIGUAL: '<=';
MAYOR: '>';
MENOR: '<';
MUL: '*' ;
DIV: '/' ;
ADD: '+' ;
SUB: '-' ;
LP       : '(';
RP       : ')';
L_LLAVE  : '{';
R_LLAVE  : '}';


WHITESPACE: [ \\\r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

