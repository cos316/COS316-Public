# VM Setup Instructions for Mac OS

### Install Vagrant

Vagrant is a tool for automatically configuring a VM using instructions given in a
single "Vagrantfile."

* Download and install [Vagrant](https://www.vagrantup.com/downloads.html).

### Install VirtualBox

VirtualBox is a VM provider (hypervisor).

* Download and install [VirtualBox platform packages](https://www.virtualbox.org/wiki/Downloads).
  During the installation process you may see one or more messages from the
  operating system that VirtualBox is requesting access to certain system features.
  This is normal. The installation may fail the first time around, and prompt you
  to allow VirtualBox this access via the Security tab of System Preferences.
  Click the "Allow" button at the bottom of the window, and run the installer again.
  This time it should finish without issue.

Note: This will also install the VirtualBox application on your computer.
You should never need to run it, though it is a useful piece of software and it
may be helpful.

### Install Git (and SSH-capable terminal on Windows)

Git is a distributed version control system. This course does not require any
knowledge of Git (all commands you need to know will be provided). However, if
you aren't familiar with Git, you should consider learning more. Good version
control is an invaluable tool for any software developer, and you will find it
extremely useful for future courses (e.g. COS333, COS426).

* Download [Git](https://git-scm.com/downloads).
  When you open the .dmg installation file, you will see a Finder window including
  a `.pkg` file, which is the installer. Opening this normally may give you a
  prompt saying it can't be opened because it is from an unidentified developer.
  To override this protection, instead right-click on that `.pkg` file and select
  "Open". This will show a prompt asking you if you are sure you want to open it.
  Select "Yes". This will take you to the (straightforward) installation.

  * Depending on your version of macOS, Terminal may also prompt you to install
    macOS's command-line developer tools from the App Store. We recommend skipping
    this and installing Git directly, so you can be sure the version of Git you
    have installed is as current as possible. That said, we will be using only
    the most basic functionality of Git in this course, so if you have already
    installed developer tools, either version should suffice.

### Install X Server

You will need an X Server to input certain commands to the virtual machine.

* Install [XQuartz](https://www.xquartz.org/).
  Log out and log back in to complete the installation
  (as mentioned by the prompt at the end).

### Log out

If you did not log out in the previous step, make sure to do so before returning
to the [main Assignment 0 instructions](README.md).
