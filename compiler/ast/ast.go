package ast

import (
        "strings"
        "fmt"

        "github.com/TaconeoMental/Slang/compiler/tokenizer"
)


// Siempre uso este sistema en todos mis parsers, pero funciona así que no me
// importa :)
var (
        astIndentation = 4
        astLast = "└" + strings.Repeat("─", astIndentation - 1)
        astMiddle = "├" + strings.Repeat("─", astIndentation - 1)
        astLine = "│" + strings.Repeat(" ", astIndentation - 1)
        astSpace = strings.Repeat(" ", astIndentation)
)

func indentation(last bool) string {
        if last {
                fmt.Print(astLast)
                return astSpace
        } else {
                fmt.Print(astMiddle)
                return astLine
        }
}

func concatps(strings ...*string) *string {
    end := ""
    for _, str := range strings {
        end += *str
    }
    return &end
}
// Acá empezamos con los nodos de AST
type Node interface {
        GetTokenType() tokenizer.TokenType
        PrintTree(*string, bool)
}

type Statement interface {
        Node
        statement()
}

type Expression interface {
        Node
        expression()
}

type Program struct {
        Statements []Statement
}

// La raiz del AST
func (p *Program) GetTokenType() tokenizer.TokenType {
        if len(p.Statements) == 0 {
                return tokenizer.TOK_EOF
        }
        return p.Statements[0].GetTokenType()
}

func (p *Program) PrintTree(indent *string, last bool) {
        fmt.Println("Program")
        for index, element := range p.Statements {
                element.PrintTree(indent, index == len(p.Statements) - 1)
        }
}

// Return de una función
type ReturnStatement struct {
        Token       tokenizer.Token // TOK_KEY_RETURN
        ReturnValue Expression
}

func (rs *ReturnStatement) statement() {}

func (rs *ReturnStatement) GetTokenType() tokenizer.TokenType {
        return rs.Token.Type
}

func (rs *ReturnStatement) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("ReturnStatement")
        fmt.Println(*indent + astMiddle + "ReturnValue")
        rs.ReturnValue.PrintTree(concatps(indent, &astLine), true)
}

// Declaración de una variable
type VarDeclarationStatement struct {
        Token tokenizer.Token
        Left  Expression
        Value Expression
}

func (vd *VarDeclarationStatement) statement() {}
func (vd *VarDeclarationStatement) GetTokenType() tokenizer.TokenType {
        return vd.Token.Type
}

func (vd *VarDeclarationStatement) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("VarDeclarationStatement")
        fmt.Println(*indent + astMiddle + "Left")
        vd.Left.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astLast + "Value")
        vd.Value.PrintTree(concatps(indent, &astSpace), true)
}


// Identificador
type Identifier struct {
        Token tokenizer.Token
        Name  string
}

func (id *Identifier) expression() {}

func (id *Identifier) GetTokenType() tokenizer.TokenType {
        return id.Token.Type
}

func (id *Identifier) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("Identifier")
        fmt.Println(*indent + astLast + id.Name)
}

// Operación binaria
type BinaryOperatorExpression struct {
        Token tokenizer.Token
        Left  Expression
        Right Expression
}

func (bo *BinaryOperatorExpression) expression() {}
func (bo *BinaryOperatorExpression) GetTokenType() tokenizer.TokenType {
        return bo.Token.Type
}

func (bo *BinaryOperatorExpression) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("BinaryOperatorExpression")
        fmt.Println(*indent + astMiddle + "Left")
        bo.Left.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astMiddle + "Right")
        bo.Right.PrintTree(concatps(indent, &astSpace), true)
}

// Operación unaria
type UnaryOperatorExpression struct {
        Token tokenizer.Token
        Expression Expression
}

func (uo *UnaryOperatorExpression) expression() {}
func (uo *UnaryOperatorExpression) GetTokenType() tokenizer.TokenType {
        return uo.Token.Type
}

func (uo *UnaryOperatorExpression) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("UnaryOperatorExpression")

        fmt.Println(*indent + astMiddle + "Expression")
        uo.Expression.PrintTree(concatps(indent, &astSpace), true)
}
