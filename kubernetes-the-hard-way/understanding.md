---
title: Understanding every aspect of Kubernetes the Hard way
---

## Bringing the machines

We are supposed to bring Virtual or Physical Machines, we will bring Virtual machines.

In order to create a virtual machine, there are a lot of ways, one things I am 
reading about right now is QEMU and KVM.

### KVM

The idea of virtualisation is simple. We want to run multiple VMs as if they have their own independent hardware, but in reality, the hypervisor is controlling and sharing the underlying hardware resources.

This is a hypervisor.

Compute Host, this is the hardware, it majorly has 3 components
- CPU
- RAM
- Network
- Storage (Optional)

> [!IMPORTANT]
> We are keeping storage as optional because the Hypervisor does not virtualise
> the storage.

We need to spread the power of the above hardware amongst multiple users, that is
where the Hypervisor comes in.

---

#### Hypervisor

A software layer, that sits on top of the compute host and virtualizes all of the functions of the host.

---

### Containers vs VMs

#### Virtual Machines

- Virtualisation on the hardware level.
- It virtualizes the host itself.
- Machine/Hardware Isolation.
- We are creating different machines on our hardware.

#### Containers

- HW -> Kernel -> Operating System (Host OS) -> Containers
- Process isolation.
- Namespaces (For separating processes), CGroups (For monitoring resources).

#### Hardware Assisted Virtualisation

