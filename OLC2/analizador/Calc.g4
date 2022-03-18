parser grammar Calc;

options {
  tokenVocab = CalcLexer;
}


@header {

    import "OLC2/analizador/ast"
    import "OLC2/analizador/ast/interfaces"
    import "OLC2/analizador/ast/expresion"
    import "OLC2/analizador/ast/expresion/Accesos"
    import "OLC2/analizador/ast/funbasica"
    import "OLC2/analizador/ast/definicion"
    import "OLC2/analizador/ast/definicion/defobjetos"
    import "OLC2/analizador/ast/definicion/defarreglos"
    import "OLC2/analizador/ast/control"
    import "OLC2/analizador/ast/expresion2"
    import "OLC2/analizador/ast/transferenciaFlujo"
    import "OLC2/analizador/entorno"
    import "OLC2/analizador/entorno/Simbolos"
    import arrayList "github.com/colegno/arraylist"
}

@members{

}
start returns [ast.Ast  ast]
    : listaClases                      { $ast = ast.NewAst( $listaClases.lista)}
;


listaClases returns [*arrayList.List lista]
@init{
    $lista = arrayList.New()
}
    : SUBLISTA =  listaClases clases        {
                                                $SUBLISTA.lista.Add( $clases.instr)
                                                $lista =  $SUBLISTA.lista
                                            }
    | clases                                { $lista.Add( $clases.instr ) }
;


clases returns[interfaces.Instruccion instr]
    : CLASS ID cuerpoClase                              {$instr = definicion.NewDefClase($ID.text, $cuerpoClase.lista)}
;



cuerpoClase returns[*arrayList.List lista]
    : L_LLAVE contenidoClase R_LLAVE                    {$lista = $contenidoClase.lista}
    | L_LLAVE R_LLAVE                                   {$lista = arrayList.New()}
;

contenidoClase returns [*arrayList.List lista]
@init{
    $lista = arrayList.New()
}
    : SUBLISTA = contenidoClase itemClase              {
                                                            $SUBLISTA.lista.Add( $itemClase.instr )
                                                            $lista = $SUBLISTA.lista
                                                        }
    | itemClase                                         {
                                                            $lista.Add( $itemClase.instr )
                                                        }
;

itemClase returns[interfaces.Instruccion instr]
    : funcionItem                               {$instr = $funcionItem.instr}
    | declaracionIni       ';'                  {$instr = $declaracionIni.instr}
    | declaracion          ';'                  {$instr = $declaracion.instr}
;

funcionItem   returns [ interfaces.Instruccion  instr]
@init{ listaParams :=  arrayList.New() }
    : funcmain                                              {$instr =  $funcmain.instr}
    | modaccess tiposvars ID '(' ')' bloque                 { $instr = Simbolos.NewFuncion($ID.text,listaParams,$bloque.lista,entorno.VOID)}
    | modaccess tiposvars ID '('  parametros ')' bloque     { $instr = Simbolos.NewFuncion($ID.text,$parametros.lista,$bloque.lista,$tiposvars.tipo)}
;

modaccess returns [entorno.TipoModAccess  modAccess]
    : PUBLIC                                                { $modAccess = entorno.PUBLIC}
    | PRIVATE                                               { $modAccess = entorno.PRIVATE}
    |                                                       { $modAccess = entorno.PRIVATE}
;

parametros returns [*arrayList.List lista]
@init{
$lista =  arrayList.New()
}
    : sublista = parametros ','  parametro                      {
                                                                    $sublista.lista.Add( $parametro.instr )
                                                                    $lista =  $sublista.lista
                                                                 }
    | parametro                                                 {
                                                                    $lista.Add( $parametro.instr)
                                                                 }
;

parametro returns [interfaces.Instruccion instr]
    :   tiposvars ID                                            {
                                                                    listaIdes := arrayList.New()
                                                                    listaIdes.Add(expresion.NewIdentificador($ID.text))
                                                                    $instr = definicion.NewDeclaracionParametro(listaIdes, $tiposvars.tipo,false)
                                                                }
    | '*' tiposvars ID                                          {
                                                                    listaIdes := arrayList.New()
                                                                    listaIdes.Add(expresion.NewIdentificador($ID.text))
                                                                    $instr = definicion.NewDeclaracionParametro(listaIdes, $tiposvars.tipo,true)
                                                                }
;



funcmain returns[interfaces.Instruccion instr]
@init{ listaParams:= arrayList.New() }
    : PUBLIC STATIC VOIDTYPE MAIN '(' STRINGARGS ARGS '[' ']' ')' bloque
    { $instr = Simbolos.NewFuncion("main",listaParams,$bloque.lista,entorno.VOID)}
;

instrucciones returns [*arrayList.List lista]
@init{ $lista =  arrayList.New() }
  : e += instruccion+                                           {
                                                                    listInt := localctx.(*InstruccionesContext).GetE()
                                                                        for _, e := range listInt {
                                                                          $lista.Add(e.GetInstr())
                                                                        }
                                                                    fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())
                                                                }

;

