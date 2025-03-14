---
title: Everything about libbpf
---

You need to know the steps needed for things to work.

`libbpf-bootstrap` is scaffolding playground setting up as much things as possible
for beginner users to get started without unnecessary hassle of setting things
up.

Keep you BPF logic and implementation as simple as possible so that it is backward
compatible and not too much reliant on latest features.

## Understanding the Makefile

```make

INCLUDES := -I$(OUTPUT) -I../../libbpf/include/uapi -I$(dir $(VMLINUX)) -I$(LIBBLAZESYM_INC)

```

The above argument tells the `CLang` compiler to look for Header files in the

- `OUTPUT` directory.
- `../../libbpf/include/uapi` directory.
- `VMLINUX` directory.
- `LIBBLAZESYM_INC` directory.

### Complex commands

- `$(dir $(VMLINUX))`: Overall, it tells the `CLang` compiler to search for headers
in the directory where `vmlinux.h` file is located.
  - `$(VMLINUX)`: Makefile variable that holds the path to the kernel header file,
  which is very crucial when it comes to `CO-RE`.
  - `$(dir ...)`: Makefile function that extracts the directory part of where a
  file is present.

```make
CLANG_BPF_SYS_INCLUDES ?= $(shell $(CLANG) -v -E - </dev/null 2>&1 | sed -n '/<...> search starts here:/,/End of search list./{ s| \(/.*\)|-idirafter \1|p }')
```

This line is designed to find the default system include directories that Clang
uses when it's targeting the BPF architecture (the architecture for eBPF programs).
Here's a more detailed breakdown:

1. `CLANG_BPF_SYS_INCLUDES ?=`: This defines a Makefile variable named `CLANG_BPF_SYS_INCLUDES`.
The ?= means that if this variable is not already defined, it will be assigned
the value on the right.

2. `$(shell ...)`: This is a Makefile function that executes the command inside the
parentheses in a shell and returns its output.

3. `$(CLANG) -v -E - </dev/null 2>&1`: This is the shell command being executed.

    - $(CLANG): This refers to the Clang compiler (which was defined earlier in
      the Makefile).
    - `v`: This Clang flag stands for "verbose". When used with other flags, it
      often causes Clang to print more information about what it's doing. In this
      case, it will include the paths where it searches for system header files.
    - `E`: This Clang flag tells the compiler to run only the preprocessor stage.
      The preprocessor is the part of the compiler that handles directives like
      `#include`. When you use `-E`, Clang will output the result of preprocessing
      your code (which in this case is an empty input from /dev/null).
    - `-`: This tells Clang to read the input from the standard input.
    - `</dev/null`: This redirects the standard input to come from /dev/null,
      which is an empty file. So, we're giving Clang no actual C code to `preprocess`.
    - `2>&1`: This redirects the standard error (where Clang might print warnings
      or other messages) to the standard output. This is done so that the entire
      output, including the header search paths, can be captured by the $(shell)
      function.

4. `sed -n '/<...> search starts here:/,/End of search list./{ s| \(/.*\)|-idirafter \1|p }'`:
This part pipes the output of the Clang command to the sed command for further
processing. sed is a stream editor used for filtering and transforming text.
    - `n`: This option tells sed not to print every line by default. We will
      explicitly tell it which lines to print.
    - `/<...> search starts here:/,/End of search list./{ ... }`: This is a sed
      command that operates on the lines between the pattern <...> search starts
      here: and the pattern End of search list.. These are the lines in Clang's
      verbose output that list the system include directories.
    - `s| \(/.*\)|-idirafter \1|p`: This is a substitution command within the block.
      - `s`: Indicates a substitution.
      - `|`: This is used as a delimiter instead of the usual / to avoid issues
        with slashes in the file paths.
      - ``: Matches a space.
      - `\(/.*\)`: This matches a space followed by a path enclosed in parentheses.
        The parentheses are escaped with a backslash \ to be treated literally. The
        .*matches any character (.) zero or more times (*), so this part captures
        the entire path inside the parentheses. The parentheses around it create
        a capturing group, which can be referenced later.
      - `idirafter \1`: This is the replacement string. It replaces the matched
        part with `-idirafter` followed by a space and \1. \1 refers to the first
        capturing group (the path we captured earlier). The `-idirafter` flag is
        similar to -I, but it tells the compiler to search these directories
        after the directories specified with -I.
      - `|`: Another delimiter.
      - `p`: This tells `sed` to print the lines where a substitution was made.

In essence, this whole line executes Clang in a verbose preprocessor mode,
captures its output, filters out the lines that list the system include
directories, and then formats these directories as -idirafter flags.

The use of -idirafter is important because it ensures that the include paths
specified directly in the INCLUDES variable (like those in libbpf/include/uapi)
are searched by the compiler before the default system include paths found by
this command. This is a common strategy when building software that needs to
use its own versions of headers or headers specific to a target environment
(like the BPF architecture) to avoid potential conflicts with system-wide headers
that might be different or incompatible . This is particularly relevant when
cross-compiling for eBPF, where the system you are building on might have
different kernel headers than the target kernel where the eBPF program will run.

## Linker

Linker combines the object files produced by the compiler after the compilation
of the code. Linker combines them into a single executable program.

### Linker Flags

These are the options that are passed to the `linker` program after the compilation
stage.

## Things to understand

- [ ] CO-RE
- [ ] BTF Support
- [ ] ELF sections
