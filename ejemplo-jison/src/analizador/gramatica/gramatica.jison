%{
    const {errores} = require("../Errores");
    const {Error_} = require('../Error');
	
	const {Primitivo} = require('../ast/expresiones/Primitivo');
	const {Operacion,operador} = require('../ast/expresiones/Operacion');
	const {tipo} = require('../ast/abstract/Retorno');
	const {Imprimir} = require('../ast/funbasicas/Imprimir')

	const {Ast} = require('../ast/Ast')
%}

/**
* Definición léxica
*
*/
%lex 
%options case-sensitive 

/*%options case-sensitive*/

number  [0-9]+\b
decimal [0-9]+("."[0-9]+)?\b
string  ([\"][^"]*[\"])
string2  ([\'][^\']*[\'])



%%
\s+
{decimal}               return 'DECIMAL';
{number}                return 'NUMBER';
{string}                return 'STRING';
{string2}               return 'STRING2';
";"                     return 'PTCOMA';
"("                     return 'PARIZQ';
")"                     return 'PARDER';
"["                     return 'CORIZQ';
"]"                     return 'CORDER';
"}"						return 'LLAVECIERRE';
"{"						return 'LLAVEABRE';

"+"                     return 'MAS';
"-"                     return 'MENOS';
"*"                     return 'POR';
"/"                     return 'DIVIDIDO';

"."						return '.';

"public"				return 'PUBLIC';
"static"				return 'STATIC';
"class"					return 'CLASS';
"int"					return 'INT_TYPE';
"boolean"				return 'BOOL_TYPE';
"float"					return 'FLOAT_TYPE';
"String"				return 'STRING_TYPE';
"system"				return 'SYSTEM';
"out"					return 'OUT';
"println"				return 'PRINTLN';
"void"					return 'VOID';
"main"					return 'MAIN';
"args"					return 'ARGS';

([a-zA-Z_])[a-zA-Z0-9_ñÑ]*	return 'ID';

[ \r\t]+			    {}
\n						{}

.                           errores.push(new Error_(yylloc.first_line, yylloc.first_column, 'Lexico','Valor inesperado ' + yytext));  

<<EOF>>		                return 'EOF';

/lex



/* Asociación de operadores - precedencia*/
%left 'MAS' 'MENOS'
%left 'POR' 'DIVIDIDO'
%left  UMENOS

/* Nodo de inicio*/
%start ini   

%%  /* Definición de la gramática */

ini
	: clase EOF { return new Ast($1); }
;

clase: 
	PUBLIC CLASS ID cuerpoClase {$$ = $4;}
;

cuerpoClase: 
	LLAVEABRE seccionClase LLAVECIERRE { $$ = $2}
;

seccionClase: 
	main {$$ = $1}
;

main: 
	PUBLIC STATIC VOID MAIN PARIZQ STRING_TYPE CORIZQ CORDER ARGS PARDER bloque  {$$ = $11;}
;

bloque:  
	LLAVEABRE instrucciones LLAVECIERRE { $$ = $2}
;

instrucciones
	: instrucciones instruccion { $1.push($2); $$ = $1 }
	| instruccion				{ $$ = [$1] }
	| error ';'     			{ console.error('Este es un error sintáctico: ' + yytext + ', en la linea: ' + this._$.first_line + ', en la columna: ' + this._$.first_column); }
;

instruccion
	: SYSTEM '.' OUT '.' PRINTLN PARIZQ expresion PARDER PTCOMA {$$ = new Imprimir($7,@1.first_line, @1.first_column)}
;


expresion
	: MENOS expresion %prec UMENOS	{ $$ = new Operacion($2,operador.MENOS,null,true,@1.first_line, @1.first_column);}
	| expresion MAS expresion		{ $$ = new Operacion($1,operador.MAS,$3,false,@1.first_line, @1.first_column); }
	| expresion MENOS expresion		{ $$ = new Operacion($1,operador.MENOS,$3,false,@1.first_line, @1.first_column);  }
	| expresion POR expresion		{ $$ = new Operacion($1,operador.MULTIPLICACION,$3,false,@1.first_line, @1.first_column);  }
	| expresion DIVIDIDO expresion	{ $$ = new Operacion($1,operador.DIVISION,$3,false,@1.first_line, @1.first_column); }
	| NUMBER						{ $$ = new Primitivo(Number($1),tipo.INTEGER); }
	| DECIMAL						{ $$ = new Primitivo(Number($1),tipo.FLOAT); }
	| STRING						{ $$ = new Primitivo($1.substring(1,$1.length-1),tipo.STRING);}
	| STRING2						{ $$ = new Primitivo($1.substring(1,$1.length-1),tipo.STRING);}
	| PARIZQ expresion PARDER		{ $$ = $2;}
;
