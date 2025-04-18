This article explains how a piece of code transforms from being written in your code editor to being executed by the CPU, focusing on the journey up to memory loading.

## From Code to Executable

You start by writing **source code** in a human-readable programming language of your choice. What happens next depends on whether you’re using a **compiled** or **interpreted** language:

- **Compiled Language**: A compiler translates the source code into low-level **machine code** (sequences of 0s and 1s). This machine code, combined with data like initial values of global variables, memory offsets, and metadata, is packaged into an **executable file**, commonly an **ELF** file on Linux.
- **Interpreted Language**: An interpreter translates the source code into machine code dynamically, executing it one statement at a time.

Since this series is focused on Go, a compiled language, we’ll focus on the compilation process.

### Executable File Sections

An executable file, such as an ELF file, is organized into distinct sections:

- **.text**: Contains the machine code instructions.
- **.data**: Stores initialized global and static variables.
- **.bss**: Holds uninitialized global and static variables.
- **Metadata**: Includes file information, symbols, and references to shared libraries.

For a deeper dive into ELF files, check out [this article]().

## Launching the Program

Once the executable is ready, the operating system (OS) steps in to run it. Here’s the process:

### 1. Request

You initiate execution by double-clicking the file or running it via the command line.

### 2. Process Creation

The OS creates a **process**, assigning it a unique **Process ID (PID)**. A process is an instance of the program in execution, managed by the OS to allocate resources like memory and CPU time.

### 3. Virtual Address Space

The OS assigns the process a **private virtual address space**, a large, linear range of *memory* addresses starting at 0. These addresses are **virtual**—an abstraction mapped to physical memory by the OS.

#### Why Virtual Address Space?

- **Isolation**: Ensures processes don’t interfere with each other’s memory, preventing errors or security issues.
- **Mapping**: The OS, aided by the CPU’s **Memory Management Unit (MMU)**, translates virtual addresses to physical ones for efficient memory use.

### 4. Loading

The OS **loader** reads the ELF file’s sections (e.g., `.text`, `.data`) and places them into the process’s virtual address space, preparing them for CPU access and execution.

### 5. Dynamic Linking

In languages like C, external libraries (e.g., `#include <stdio.h>`) are common. During **dynamic linking**, the OS loads these libraries into the virtual address space, making functions like `printf` available when called.

This marks the end of the journey from code to memory loading, setting the stage for execution.

This article is all you need to understand how a piece of code goes from being in your code editor to being executed by the CPU in the kernel.

## From Code to Executable

You write your code in a human readable format, and in your preferred programming language. This is called source code.

Now, depending on whether you used a compiled or an interpreted language, the next step can be either of the two listed below.

- **Compiled Language**: A compiler translates your source code into low-level machine code (sequences of 0s and 1s). This machine code is combined with additional data like initial values of global variables, memory offsets, metadata etc. are stored into an **executable file**, which in Linux is mostly an **ELF** file.
- **Interpreted Language**: An interpreter translates the source code into machine code on the fly, one argument at a time.

Since this series is more focused on Go, which is a compiled language, so we will pay more attention to that compiled language process.
### Executable File Sections
This is a brief introduction about the format of an executable file, a more in-depth article can be found [here](). The executable file typically contains distinct sections:
- **.text**: The actual machine code instructions.
- **.data**: Initialized global and static variables.
- **.bss**: Un-initialized global and static variables.
- **Metadata**: Information about the file, symbols, required shared libraries etc.

## Launching the Program

At this step, the executable defined above is executed and the Operating system takes over. Below are the steps in the flow:

#### Request

You double click or execute the binary from CLI to run the program.

#### Process creation

The OS creates a **process**. A process in itself is something we can go in great depths about, but for this article all you need to understand that OS creates it and assigns a unique `PID` to it.

#### Virtual Address Space

This is a crucial step, as we know that for a process to be in the running state, it needs to be loaded into the RAM. So, OS provides the newly created process a **private virtual address space**. 
  
It is a large, linear range of memory addresses that belong to the process. They are called virtual because their numbering is local (starting from 0 for all processes) and are mapped to actual addresses in the memory.

##### Why Virtual Address Space

There are two major reasons for creating private virtual address space for each process.

- Isolation: This virtual space prevents process from accidentally (or maliciously) interfering with each other's memory. Because this would be disastrous as a simple variable holding value `4` for a process might be replaced with `5` by another and all of a sudden, your program starts suggesting `2+2=5`.
- Mapping: The OS, with the help of CPU's Memory Management Unit maps virtual addresses in this private space, to actual physical addresses in the memory.

#### Loading

The OS loader, reads different sections of the ELF file (.text, .data) and copies them into the appropriate locations within the process's virtual address space, so they can be accessed and executed by the CPU.

#### Dynamic Linking

If you are working with a programming language like `C`, linking process becomes very crucial. Whenever you write a software, you often rely on external libraries and packages. For example, you use `#include<stdio.h>` pre-processor, that makes sure your compiled software also includes the entire code present in the `stdio.h` header file.

Then, while writing your code, you can use functions like `printf` which are provided by the `stdio.h` header file. It is actually at the linking step, that all the external libraries and dependencies are loaded into the virtual memory, to make sure the actual implementation of the `printf` function is available when it is used anywhere.

This concludes the life-cycle of a program all the way from being written to being loaded into the memory.

Next articles in the series go more in-depth. They are about how the execution works. This would involve understanding things like

- Threads and their execution flow.
	- Single Threaded vs Multi threaded applications
	- Thread components including *program counter*, *temporary storage*, *CPU registers*, *Stack* etc.
	- Shared resources: How resources are shared between multiple threads running under the same process.
	- Multi-threading and Go Routines.
	- CPU scheduling algorithms, interrupts and time sharing.
- Execution Stack in detail.
	- Managing function calls, passing arguments and handling returns.
	- Application Binary Interface (ABI) of a language.
	- Stack Frame.
- The Heap in detail (Dynamic Memory).