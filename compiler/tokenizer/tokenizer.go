package tokenizer

type Tokenizer struct {
        source              []rune

        currentLineNumber   int
        currentColumnNumber int

        tokenArray          []Token

        currentChar         rune
        peekChar            rune
}

func New(source string) *Tokenizer {
        t := new(Tokenizer)
        t.source = []rune(string)
        t.peekChar = l.source[l.currentColumnNumber]

        t.pushChar()
        return t
}

func (t *Tokenizer) ReadToken() Token {
        token, found := t.checkWord()
        if found {
                return token
        }

        token, found = t.checkOperator()
        if found {
                return token
        }

        token, found = t.checkLiteral()
        if found {
                return token
        }

        token, found = t.checkComment()
        if found {
                return token
        }
        return Token{TOK_UNKNOWN, t.currentLineNumber, t.currentColumnNumber}
}
