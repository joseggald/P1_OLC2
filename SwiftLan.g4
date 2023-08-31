grammar SwiftLan;

inicio: sentencias;

sentencias: (statement)*;

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
;

reasigMatriz: Id '[' expression ']' '[' expression ']' '=' expression		# FuncionReasignMatriz
			|	Id '[' expression ']' '[' expression ']' '[' expression ']' '=' expression	 	# FuncionReasignMatriz3D
;

matrizAsign: Var Id ':' '[' '[' tiposAsign ']' ']' '=' '[' exprListMatrixDecla ']' 			# FuncionAsignarMatrizNormal
|	Var Id ':' '[' '[''['  tiposAsign ']' ']' ']' '=' '[' ( '[' exprListMatrixDecla ']' )* ']' 	# FuncionAsignarMatriz3D
|	Var Id ':' '[' '[' '['  tiposAsign ']' ']' ']' '=' '[' '[''['  tiposAsign ']' ']' ']' '(' REPEATING ':' '[' 
'[' tiposAsign ']' ']' '(' REPEATING ':' '[' tiposAsign ']' ',' '(' REPEATING ':' expression ',' COUNT ':' 
expression ')' ','  COUNT ':' expression')' ',' COUNT ':' expression ')'  # FuncionAsignarM3D
;

defStruct: STRUCT IdMayus '{' (atributosLista)*  '}' # FuncionDefStruct
;

atributosLista:	op=(Let|Var) (Id|IdMayus) (':' tiposAsign) # FuncionAtributosListTipo
| op=(Let|Var) (Id|IdMayus) ('=' expression) 	# FuncionAtributosListExp
| op=(Let|Var) (Id|IdMayus)  ':' tiposAsign '=' Id '(' exprListStruct ')'	# FuncionAtributosStruct
;

incremento: (Id|IdMayus) '+''=' expression # FuncionIncremento
;

decremento: (Id|IdMayus) '-''=' expression # FuncionDecremento
;

removeVec:tiposId '.' REMOVE '(' AT ':' expression')' # FuncionRemoveVec
;

vecReasig: tiposId'[' expression ']' '=' expression # FuncionVecReasig
;

removeLastVec:tiposId'.' REMOVELAST '('')' # FuncionRemoveLastVec
;

appendVec: tiposId '.'APEND '('expression')' # FuncionAppendVector
;

vectorAsign: Var tiposId ':' '[' tiposAsign ']' '=' '[' (exprVector)? ']' # FuncionVectorAsig
	| Var tiposId ':' '[' tiposAsign ']' '=' tiposId # FuncionVectorAsigVar
;

reasignacion: (Id|IdMayus) '=' expression # FuncionReasign
;

forstmt: FOR tiposId  'in' EnteroRange '{' sentencias '}' # FuncionForstmt
| FOR tiposId  'in' String '{' sentencias '}' # FuncionForExpstmt
|  FOR tiposId 'in' tiposId '{' sentencias '}' # FuncionForIdstmt
;

guardstmt: GUARD expression ELSE '{' sentencias '}'
;

EnteroRange: (Entero|Id|IdMayus) RANGE (Entero|Id|IdMayus)
;

whilestmt: WHILE expression '{' sentencias '}' # FuncionWhilestmt
;

switchstmt: SWITCH expression '{' (bloqueCase)* DEFAULT ':' (sentencias)? '}' # FuncionSwitchstmt
;

bloqueCase: CASE expression ':' (sentencias)?
;

callFuncstmt: Id '(' (exprListCallFunc)? ')' # FuncionCallFunc
| Id '(' (exprVector)? ')' # FuncionCallFunc2
;

funcstmt: FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '->' tiposAsign '{' sentencias '}' # FuncionDeclaFunc
	| FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '{' sentencias '}' # FuncionDeclaFunc2
;

ifstmt: ifstat ((elseifstmt)*)? (elsestmt)?
;

ifstat: IF  expression  '{' sentencias (breakstmt|retStmt)?  '}' 
;

elseifstmt: ELSE IF  expression  '{' sentencias (breakstmt|retStmt)?  '}'
;

elsestmt:ELSE '{' sentencias (breakstmt|retStmt)? '}' 
;

retStmt: RETURN expression # FuncionReturnVal
| RETURN # FuncionReturnVoid
;

breakstmt: BREAK # FuncionBreak
;

continuestmt: CONTINUE # FuncionContinue
;

asignacion: 
	tipoInit tiposId '=' structAsig # funcionAsigStruct
	| tipoInit tiposId ':' tiposAsign '=' expression # funcionAsigTipoExp
	| tipoInit tiposId  ':' tiposAsign '?' # funcionAsigTipoNil
	| tipoInit tiposId  '=' expression 	# funcionAsigExp
;

structAsig: IdMayus '(' (exprListStruct) ')'						# defStructExpression
;

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

fPrint:
	Print '(' expression? ')'	# funcionPrint
;

exprListStruct: tiposId ':' (expression|structAsig) ( ',' tiposId ':' (expression|structAsig))* # listAtibStruct
;

tiposId:Id
|IdMayus;

exprListFunc: tiposId tiposId ':' tiposAsign ( ',' tiposId tiposId ':' tiposAsign)*;
exprListFuncBajo: '_' tiposId ':' tiposAsign ( ',' '_' tiposId ':' tiposAsign)*;
exprListMatrixDecla: '[' exprVector ']' (',' '[' exprVector ']')* 	# exprListMatrix;
exprListCallFunc: tiposId ':' expression ( ',' tiposId ':' expression )* ;


exprVector: expression ( ',' expression )*
;

expression:
	'-' expression											# funcionUnariaExp
	| <assoc = right> expression '^' expression				# funcionPowExp
	| expression op = ('*' | '/' | '%') expression			# expressionMultDivMod
	| expression op = ('+' | '-') expression				# expressionSumRes
	| expression op = ('>=' | '<=' | '>' | '<') expression	# funcionCompExp
	| expression op = ('==' | '!=') expression				# funcionEqExp
	| expression '&&' expression							# funcionAndExp
	| expression '||' expression							# funcionOrExp
	| expression '?' expression ':' expression				# funcionTernaryExp
	| Nil													# nilExpression
	| Float													# floatExpression
	| Entero												# enteroExpression
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
	| tiposId '.' tiposId									# structExpression
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
//Valores
BoolVal: 'true' | 'false';
Float: [0-9]+ '.' Digit*;
Entero: [0-9]+;
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
