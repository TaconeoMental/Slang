package compiler

import (
        "fmt"
        "os"

        "github.com/TaconeoMental/Slang/compiler/tokenizer"
        "github.com/TaconeoMental/Slang/compiler/parser"
        "github.com/TaconeoMental/Slang/compiler/codegen"
        "github.com/TaconeoMental/Slang/compiler/common"
)

type SlangCompiler struct {
        mainFilePath  string
        mainFileBytes []byte
        debug          bool
}

func NewSlangCompiler(name string, filebytes []byte, debug bool) (*SlangCompiler, error) {
        cmp := new(SlangCompiler)

        err := common.CheckMainFilePath(name)
        if err != nil {
                return nil, err
        }
        cmp.mainFilePath = name

        cmp.mainFileBytes = filebytes
        cmp.debug = debug
        return cmp, nil
}

func (sc *SlangCompiler) Compile() (int, error) {
        sc.DebugPrint("Compile() called")
        t := tokenizer.New(string(sc.mainFileBytes))

        p := parser.New(t)
        ast := p.ParseProgram()
        sc.DebugPrint("AST: %v", ast)

        if sc.debug {
                emptys := ""
                ast.PrintTree(&emptys, true)
        }

        cg, err := codegen.New(ast, sc.mainFilePath)
        if err != nil {
                return -1, err
        }

        cg.Generate()
        return 0, nil
}

func (sc *SlangCompiler) DebugPrint(s string, params ...interface{}) {
        switch {
        case !sc.debug:
                return
        case len(params) == 0:
                fmt.Print("[DEBUG] ")
                fmt.Fprint(os.Stdout, s)
                fmt.Println()
        default:
                fmt.Print("[DEBUG] ")
                fmt.Fprintf(os.Stdout, s, params...)
                fmt.Println()
        }
}


