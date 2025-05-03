BPF maps are great to store the state of your eBPF program. They are quite powerful because they help two or more eBPF programs to talk to each other. They also are used to send data from the kernel space to the user space.

BPF maps are defined in the eBPF programs, using BTF (modern way of defining maps). They are placed in the `maps` section of the ELF file (created by compiling the program).

When the user space programs loads the eBPF program using system calls, the kernel creates those maps (if verefier accepts them) and returns file descriptors to the user space program.

The user space program uses these to read/write data to/from the maps which are present in the kernel space.

One of the biggest difference between using BTF types vs old way is that, the old way only tells the size of the data, and not what it represents or holds, even the structure, which is not the case in terms of BTF types.

The BPF_ANNOTATE_KV_PAIR, just does some background magic to add thing like metadata information to the struct definition, but it still doesn't solve the problem about 

### How BTF allows Libbpf to understand things.

When you compile your code using Clang with a flag `-g`, it tells clang to embed BTF information and create separate `.BTF` section in the compiled ELF file along with other things like `.map` and text to hold the code.

Now, when the user space tries to load that program into the kernel, it starts reading and parsing the information. First it reads the `.BTF` section and creates a database for all the types.

Then it starts reading the `.maps` section to start and create maps in the memory. For every field of a map that uses a type or an external value, it searches for its type in the database and gets the necessary information like size etc.