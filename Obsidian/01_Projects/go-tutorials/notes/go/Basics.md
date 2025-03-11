for getting started, there are two ways go programs can be written
- You need dependencies
- You do no need external dependencies

### No dependency

Two things are absolutely necessary
- Package main
- function main
Any program that needs to be compiled and run must be inside package main and the program execution starts from the function main.

Presence of package main tells the Go toolchain that the package should compile to an executable binary and is not a part of a shared library. and inside the package main, func main is the main entry point.

To effectively understand Go and teach it I need to do two things:

- Follow through the getting started documentation
- Follow through the effective Go

make sense of these two things and you are done.

## Go Specific things to cover

- Modules (used for dependency tracking)
- package main and func main
- Importing functions from other modules
- Managing dependencies
- Authenticating modules
- Developing and publishing modules
- using local and external modules
- module versioning
- Effective Go
- 