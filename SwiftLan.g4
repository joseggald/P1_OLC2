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
	| mmstmt
;

vectorAsign: tipoInit Id ':' '[' tiposAsign ']' '='
;

reasignacion: Id '=' expression # FuncionReasign
;

mmstmt: expression op = ('+='|'-=') expression # FuncionMM
;

forstmt: FOR Id 'in' EnteroRange '{' sentencias '}' # FuncionForstmt
;

EnteroRange: Entero RANGE Entero
;

whilestmt: WHILE expression '{' sentencias '}' # FuncionWhilestmt
;

switchstmt: SWITCH expression '{' (bloqueCase)* DEFAULT ':' (sentencias)? '}' # FuncionSwitchstmt
;

bloqueCase: CASE expression ':' (sentencias)?
;

callFuncstmt: Id '(' (exprListCallFunc)? ')' # FuncionCallFunc
| Id '(' (exprListCallFunc2)? ')' # FuncionCallFunc2
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
	tipoInit Id '=' expression 	# funcionAsigExp
	| tipoInit Id ':' tiposAsign '=' expression # funcionAsigTipoExp
	| tipoInit Id ':' tiposAsign '?' # funcionAsigTipoNil
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


exprListFunc: Id Id ':' tiposAsign ( ',' Id Id ':' tiposAsign)*;
exprListFuncBajo: '_' Id ':' tiposAsign ( ',' '_' Id ':' tiposAsign)*;
exprListCallFunc: Id ':' expression ( ',' Id ':' expression )* ;
exprListCallFunc2: expression ( ',' Id ':' expression )*;

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
	| Float												# floatExpression
	| Entero												# enteroExpression
	| BoolVal													# boolExpression
	| Id									# idExpression
	| String											# stringExpression								
	| '(' expression ')'									# expressionExpression
	| callFuncstmt									#exprCalFunc
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
RETURN:'return';
ELSE:'else';
FOR: 'for';
WHILE: 'while';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
BREAK:'break';
RANGE: '...';
CONTINUE:'continue';
//Valores
BoolVal: 'true' | 'false';
Float: [0-9]+ '.' Digit*;
Entero: [0-9]+;
Nil:'nil';
Id: [a-zA-Z] [a-zA-Z_0-9]*;
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