instruccion returns [interfaces.Instruccion instr]
  : if_instr                                                    {$instr = $if_instr.instr}
  | consola                                  ';'                {$instr = $consola.instr}
  | declaracionIni                           ';'                {$instr = $declaracionIni.instr}
  | declaracion                              ';'                {$instr = $declaracion.instr}
  | llamada                                  ';'                {$instr = $llamada.instr}
  | retorno                                  ';'                {$instr = $retorno.instr}
  | dec_arr                                  ';'                {$instr = $dec_arr.instr}
  | dec_objeto                               ';'                {$instr = $dec_objeto.instr}
  | asignacion                               ';'                {$instr = $asignacion.instr}
;

// SECCIÓN CLASES

dec_objeto returns[interfaces.Instruccion instr]
    : ID listides '=' expresion                                {$instr = defobjetos.NewDeclararObjeto( $ID.text, $listides.lista, $expresion.expr)}
;



//SECCIÓN ARREGLOS
dec_arr returns [interfaces.Instruccion instr]
    : tiposvars  dimensiones ID '=' expresion                  {$instr = defarreglos.NewDeclaracionArray($dimensiones.tam,$ID.text,$expresion.expr,$tiposvars.tipo)}
;


dimensiones returns [int tam]
@init{ $tam = 0}
    :  tamano = dimensiones dimension                           {

                                                                    $tam = $tamano.tam + 1
                                                                 }
    |  dimension                                                {$tam = 1}
;

dimension
    : '[' ']'
;


listanchos returns[*arrayList.List lista]
@init{
    $lista = arrayList.New()
}
    :  sublist = listanchos ancho                                          {
                                                                                $sublist.lista.Add($ancho.expr)
                                                                                $lista = $sublist.lista
                                                                            }
    |  ancho                                                                {$lista.Add($ancho.expr)}
;

ancho   returns [interfaces.Expresion expr]
    :   '[' expresion ']'                                                  {$expr = $expresion.expr}
;

//SECCION ASIGNACION

asignacion returns[interfaces.Instruccion instr]
    : ID '=' expresion                                                     {$instr = definicion.NewAsignacion($ID.text, $expresion.expr)}
;


// SECCION IF

if_instr  returns [interfaces.Instruccion instr]
    : IF_TOK LP expresion RP bloque                                        {$instr = control.NewIfInstruccion($expresion.expr,$bloque.lista,nil,nil)}
    | IF_TOK LP expresion RP bprincipal = bloque ELSE  belse = bloque      {$instr = control.NewIfInstruccion($expresion.expr,$bprincipal.lista,nil,$belse.lista)}
    | IF_TOK LP expresion RP bprincipal = bloque listaelseif ELSE  belse = bloque {
        $instr = control.NewIfInstruccion($expresion.expr,$bprincipal.lista,$listaelseif.lista,$belse.lista)
    }
;

listaelseif returns [*arrayList.List lista]
@init{ $lista = arrayList.New()}
: list += else_if+ {
                                                                            listInt := localctx.(*ListaelseifContext).GetList()
                                                                            for _, e := range listInt {
                                                                                $lista.Add(e.GetInstr())
                                                                            }
    }
;

else_if returns [interfaces.Instruccion instr]
    : ELSE IF_TOK LP expresion RP bloque                               {$instr = control.NewIfInstruccion($expresion.expr,$bloque.lista,nil,nil)}
;


bloque returns [ *arrayList.List  lista]
    : L_LLAVE instrucciones R_LLAVE                                     {$lista = $instrucciones.lista }
    | L_LLAVE R_LLAVE                                                   {$lista = arrayList.New()}
;

//SECCIÓN SYSTEM . OUT . PRINTLN

consola returns [interfaces.Instruccion instr]
    : SYSTEM '.' OUT '.' PRINTLN LP expresion RP                    {$instr = funbasica.NewImprimir($expresion.expr)}
;

//SECCIÓN LLAMADA
llamada returns [interfaces.Instruccion instr, interfaces.Expresion expr]
    : ID '(' ')'                                                    {
                                                                        $instr = expresion2.NewLlamada($ID.text, arrayList.New())
                                                                        $expr = expresion2.NewLlamada($ID.text, arrayList.New())}
    | ID '(' listaExpresiones ')'                                   {
                                                                        $instr = expresion2.NewLlamada($ID.text, $listaExpresiones.lista)
                                                                        $expr = expresion2.NewLlamada($ID.text, $listaExpresiones.lista)}
;

listaExpresiones returns [*arrayList.List lista]
@init{
    $lista = arrayList.New()
}
    : LISTA = listaExpresiones ',' expresion                        {
                                                                        $LISTA.lista.Add( $expresion.expr )
                                                                        $lista =  $LISTA.lista
                                                                     }
    | expresion                                                    {
                                                                        $lista.Add( $expresion.expr )
                                                                     }
;

//SECCIÓN DECLARACIÓN INICIALIZADA
declaracionIni returns [interfaces.Instruccion instr]
    : tiposvars listides '=' expresion                              {
                                                                        $instr = definicion.NewDeclaracionInicializacion($listides.lista,$tiposvars.tipo,$expresion.expr)
                                                                     }
;

