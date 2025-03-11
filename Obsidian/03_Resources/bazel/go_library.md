```
go_library(name, cdeps, cgo, clinkopts, copts, cppopts, cxxopts, data, deps, embed, embedsrcs, gc_goopts, importmap, importpath, importpath_aliases, srcs, x_defs)
```
This builds a Go library from a set of source files that are all part of the same package.

- **name**: A unique name of the target
- **srcs**: List of Go source files that are compiled to create the package.
- **importpath**: The source import path of the library. Other libraries can import this library using the path.

## Providers

- GoLibrary
- GoSource
- GoArchive