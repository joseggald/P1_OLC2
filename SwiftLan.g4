grammar SwiftLan;

inicio: sentencias;

sentencias: (statement)*;

//Sentencias
statement:
	asignacion
	| reasignacion
	| vectorAsign
	| funcstmt
	| fPrint
	| ifstmt
	| callFuncstmt
	| forstmt
	| switchstmt
	| whilestmt
	| retStmt
	| breakstmt
	| continuestmt
	| appendVec
	| removeVec
	| removeLastVec
	| vecReasig
	| guardstmt
	| incremento
	| decremento
	| matrizAsign
	| reasigMatriz
	| defStruct
	| structObj
	| callFuncStruct
	| reasigVarStruct
;

//Declaracion varaibles, vectores, matrices,funciones y structs.
asignacion: tipoInit tiposId '=' structAsig # funcionAsigStruct
	| op=(Let|Var) tiposId ':' tiposAsign '=' expression # funcionAsigTipoExp
	| op=(Let|Var) tiposId  ':' tiposAsign '?' # funcionAsigTipoNil
	| op=(Let|Var) tiposId  '=' expression 	# funcionAsigExp
;

matrizAsign: Var Id ':' '[' '[' tiposAsign ']' ']' '=' '[' exprListMatrixDecla ']' 			# FuncionAsignarMatrizNormal
|	Var Id ':' '[' '[''['  tiposAsign ']' ']' ']' '=' '[' ( '[' exprListMatrixDecla ']' )* ']' 	# FuncionAsignarMatriz3D
|	Var Id ':' '[' '[' '['  tiposAsign ']' ']' ']' '=' '[' '[''['  tiposAsign ']' ']' ']' '(' REPEATING ':' '[' 
'[' tiposAsign ']' ']' '(' REPEATING ':' '[' tiposAsign ']' ',' '(' REPEATING ':' expression ',' COUNT ':' 
expression ')' ','  COUNT ':' expression')' ',' COUNT ':' expression ')'  # FuncionAsignarM3D;

defStruct: STRUCT IdMayus '{' (atributosLista| atributosLista2)* (funcStructs)* '}' # FuncionDefStruct;

vectorAsign: Var tiposId ':' '[' tiposAsign ']' '=' '[' (exprVector)? ']' # FuncionVectorAsig
	| Var tiposId ':' '[' tiposAsign ']' '=' tiposId # FuncionVectorAsigVar
	| Var tiposId '=' '[' IdMayus ']' '(' ')' # FuncionVectorAsigVarStruct
	| Var tiposId '=' '[' structAsig (',' structAsig)* ']' # FuncionVectorAsigVarObjs;

funcstmt: FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '->' tiposAsign '{' sentencias '}' # FuncionDeclaFunc
	| FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '{' sentencias '}' # FuncionDeclaFunc2;

//Manejo de Variables
reasignacion: tiposId '=' expression # FuncionReasign;
incremento: tiposId '+''=' expression # FuncionIncremento;
decremento: tiposId '-''=' expression # FuncionDecremento;
tipoInit: Var |
Let;
tiposAsign:
	IntDecla
	| FloatDecla
	| StringDecla
	| BoolDecla
	| CharDecla
	| IdMayus
;

//Llamada de funciones
callFuncstmt: Id '(' (exprListCallFunc)? ')' # FuncionCallFunc
| Id '(' (exprVector)? ')' # FuncionCallFunc2
;
callFuncStruct: tiposId '.' tiposId'('')' # FuncionCallFuncStrcut;
reasigVarStruct: SELF '.' tiposId tipoIgual expression	# FuncionSelfReasig;
tipoIgual:'='
|'+''='
|'-''='
;

funcStructs: FUNC tiposId '('')' '{' sentencias '}' 	# FuncionCrearFunc
| MUTANT FUNC tiposId '('')' '{' sentencias '}'			# FuncionCrearFuncMut
;

//Lista expresiones
exprListStruct: tiposId ':' expr_struct ( ',' tiposId ':' expr_struct )* # listAtibStruct;
exprListFunc: dataFuncTipo (','dataFuncTipo)*	# listaFuncConTipo; 
dataFuncTipo:tiposId tiposId ':'(INOUT)? expr_llave? tiposAsign expr_llave2? # FuncionDataFuncTipo;
exprListFuncBajo: dataFuncBajo (',' dataFuncBajo)*	# listaFuncConBarra;
dataFuncBajo:'_' tiposId ':'(INOUT)? expr_llave? tiposAsign expr_llave2? # FuncionDataFuncBajo;
expr_llave:('[')*;
expr_llave2:(']')*;
exprListMatrixDecla: '[' exprVector ']' (',' '[' exprVector ']')* 	# exprListMatrix;
exprListCallFunc: tiposId ':' expression ( ',' tiposId ':' expression )* ;
exprVector: expression ( ',' expression )*;
expr_struct:(expression |structAsig) # retornoExpStruct
;


//Funciones matrices y vectores
reasigMatriz: Id '[' expression ']' '[' expression ']' '=' expression		# FuncionReasignMatriz
			|	Id '[' expression ']' '[' expression ']' '[' expression ']' '=' expression	 	# FuncionReasignMatriz3D;

removeVec:tiposId '.' REMOVE '(' AT ':' expression')' # FuncionRemoveVec;

vecReasig: tiposId'[' expression ']' '=' expression # FuncionVecReasig;

removeLastVec:tiposId'.' REMOVELAST '('')' # FuncionRemoveLastVec;

appendVec: tiposId '.' APEND '('expression')' # FuncionAppendVector
| tiposId '.' APEND '('structAsig')' # FuncionAppendVectorStr;

