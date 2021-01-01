package parser

import (
        "fmt"
        "errors"
        "strconv"

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
        p.currentToken = tokenizer.Token{tokenizer.TOK_UNKNOWN, "", -1, -1}

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
        message := fmt.Sprintf("Expected '%s', but got '%s' instead", tokenizer.TokenTypeString(tt), tokenizer.TokenTypeString(p.peekToken.Type))
        fmt.Println(p.ParseError(message, p.peekToken))
        return errors.New(p.ParseError(message, p.peekToken))
}

func (p *Parser) ParseProgram() *ast.Program {
        program := &ast.Program{
                Statements: []ast.Statement{},
        }

        for !p.peekTokenEquals(tokenizer.TOK_EOF) {
                statement := p.parseStatement()
                if statement != nil {
                        program.Statements = append(program.Statements, statement)
                }
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
        switch p.peekToken.Type {
        case tokenizer.TOK_IDENTIFIER:
                operand := p.parsePrimaryExpression()

                if p.peekTokenEquals(tokenizer.TOK_OP_ASSIGN) {
                        return p.parseAssignStatement(operand)
                } else {
                        return operand.(ast.Statement)
                }
        case tokenizer.TOK_KEY_FN:
                return p.parseFunctionDeclarationStatement()
        case tokenizer.TOK_KEY_STRUCT:
                return p.parseStructDeclarationStatement()
        case tokenizer.TOK_KEY_RETURN:
                return p.parseReturnStatement()
//      case tokenizer.TOK_KEY_IF:
//              return p.parseIfStatement()

        }
        return nil
}

func (p* Parser) parseFunctionDeclarationStatement() *ast.FunctionDeclarationStatement {
        p.consumePeek(tokenizer.TOK_KEY_FN)

        statement := new(ast.FunctionDeclarationStatement)

        p.consumePeek(tokenizer.TOK_IDENTIFIER)

        statement.Identifier = ast.Identifier{p.currentToken, p.currentToken.Value}

        statement.Parameters = p.parseFunctionParameters()

        statement.Body = p.parseBlock()

        return statement
}

func (p *Parser) parseFunctionParameters() []ast.Identifier {
        identifiers := make([]ast.Identifier, 0)

        if p.peekTokenEquals(tokenizer.TOK_OP_OPEN_CURLYBRACES) {
                return identifiers
        }

        p.consumePeek(tokenizer.TOK_OP_OPEN_PARENTHESIS)

        // Cuando no hay parámetros
        if p.peekTokenEquals(tokenizer.TOK_OP_CLOSE_PARENTHESIS) {
                p.pushToken()
                return identifiers
        }

        p.consumePeek(tokenizer.TOK_IDENTIFIER)
        identifiers = append(identifiers, ast.Identifier{p.currentToken, p.currentToken.Value})

        for p.peekTokenEquals(tokenizer.TOK_OP_COMMA) {
                p.pushToken() // TOK_OP_COMMA
                p.consumePeek(tokenizer.TOK_IDENTIFIER)
                identifiers = append(identifiers, ast.Identifier{p.currentToken, p.currentToken.Value})
        }

        p.consumePeek(tokenizer.TOK_OP_CLOSE_PARENTHESIS)

        return identifiers
}

func (p *Parser) parseStructDeclarationStatement() *ast.StructDeclarationStatement {
        p.consumePeek(tokenizer.TOK_KEY_STRUCT)

        statement := new(ast.StructDeclarationStatement)

        p.consumePeek(tokenizer.TOK_IDENTIFIER)

        statement.Identifier = ast.Identifier{p.currentToken, p.currentToken.Value}

        statement.Body = p.parseBlock()

        return statement
}

func (p *Parser) parseBlock() []ast.Statement {
        p.consumePeek(tokenizer.TOK_OP_OPEN_CURLYBRACES)

        stmts := make([]ast.Statement, 0)

        for !p.peekTokenEquals(tokenizer.TOK_OP_CLOSE_CURLYBRACES) {
                statement := p.parseStatement()
                if statement != nil {
                        stmts = append(stmts, statement)
                }
        }
        p.consumePeek(tokenizer.TOK_OP_CLOSE_CURLYBRACES)
        return stmts
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
        p.consumePeek(tokenizer.TOK_KEY_RETURN)

        stmt := new(ast.ReturnStatement)

        stmt.Token = p.currentToken
        stmt.ReturnValue = p.parseExpression()

        return stmt
}
/*
func (p *Parser) parseIfStatement() *ast.IfStatement {
        p.consumePeek(tokenizer.TOK_KEY_IF)

        stmt := new(ast.ReturnStatement)

        stmt.Token = p.currentToken
        stmt.ReturnValue = p.parseExpression()

        return stmt
}
*/
func (p *Parser) parseAssignStatement(operand ast.Expression) *ast.VarDeclarationStatement {
        statement := new(ast.VarDeclarationStatement)
        statement.Left = operand

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
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseBooleanAndExpression()}
        }
        return expr
}

