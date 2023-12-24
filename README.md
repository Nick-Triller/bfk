# bfk

![Test](https://github.com/Nick-Triller/bfk/actions/workflows/WORKFLOW-FILE/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Nick-Triller/bfk)](https://goreportcard.com/report/github.com/Nick-Triller/bfk)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)

bfk is an interpreter for the esoteric programming 
language "Brainfuck" implemented in Go.

The interpreter uses [this specification](https://github.com/brain-lang/brainfuck/blob/master/brainfuck.md) 
as a guideline.

## Usage

```
bfk is a brainfuck interpreter written in go.

Usage:
  bfk [command]

Available Commands:
  help        Help about any command
  minify      Minifies a program.
  run         Run a program.

Flags:
  -h, --help      help for bfk
      --version   version for bfk

Use "bfk [command] --help" for more information about a command.
```

Run a program:
```
bfk run programs/helloworld.bfk
```

## Memory layout

The VM simulates an infinite number of cells for positive and 
negative adresses via a paging mechanism. 
Each cell consists of an unsigned 8 bit integer.

## Instructions

| Instruction | Description                                                                            |
| ----------- | -------------------------------------------------------------------------------------- |
| `<`         | Moves the memory pointer one cell to the left.                                         |
| `>`         | Moves the memory pointer one cell to the right.                                        |
| `+`         | Increments the current cell by one.                                                    |
| `-`         | Decrements the current cell by one.                                                    |
| `,`         | Reads one byte of data from stdin.                                                     |
| `.`         | Writes the value of the current cell to stdout as ASCII.                               |
| `[`         | Jumps forward to the instruction after the matching `]` if the current cell is 0.      |
| `]`         | Jumps backwards to the instruction after the matchin `[` if the current cell is not 0. |

## Build

```
set GO111MODULE=on
go build
```

## Test

```
go test ./...
```

## Dependencies

bfk uses [cobra](https://github.com/spf13/cobra) for CLI parsing.