//Struct
structAsig: IdMayus '(' (exprListStruct) ')'						# defStructExpression;
structObj: tiposId ('.' tiposId)+ '=' expression 			# FuncionReasigObj;
atributosLista2: op=(Let|Var) tiposId  ':' IdMayus	# FuncionAtributosStruct;
atributosLista:	op=(Let|Var) tiposId (':' tiposAsign) # FuncionAtributosListTipo
| op=(Let|Var) tiposId ('=' expression) 	# FuncionAtributosListExp;

//Ciclos
forstmt: FOR (tiposId|'_')  'in' EnteroRange (expression)'{' sentencias '}' # FuncionForstmt
| FOR tiposId  'in' String '{' sentencias '}' # FuncionForExpstmt
|  FOR tiposId 'in' tiposId '{' sentencias '}' # FuncionForIdstmt;

EnteroRange: (Number|Id|IdMayus) RANGE ;
EnteroRange2: (Number|Id|IdMayus) RANGE (Number);
whilestmt: WHILE expression '{' sentencias '}' # FuncionWhilestmt;
switchstmt: SWITCH expression '{' (bloqueCase)* DEFAULT ':' (sentencias)? '}' # FuncionSwitchstmt;
bloqueCase: CASE expression ':' (sentencias)?;

//Condicionales
ifstmt: ifstat ((elseifstmt)*)? (elsestmt)?;
ifstat: IF  expression  '{' sentencias (breakstmt|retStmt)?  '}' ;
elseifstmt: ELSE IF  expression  '{' sentencias (breakstmt|retStmt)?  '}';
elsestmt:ELSE '{' sentencias (breakstmt|retStmt)? '}' ;
guardstmt: GUARD expression ELSE '{' sentencias '}';

//Funciones de retorno o queibre
retStmt: RETURN expression # FuncionReturnVal
| RETURN  # FuncionReturnVoid;
breakstmt: BREAK # FuncionBreak;
continuestmt: CONTINUE # FuncionContinue;


//Area de expresiones y funciones ya definidas
fPrint:
	Print '(' (expression|concaExp) ')'	# funcionPrint
;

concaExp: expression  (',' expression)+							# concatenarExpression;

tiposId:Id
|IdMayus;

expression: '!' expression										# funcionNot
	| '-' expression											# funcionUnariaExp
	| <assoc = right> expression '^' expression				# funcionPowExp
	| expression op = ('*' | '/' | '%') expression			# expressionMultDivMod
	| expression op = ('+' | '-') expression				# expressionSumRes
	| expression op = ('>=' | '<=' | '>' | '<') expression	# funcionCompExp
	| expression op = ('==' | '!=') expression				# funcionEqExp
	| expression '&&' expression							# funcionAndExp
	| expression '||' expression							# funcionOrExp
	| expression '?' expression ':' expression				# funcionTernaryExp
	| Nil													# nilExpression
	| Number												# numberExpression
	| BoolVal												# boolExpression
	| Id													# idExpression
	| String												# stringExpression								
	| '(' expression ')'									# expressionExpression
	| callFuncstmt											# exprCalFunc
	| tiposId '.' COUNT											# countExpression
	| tiposId '.' 'IsEmpty'										# emptyVecExpression
	| tiposId '[' expression ']'								 	# vecCallExpression
	| tiposId '[' expression ']' '[' expression ']'				# matrizCallExpression
	| tiposId '[' expression ']' '[' expression ']' '[' expression ']'	# matriz3DCallExpression
	| tiposId ('.' tiposId)+								# callVarStructExpression
	| tiposAsign '(' expression ')'										# funcionesEmbeExpression
	| '&'tiposId											# callArray
	| '&'tiposId '[' expression ']'	'[' expression ']'		# callMatriz
	| SELF '.' tiposId 										# callSelfExp
	| tiposId '[' expression ']''.' tiposId								 	# vecStrCallExpression
;

//Reservadas
Var:'var';
Let:'let';
Print: 'print';
IntDecla: 'Int';
FloatDecla: 'Float';
BoolDecla: 'Bool';
StringDecla: 'String';
CharDecla: 'Character'; 
IF:'if';
FUNC:'func';
INOUT:'inout';
STRUCT:'struct';
RETURN:'return';
ELSE:'else';
FOR: 'for';
WHILE: 'while';
SWITCH: 'switch';
APEND:'append';
CASE: 'case';
DEFAULT: 'default';
BREAK:'break';
GUARD:'guard';
COUNT:'count';
REMOVE:'remove';
REMOVELAST:'removeLast';
AT:'at';
RANGE: '...';
CONTINUE:'continue';
REPEATING:'repeating';
MUTANT:'mutating';
SELF:'self';
//Valores
BoolVal: 'true' | 'false';
Number: Int ( '.' Digit*)?;
entero: Int (Int*);
Nil:'nil';
Id: [a-z][a-zA-Z_0-9]*;
IdMayus: [A-Z][a-zA-Z_0-9]*;
String:
	["] (~["\r\n\\] | '\\' ~[\r\n])* ["]
	| ['] ( ~['\r\n\\] | '\\' ~[\r\n])* [']
;

fragment Int: [1-9] Digit* | '0';

fragment Digit: [0-9];

Or: '||';
And: '&&';
Equals: '==';
NotEquals: '!=';
MayQEquals: '>=';
MinQEquals: '<=';
Power: '^';
Suma: '+';
Resta: '-';
Mult: '*';
Div: '/';
Mod: '%';
MayorQue: '>';
MenorQue: '<';

// skip
Space: [ \t\r\n\u000C] -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;
