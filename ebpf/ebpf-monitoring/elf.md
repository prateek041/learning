---
title: "Everything about ELF"
---

## What is ELF

// Write about ELF Here.

## ELF Sections

Sections in ELF file are like labeled containers that organise different types
of data in a program.

Some common sections

- `.text`: The machine code (instructions) your program executes.
- `.data`: Initialised global or static variables (e.g. `int x = 10;`).
- `.bss`: Uninitialised global or static variables (`int y;`).
- `.rodata`: Read-only data.
- `.symtab`: A symbol table listing functions and variables with their addresses.
- `.rela.text`: Relocation entries for the `.text` section.
- `.strtab`: A string table with names of symbols.

## Why do we need them

- **Organization**: Sections group related data, so tools know where to find
specific pieces (e.g., code vs. variables).
- **Efficiency**: Each section has attributes (e.g., executable, writable),
letting the operating system load them into memory with the right permissions.
- **Flexibility**: During linking, sections from multiple files can be combined
or rearranged to build the final program.

## Different types of ELF files

The ELF header’s type field tells us what the file is for. Each type serves
a unique purpose in the process of creating and running programs. Here are
the common types with examples:

- **REL** (Relocatable File, value 1):
  - What: An object file (.o) from the compiler, not yet linked.
  - Use: Contains code and data with relocations, ready to be linked into an executable.
  - Example: `gcc -c myfile.c -o myfile.o` creates a `REL` file.
- **EXEC** (Executable File, value 2):
  - What: A fully linked program you can run.
  - Use: Loaded directly into memory for execution.
  - Example: `gcc myfile.c -o myprogram` creates an `EXEC` file.
- **DYN** (Shared Object File, value 3):
  - What: A shared library (.so) for dynamic linking.
  - Use: Provides reusable code loaded at runtime (e.g., by multiple programs).
  - Example: `libc.so` (the C standard library) is a `DYN` file.
- **CORE** (value 4):
  - What: A core dump file from a crashed program.
  - Use: Helps debug what went wrong by showing the program’s state at the crash.
  - Example: If `myprogram` crashes, the kernel might create core or `core.pid`.

### Why Different Types?

Each type matches a step or role in the process—compiling (REL), running (EXEC),
sharing code (DYN), or debugging crashes (CORE). Tools like readelf use this
field to know how to handle the file.

## Why do we need relocation

Relocations make it possible to build programs from multiple pieces, even when
their addresses aren’t known until later.

Relocations are placeholders in an ELF file that mark spots where addresses need
to be filled in later.

In an object file (.o), if your code calls an external function like `fprintf`
(from `libc`), the compiler can’t set its address yet. Instead, it puts a
temporary value in the `.text` section and adds a relocation entry (e.g., in
`.rela.text`) saying:

> At this spot in the code, insert the address of `fprintf`.
> It includes details like the symbol name and how to calculate the address.

### The Process: From Compilation to Execution

Here’s how a C program goes from source code to running in memory, focusing
on what happens after compilation:

- **Step 1: Compilation**

What Happens: When you compile a `.c` file (e.g., `gcc -c myfile.c -o myfile.o`)
, you get a `.o` file, which is an ELF REL (relocatable) file.

**What’s Inside**: Sections like `.text`  (code), `.data` (variables), `.bss`, and
`.rela.text` (relocations). No segments yet, since it’s not executable.

**Purpose**: This file has your code but with unresolved references (e.g., to `printf`).

- **Step 2: Linking**

**What Happens**: The linker (`ld`) combines one or more `.o` files and libraries
into an executable (e.g., `gcc myfile.o -o myprogram`).

**Process**: Combines sections (e.g., merges all `.text` sections into one).
**Resolves relocations**:
    - Static linking: Fills in all addresses (e.g., `printf`’s location in the executable).
    - Dynamic linking: Leaves some for runtime (e.g., `printf` from `libc.so`).
Adds segments (program headers) for loading into memory.

**Output**: An ELF EXEC file (`myprogram`) with sections and segments.

- **Step 3: Loading and Execution**

**What Happens** : When you run `./myprogram`, the OS loads it into memory.

- **Loading**: The OS reads the ELF file’s segments (program headers).
  - For each LOAD segment:
    - Allocates memory with the right permissions
    (e.g., executable for `.text`).
    - Copies the segment’s data from the file.
    - Zeroes out `.bss` (since it’s uninitialized and takes no file space).
  - If dynamically linked:
    - The dynamic linker (`ld.so`) loads libraries (e.g., `libc.so`).
    - Resolves remaining relocations, updating addresses in memory.
