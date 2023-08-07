grammar SwiftLan;

inicio: sentencias;

sentencias: (statement)*;

statement:
	asignacion
	| llamadaPrint;

asignacion: Declarar Id '=' expression;

llamadaPrint:
	Print '(' expression? ')'	# printLlamadaPrint
;

expression:
	'-' expression											# unaryMinusExpression
	| <assoc = right> expression '^' expression				# powerExpression
	| expression op = ('*' | '/' | '%') expression			# multExpression
	| expression op = ('+' | '-') expression				# addExpression
	| expression op = ('>=' | '<=' | '>' | '<') expression	# compExpression
	| expression op = ('==' | '!=') expression				# eqExpression
	| expression '&&' expression							# andExpression
	| expression '||' expression							# orExpression
	| expression '?' expression ':' expression				# ternaryExpression
	| Number												# numberExpression
	| BoolVal													# boolExpression
	| llamadaPrint									# llamadaFuncionExpression
	| Id indexes?									# idExpression
	| String											# stringExpression
	| '(' expression ')'									# expressionExpression;

indexes: ( '[' expression ']')+;

//Reservadas
Declarar:'var';
Print: 'print';
IntDecla: 'Int';
FloatDecla: 'Float';
BoolDecla: 'Bool';
StringDecla: 'String';
CharDecla: 'Character'; 

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
Add: '+';
Minus: '-';
Mult: '*';
Div: '/';
Mod: '%';
MayorQue: '>';
MenorQue: '<';

// skip
Space: [ \t\r\n\u000C] -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

