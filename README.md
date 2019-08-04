# Lyte
[![](https://img.shields.io/github/languages/top/llusx/Lyte.svg)]()
[![](https://img.shields.io/github/last-commit/llusx/Lyte)]()
[![](https://img.shields.io/github/license/llusx/Lyte)]()
[![](https://img.shields.io/maintenance/yes/2019)]()

A Lightweight Turing Machine Code Interpreter written in Go
---
Simple interpreter (for now) of the language used in the website [Online Turing Simulator](https://turingmachinesimulator.com/), currently supports code written for a machine with one tape.
## Building
To compile from source, you should have *go* set up on your machine.
Clone the repo, and run `go build`. This will generate a binary named `Lyte`. Make it executable (if needed) and run it against your Turing Machine code.

`./Lyte decimaltobinary.lyte` 

## Stages of Interpreting
Given the path to your code is valid, Lyte will do do the following.

+ Load up your code, and try to parse it.
+ Inform you if the parsing process was successful.
+ Ask you to provide initial values for the tape.
+ Execute the program.
+ Inform you whether the machine halted on an accepting state.
+ Display a representation of the tape after halting.
