Modules are how go manages dependencies.
> Thread id thread_TSo5OkH3yKLqCLaJIi4KyCpj

> External modules imported into the current module are called dependencies, they are listed into the go.mod file.

A module is a collection of packages, where each package is a collection of go source code files.
> **Package** is a collection of Go source files i.e. every file ending with `.go` extension.

The go.mod file is very important. it stores the unique identification of every go module along with all the other go modules it depends on.
> Root directory of the module is where the go.mod file is present.

### Module Paths


