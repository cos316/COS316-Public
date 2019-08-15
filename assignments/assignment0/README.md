# COS316, Assignment 0: Virtual Machines

## Due: [Date TBD] at [Time TBD]

Welcome to COS 316: Principles of Computer System Design!

Through this and the following  assignments, you will gain hands-on experience with real-world systems and network programming:

* You will write a program that allows your computer to communicate with another,
  be it across the room or across the world.
* You will program a web framework to process HTTP requests and route them to
  controllers.
* You will implement an in-memory cache to understand how systems efficiently
  utilize their limited resources.
***TODO: Information about the other assignments of the course, as this becomes available.***

The programming assignments are designed to be challenging but manageable in the
time allotted.
If you have questions, want to suggest clarifications, or are struggling with
any of the assignments this semester, please come to office hours,
ask questions on Piazza, or talk to an instructor before or after class.

You are not allowed to copy or look at code from other students. However, you are
welcome to discuss the assignments with other students without sharing code.

Let's get started!

## Set up a virtual machine

The first part of this assignment is to set up the virtual machine (VM) you will
use for the rest of the course. This will make it easy to install all dependencies
for the programming assignments, saving you the tedium of installing individual packages and ensuring your development environment is correct.

### Install Vagrant

Vagrant is a tool for automatically configuring a VM using instructions given in a
single "Vagrantfile."

