grammar SwiftLan;

inicio: sentencias;

sentencias: (statement)*;

statement:
	asignacion
	| funcstmt
	| fPrint
	| ifstmt
	| retStmt
	| callFuncstmt
	| forstmt
	| switchstmt
	| whilestmt
;

forstmt: FOR Id 'in' expression '{' sentencias'}' # FuncionForstmt
;

whilestmt: WHILE expression '{' sentencias'}' # FuncionWhilestmt
;

switchstmt: SWITCH expression '{' bloqueCase '}' 
;

bloqueCase:
;

callFuncstmt: Id '(' (exprListCallFunc)? ')' # FuncionCallFunc
| Id '(' (exprListCallFunc2)? ')' # FuncionCallFunc2
;

funcstmt: FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '->' tiposAsign '{' sentenciasFunc'}' # FuncionDeclaFunc
	| FUNC Id '(' (exprListFunc|exprListFuncBajo)? ')' '{' sentenciasFunc '}' # FuncionDeclaFunc2
;

sentenciasFunc: (statement)*
;

ifstmt: ifstat ((elseifstmt)*)? (elsestmt)?
;

ifstat: IF  expression  '{' sentencias '}' 
;

elseifstmt: ELSE IF  expression  '{' sentencias '}'
;

elsestmt:ELSE '{' sentencias '}' 
;

retStmt: RETURN expression # FuncionReturnVal
| RETURN # FuncionReturnVoid
;

asignacion: 
	Var Id '=' expression 	# funcionAsigExp
	| Var Id ':' tiposAsign '=' expression # funcionAsigTipoExp
	| Var Id ':' tiposAsign '?' # funcionAsigTipoNil
;

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
	| Number												# numberExpression
	| BoolVal													# boolExpression
	| Id									# idExpression
	| String											# stringExpression									
	| '(' expression ')'									# expressionExpression
	| callFuncstmt									#exprCalFunc
;

indexes: ( '[' expression ']')+;

//Reservadas
Var:'var';
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

//Valores
BoolVal: 'true' | 'false';
Number: Int ( '.' Digit*)?;
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

