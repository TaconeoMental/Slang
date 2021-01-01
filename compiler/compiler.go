package compiler

import (
        "fmt"
        "os"

        "github.com/TaconeoMental/Slang/compiler/tokenizer"
        "github.com/TaconeoMental/Slang/compiler/parser"
)

type SlangCompiler struct {
        mainFilePath  string
        mainFileBytes []byte
        debug          bool
}

func NewSlangCompiler(name string, filebytes []byte, debug bool) *SlangCompiler {
        cmp := new(SlangCompiler)
        cmp.mainFilePath = name
        cmp.mainFileBytes = filebytes
        cmp.debug = debug
        return cmp
}

func (sc *SlangCompiler) Compile() (int, error) {
        sc.DebugPrint("Compile() called")
        t := tokenizer.New(string(sc.mainFileBytes))

        p := parser.New(t)
        ast := p.ParseProgram()
        sc.DebugPrint("AST: %v", ast)
        emptys := ""
        ast.PrintTree(&emptys, true)
        // TODO
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