func (p *Parser) parseBooleanAndExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseBooleanNegExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_AND) {
                p.pushToken()
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseBooleanNegExpression()}
        }
        return expr
}

func (p *Parser) parseBooleanNegExpression() ast.Expression {
        if p.peekTokenEquals(tokenizer.TOK_OP_NOT) {
                p.pushToken()
                op := p.currentToken
                return &ast.UnaryOperatorExpression{op, p.parseBooleanNegExpression()}
        }
        return p.parseEqualityExpression()
}

func (p *Parser) parseEqualityExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseRelationalExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_EQUAL) || p.peekTokenEquals(tokenizer.TOK_OP_NOT_EQUAL) {
                p.pushToken()
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseRelationalExpression()}
        }
        return expr
}

func (p *Parser) parseRelationalExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseArithmeticExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_GREATER_THAN) || p.peekTokenEquals(tokenizer.TOK_OP_GREATER_EQUAL) ||
               p.peekTokenEquals(tokenizer.TOK_OP_LESS_THAN) || p.peekTokenEquals(tokenizer.TOK_OP_LESS_EQUAL) {
                p.pushToken()
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseArithmeticExpression()}
        }
        return expr
}

func (p *Parser) parseArithmeticExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseMultiplicationExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_ADD) || p.peekTokenEquals(tokenizer.TOK_OP_MINUS) {
                p.pushToken()
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseMultiplicationExpression()}
        }
        return expr
}

func (p *Parser) parseMultiplicationExpression() ast.Expression {
        var expr ast.Expression

        expr = p.parseNegativeExpression()
        for p.peekTokenEquals(tokenizer.TOK_OP_MUL) || p.peekTokenEquals(tokenizer.TOK_OP_DIV) || p.peekTokenEquals(tokenizer.TOK_OP_MOD) {
                p.pushToken()
                op := p.currentToken
                expr = &ast.BinaryOperatorExpression{op, expr, p.parseNegativeExpression()}
        }
        return expr
}

func (p *Parser) parseNegativeExpression() ast.Expression {
        if p.peekTokenEquals(tokenizer.TOK_OP_MINUS) {
                p.pushToken()
                op := p.currentToken
                return &ast.UnaryOperatorExpression{op, p.parseNegativeExpression()}
        }
        return p.parsePrimaryExpression()
}

func (p *Parser) parseFunctionCall(operand ast.Expression) *ast.FunctionCallExpression {
        expr := new(ast.FunctionCallExpression)

        expr.Identifier = operand

        arguments := make([]ast.Expression, 0)

        p.consumePeek(tokenizer.TOK_OP_OPEN_PARENTHESIS)

        if p.peekTokenEquals(tokenizer.TOK_OP_CLOSE_PARENTHESIS) {
                p.pushToken()
                expr.Arguments = arguments
                return expr
        }

        expression := p.parseExpression()
        if expr != nil {
                arguments = append(arguments, expression)
        }

        for p.peekTokenEquals(tokenizer.TOK_OP_COMMA) {
                p.pushToken()
                expression := p.parseExpression()
                if expr != nil {
                        arguments = append(arguments, expression)
                }
        }

        expr.Arguments = arguments

        p.consumePeek(tokenizer.TOK_OP_CLOSE_PARENTHESIS)

        return expr
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
        var operand ast.Expression

        operand = p.parseOperand()

        for p.peekTokenEquals(tokenizer.TOK_OP_DOT) {
                dotToken := p.peekToken
                p.pushToken()

                p.consumePeek(tokenizer.TOK_IDENTIFIER)
                operand = &ast.AttributeAccessExpression{dotToken, operand, &ast.Identifier{p.currentToken, p.currentToken.Value}}
        }

        if p.peekTokenEquals(tokenizer.TOK_OP_OPEN_PARENTHESIS) {
                operand = p.parseFunctionCall(operand)
        }

        return operand
}

func (p *Parser) parseOperand() ast.Expression {
        switch p.peekToken.Type {
        case tokenizer.TOK_IDENTIFIER:
                p.pushToken()
                return &ast.Identifier{p.currentToken, p.currentToken.Value}
        case tokenizer.TOK_LIT_INTEGER:
                p.pushToken()
                i, _ := strconv.ParseInt(p.currentToken.Value, 10, 64)
                return &ast.IntLiteral{p.currentToken, i}
        case tokenizer.TOK_LIT_STRING:
                p.pushToken()
                return &ast.StringLiteral{p.currentToken, p.currentToken.Value}
        }
        return nil
}
