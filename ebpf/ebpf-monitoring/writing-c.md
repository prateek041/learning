---
title: Writing C programs
description: Basics you need to know to understand how to write software in C
---

## What happens when you Compile

[Read This](https://www.cs.tufts.edu/cs/40/docs/compiling-notes.html)

## What are Header files

Header files end with `.h` extension contain declarations of functions,
structures and other definitions that your C code might use.

You can use the `#include` pre-processor directive to tell the compiler
to include the content of these files and the compiler will check these files
in **designated** places.

When your code compiles, the directive is replaced with the content of the
actual file.

## Essentials of building a software in C

When building software with C, like parts of the `eBPF` ecosystem you will
encounter different kinds of files during it's compilation.

- `.o files`: These are called object files which are generated when a
  compiler like `Clang` processes a source code. This file is not complete
  program yet. It misses connections to other parts of the project.

- `.a files`: These are **archive** files also known as **static libraries**.
  These are *pre-compiled* object files bundled together into a single
  file. When a program needs functions or data from these files, the
  linker can pull the necessary object files from the `.a` files and
  link them into the final program.

## Specifying where to pull header files from

We use the `-I` flag to tell a compiler like `CLang` where to pull the header
files from.
