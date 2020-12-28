package tokenizer

type Token uint8
const (
        // Operators
        TOK_ADD
        TOK_MINUS
        TOK_MUL
        TOK_DIV
        TOK_MOD
        TOK_POW
        TOK_LESS_THAN
        TOK_LESS_EQUAL
        TOK_GREATER_THAN
        TOK_GREATER_EQUAL
        TOK_EQUAL
        TOK_NOT_EQUAL
        TOK_AND
        TOK_OR
        TOK_NOT

        // Keywords
        TOK_FN
        TOK_STRUCT
        TOK_SELF
        TOK_IF
        TOK_ELSE
)
