package ast

import (
        "strings"
        "fmt"
        "strconv"

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

// Concatenar pointers a string y devolver un pointer a string
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
                *indent = ""
                element.PrintTree(indent, index == (len(p.Statements) - 1))
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
        fmt.Println(*indent + astLast + "ReturnValue")
        rs.ReturnValue.PrintTree(concatps(indent, &astSpace), true)
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

type FunctionDeclarationStatement struct {
        Token      tokenizer.Token
        Identifier Identifier
        Parameters []Identifier
        Body       []Statement
}

func (fd *FunctionDeclarationStatement) statement() {}

func (fd *FunctionDeclarationStatement) GetTokenType() tokenizer.TokenType {
        return fd.Token.Type
}

func (fd *FunctionDeclarationStatement) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("FunctionDeclarationStatement")
        fmt.Println(*indent + astMiddle + "Identifier")
        fd.Identifier.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astMiddle + "Parameters")
        for index, element := range fd.Parameters {
                element.PrintTree(concatps(indent, &astLine), index == (len(fd.Parameters) - 1))
        }

        fmt.Println(*indent + astLast + "Body")
        for index, element := range fd.Body {
                element.PrintTree(concatps(indent, &astSpace), index == (len(fd.Body) - 1))
        }
}

type StructDeclarationStatement struct {
        Token      tokenizer.Token
        Identifier Identifier
        Body       []Statement
}

func (sd *StructDeclarationStatement) statement() {}

func (sd *StructDeclarationStatement) GetTokenType() tokenizer.TokenType {
        return sd.Token.Type
}

func (sd *StructDeclarationStatement) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("StructDeclarationStatement")
        fmt.Println(*indent + astMiddle + "Identifier")
        sd.Identifier.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astLast + "Body")
        for index, element := range sd.Body {
                element.PrintTree(concatps(indent, &astSpace), index == (len(sd.Body) - 1))
        }
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

        fmt.Println(*indent + astMiddle + "Operator")
        fmt.Println(*indent + astLine + astLast + tokenizer.TokenTypeString(bo.Token.Type))


        fmt.Println(*indent + astLast + "Right")
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

        fmt.Println(*indent + astMiddle + "Operator")
        fmt.Println(*indent + astLine + astLast + tokenizer.TokenTypeString(uo.Token.Type))

        fmt.Println(*indent + astLast + "Expression")
        uo.Expression.PrintTree(concatps(indent, &astSpace), true)
}

type AttributeAccessExpression struct {
        Token     tokenizer.Token
        Object    Expression
        Attribute Expression
}

func (ae *AttributeAccessExpression) expression() {}

func (ae *AttributeAccessExpression) GetTokenType() tokenizer.TokenType {
        return ae.Token.Type
}

func (ae *AttributeAccessExpression) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("AttributeAccessExpression")

        fmt.Println(*indent + astMiddle + "Object")
        ae.Object.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astLast + "Attribute")
        ae.Attribute.PrintTree(concatps(indent, &astSpace), true)
}

type FunctionCallExpression struct {
        Token      tokenizer.Token
        Identifier Expression
        Arguments  []Expression
}

func (fc *FunctionCallExpression) statement() {}
func (fc *FunctionCallExpression) expression() {}

func (fc *FunctionCallExpression) GetTokenType() tokenizer.TokenType {
        return fc.Token.Type
}

func (fc *FunctionCallExpression) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("FunctionCallExpression")
        fmt.Println(*indent + astMiddle + "Identifier")
        fc.Identifier.PrintTree(concatps(indent, &astLine), true)

        fmt.Println(*indent + astLast + "Arguments")
        for index, element := range fc.Arguments {
                element.PrintTree(concatps(indent, &astSpace), index == (len(fc.Arguments) - 1))
        }
}

///// LITERALES /////

type IntLiteral struct {
        Token tokenizer.Token
        Value int64
}

func (il *IntLiteral) expression() {}

func (il *IntLiteral) GetTokenType() tokenizer.TokenType {
        return il.Token.Type
}

func (il *IntLiteral) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("IntLiteral")
        fmt.Println(*indent + astLast + strconv.FormatInt(il.Value, 10))
}

type StringLiteral struct {
        Token tokenizer.Token
        Value string
}

func (sl *StringLiteral) expression() {}

func (sl *StringLiteral) GetTokenType() tokenizer.TokenType {
        return sl.Token.Type
}

func (sl *StringLiteral) PrintTree(indent *string, last bool) {
        fmt.Print(*indent)
        *indent += indentation(last)

        fmt.Println("StringLiteral")
        fmt.Println(*indent + astLast + sl.Value)
}
