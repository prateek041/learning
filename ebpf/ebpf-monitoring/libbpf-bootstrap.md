---
title: Everything about libbpf
---

You need to know the steps needed for things to work.

`libbpf-bootstrap` is scaffolding playground setting up as much things as possible
for beginner users to get started without unnecessary hassle of setting things
up.

Keep you `BPF` logic and implementation as simple as possible so that it is backward
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
uses when it's targeting the `BPF` architecture (the architecture for `eBPF` programs).
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
This part pipes the output of the Clang command to the `sed` command for further
processing. `sed` is a stream editor used for filtering and transforming text.
    - `n`: This option tells `sed` not to print every line by default. We will
      explicitly tell it which lines to print.
    - `/<...> search starts here:/,/End of search list./{ ... }`: This is a `sed`
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
directories, and then formats these directories as `-idirafter` flags.

The use of `-idirafter` is important because it ensures that the include paths
specified directly in the INCLUDES variable (like those in `libbpf/include/uapi`)
are searched by the compiler before the default system include paths found by
this command. This is a common strategy when building software that needs to
use its own versions of headers or headers specific to a target environment
(like the `BPF` architecture) to avoid potential conflicts with system-wide headers
that might be different or incompatible . This is particularly relevant when
cross-compiling for `eBPF`, where the system you are building on might have
different kernel headers than the target kernel where the `eBPF` program will run.

## Linker

Linker combines the object files produced by the compiler after the compilation
of the code. Linker combines them into a single executable program.

### Linker Flags

These are the options that are passed to the `linker` program after the compilation
stage.

## Minimal App

## No parent section

### Finding all the available events

```bash
find /sys/kernel/tracing/events -type d
```

## Important Commands

- `file`: used to get some information about the binary/object file.
- `readelf`: used to read more in-depth information about the binary/object file.

---

## Little Endian and Big Endian systems

### The hardware limitations

In most computer systems, RAM (Random Access Memory) is designed to store and
retrieve data in units called *bytes*. A byte is 8 *bits* and it's the smallest
*addressable* unit in the memory. This means when the CPU talk to the RAM.
It cannot grab anything smaller than 1 byte at a time (like 2/3/4... `bits`), it
always has to work with at least 1 byte at a time.

So, if you store anything smaller than 1 byte in your memory, it still takes up
the entire space. For example:

- A single bit for `true` might be stored as `0x01` (binary: 00000001), where only
1 bit is used and other 7 are zeros.
- A 4 bit number like 5 (binary: 0101) becomes `0x05` (binary: 00000101), with
4 bits unused.

### The number system

There are multiple numbers systems to represent a number in a computer. Some of
them are:

- Binary (base 2)
- Hexadecimal (base 16)
- Decimal (base 10)

Let's try to understand how they interact with each other.

#### Hexadecimal and Binary

To represent a `hex` number in `binary`, we need 4 bits. Therefore the numbers
mapping would look like this:

- 0 -> 0000
- 1 -> 0001
- 2 -> 0010
- 3 -> 0011
- 4 -> 0100
- 5 -> 0101
- 6 -> 0110
- 7 -> 0111
- 8 -> 1000
- 9 -> 1001
- A -> 1010
- B -> 1011
- C -> 1100
- D -> 1101
- E -> 1110
- F -> 1111

Now, since a byte is 8 bits, and each hex digit is 4 bits, it takes two hex digits
to represent one byte. For example:

##### Multiple Byte integers

- `0x12` (Hexadecimal) = `0001 0010` (Binary) = 1 byte.

- `0x12345678` (Hexadecimal) = `0001 0010 0011 0100 0101 0110 0111 1000` (Binary)
= 32 bits = 4 bytes
  - Byte 1: `0x12`
  - Byte 2: `0x34`
  - Byte 3: `0x56`
  - Byte 4: `0x78`

When computer stores numbers bigger than a single byte (8 bits), it needs to break
them down into smaller pieces (bytes) and decide how to arrange bytes in memory.
`Endianess` is the rule that decides the order of those bytes.

There are two systems for that: **little endian** and **big endian**

1.**Big Endian: "Big end First"**

In a big endian system, the most significant byte (the big part) gets stored at the
lowest memory address, just like how humans read numbers (left to right). Example:

`x12345678` (a 4 byte integer in hex) can be stored in the memory like:

- `0x12` is the most significant byte (MSB).
- `0x78` is the least significant byte (LSB).

In memory, a big endian system stores it like this:

- Address 0: `0x12` (MSB)
- Address 1: `0x34`
- Address 2: `0x56`
- Address 3: `0x78` (LSB)

2.**Little Endian: "Little End First"**

A little endian system is just the opposite of a big endian system. It stores the
`LSB` in the lowest memory address. So the above example becomes:

- Address 3: `0x12` (MSB)
- Address 2: `0x34`
- Address 1: `0x56`
- Address 0: `0x78` (LSB)

---

## ELF Files

ELF is the abbreviation for Executable and Linkable Format and defines the structure
for binaries, libraries, and core files. ELF files are typically the output of a
compiler or linker and are a binary format.

### From source to process

So whatever operating system we run, it needs to translate common functions to
the language of the CPU, also known as machine code. A function could be something
basic like opening a file on disk or showing something on the screen. Instead of
talking directly to the CPU, we use a programming language, using internal functions.
A compiler then translates these functions into object code. This object code is
then linked into a full program, by using a linker tool. The result is a binary
file, which then can be executed on that specific platform and CPU type.

### Structure of an ELF file

There are two major sections in an ELF file

- ELF Headers
- File data

#### ELF Header

If you use the command `readelf -a <filename>` you will get to see something like
this.

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
  Start of section headers:          15224 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           0 (bytes)
  Number of program headers:         0
  Size of section headers:           64 (bytes)
  Number of section headers:         23
  Section header string table index: 22
```

Let's understand each individual parts.

- **Magic** : This ELF header magic provides information about the file. The first
four hexadecimal parts define that this is an ELF file (45=E,4c=L,46=F),
prefixed with the 7f value.
- **Class**: This defines the architecture of the file i.e. whether it is
**32 bit (=01)** or **64 bit (=02)**. You can see in the Magic section, **02**
which `readelf` command translates to `ELF64`.
- **Data**: The next section tells whether the system follows `big endian` or
`little endian` architecture. Where `01` is for little endian.

## Kernel Events and their types

- `syscalls/` → System call events (sys_enter_*, sys_exit_*).
- `sched/` → Scheduler events (sched_switch, sched_wakeup).
- `irq/` → Interrupt events (irq_handler_entry, irq_handler_exit).
- `workqueue/` → Workqueue events (workqueue_execute_start, workqueue_execute_end).
- `net/` → Networking events (net_dev_xmit, netif_receive_skb).
- `block/` → Block I/O events (block_rq_issue, block_rq_complete).
- `kmem/` → Memory allocation events (kmalloc, kfree).
- `sound/` → Sound-related events (snd_soc_jack_notify, etc.).

## Things to understand

- [ ] CO-RE
- [ ] BTF Support
- [ ] ELF sections
