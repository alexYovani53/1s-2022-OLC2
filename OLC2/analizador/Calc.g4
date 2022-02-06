parser grammar Calc;



options {
  tokenVocab = CalcLexer;
}

@header {
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/funbasica"
    import arrayList "github.com/colegno/arraylist"
}

@members{
    type  Prueba2 struct {
        Op1 int
        Operador string
        Op2 int
    }
}



// start  returns[interfaces.Expresion p]
//     @init{
//         fmt.Println("INICIO DE ANALISIS")
//     }
//     : expre = expression EOF {$p = $expre.p};

start returns [*arrayList.List lista]
  : L_LLAVE instrucciones R_LLAVE {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instruccion*  {
      listInt := localctx.(*InstruccionesContext).GetE()
      		for _, e := range listInt {
            $l.Add(e.GetInstr())
          }
      fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())
  }
;


// instrucciones returns [*arrayList.List l]
//   @init{
//     $l =  arrayList.New()
//   }
//   : p=instrucciones instruccion {
//              listInt := localctx.(*InstruccionesContext).GetL()
//               fmt.Printf("hf %T",listInt )
//       $l.Add($instruccion.instr)
//   }
//   | instruccion {
//       $l.Add($instruccion.instr)
//   }
// ;


instruccion returns [interfaces.Instruccion instr]
  : SYSTEM '.' OUT '.' PRINTLN LP expression RP ';' {$instr = funbasica.NewImprimir($expression.p)}
;

expression returns[interfaces.Expresion p]
    : expr_rel    {$p = $expr_rel.p}
    | expr_arit    {$p = $expr_arit.p}
;

expr_rel returns[interfaces.Expresion p]
    : opIz = expr_rel op=( MAYORIGUAL | MENORIGUAL | MENOR | MAYOR ) opDe = expr_rel {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | expr_arit  {$p = $expr_arit.p}
;

expr_arit returns[interfaces.Expresion p]
    : opIz = expr_arit op=('*'|'/') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}     
    | primitivo {$p = $primitivo.p} 
    | LP expression RP {$p = $expression.p}
;

primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.NewPrimitivo (num,interfaces.INTEGER)
       } 
    | STRING {
      str:= $STRING.text[1:len($STRING.line)-1]
      $p = expresion.NewPrimitivo(str,interfaces.STRING)}
;

// NUMBER: [0-9]+;

// DIFERENTE: '!';
// MAYORIGUAL: '>=';
// MENORIGUAL: '<=';
// MAYOR: '>';
// MENOR: '<';
// MUL: '*' ;
// DIV: '/' ;
// ADD: '+' ;
// SUB: '-' ;
// LP       : '('   ;
// RP       : ')'   ;

// WHITESPACE: [ \r\n\t]+ -> skip;

// fragment
// ESC_SEQ
//     :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
//     ;

/*
BinaryExpression:
  'or'? AndOp; //or op

AndOp:
  'and'? ComparisonOp;

ComparisonOp:
  ('>'|'<'|'>='|'<='|'=='|'~=')? ConcatOp;

ConcatOp:
  '..'? AddSubOp;

AddSubOp:
  ('+' | '-')? MultDivOp;

MultDivOp:
  ('*' | '/')? ExpOp;

ExpOp:
  '^'? expr=Expression; */