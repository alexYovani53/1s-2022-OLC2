parser grammar Calc;

options {
  tokenVocab = CalcLexer;
}


@header {
    import "OLC2/analizador/ast"
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
    import "OLC2/analizador/ast/control"
    import "OLC2/analizador/entorno"
    import arrayList "github.com/colegno/arraylist"
}

@members{
    type  Prueba2 struct {
        Op1 int
        Operador string
        Op2 int
    }
}


start returns [ast.Ast  ast]
  : L_LLAVE instrucciones R_LLAVE                               {$ast = ast.NewAst($instrucciones.l)}
;

prueba
    : L_LLAVE R_LLAVE ';'
;

/*
    {
            FORMA 1
            instruccion instruccion  instruccion instruccion instruccion 

             recolector +=  instruccion+


            FORMA 2 
            list return [retornoLista]
              : pivote = list listaItem   pivote.retornoLista.Add( ....... )
                                          retornoLista = pivote

              | listaItem                 retornoLista.Add(listaItem.instruccion)
            ;

            listaItem return [instruccion]
              : instruccion
            ;


            System.println("hola")   instruccion
            System.println("hola")   instruccion
            System.println("hola")   instruccion
            System.println("hola")   instruccion


            mylista := [instruccion,instruccion,instruccion,instruccion]

            for _, elemntoInstruccionConContexto = range recolector {

                $l.Add(elemntoInstruccionConContexto.instr)

            }

    }

*/

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instruccion+                                           {
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
  : consola ';'                                                 {$instr = $consola.instr}
  | declaracion ';'                                             {$instr = $declaracion.instr}
  | if_instr                                                    {$instr = $if_instr.instr}
;

consola returns [interfaces.Instruccion instr]
    : SYSTEM '.' OUT '.' PRINTLN LP expression RP             {$instr = funbasica.NewImprimir($expression.p)}
;

declaracion returns [interfaces.Instruccion instr]
    : tiposvars listides '=' expression                         {
                                                                    $instr = definicion.NewDeclaracionInicializacion($listides.lista,$tiposvars.tipo,$expression.p)
                                                                }
;


/**
    if(expresion){
      instr
      inst
      inst.....
    }

    if(expresion){

    }else {

    }


    if(expresion){

    }else if(expresion) {
      
    }else if(expresion) {
      
    }else if(expresion) {
      
    }

    
    if(expresion){

    }else if(expresion) {
      
    }else{
        inst.
        inst.
        inst.
    }

 */

if_instr returns [interfaces.Instruccion instr]
  : IF_T LP expression RP bloque                                    {$instr = control.NewIfInstruccion($expression.p,$bloque.contenido,nil,nil)}
  | IF_T LP expression RP bprincipal= bloque ELSE_T belse = bloque  {$instr = control.NewIfInstruccion($expression.p,$bprincipal.contenido,nil,$belse.contenido)}
  | IF_T LP expression RP bprincipal= bloque listaelseif            {$instr = control.NewIfInstruccion($expression.p,$bloque.contenido,$listaelseif.lista,nil)}
  | IF_T LP expression RP bprincipal= bloque listaelseif ELSE_T belse = bloque {$instr = control.NewIfInstruccion($expression.p,$bprincipal.contenido,$listaelseif.lista,$belse.contenido)}
;

listaelseif returns[*arrayList.List lista]
@init{
  $lista =  arrayList.New()
}
    : list += else_if+   {

                          pivoteLista := localctx.(*ListaelseifContext).GetList()

                          for _ , e := range pivoteLista {
                            $lista.Add(e.GetInstr())
                          }

    }
;

else_if returns[interfaces.Instruccion instr]
    : ELSE_T IF_T LP expression RP bloque        {$instr = control.NewIfInstruccion($expression.p,$bloque.contenido,nil,nil)}
;

bloque returns [*arrayList.List contenido]
  : L_LLAVE instrucciones R_LLAVE       {$contenido = $instrucciones.l}
  | L_LLAVE  R_LLAVE                    {$contenido = arrayList.New() }
;

/*

  {

    inst
    inst

  }


 */


/*
    RETORNO TIPO LISTA   [lista]


    ID , ID , ID

    : lista0 = ID , lista1 = ID
    | ID

    - MAS DE UN ID
    - UN SOLO ID Lista.Add($ID.text)


*/

listides returns [*arrayList.List lista]
  @init{
    $lista =  arrayList.New()
  }
    : sub = listides ',' ID                                     {
                                                                    $sub.lista.Add(expresion.NewIdentificador($ID.text))
                                                                    $lista = $sub.lista
                                                                }
    | ID                                                        {$lista.Add(expresion.NewIdentificador($ID.text))}
;

tiposvars returns[entorno.TipoDato tipo]
    : INTTYPE                                                   {$tipo = entorno.INTEGER}
    | STRINGTYPE                                                {$tipo = entorno.STRING}
    | FLOATTYPE                                                 {$tipo = entorno.FLOAT}
    | BOOLTYPE                                                  {$tipo = entorno.BOOLEAN}
;

expression returns[interfaces.Expresion p]
    : expr_rel                                                  {$p = $expr_rel.p}
    | expr_arit                                                 {$p = $expr_arit.p}

    // localctx.(ExpressionContext).GetExpr_arit().GetStart().GetLine()
    // localctx.(ExpressionContext).GetExpr_arit().GetStart().GetColumn()
;


/*

  {

    system.println("hola")                  INSTRUCCION      *
    system.println("hola")                  INSTRUCCION      *
    system.println("hola")                  INSTRUCCION      * 
    if(true){                               INSTRUCCION      * 
          system.println("hola")                      INSTRUCCION     *
          system.println("hola")                      INSTRUCCION     *
          for (x -> x + 1){                           INSTRUCCION     * 
              system.println("hola")                              INSTRUCCION    *
                                                                            exprexion
              system.println("hola")                              INSTRUCCION    * 
                                                                            exprexion 
          }
          system.println("hola")                      INSTRUCCION     
          system.println("hola")                      INSTRUCCION     
    }
    system.println("hola")                  INSTRUCCION      
    system.println("hola")                  INSTRUCCION      
    system.println("hola")                  INSTRUCCION       


  }



    expr
        : expr  '*' expr
        | expr  '+' expr
        | subexpr
    ;

    subexpr
        : expr
    ;



*/




expr_rel returns[interfaces.Expresion p]
    : opIz = expr_rel op=( MAYORIGUAL | MENORIGUAL | MENOR | MAYOR ) opDe = expr_rel {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | expr_arit  {$p = $expr_arit.p}
;

expr_arit returns[interfaces.Expresion p]
    : '-' opU = expression                                      {$p = expresion.NewOperacion($opU.p,"-",nil,true)}
    | opIz = expr_arit op=('*'|'/') opDe = expr_arit            {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit            {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | expr_valor                                                {$p = $expr_valor.p}
    | LP expression RP                                          {$p = $expression.p}



;

expr_valor returns[interfaces.Expresion p]
  : primitivo                                                   {$p = $primitivo.p}
;

primitivo returns[interfaces.Expresion p]
    :NUMBER                                                     {
                                                                    num,err := strconv.Atoi($NUMBER.text)
                                                                    if err!= nil{
                                                                        fmt.Println(err)
                                                                    }
                                                                    $p = expresion.NewPrimitivo (num,entorno.INTEGER)
                                                                }
    | FLOAT                                                     {
                                                                     num,err := strconv.ParseFloat($FLOAT.text,64)
                                                                     if err!= nil{
                                                                         fmt.Println(err)
                                                                     }
                                                                     $p = expresion.NewPrimitivo (num,entorno.FLOAT)
                                                                }

    | STRING                                                    {
                                                                    str:= $STRING.text[1:len($STRING.text)-1]
                                                                    $p = expresion.NewPrimitivo(str,entorno.STRING)
                                                                }
    | ID                                                        { $p = expresion.NewIdentificador($ID.text)}
    | TRUE                                                      { $p = expresion.NewPrimitivo(true,entorno.BOOLEAN)}
    | FALSE                                                     { $p = expresion.NewPrimitivo(false,entorno.BOOLEAN)}
    
;


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