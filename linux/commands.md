## Here are the commands we ran so far

Iplink
iwctl
Station list
station wlan0 get-networks
station wlan0 show
systemctl status sshd.service
Cat /etc/os-release
cat /sys/firmware/efi/fw_platform_size
Lsblk : to check the current partition of the file system

we are using logical volumes instead

- fdisk nvme0n1
- do the partitions
- then we move on to encryption
- learn about what dm-crypt is and what it does.
- what is the use of mapper here?
- learn about mounting, how it works in the linux system.
- check out fdisk more in depth.
- there is a difference with setting things up using lsblk vs df -h.
- then we go systemd
- mkinit
- initramfs
- mkinitcpio -P
-

- pvcreate /dev/mapper/cryptlvm
- vgcreate ganesh /dev/mapper/cryptlvm
- lvcreate -L 8G ganesh -n swap
- lvcreate -L 32G ganesh -n root
- lvcreate -l 100%FREE -n home

- mkfs.ext4 /dev/ganesh/root
- mkfs.ext4 /dev/ganesh/home
- mkfs.ext4 /dev/ganesh/swap
- mkswap /dev/ganesh/swap

- mkfs.fat -F32 /dev/nvme0n1p1
- mount --mkdir /dev/nvme0n1p1 /mnt/boot

- mount /dev/ganesh/root /mnt
