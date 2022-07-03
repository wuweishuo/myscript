parser grammar MyScriptParser;               // 定义文法的名字

options {
    tokenVocab = MyScriptLexer;
}

// RULE
statementList: (statement)+;

statement:
	assignStmt
	| ifStmt
	| funcStmt
	;

assignStmt: IDENTIFIER ASSIGN (arithmetic|relation|NUMBER|BOOL);

ifStmt:
	IF logicStmt block (ELSE (ifStmt | block))?;

block: L_CURLY statementList R_CURLY;

arithmetic:
    arithmetic op=(MUL|DIV) arithmetic
    | arithmetic op=(ADD|SUB) arithmetic
    | L_PAREN arithmetic R_PAREN
    | NUMBER
    | IDENTIFIER
    ;

logicStmt:
    relation (op=(LOGICAL_OR|LOGICAL_AND) relation)*;

relation:
    EXCLAMATION relation
    | arithmetic op=(EQUALS|NOT_EQUALS|LESS|LESS_OR_EQUALS|GREATER|GREATER_OR_EQUALS) arithmetic
    | L_PAREN (relation) R_PAREN
    | BOOL
    | IDENTIFIER
    ;

funcStmt: IDENTIFIER L_PAREN paramList R_PAREN;

paramList: param (COMMA param)*;

param: IDENTIFIER|BOOL|NUMBER;