package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "flag"

    "github.com/TaconeoMental/Slang/compiler"
)

func main() {
        debug := flag.Bool("debug", false, "Run in debug mode.")
        compileFile := flag.Bool("compile", false, "Compile a .sl Slang file into a .hbc Hiru bytecode file.")

        // Definimos nuestro propio usage del CLI para mostrar la posición de
        // [filenmae]
        flag.Usage = func() {
            fmt.Fprintf(os.Stderr, "Usage: %s [Options] [Filename]\n", os.Args[0])
            fmt.Fprintln(os.Stderr, "Options:")

            flag.VisitAll(func(f *flag.Flag) {
                fmt.Fprintf(os.Stderr, "    -%v,\t%v\n", f.Name, f.Usage) // f.Name, f.Value
            })
        }

        flag.Parse()

        if len(flag.Args()) == 0 {
                fmt.Fprintln(os.Stderr, "Slang: No file specified")
        }

        if *compileFile {
                fileBytes, err := ioutil.ReadFile(flag.Arg(0))
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Slang: Error opening file: %T:%v\n", err, err)
                }

                slangCompiler := compiler.NewSlangCompiler(flag.Arg(0), fileBytes, *debug)
                _, err = slangCompiler.Compile()
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Slang: Error compiling file: %T:%v\n", err, err)
                }
        }
}
