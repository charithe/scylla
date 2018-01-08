Scylla
======

A **highly experimental** GraphQL parser utilising Ragel state machines.

**Why?**

- There are several excellent GraphQL libraries available for Go. However, they either don't implement the full spec or don't provide full access to the parse tree. Having access to the parse tree is quite useful if you are attempting to deal with dynamically loaded schemas or if you are trying to write a tool such as a code generator. 
- Usually Ragel is used as a lexer to produce the tokens that get consumed by a parser generator such as `goyacc`. Since parser generators are (mostly) state machines themselves, I wanted to see whether Ragel alone can be used to generate a parser.


Quick Start
-----------

Ensure Ragel is installed and available on $PATH

Run `make clean test`
