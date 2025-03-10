## Replacing the hostname assignment command

This command replaces the entry if present, and writes the entry if does not exist.

```shell
while read IP FQDN HOST SUBNET; do
    CMD="grep -q '^127.0.1.1' /etc/hosts || echo '127.0.1.1 ${FQDN} ${HOST}' >> /etc/hosts; sed -i 's/^127.0.1.1.*/127.0.1.1\t${FQDN} ${HOST}/' /etc/hosts"
    ssh -n root@${IP} "$CMD"
    ssh -n root@${IP} hostnamectl hostname ${HOST}
done < machines.txt

```

## TODOs

- Learn more about how the whole resolution process works here.
- I remember it being mentioned somewhere that Kubernetes nodes
register themselves with the API server, so since they are
communicating with 
- What is provided pre-written and what are we configuring

## Here is what I did so far.

- I created 4 VMs with Debian 12 image, using `lima user-v2` network,
to enable **vm->vm** communication.
- Then I enabled ssh, using keys as well as password so that these
machines can talk to each other.
- Then I generated a Key pair i.e. public/private key on the `jumpbox`
machine.
- Then I got the IP addresses of each machine, using `ip a` command
and wrote them in a `machines.txt` file.
- Then use the `ssh-copy-id` command to share the public id with
all the other three machines. This command copies the public SSH
key to the `/root/.ssh/authorized_keys` on the remote system.
- Then we set hostnames and Fully qualified domain name by changing the `/etc/hosts` file.
- Then we define things properly in the lookup tables, which is the 
`/etc/hosts`.

Read and summarise this thread
[PKI decryption and verification](https://chatgpt.com/c/6769b08f-4d5c-800c-8bf2-7d8b9bb50db3)

