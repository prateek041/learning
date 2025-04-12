# Understanding Linux Virtualisation

## Types of Virtualisation

We are only going to talk about software (hypervisor-based) virtualisation
which is the process of hiding the underlying physical hardware so that it
can be shared and used by multiple operating systems. This is done by introducing
a layer called a **Hypervisor/VMM** between the underlying hardware and the
operating systems running on it.

### Operating system virtualisation

Here is the mental model you can build for this. Operating system virtualisation
is focused on isolating processes and applications from each other. It is also
termed as `containerisation`. In this case, the host kernel (OS) is shared between
all the guest Operating systems, they don't bring their own kernel with them.
Therefore there is serious limitations with them.

You cannot run two Operating systems that depend on different kernels, windows 7
cannot run on a windows 10 OS virtualised machine because they have different
kernel levels that mismatch. Good side is, if your host is Linux, you can run
different distributions (as containers) on top of it like Debian, Ubuntu, Fedora
etc.

Same physical host servers different workloads and isolate each of the workloads.
This allows same physical system to run multiple isolated instances of the same
operating system, called containers. There is no emulation involved here, the
host operating system does not need to emulate system call interfaces for
alternative operating systems. There is just one operating system kernel. The
kernel just allows us to run multiple versions of it.

Example: we can not run windows 7 on windows 10, or macos on the linux host OS.

### Why so many issues with the rings

For anything to access hardware, they need to be in ring-0. So normally, the host
operating system kernel runs there. But, when it comes to virtualisation using
hypervisor/VMM, it also needs access to the hardware, therefore it also needs to
be run in ring-0. This causes an issue, because only one of them can run in the
ring-0 (only one entity can run ring-0 at a time). So, there needs to be a
way to solve this. This gives rise to two types virtualisation methods:

- Full Virtualisation: Run the VM OS in ring-1, and the VMM in ring-0. Then for every
privileged access request, the guest OS has to go through the VMM. This is called
Binary translation where the VMM traps and virtualises the execution of instructions.
The guest OS has no idea if anything it is running in a virtual machine.
- Para Virtualisation: Here the guest OS is aware of the virtualisation and the
Hypervisor. It interacts with the APIs provided by the VMM to interact with hardware.
This is more performant.

### Hardware Assisted Virtualisation

This is full virtualisation with hardware capabilities. This is done by creating
an even more privileged ring i.e. ring -1, where the hypervisor can run so that
the guest virtual machines can run in ring-0. So, the system can go back to it's
default configurations with VMM providing access to the hardware.

## VMM/Hypervisor

This is a Hardware assisted full virtualisation. Depending on whether the host system
has a underlying operating system or not, there can be two types of Hypervisors:

- Type 1: No Host OS, Hypervisor directly runs on the Hardware.
- Type 2: Host OS, between Hypervisor Software and Hardware.

## KVM

Kernel-based Virtual Machine (KVM). This virtualises the hardware like processor,
disk, network, VGA, PCI, USB, serial/parallel ports and so on to build a complete
virtual hardware on which guest operating systems can be installed. For I/O emulation
it relies on QEMU.

### High level overview of KVM

`libvirtd` launches a separate `qemu-kvm` process for each VM running on the system.
VM configurations are defined using a separate XML files located in the `/etc/libvirt/qemu/`.
Clouds use this tech very much.
