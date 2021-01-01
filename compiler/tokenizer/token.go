package tokenizer

type TokenType uint8
const (
        // MISC
        TOK_EOF = iota
        TOK_IDENTIFIER
        TOK_UNKNOWN

        // Literals
        TOK_LIT_INTEGER
        TOK_LIT_FLOAT
        TOK_LIT_STRING

        // Operators
        TOK_OP_ADD
        TOK_OP_MINUS
        TOK_OP_MUL
        TOK_OP_DIV
        TOK_OP_MOD
        TOK_OP_POW
        TOK_OP_LESS_THAN
        TOK_OP_LESS_EQUAL
        TOK_OP_GREATER_THAN
        TOK_OP_GREATER_EQUAL
        TOK_OP_EQUAL
        TOK_OP_NOT_EQUAL
        TOK_OP_AND
        TOK_OP_OR
        TOK_OP_NOT
        TOK_OP_ASSIGN

        TOK_OP_COMMA
        TOK_OP_DOT
        TOK_OP_OPEN_PARENTHESIS
        TOK_OP_CLOSE_PARENTHESIS
        TOK_OP_OPEN_BRACKETS
        TOK_OP_CLOSE_BRACKETS
        TOK_OP_OPEN_CURLYBRACES
        TOK_OP_CLOSE_CURLYBRACES

        // Keywords
        TOK_KEY_FN
        TOK_KEY_STRUCT
        TOK_KEY_SELF
        TOK_KEY_IF
        TOK_KEY_ELSE
        TOK_KEY_FOR
        TOK_KEY_TRUE
        TOK_KEY_FALSE
        TOK_KEY_NULL
        TOK_KEY_RETURN
)

type Token struct {
        Type         TokenType
        Value        string
        LineNumber   int
        ColumnNumber int
}

func TokenTypeString(tok TokenType) string {
        switch tok {
        case TOK_EOF:
                return "EOF"
        case TOK_IDENTIFIER:
                return "IDENTIFIER"
        case TOK_UNKNOWN:
                return "UNKNOWN"

        case TOK_LIT_INTEGER:
                return "INTEGER"
        case TOK_LIT_FLOAT:
                return "FLOAT"
        case TOK_LIT_STRING:
                return "STRING"

        case TOK_OP_ADD:
                return "+"
        case TOK_OP_MINUS:
                return "-"
        case TOK_OP_MUL:
                return "*"
        case TOK_OP_DIV:
                return "/"
        case TOK_OP_MOD:
                return "%"
        case TOK_OP_POW:
                return "^"
        case TOK_OP_LESS_THAN:
                return "<"
        case TOK_OP_LESS_EQUAL:
                return "<="
        case TOK_OP_GREATER_THAN:
                return ">"
        case TOK_OP_GREATER_EQUAL:
                return ">="
        case TOK_OP_EQUAL:
                return "=="
        case TOK_OP_NOT_EQUAL:
                return "¬="
        case TOK_OP_AND:
                return "&"
        case TOK_OP_OR:
                return "|"
        case TOK_OP_NOT:
                return "¬"
        case TOK_OP_ASSIGN:
                return "="

        case TOK_OP_COMMA:
                return ","
        case TOK_OP_DOT:
                return "."
        case TOK_OP_OPEN_PARENTHESIS:
                return "("
        case TOK_OP_CLOSE_PARENTHESIS:
                return ")"
        case TOK_OP_OPEN_BRACKETS:
                return "["
        case TOK_OP_CLOSE_BRACKETS:
                return "]"
        case TOK_OP_OPEN_CURLYBRACES:
                return "{"
        case TOK_OP_CLOSE_CURLYBRACES:
                return "}"

        case TOK_KEY_FN:
                return "fn"
        case TOK_KEY_STRUCT:
                return "struct"
        case TOK_KEY_SELF:
                return "self"
        case TOK_KEY_IF:
                return "if"
        case TOK_KEY_ELSE:
                return "else"
        case TOK_KEY_FOR:
                return "for"
        case TOK_KEY_TRUE:
                return "true"
        case TOK_KEY_FALSE:
                return "false"
        case TOK_KEY_NULL:
                return "null"
        case TOK_KEY_RETURN:
                return "return"
        }
        return "UNREACHEABLE"
}
