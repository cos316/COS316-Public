# COS316, Assignment 0: Virtual Machines

## Due: September 12, 2019 at 23:00

Welcome to COS 316: Principles of Computer System Design!

Through this and the following  assignments, you will gain hands-on experience
with real-world systems and network programming:

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
for the programming assignments, saving you the tedium of installing individual
packages and ensuring your development environment is correct.

### Install Required Software

Before you can set up your VM, you will have to install several required software
packages. Follow the instructions for your operating system below:

* [Windows](WINDOWS.md)
* [Mac OS](MAC.md)
* [Linux](LINUX.md)

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

## FAQ & Troubleshooting:

* **I'm getting an error when I run the command `vagrant up`. What do I do?**

  Many errors/warnings are not a problem and do not need to be addressed, such as
  `==> default: stdin: is not a tty`.
  Usually, errors starting with `==> default` should not concern you,
  but others should, in particular if they cause the process to be aborted.

  Use `vagrant status` to see if the VM is running after `vagrant up`;
  if it is not, then there is a real problem. Here are some known errors and how
  to fix them:

    * **"A Vagrant environment or target machine is required to run this command..."**:

      You must run `vagrant up` from a subdirectory of the directory containing
      the Vagrantfile (in this case, `COS316-Public`).

    * **"Vagrant cannot forward the specified ports on this VM, since they would
      collide with some other application that is already listening on these ports..."**:

      Perhaps you cloned the repository twice and the VM is already running on
      one of them. Since they both use the same port, they cannot run at the same
      time. You may also have some other application using port 8888.

      To help find what is using it, follow
      [these](http://osxdaily.com/2014/05/20/port-scanner-mac-network-utility/) instructions for macOS,
      [these](https://techtalk.gfi.com/scan-open-ports-in-windows-a-quick-guide/) for Windows and
      [these](https://wiki.archlinux.org/index.php/Nmap#Port_scan) for Linux (you may have to install `nmap`).
      Use 127.0.0.1 as the IP and 8888-8888 as the port range in your port scan.

    * **"VBoxManage.exe: error: Details: code E_FAIL (0x80004005), component
    MachineWrap, interface IMachine"**:

      Try rebooting your host machine (i.e. your laptop). If this fails to fix the
      issue, you may have other virtualization software installed that conflicts
      with Vagrant or VirtualBox, which you will need to disable.

    * **"Windows detected. Remember to set DISPLAY"**:

      This relates to the X11 server you installed in a previous step.
      We are not using the X11 server for this assignment, so you can ignore this
      message for now. If you are not using Windows, this message is erroneous
      and should be ignored.

  If this did not help you fix the problem, please ask on Piazza or at office hours.

* **I'm using Cmder on Windows, and the terminal window is flickering rapidly when
  I'm connected to the virtual machine**

  This is a [known issue](https://github.com/Maximus5/ConEmu/issues/1625) that
  is especially pronounced when trying to use a terminal application like Emacs or vim.
  The root cause is an incompatibility between the Windows OpenSSH implementation
  (which is likely preinstalled on your machine) and the ConEmu software underlying
  Cmder. You will have to uninstall Windows' version of OpenSSH by following
  [these directions](https://github.com/Maximus5/ConEmu/issues/1625#issuecomment-413459750).

  This is no great loss, as you've already installed a more reliable SSH client
  as part of the VM setup steps above.
