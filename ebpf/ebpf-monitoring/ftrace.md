---
title: Everything you need to know about Ftrace
---

In the Linux kernel there are `tracepoints`, they are provided by the
kernel travels which work as an instrumentation hook points for you
to attach your custom logic/code, which triggers when `tracepoint` related
to that function is run.

`ftrace` builds on top of it. I uses the `tracefs` file system to configure
how the system gets monitored. The `/sys/kernel/tracing/events/` file
contains all the `tracepoints` supported by the kernel.

There are `tracers` which are part of the `ftrace` framework which are
pre-written, kernel source code which handle the logic of processing
these traces when they are enabled. They can be enabled by just converting
`0->1`.

`tracepoints` are not related to `tracers` in the sense that `tracepoints`
are a linux kernel thing, while `tracers` are a `ftrace` thing.

## Flow of Information

The `kernel` code writes the traces into a ring buffer, which is accessible
through an interface (a file), which is `/sys/kernel/tracing/trace`, when
read.

> [!IMPORTANT]
> The file `tracepoint.h` holds the definition of `TRACE_EVENT` which is
> used to define a trace event. Similarly, the
> [/trace/events](<https://elixir.bootlin.com/linux/v6.2.2/>> source/include/trace/events)
> contains the pre-defined kernel trace events.

### Learn about

- [ ] BPF Type Format (BTF) [link](https://docs.kernel.org/bpf/btf.html)
- [ ] ELF file and its types [ELF files](https://linux-audit.com/elf-binaries-on-linux-understanding-and-analysis/)
