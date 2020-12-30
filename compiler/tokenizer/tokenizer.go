package tokenizer

import (
        "regexp"
        "strconv"
        "unicode"
        "fmt"
)

type Tokenizer struct {
        source              []rune

        currentLineNumber   int
        currentColumnNumber int

        currentChar         rune
        peekChar            rune
}

func New(source string) *Tokenizer {
        t := new(Tokenizer)
        t.source = []rune(source)
        t.currentColumnNumber = 0
        t.peekChar = t.source[t.currentColumnNumber]

        return t
}

func (t *Tokenizer) NewToken(ttype TokenType, value string) Token {
        return Token{ttype, value, t.currentLineNumber, t.currentColumnNumber}
}

func (t *Tokenizer) readUntil(ch string) string {
        var res string
        for string(t.currentChar) != ch {
                res += string(t.currentChar)
                t.pushChar()
        }
        return res
}

func (t *Tokenizer) pushChar() {
        t.currentChar = t.peekChar
        t.currentColumnNumber += 1

        if t.currentColumnNumber >= len(t.source) {
                t.peekChar = TOK_EOF
        } else {
                t.peekChar = t.source[t.currentColumnNumber]
        }

        if string(t.currentChar) == "\n" {
                t.currentLineNumber += 1
        }
}

func isValidIdChar(first bool, char rune) bool {
        var regexs string
        if first {
                regexs = "[a-zA-Z_]"
        } else {
                regexs = "[a-zA-Z_0-9]"
        }
        match, _ := regexp.MatchString(regexs, string(char))
        return match
}

var keywords = map[string]TokenType {
        "fn": TOK_KEY_FN,
        "struct": TOK_KEY_STRUCT,
        "self": TOK_KEY_SELF,
        "if": TOK_KEY_IF,
        "else": TOK_KEY_ELSE,
        "for": TOK_KEY_FOR,
        "true": TOK_KEY_TRUE,
        "false": TOK_KEY_FALSE,
        "null": TOK_KEY_NULL,
        "return": TOK_KEY_RETURN,
}

func isKeyword(name string) bool {
        _, ok := keywords[name]
        return ok
}


func (t *Tokenizer) checkWord() (Token, bool) {
        var word string
        var tokenType TokenType
        if isValidIdChar(true, t.currentChar) {
                word += string(t.currentChar)
                for isValidIdChar(false, t.peekChar) {
                        t.pushChar()
                        word += string(t.currentChar)
                }

                if isKeyword(word) {
                        tokenType = keywords[word]
                } else {
                        tokenType = TOK_IDENTIFIER
                }
                return t.NewToken(tokenType, word), true
        }
        return t.NewToken(TOK_UNKNOWN, ""), false
}

func (t *Tokenizer) checkOperator() (Token, bool) {
        var tokenType TokenType
        switch string(t.currentChar) {
        case "+":
                tokenType = TOK_OP_ADD
        case "-":
                tokenType = TOK_OP_MINUS
        case "*":
                tokenType = TOK_OP_MUL
        case "/":
                tokenType = TOK_OP_DIV
        case "%":
                tokenType = TOK_OP_MOD
        case "^":
                tokenType = TOK_OP_POW
        case "<":
                if string(t.peekChar) == "=" {
                        t.pushChar()
                        tokenType = TOK_OP_LESS_EQUAL
                } else {
                        tokenType = TOK_OP_LESS_THAN
                }
        case ">":
                if string(t.peekChar) == "=" {
                        t.pushChar()
                        tokenType = TOK_OP_GREATER_EQUAL
                } else {
                        tokenType = TOK_OP_GREATER_THAN
                }
        case "=":
                if string(t.peekChar) == "=" {
                        t.pushChar()
                        tokenType = TOK_OP_EQUAL
                } else {
                        tokenType = TOK_OP_ASSIGN
                }
        case "!":
                if string(t.peekChar) == "=" {
                        t.pushChar()
                        tokenType = TOK_OP_NOT_EQUAL
                } else {
                        return t.NewToken(TOK_UNKNOWN, ""), false
                }
        case "&":
                tokenType = TOK_OP_AND
        case "|":
                tokenType = TOK_OP_OR
        case "¬":
                tokenType = TOK_OP_NOT
        case ",":
                tokenType = TOK_OP_COMMA
        case ".":
                tokenType = TOK_OP_DOT
        case "(":
                tokenType = TOK_OP_OPEN_PARENTHESIS
        case ")":
                tokenType = TOK_OP_CLOSE_PARENTHESIS
        case "[":
                tokenType = TOK_OP_OPEN_BRACKETS
        case "]":
                tokenType = TOK_OP_CLOSE_BRACKETS
        case "{":
                tokenType = TOK_OP_OPEN_CURLYBRACES
        case "}":
                tokenType = TOK_OP_CLOSE_CURLYBRACES
        }
        return t.NewToken(tokenType, ""), true
}

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func (t *Tokenizer) checkLiteral() (Token, bool) {
        var literal string
        var tokenType TokenType

        if string(t.currentChar) == "\"" {
                t.pushChar()
                tokenType = TOK_LIT_STRING
                literal = t.readUntil("\"")
                return t.NewToken(tokenType, literal), true
        } else if _, err := strconv.Atoi(string(t.currentChar)); err == nil {
                tokenType = TOK_LIT_INTEGER

                literal += string(t.currentChar)

                for isInt(string(t.peekChar)) {
                        t.pushChar()
                        literal += string(t.currentChar)

                }
                return t.NewToken(tokenType, literal), true
        }
        return t.NewToken(TOK_UNKNOWN, ""), false
}

func (t *Tokenizer) checkComment() bool {
        if string(t.currentChar) == "/" && string(t.peekChar) == "/" {
                t.readUntil("\n")
                return true
        }
        return false
}

func (t *Tokenizer) ReadToken() (Token, error) {
        t.pushChar()
        var token Token
        var found bool = false

        for unicode.IsSpace(t.currentChar) {
                t.pushChar()
        }

        for t.checkComment() {
                t.pushChar()
        }

        for unicode.IsSpace(t.currentChar) {
                t.pushChar()
        }

        token, found = t.checkLiteral()
        if found {
                return token, nil
        }

        token, found = t.checkWord()
        if found {
                return token, nil
        }

        token, found = t.checkOperator()
        if found {
                return token, nil
        }

        return t.NewToken(TOK_UNKNOWN, ""), fmt.Errorf("%d| Unknown char '%s'", t.currentLineNumber, string(t.currentChar))
}
