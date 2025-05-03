- Sometimes we want the applications to know that they have moved.
- To move a VM, you need to move 4 things, CPU state, RAM, Disk and Network.
- Ask about the `veth` thing, what it does and how is it actually working.
- I think conduit is heavily reliant on XDP for redirecting packets while the migration is happening, but what about once the migration is done? How does the client or the server knows/upgrade with the new Connections? at some point redirection should stop right?
  
  I mostly wanna know this because you said that you are not touching the VM or the services running inside.
- I was wondering how exactly do you migrate stateful connections like TCP? how do you copy it? Are you using CRIU? Are we transferring the entire TCB?
- What happens to the packets being transferred during the 50 ms downtime? are they dropped or you have implemented some buffer mechanisms for it? is the data lost? or TCP just handles it through re-transmission (since there was no ACK).
- How do you handle the the process of global steering of packets? are we using Global Load balancers for Anycast IP Addresses?
- Have we achieved the GCP (US) to AWS (Spain) transmission.

- XDP dump
- PWRU
- Veth interfaces

- Traffic coming from the linux kernel is not considered ingress traffic, so it does not trigger an ingress hook.
- He mentions about "when you load eBPF programs you have to mention if they are generic mode or not", what is that about?
- What is a full checksum re-calculation and what is partial re-calculation.
- What is the LLVM thing he is talking about? what is the looping issue he is talking about? and why is it causing verifier issues in checksum calculations? Read the transcript and explain it to me.
- Explain the full checksum calculation process.
- A packet appeared and it didn't come from the kernel routing table. The veth is implemented, there is no routing table between veth internal and veth external.

- Line Rate encrypting VPNs in XDP. Ktls functions in XDP hooks. Handshake outside. Encrypted point to point VPN where the data doesn't have to ever come up to user space.