- **Execution**: The OS jumps to the program’s entry point (e.g., `_start`),
which calls main.

### Runtime Use Case for Sections

Even after loading, sections can be useful:

A debugger like `gdb` uses `.symtab` and `.debug_*` sections to show you
variable values and source code lines while the program runs.

## An ELF File and Inspecting it

Let's say we have this example C program:

```c
#include <stdio.h>
int main(){
  printf("Hello, World!\n");
  return 0;
}
```

We will compile this in two stages:

- `gcc -c main.c -o main.o`: Creates a REL (relocatable) ELF file.
- `gcc main.o -o main`: Links it into an EXEC (executable) ELF file.

### Inspecting REL File

Running `readelf -a main.o` reveals a few things:

```bash
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              REL (Relocatable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x0
  Start of program headers:          0 (bytes into file)
  Start of section headers:          600 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           0 (bytes)
  Number of program headers:         0
  Size of section headers:           64 (bytes)
  Number of section headers:         14
  Section header string table index: 13
```

- **ELF Header**:
  - Magic: `7f 45 4c 46 ...` (identifies it as an ELF file).
  - Type: `REL` (relocatable, not ready to run yet).
  - Machine: `x86_64` (or your architecture).
  - Entry Point: None (not applicable for REL files).
  - Section Header Offset: Points to where sections are listed.
  - Program Header Offset: 0 (no segments yet).

  ```txt
  Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .text             PROGBITS         0000000000000000  00000040
       000000000000001a  0000000000000000  AX       0     0     1
  [ 2] .rela.text        RELA             0000000000000000  00000198
       0000000000000030  0000000000000018   I      11     1     8
  [ 3] .data             PROGBITS         0000000000000000  0000005a
       0000000000000000  0000000000000000  WA       0     0     1
  [ 4] .bss              NOBITS           0000000000000000  0000005a
       0000000000000000  0000000000000000  WA       0     0     1
  [ 5] .rodata           PROGBITS         0000000000000000  0000005a
       000000000000000d  0000000000000000   A       0     0     1
  [ 6] .comment          PROGBITS         0000000000000000  00000067
       000000000000001c  0000000000000001  MS       0     0     1
  [ 7] .note.GNU-stack   PROGBITS         0000000000000000  00000083
       0000000000000000  0000000000000000           0     0     1
  [ 8] .note.gnu.pr[...] NOTE             0000000000000000  00000088
       0000000000000030  0000000000000000   A       0     0     8
  [ 9] .eh_frame         PROGBITS         0000000000000000  000000b8
       0000000000000038  0000000000000000   A       0     0     8
  [10] .rela.eh_frame    RELA             0000000000000000  000001c8
       0000000000000018  0000000000000018   I      11     9     8
  [11] .symtab           SYMTAB           0000000000000000  000000f0
       0000000000000090  0000000000000018          12     4     8
  [12] .strtab           STRTAB           0000000000000000  00000180
       0000000000000012  0000000000000000           0     0     1
  [13] .shstrtab         STRTAB           0000000000000000  000001e0
       0000000000000074  0000000000000000           0     0     1
  ```

- `.text`: Contains the compiled machine code for main.
  - Example: call `0x0` (placeholder for `printf`).
- `.rela.text`: Relocation entries for unresolved references.
  - Example: “At offset 0x5, fix the call to `printf`.”
- `.rodata`: Read-only data, like the string `"Hello, World!\n"`.
- `.data`: Initialized global variables (none here).
- `.bss`: Uninitialized global variables (none here).
- `.symtab`: Symbol table.
  - `main`: Defined, local, in `.text`.
  - `printf`: Undefined (needs linking).
- `.strtab`: String table for symbol names.

### Inspecting EXEC File

```txt
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              DYN (Position-Independent Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x1040
  Start of program headers:          64 (bytes into file)
  Start of section headers:          13496 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         14
  Size of section headers:           64 (bytes)
  Number of section headers:         30
  Section header string table index: 29
```

Headers are similar, except there is a starting point along with file type
being executable.

