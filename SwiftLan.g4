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

defStruct: STRUCT Id '{' (atributosLista)*  '}' # FuncionDefStruct
;

atributosLista:	op=(Let|Var) Id (':' tiposAsign) # FuncionAtributosListTipo
| op=(Let|Var) Id ('=' expression) 	# FuncionAtributosListExp
| op=(Let|Var) Id ':' tiposAsign '=' Id '(' exprListStruct ')'	# FuncionAtributosStruct
;

incremento: Id '+''=' expression # FuncionIncremento
;

decremento: Id '-''=' expression # FuncionDecremento
;

removeVec:Id '.' REMOVE '(' AT ':' expression')' # FuncionRemoveVec
;

vecReasig: Id '[' expression ']' '=' expression # FuncionVecReasig
;

removeLastVec:Id '.' REMOVELAST '('')' # FuncionRemoveLastVec
;

appendVec: Id'.'APEND '('expression')' # FuncionAppendVector
;

vectorAsign: Var Id ':' '[' tiposAsign ']' '=' '[' (exprVector)? ']' # FuncionVectorAsig
	| Var Id ':' '[' tiposAsign ']' '=' Id # FuncionVectorAsigVar
;

reasignacion: Id '=' expression # FuncionReasign
;

forstmt: FOR Id 'in' EnteroRange '{' sentencias '}' # FuncionForstmt
| FOR Id 'in' String '{' sentencias '}' # FuncionForExpstmt
|  FOR Id 'in' Id '{' sentencias '}' # FuncionForIdstmt
;

guardstmt: GUARD expression ELSE '{' sentencias '}'
;

EnteroRange: (Entero|Id) RANGE (Entero|Id)
;

whilestmt: WHILE expression '{' sentencias '}' # FuncionWhilestmt
;

switchstmt: SWITCH expression '{' (bloqueCase)* DEFAULT ':' (sentencias)? '}' # FuncionSwitchstmt
;

bloqueCase: CASE expression ':' (sentencias)?
;

callFuncstmt: IdMinus '(' (exprListCallFunc)? ')' # FuncionCallFunc
| IdMinus '(' (exprVector)? ')' # FuncionCallFunc2
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
	tipoInit Id '=' structAsig # funcionAsigStruct
	| tipoInit Id ':' tiposAsign '=' expression # funcionAsigTipoExp
	| tipoInit Id ':' tiposAsign '?' # funcionAsigTipoNil
	| tipoInit Id '=' expression 	# funcionAsigExp
;

structAsig: Id '(' (exprListStruct) ')'						# defStructExpression
;

tipoInit: Var |
Let;

tiposAsign:
	IntDecla
	| FloatDecla
	| StringDecla
	| BoolDecla
	| CharDecla
;

fPrint:
	Print '(' expression? ')'	# funcionPrint
;

exprListStruct:Id ':' (expression|structAsig) ( ',' Id ':' (expression|structAsig))*
;

exprListFunc: Id Id ':' tiposAsign ( ',' Id Id ':' tiposAsign)*;
exprListFuncBajo: '_' Id ':' tiposAsign ( ',' '_' Id ':' tiposAsign)*;
exprListMatrixDecla: '[' exprVector ']' (',' '[' exprVector ']')* 	# exprListMatrix;
exprListCallFunc: Id ':' expression ( ',' Id ':' expression )* ;


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
	| Id '.' COUNT											# countExpression
	| Id '.' 'IsEmpty'										# emptyVecExpression
	| Id '[' expression ']'								 	# vecCallExpression
	| Id '[' expression ']' '[' expression ']'				# matrizCallExpression
	| Id '[' expression ']' '[' expression ']' '[' expression ']'	# matriz3DCallExpression
;

indexes: ( '[' expression ']')+;

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
Id: [a-zA-Z] [a-zA-Z_0-9]*;
IdMinus: [a-z] [a-zA-Z_0-9]*;
IdMayus: [A-Z] [a-zA-Z_0-9]*;
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
