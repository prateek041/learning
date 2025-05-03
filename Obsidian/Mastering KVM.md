## Chapter 1

The main computer on which everything runs is called the host machine and the virtual machines installed on top of that host are called Guests.

Virtualization is the process of duplicating the functions of a physical hardware and present it to the operating system.

## Types of Virtualization

As mentioned earlier, the virtualization process can be used to virtualise anything including the network, storage, access etc. There are a lot of options to what we can actually do here.

There are dedicated software for virtualizing each component, but we can focus on one thing as well if needed.

For our specific use-case the virtualisation we are going to work with is about hiding the underlying hardware so it can be shared by multiple software and guest OS(s).

This is also sometimes called as Platform virtualisation.

### Operating System Virtualisation

This is often also knows as Workload vertualisation because the same host OS is running all the workloads and processes. Where the workloads are isolated from each other through separate FS, memory etc.

This is faster in terms of all the workloads being run by the same host OS. This also called container based virtualisation.

