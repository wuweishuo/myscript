lexer grammar MyScriptLexer;               // 定义文法的名字

// Keywords
ELSE                   : 'else';
IF                     : 'if';

// Punctuation
L_PAREN                : '(';
R_PAREN                : ')';
L_CURLY                : '{';
R_CURLY                : '}';
ASSIGN                 : '=';
COMMA                  : ',';

// Arithmetic operators
MUL                    : '*';
DIV                    : '/';
ADD                    : '+';
SUB                    : '-';

// Relation operators
EQUALS                 : '==';
NOT_EQUALS             : '!=';
LESS                   : '<';
LESS_OR_EQUALS         : '<=';
GREATER                : '>';
GREATER_OR_EQUALS      : '>=';

// Logical
LOGICAL_OR             : '||';
LOGICAL_AND            : '&&';

// Unary operators
EXCLAMATION            : '!';

NUMBER                 : [0-9]+;
IDENTIFIER             : [a-zA-Z]+[_0-9]*;
BOOL                   : 'true'|'false';
WHITESPACE             : [ \t\r\n]+ -> skip ;    // 跳过空格、制表符、回车符和换行符