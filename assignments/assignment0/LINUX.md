# VM Setup Instructions for Linux

## The short way:

If you know what you are doing, and are just looking for commands to paste in a
terminal, here you go:

1. Debian, Ubuntu, etc.:

    ```bash
    sudo apt-get update
    sudo apt-add-repository universe
    sudo apt-get install vagrant virtualbox git
    ```

1. Arch Linux, Manjaro, etc.:

    ```bash
    sudo pacman -S vagrant virtualbox git
    ```

If you would like further explanation as to what these commands do, read on:

## The detailed way:

### Install Vagrant

> For Arch Linux / Manjaro users, please refer to [Vagrant - ArchWiki](https://wiki.archlinux.org/index.php/Vagrant).

Vagrant is a tool for automatically configuring a VM using instructions given in a
single "Vagrantfile."

* First, make sure your package installer is up to date by
  running the command `sudo apt-get update`. To install Vagrant, you must have the
  "Universe" repository on your computer; run `sudo apt-add-repository universe`
  to add it. Finally, run `sudo apt-get install vagrant` to install vagrant.

### Install VirtualBox

> For Arch Linux / Manjaro users, please refer to [VirtualBox - ArchWiki](https://wiki.archlinux.org/index.php/VirtualBox).

VirtualBox is a VM provider (hypervisor).

* Run the command `sudo apt-get install virtualbox`.

Note: This will also install the VirtualBox application on your computer.
You should never need to run it, though it is a useful piece of software and
it may be helpful.

### Install Git (and SSH-capable terminal on Windows)

> For Arch Linux / Manjaro users, please refer to [Git - ArchWiki](https://wiki.archlinux.org/index.php/Git).

Git is a distributed version control system. This course does not require any
knowledge of Git (all commands you need to know will be provided). However, if
you aren't familiar with Git, you should consider learning more. Good version
control is an invaluable tool for any software developer, and you will find it
extremely useful for future courses (e.g. COS333, COS426).

* Run `sudo apt-get install git`.

### Install X Server

> For Arch Linux / Manjaro users, please refer to [Xorg - ArchWiki](https://wiki.archlinux.org/index.php/Xorg).

You will need an X Server to input certain commands to the virtual machine.

Luckily, the X Server is pre-installed on Linux, so this step is taken care of.

### Return to the Assignment 0 Instructions

Now that you have installed the required software, continue following the
[main Assignment 0 instructions](README.md).
