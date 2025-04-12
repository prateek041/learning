# KVM Internals

What will this chapter cover

- The internal working of libvirt, QEMU and KVM.
- Important data structures and code paths of libvirt, QEMU and KVM.
- Execution flow of vCPUs
- Communication between them.

## Libvirt

This is management layer, an API that is between your use space command lines
like `virsh` or `virt-manager`. It gives them access to interact with different
types of hypervisors like QEMU, XEN etc.

The goal of the libvirt library is to provide a common and stable
layer to manage VMs running on a hypervisor. In short, as a management layer it
is responsible for providing the API that does management tasks such as virtual
machine provision, creation, modification, monitoring, control, migration, and so
on.
