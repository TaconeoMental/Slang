<p align="center">
  <img height="150x" src="https://raw.githubusercontent.com/TaconeoMental/Slang/main/assets/slang_logo.png" />
</p>

![Status Badge](https://img.shields.io/badge/status-Work%20in%20progress-yellow)

Slang is a simple and compact programming language designed to be used with the [Hiru Virtual machine](https://github.com/TaconeoMental/Hiru-VM). It has a familiar syntax and all the neccessary features to write basic programs.
It uses the Hiru assembly language as an IR which then compiles directly to bytecode.

## car.sl
Here is a little program showing some of Slangs' syntax.
<div style="text-align:center"><img height="700em" src="https://raw.githubusercontent.com/TaconeoMental/Slang/main/assets/code_example.png" /></div>

## Building source
```bash
git clone https://github.com/TaconeoMental/Slang
cd Slang
go build
```

## Compiling
This command will produce the bytecode file```car.hbc```.
```bash
./slang -compile car.sl
```

## Running
To execute a __.hbc__ compiled bytecode file:
```bash
./slang car.hbc
20
Too fast!
20
0
```