* **Mac users**: Download and install [Vagrant](https://www.vagrantup.com/downloads.html).

* **Windows users**: Download and install [Vagrant](https://www.vagrantup.com/downloads.html).
  You will be asked to restart your computer at the end of the installation.
  Click Yes to do so right away, or restart manually later, but don't forget to
  do so or Vagrant will not work! (We will prompt you again, later in this
  document, when it's necessary for you to restart your computer.)

* **Linux users**: First, make sure your package installer is up to date by
running the command `sudo apt-get update`. To install Vagrant, you must have the
"Universe" repository on your computer; run `sudo apt-add-repository universe`
to add it. Finally, run `sudo apt-get install vagrant` to install vagrant.

### Install VirtualBox

VirtualBox is a VM provider (hypervisor).

* **Mac users**: Download and install [VirtualBox platform packages](https://www.virtualbox.org/wiki/Downloads).
  During the installation process you may see one or more messages from the
  operating system that VirtualBox is requesting access to certain system features.
  This is normal. The installation may fail the first time around, and prompt you
  to allow VirtualBox this access via the Security tab of System Preferences.
  Click the "Allow" button at the bottom of the window, and run the installer again.
  This time it should finish without issue.

* **Windows users**: Download and install [VirtualBox platform packages](https://www.virtualbox.org/wiki/Downloads).
  Use all the default installation settings, but you can uncheck the "Start Oracle VirtualBox 6.x.x after installation" checkbox.

* **Linux users**: Run the command `sudo apt-get install virtualbox`.

Note: This will also install the VirtualBox application on your computer. You should never need to run it, though it is a useful piece of software and it may be helpful.

### Install Git (and SSH-capable terminal on Windows)

Git is a distributed version control system. This course does not require any
knowledge of Git (all commands you need to know will be provided). However, if
you aren't familiar with Git, you should consider learning more. Good version
control is an invaluable tool for any software developer, and you will find it
extremely useful for future courses (e.g. COS333, COS426).

* **Mac users:** Download [Git](https://git-scm.com/downloads).
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

* **Windows users:** There are several Git and SSH-capable terminals available
  for Windows. Which you choose is a matter of personal taste.
  My personal recommendation is to use Cmder, a terminal supporting several
  shells that may make your life as a Windows developer less painful. Whichever
  option you choose, *you only need to install Git __once__, using __one__ of the
  options below.*

  * **Installing vanilla Git:** You will be given many options to choose from
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
    You may use these if you can get them working, but your mileage may vary, and course staff can't guarantee they will work as expected.

* **Linux users:** Run `sudo apt-get install git`.

### Install X Server

You will need an X Server to input certain commands to the virtual machine.

* **Mac users:** Install [XQuartz](https://www.xquartz.org/).
  Log out and log back in to complete the installation
  (as mentioned by the prompt at the end).

* **Windows users:** Install [Xming](https://sourceforge.net/projects/xming/files/Xming/6.9.0.31/Xming-6-9-0-31-setup.exe/download).
  Use default options and uncheck "Launch Xming" at the end.

* **Linux users:** The X server is pre-installed!

### Restart your computer

If you are a Windows user and you have not restarted your computer in a previous step, do so now, or your newly-installed software may not work.

If you are a Mac user and you have not logged out and back in the previous step, do so now.

### Clone course Git repository

Open your terminal (use the one mentioned in step 3 if using Windows) and `cd`
to wherever you want to keep files for this course on your computer.

Run `git clone https://github.com/COS316/COS316-Public`to download
the course files from GitHub.

`cd COS316-Public` so that you are inside the directory just downloaded.

### Provision virtual machine using Vagrant

From the `COS316-Public` directory you just entered, run the command  `vagrant up`
to start the VM and  provision it according to the Vagrantfile.
**Creating the VM for the first time may take 20 minutes or more**
(it requires the download of several very large files),
but status messages should be consistently printing to the console.
You may see warnings/errors in red, such as "default: stdin: is not a tty", but
you shouldn't have worry about them.

* **Note 1**: The following commands will allow you to stop the VM at any point
  (such as when you are done working on an assignment for the day):

  * `vagrant suspend` will save the state of the VM and stop it.
  * `vagrant halt` will gracefully shutdown the VM operating system and power
    down the VM.
  * `vagrant destroy` will remove all traces of the VM from your system.
    If you have important files saved on the VM (like your assignment solutions),
    **DO NOT** use this command.

  Additionally, the command `vagrant status` will allow you to check the status
  of your machine in case you are unsure (e.g. running, powered off, saved...).
  You must be in some subdirectory of the directory containing the Vagrantfile to
  use any of the commands above, otherwise Vagrant will not know which VM you are
  referring to.

* **Note 2**: The VirtualBox application that was installed in Step 2 provides a
  visual interface as an alternative to these commands, where you can see the
  status of your VM and power it on/off or save its state.
  However, we recommend you not use this interface since it is not integrated with
  Vagrant, and it does not make typing commands any easier.
  It is also not an alternative to the initial `vagrant up` since this creates the VM.

### Test SSH to VPN

Run `vagrant ssh` from your terminal.
This is the command you will use every time you want to access the VM.
If it works, your terminal prompt will change to `vagrant@cos316:~$`.
All further commands will execute on the VM.
You can then run `cd /vagrant` to get to the course directory that's shared
between your regular OS and the VM.

Vagrant is especially useful because it allows for a shared directory structure.
You don't need to copy files to and from the VM.
Any file or directory in the `COS316-Public` directory where the `Vagrantfile` is
located (or in any other synced folder) is automatically shared between your
computer and the virtual machine.
This means you can use your IDE of choice from outside the VM to write your code
(but you should still build and execute from within the VM).

The command `logout` will stop the SSH connection at any point.
You may also press `Ctrl-D` to send an EOF and stop the connection.

### Extra Note for Windows users

Line endings are symbolized differently in DOS (Windows) and Unix (Linux/MacOS).
In the former, they are represented by a carriage return and line feed
(CRLF, or "\r\n"), and in the latter, just a line feed (LF, or "\n").
Given that you ran `git pull` from Windows, git detects your operating system and
adds carriage returns to files when downloading. This can lead to parsing problems
within the VM, which runs Ubuntu (Unix).

Fortunately, this only seems to affect shell scripts (\*.sh files) we may provide
for testing. The `Vagrantfile` is set to automatically convert all files back to
Unix format, so **you shouldn't have to worry about this**.

**However**, if you want to write/edit shell scripts to help yourself with testing,
or if you encounter this problem with some other type of file, use the preinstalled
program `dos2unix`.
Run `dos2unix [file]` to convert it to Unix format (before editing/running in VM),
and run `unix2dos [file]` to convert it to DOS format (before editing on Windows).
A good hint that you need to do this when running from the VM is some error message
involving `^M` (carriage return).
A good hint you need to do this when editing on Windows is the lack of new lines.
Remember, doing this should only be necessary if you want to edit shell scripts.
You can also configure many modern editors (Atom, VSCode, Sublime) to use the correct
line endings.

### Go take a break. You've earned it!

You're all set to start working on assignment 1!