```txt
Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .note.gnu.pr[...] NOTE             0000000000000350  00000350
       0000000000000040  0000000000000000   A       0     0     8
  [ 2] .note.gnu.bu[...] NOTE             0000000000000390  00000390
       0000000000000024  0000000000000000   A       0     0     4
  [ 3] .interp           PROGBITS         00000000000003b4  000003b4
       000000000000001c  0000000000000000   A       0     0     1
  [ 4] .gnu.hash         GNU_HASH         00000000000003d0  000003d0
       000000000000001c  0000000000000000   A       5     0     8
  [ 5] .dynsym           DYNSYM           00000000000003f0  000003f0
       00000000000000a8  0000000000000018   A       6     1     8
  [ 6] .dynstr           STRTAB           0000000000000498  00000498
       000000000000008d  0000000000000000   A       0     0     1
  [ 7] .gnu.version      VERSYM           0000000000000526  00000526
       000000000000000e  0000000000000002   A       5     0     2
  [ 8] .gnu.version_r    VERNEED          0000000000000538  00000538
       0000000000000030  0000000000000000   A       6     1     8
  [ 9] .rela.dyn         RELA             0000000000000568  00000568
       00000000000000c0  0000000000000018   A       5     0     8
  [10] .rela.plt         RELA             0000000000000628  00000628
       0000000000000018  0000000000000018  AI       5    23     8
  [11] .init             PROGBITS         0000000000001000  00001000
       000000000000001b  0000000000000000  AX       0     0     4
  [12] .plt              PROGBITS         0000000000001020  00001020
       0000000000000020  0000000000000010  AX       0     0     16
  [13] .text             PROGBITS         0000000000001040  00001040
       0000000000000113  0000000000000000  AX       0     0     16
  [14] .fini             PROGBITS         0000000000001154  00001154
       000000000000000d  0000000000000000  AX       0     0     4
  [15] .rodata           PROGBITS         0000000000002000  00002000
       0000000000000011  0000000000000000   A       0     0     4
  [16] .eh_frame_hdr     PROGBITS         0000000000002014  00002014
       0000000000000024  0000000000000000   A       0     0     4
  [17] .eh_frame         PROGBITS         0000000000002038  00002038
       000000000000007c  0000000000000000   A       0     0     8
  [18] .note.ABI-tag     NOTE             00000000000020b4  000020b4
       0000000000000020  0000000000000000   A       0     0     4
  [19] .init_array       INIT_ARRAY       0000000000003dd0  00002dd0
       0000000000000008  0000000000000008  WA       0     0     8
  [20] .fini_array       FINI_ARRAY       0000000000003dd8  00002dd8
       0000000000000008  0000000000000008  WA       0     0     8
  [21] .dynamic          DYNAMIC          0000000000003de0  00002de0
       00000000000001e0  0000000000000010  WA       6     0     8
  [22] .got              PROGBITS         0000000000003fc0  00002fc0
       0000000000000028  0000000000000008  WA       0     0     8
  [23] .got.plt          PROGBITS         0000000000003fe8  00002fe8
       0000000000000020  0000000000000008  WA       0     0     8
  [24] .data             PROGBITS         0000000000004008  00003008
       0000000000000010  0000000000000000  WA       0     0     8
  [25] .bss              NOBITS           0000000000004018  00003018
       0000000000000008  0000000000000000  WA       0     0     1
  [26] .comment          PROGBITS         0000000000000000  00003018
       000000000000001b  0000000000000001  MS       0     0     1
  [27] .symtab           SYMTAB           0000000000000000  00003038
       0000000000000240  0000000000000018          28     6     8
  [28] .strtab           STRTAB           0000000000000000  00003278
       0000000000000126  0000000000000000           0     0     1
  [29] .shstrtab         STRTAB           0000000000000000  0000339e
       0000000000000116  0000000000000000           0     0     1
```

- **Sections**

  - `.text`: Now includes main’s code plus startup code from `libc`.
  - `.rodata`: Still has "Hello, World!\n".
  - `.plt`: Procedure Linkage Table for dynamic calls (e.g., to `printf`).
  - `.got`: Global Offset Table for dynamic addresses.
  - `.data`, `.bss`: As before, plus any from linked libraries.
  - `.dynamic`: Dynamic linking info (e.g., needed libraries like `libc.so`).
  - `.symtab`: Symbols, though often stripped in release builds.

  To read in depth description of each section, read [this](https://www.muppetlabs.com/~breadbox/software/ELF.txt).