//SECCIÓN DECLARACIÓN SIN INICIAR
declaracion returns [interfaces.Instruccion instr]
    : tiposvars listides                                            {
                                                                        $instr = definicion.NewDeclaracion($listides.lista,$tiposvars.tipo)
                                                                    }
;

//SECCIÓN RETURN
retorno returns [interfaces.Instruccion instr]
    : RETURN_P                                                      { $instr = transferenciaFlujo.NewReturn(entorno.VOID,nil)}
    | RETURN_P  expresion                                          { $instr = transferenciaFlujo.NewReturn(entorno.NULL,$expresion.expr)}
;




listides returns [*arrayList.List lista]
  @init{  $lista =  arrayList.New() }
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
    | VOIDTYPE                                                  {$tipo = entorno.VOID}
;

expresion returns[interfaces.Expresion expr]
    : expr_rel                                                  {$expr = $expr_rel.expr}
    | expr_arit                                                 {$expr = $expr_arit.expr}
    | instancia                                                 {$expr = $instancia.expr}
    | arraydata                                                 {$expr = $arraydata.expr}
    | instanciaClase                                             {$expr = $instanciaClase.expr}
;

arraydata returns [interfaces.Expresion expr]
    : L_LLAVE listaExpresiones R_LLAVE                          {$expr = expresion2.NewValorArreglo($listaExpresiones.lista)}
;


instancia returns[interfaces.Expresion expr]
    : NEW_ tiposvars listanchos                                 {$expr = expresion2.NewInstanciaArreglo($tiposvars.tipo, $listanchos.lista )}
;

instanciaClase returns[interfaces.Expresion expr]
    : NEW_ ID '(' ')'                                           {$expr = expresion2.NewInstanciaObjeto($ID.text )}
;


accesoarr returns[interfaces.Expresion expr]
    : ID listanchos                                             {$expr = Accesos.NewAccessoArr($ID.text,$listanchos.lista)}
;

accesoObjeto returns[interfaces.Expresion expr]
    :  listaAccesos                                             {$expr = Accesos.NewAccesoObjeto( $listaAccesos.lista)}
;

listaAccesos returns[*arrayList.List lista]
@init{
    $lista = arrayList.New()
}
    :  SUBLISTA =  listaAccesos '.' acceso       {
                                                    $SUBLISTA.lista.Add( $acceso.expr)
                                                    $lista = $SUBLISTA.lista
                                                }
    | acceso                                    {   $lista.Add($acceso.expr)}
;

acceso  returns [interfaces.Expresion expr]
    : ID                                        { $expr = expresion.NewIdentificador($ID.text)}
    | accesoarr                                 { $expr = $accesoarr.expr}
;



expr_rel returns[interfaces.Expresion expr]
    : opIz = expr_rel op=( MAYORIGUAL | MENORIGUAL | MENOR | MAYOR ) opDe = expr_rel {$expr = expresion.NewOperacion($opIz.expr,$op.text,$opDe.expr,false)}
    | expr_arit  {$expr = $expr_arit.expr}
;

expr_arit returns[interfaces.Expresion expr]

    : '-' opU = expresion                                      {$expr = expresion.NewOperacion($opU.expr,"-",nil,true)}
    | opIz = expr_arit op=('*'|'/') opDe = expr_arit            {$expr = expresion.NewOperacion($opIz.expr,$op.text,$opDe.expr,false)}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit            {$expr = expresion.NewOperacion($opIz.expr,$op.text,$opDe.expr,false)}
    | expr_valor                                                {$expr = $expr_valor.expr}
    | LP expresion RP                                          {$expr = $expresion.expr}
;

expr_valor returns[interfaces.Expresion expr]
  : primitivo                                                   {$expr = $primitivo.expr}
  | llamada                                                     {$expr = $llamada.expr}
  | accesoarr                                                   {$expr = $accesoarr.expr}
  | accesoObjeto                                                {$expr = $accesoObjeto.expr}
;

primitivo returns[interfaces.Expresion expr]
    :NUMBER                                                     {
                                                                    num,err := strconv.Atoi($NUMBER.text)
                                                                    if err!= nil{
                                                                        fmt.Println(err)
                                                                    }
                                                                    $expr = expresion.NewPrimitivo (num,entorno.INTEGER)
                                                                }
    | FLOAT                                                     {
                                                                     num,err := strconv.ParseFloat($FLOAT.text,64)
                                                                     if err!= nil{
                                                                         fmt.Println(err)
                                                                     }
                                                                     $expr = expresion.NewPrimitivo (num,entorno.FLOAT)
                                                                }

    | STRING                                                    {
                                                                    str:= $STRING.text[1:len($STRING.text)-1]
                                                                    $expr = expresion.NewPrimitivo(str,entorno.STRING)
                                                                }

    | ID                                                        { $expr = expresion.NewIdentificador($ID.text)}

    | TRUE                                                      { $expr = expresion.NewPrimitivo(true,entorno.BOOLEAN)}
    | FALSE                                                     { $expr = expresion.NewPrimitivo(false,entorno.BOOLEAN)}

;


/*
Binaryexpresion:
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
  '^'? expr=expresion; */