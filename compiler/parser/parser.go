package parser

import (
        "fmt"
        "errors"

        "github.com/TaconeoMental/Slang/compiler/tokenizer"
        "github.com/TaconeoMental/Slang/compiler/ast"
)

type Parser struct {
        t *tokenizer.Tokenizer

        currentToken tokenizer.Token
        peekToken    tokenizer.Token
}

func New(t *tokenizer.Tokenizer) *Parser {
        p := new(Parser)
        p.t = t

        p.pushToken()
        p.pushToken()
        return p
}

func (p *Parser) pushToken() {
        p.currentToken = p.peekToken
        pt, err := p.t.ReadToken()
        if err != nil {
                fmt.Println(err.Error())
                return
        } else {
                p.peekToken = pt
        }
}

func (p *Parser) peekTokenEquals(tt tokenizer.TokenType) bool {
        return p.peekToken.Type == tt
}

func (p *Parser) currentTokenEquals(tt tokenizer.TokenType) bool {
        return p.currentToken.Type == tt
}

func (p *Parser) ParseError(description string, token tokenizer.Token) string {
        return fmt.Sprintf("%d:%d| %s", token.LineNumber, token.ColumnNumber, description)
}

func (p *Parser) consumePeek(tt tokenizer.TokenType) (error) {
        if p.peekTokenEquals(tt) {
                p.pushToken()
                return nil
        }
        message := fmt.Sprintf("Expected '%s', but got '%s' instead", tokenizer.TokenTypeString(p.peekToken.Type), tokenizer.TokenTypeString(tt))
        return errors.New(p.ParseError(message, p.peekToken))
}

func (p *Parser) ParseProgram() *ast.Program {
        program := &ast.Program{
                Statements: []ast.Statement{},
        }

        for !p.currentTokenEquals(tokenizer.TOK_EOF) {
                statement := p.parseStatement()
                if statement != nil {
                        program.Statements = append(program.Statements, statement)
                }
                p.pushToken()
        }

        return program
}

func (p *Parser) parseStatement() ast.Statement {
        // ESTAMENTO        = DECL_VARIABLE
        //                  | DECL_FN
        //                  | DECL_STRUCT
        //                  | RETURN_STMT
        //                  | IF_STMT
        //                  | FOR_STMT
        //                  | EXPR;
        switch p.currentToken.Type {
        case tokenizer.TOK_IDENTIFIER:
                if p.peekTokenEquals(tokenizer.TOK_OP_ASSIGN) {
                        return p.parseAssignStatement()
                }
        }
        return nil
}

func (p *Parser) parseAssignStatement() *ast.VarDeclarationStatement {
        statement := new(ast.VarDeclarationStatement)
        statement.Left = &ast.Identifier{p.currentToken, p.currentToken.Value}

        p.pushToken() // Signo =

        statement.Value = p.parseExpression()

        return statement
}

func (p *Parser) parseExpression() ast.Expression {
        return p.parseBooleanOrExpression()
}

func (p *Parser) parseBooleanOrExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseBooleanAndExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_OR) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseBooleanAndExpression()}
        }
        return expr
}

func (p *Parser) parseBooleanAndExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseBooleanNegExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_AND) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseBooleanNegExpression()}
        }
        return expr
}

func (p *Parser) parseBooleanNegExpression() ast.Expression {
        if p.peekTokenEquals(tokenizer.TOK_OP_NOT) {
                p.pushToken()
                return &ast.UnaryOperatorExpression{p.currentToken, p.parseBooleanNegExpression()}
        }
        return p.parseEqualityExpression()
}

func (p *Parser) parseEqualityExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseRelationalExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_EQUAL) || p.peekTokenEquals(tokenizer.TOK_OP_NOT_EQUAL) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseRelationalExpression()}
        }
        return expr
}

func (p *Parser) parseRelationalExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseArithmeticExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_GREATER_THAN) || p.peekTokenEquals(tokenizer.TOK_OP_GREATER_EQUAL) ||
               p.peekTokenEquals(tokenizer.TOK_OP_LESS_THAN) || p.peekTokenEquals(tokenizer.TOK_OP_LESS_EQUAL) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseArithmeticExpression()}
        }
        return expr
}

func (p *Parser) parseArithmeticExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseMultiplicationExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_ADD) || p.peekTokenEquals(tokenizer.TOK_OP_MINUS) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseMultiplicationExpression()}
        }
        return expr
}

func (p *Parser) parseMultiplicationExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseNegativeExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_MUL) || p.peekTokenEquals(tokenizer.TOK_OP_DIV) || p.peekTokenEquals(tokenizer.TOK_OP_MOD) {
                p.pushToken()
                expr = &ast.BinaryOperatorExpression{p.currentToken, expr, p.parseNegativeExpression()}
        }
        return expr
}

func (p *Parser) parseNegativeExpression() ast.Expression {
        if p.peekTokenEquals(tokenizer.TOK_OP_MINUS) {
                p.pushToken()
                return &ast.UnaryOperatorExpression{p.currentToken, p.parseNegativeExpression()}
        }
        //return p.parsePrimaryExpression()
        return p.parseOperand()
}

func (p *Parser) parseOperand() ast.Expression {
        switch p.currentToken.Type {
        case tokenizer.TOK_IDENTIFIER:
                return &ast.Identifier{p.currentToken, p.currentToken.Value}
        }
        return nil
}
