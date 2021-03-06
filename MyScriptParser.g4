parser grammar MyScriptParser;               // 定义文法的名字

options {
    tokenVocab = MyScriptLexer;
}

// RULE
statementList: (statement)+;

statement:
	assignStmt
	| ifStmt
	;

assignStmt: IDENTIFIER ASSIGN singleExpression;

ifStmt:
	IF singleExpression block (ELSE (ifStmt | block))?;

block: L_CURLY statementList R_CURLY;

expressionSequence: singleExpression (COMMA singleExpression)*;

singleExpression:
    EXCLAMATION singleExpression                                                                        # NotExpression
    | singleExpression (MUL | DIV | MOD) singleExpression                                               # MultiplicativeExpression
    | singleExpression (ADD | SUB) singleExpression                                                     # AdditiveExpression
    | singleExpression (GREATER | LESS | GREATER_OR_EQUALS | LESS_OR_EQUALS) singleExpression           # RelationalExpression
    | singleExpression (EQUALS | NOT_EQUALS) singleExpression                                           # EqualityExpression
    | singleExpression LOGICAL_AND singleExpression                                                     # LogicalAndExpression
    | singleExpression LOGICAL_OR singleExpression                                                      # LogicalOrExpression
    | literal                                                                                           # LiteralExpression
    | IDENTIFIER L_PAREN expressionSequence R_PAREN                                                     # FuncExpression
    | IDENTIFIER                                                                                        # IdentifierExpression
    | L_PAREN singleExpression R_PAREN                                                                  # ParenthesizedExpression
    ;

literal:
    BOOL
    | NUMBER
    ;