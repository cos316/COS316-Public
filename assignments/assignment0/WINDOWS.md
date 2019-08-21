# VM Setup Instructions for Windows

### Install Vagrant

Vagrant is a tool for automatically configuring a VM using instructions given in a
single "Vagrantfile."

* Download and install [Vagrant](https://www.vagrantup.com/downloads.html).
  You will be asked to restart your computer at the end of the installation.
  Click Yes to do so right away, or restart manually later, but don't forget to
  do so or Vagrant will not work! (We will prompt you again, later in this
  document, when it's necessary for you to restart your computer.)

### Install VirtualBox

VirtualBox is a VM provider (hypervisor).

* Download and install [VirtualBox platform packages](https://www.virtualbox.org/wiki/Downloads).
  Use all the default installation settings, but you can uncheck the "Start Oracle
  VirtualBox after installation" checkbox.

Note: This will also install the VirtualBox application on your computer.
You should never need to run it, though it is a useful piece of software and
it may be helpful.

### Install Git (and SSH-capable terminal on Windows)

Git is a distributed version control system. This course does not require any
knowledge of Git (all commands you need to know will be provided). However, if
you aren't familiar with Git, you should consider learning more. Good version
control is an invaluable tool for any software developer, and you will find it
extremely useful for future courses (e.g. COS333, COS426).

There are several Git and SSH-capable terminals available for Windows.
Like choosing a text editor, choosing a terminal is a matter of personal taste.
My personal recommendation is to use Cmder, a terminal supporting several
shells that may make your life as a Windows developer less painful. Whichever
option you choose, *you only need to install Git __once__, using __one__ of the
options below.*

* **Installing vanilla Git:** Download [Git](https://git-scm.com/downloads).
  You will be given many options to choose from
  during the installation; using all the defaults will be sufficient for this
  course (you can uncheck "View release notes" at the end).
  The installation includes an SSH-capable Bash terminal usually located at
  `C:\Program Files\Git\bin\bash.exe`.

  You should use this as your terminal in this class, unless you prefer another SSH-
  capable terminal (the Windows command prompt `cmd.exe` will not work). Feel free to
  create a shortcut to it, but copying and pasting the executable somewhere else will
  not work.
  You may also right-click within a directory in Windows Explorer and select
  "Git Bash here" to open a bash terminal in that directory.

* **Installing Git via Cmder:** Download [Cmder](https://cmder.net/).
  Be sure to select the *Full installation*, and *not* the Mini installation,
  which does not include Git. Follow the directions on the site; that is,
  extract the zip file to a directory of your choosing. If in doubt, `C:\cmder`
  will suffice. Feel free to make a shortcut to `cmder.exe` for easier access.

  At this point you are done, but if you so desire, you may optionally
  [configure your `PATH` and other environment variables](https://github.com/cmderdev/cmder/wiki/Setting-up-Environment-Variables)
  and/or [add Cmder to your right-click context menu](https://github.com/cmderdev/cmder/blob/master/README.md#context-menu-integration).

* **Other options for Windows:** In recent years Microsoft has taken steps to make
  Windows more "developer-friendly", including initiatives like
  [Windows Subsystem for Linux (WSL)](https://docs.microsoft.com/en-us/windows/wsl/about),
  [Windows Terminal 2019](https://devblogs.microsoft.com/commandline/introducing-windows-terminal/),
  and a native [Windows OpenSSH-based SSH client](https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse).
  You may use these if you can get them working, but your mileage may vary,
  and course staff can't guarantee they will work as expected.

### Install X Server

You will need an X Server to input certain commands to the virtual machine.

* Install [Xming](https://sourceforge.net/projects/xming/files/Xming/6.9.0.31/Xming-6-9-0-31-setup.exe/download).
  Use default options and uncheck "Launch Xming" at the end.

### Restart your computer

If you have not restarted your computer in a previous step, do so now, or your
newly-installed software may not work.

When you have done this, you may continue following the
[main Assignment 0 instructions](README.md).
