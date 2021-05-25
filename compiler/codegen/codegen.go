package codegen

import (
        "fmt"
        "os"

        "github.com/TaconeoMental/Slang/compiler/ast"
        "github.com/TaconeoMental/Slang/compiler/common"
)

type CodeGenerator struct {
        ast        *ast.Program
        asmFilePath string
}

func New(ast *ast.Program, mainFileName string) (*CodeGenerator, error) {
        cg := new(CodeGenerator)
        cg.ast = ast
        asmFilePath, err := createAsmFile(mainFileName)
        if err != nil {
                return nil, fmt.Errorf("Error creating bytecode file '%s", asmFilePath)
        }
        cg.asmFilePath = asmFilePath

        return cg, nil
}

func createAsmFile(slangPath string) (string, error) {
        fileName := common.GetFileNoExt(slangPath)
        hbcFilePath := fileName + ".asm"

        f, err := os.Create(hbcFilePath)
        defer f.Close()

        _, err = f.WriteString("[HIRU]\n")

        return hbcFilePath, err
}

func (*CodeGenerator) Generate() {}
