# COS316, Assignment 1: Virtual Machines & Socket Programming

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

### Step 1: Install Vagrant

Vagrant is a tool for automatically configuring a VM using instructions given in a
single "Vagrantfile."

* **Mac users**: Download and install [Vagrant](https://www.vagrantup.com/downloads.html).

* **Windows users**: Download and install [Vagrant](https://www.vagrantup.com/  downloads.html).
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

***TODO: Update repo URL***
Run `git clone https://github.com/PrincetonUniversity/COS316-Public`to download
the course files from GitHub.

`cd COS316-Public/assignments` to enter the course assignment directory.

### Provision virtual machine using Vagrant

From the `assignments` directory you just entered, run the command  `vagrant up`
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
Any file or directory in the `assignments` directory where the `Vagrantfile` is
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

## Socket Programming

**The remainder of this assignment must be done individually.**

As discussed in lecture, socket programming is the standard way to write programs
that communicate over a network. While originally developed for Unix computers
programmed in C, the socket abstraction is general and not tied to any specific
operating system or programming language. This allows programmers to use the socket
mental model to write correct network programs in many contexts.

This part of the assignment will give you experience with basic socket programming.
You will write two programs for transmitting text messages over the Internet: a
client for sending and a server for receiving.
Both the client and server must be written in Go.

The client and server programs should meet the following specifications.
Be sure to read these meticulously before and after programming to make sure your
implementation fulfills them:

### Server specification
* The server program should listen on a socket, wait for a client to connect,
  receive a message from the client, print the message to stdout, and then wait
  for the next client indefinitely.
* The server should take one command-line argument: the port number to listen on
  for client connections.
* The server should accept and process client communications in an infinite loop,
  allowing multiple clients to send messages to the same server. The server should
  only exit in response to an external signal (e.g. SIGINT from pressing `ctrl-c`).
* The server should gracefully handle error values potentially returned by socket
  programming library functions (see specifics below). Errors related to handling
  client connections should not cause the server to exit after handling the error;
  all others should.
* The server should maintain a client queue and handle multiple client connection
  attempts sequentially. The good news is, this is the default behavior if you are
  using `net.Listen()`. **You do not need to do anything extra to satisfy this
  requirement.** In real applications, a TCP server would fork a new process to handle
  each client connection concurrently, but that is **not necessary** for this
  assignment.

### Client specification

* The client program should contact a server, read a message from stdin, send the     
  message, and exit.
* The client should read and send the message *exactly* as it appears in stdin
  until reaching an EOF (end-of-file).
* The client should take two command-line arguments: the IP address of the server
  and the port number of the server.
* The client must be able to handle arbitrarily large messages by iteratively
  reading and sending chunks of the message, rather than reading the whole message
  into memory first.
* The client should handle partial sends (when a socket only transmits part of
  the data given in the last send operation) by attempting to re-send the rest of
  the data until it has all been sent.
* The client should gracefully handle error values potentially returned by socket     
  programming library functions.

### Error Handling

Generally speaking, there are several reasonable actions that a program might take upon realizing that it has encountered an error, and you may (or may not) sometimes need to take one or more of these actions.

* **Attempt to recover:**
  Some errors may arise due to chance events like a busy or noisy network, and in these
  cases it is possible (and desirable) to try to recover gracefully, perhaps by trying
  the exact same operation again, or by tweaking some values first and then retrying.
* **Crash:**
  On the other hand, some errors cannot be recovered from at runtime. If the user
  requests access to some resource that is already being used by another process,
  there is no straightforward way to recover, and crashing (with an informative
  message) would be an acceptable response.
* **Print error message:**
  Especially for fatal errors that cause your program to crash, it is good style to print out a message indicating what has gone wrong. For non-fatal errors, you may find it useful to print messages for debugging purposes. For your final submit, you should make an effort to minimize output by commenting out any debugging statements.
* **Do nothing:**
  It is generally poor style to leave potential errors unhandled, as your program might continue executing, believing everything to be OK, only to crash later on in a way that will be much harder to debug. Your program should make an effort to handle all reasonable errors that may arise.

#### Error Handling in Go

Go has several error handling functions that may be of use to you:

* `log.Fatal(message string)`<sup>[(docs)](http://golang.org/pkg/log/#Fatal)</sup>:
  Print message to `os.Stderr` and terminate the program with a return code of 1.

* `log.Print(message string)`<sup>[(docs)](https://golang.org/pkg/log/#Print)</sup>:
  Prints message to stderr.

* `log.Panic(message string)`<sup>[(docs)](https://golang.org/pkg/log/#Panic)</sup>:
  Prints the error message, and then calls `panic()`, which propagates the error, and
  prints a stack trace if unhandled. `panic()` is similar to Java's `throw`.
  It differs from `log.Fatal()` in that deferred functions are executed before the
  program exits (perhaps freeing resources or flushing buffers).
  An interesting note is that callers can recover from a panic using `recover()` (analagous to `catch` in Java), but you will *not* need to make use of this functionality for this assignment. See [this blog post](https://blog.golang.org/defer-panic-and-recover) for more on `defer()`, `panic()`, and `recover()`.

### Getting started

Do all building and testing on the Vagrant VM. You may either write your code on
the Vagrant VM (both Emacs and Vim text editors are pre-installed) or directly on
your OS (allowing you to use any editor you have installed).

***TODO: Revise this once the Vagrant directory structure is finalized***
After running `vagrant ssh` from your terminal, run `cd /vagrant` to get to the course
directory. We have provided scaffolding code in the `assignment1` directory. *You
should read and understand this code before starting to program.*

You should program only in the locations of the provided files marked with `TODO` comments. There is one `TODO` section for the client and one for the server.

You can add functions if you wish, but do not change file names, as they will be used for automated testing.

The following section provides details for implementing the client and server programs in Go.

### Go

The documentation for Go socket programming is located [here](https://golang.org/pkg/net/).  
The overview at the top and the section on the [Conn type](https://golang.org/pkg/net/#Conn) will be most relevant.
You may also find the buffered [Reader](https://golang.org/pkg/bufio/#Reader)
and [Writer](https://golang.org/pkg/bufio/#Writer) types to be useful, but you
aren't required to use them, and you can construct a working solution without them.

The Go language (Golang) documentation can be cryptic, so be sure to familiarize
yourself with the language a bit first, especially if you are new to Go. You may
find the [Tour of Go](https://tour.golang.org/list) documentation useful if you
have never used Go before.

The files `client.go` and `server.go` contain the scaffolding code. You will need
to add socket programming code in the locations marked `TODO`. The reference
solutions have roughly 40  (well commented and spaced) lines of code in the
`TODO` sections of each file. Your implementations may be shorter or longer.

You should build your solution by running `make` in the `assignment1` directory. Your code *must* build using the provided Makefile.
The server should be run as `./server [port] > [output file]`.
The client should be run as `./client [server IP] [server port] < [message file]`.
See "Testing" for more details.

### Testing

You should test your implementation by attempting to send messages from your
client to your server.
The server can be run in the background (append a `&` to the command) or in a separate
SSH window.
You should use `127.0.0.1` (i.e. the "localhost", or "loopback", address) as the
server IP and a high server port number between 10000 and 60000.

You can kill a background server with the command `fg` to bring it to the foreground then `ctrl-c`. Conversely, you can send a foreground process to the background by hitting `ctrl-z` to suspend the process, and typing the command `bg` to resume the process in the background.

The Bash script `test_client_server.sh` will test your implementation by attempting to send several different messages between your client and server.
The messages are the following:

0. The short message "Go Tigers!\n"
0. A long, randomly generated alphanumeric message
0. A long, randomly generated binary message
0. Several short messages sent sequentially from separate clients to one server
0. Several long, random alphaumeric messages sent concurrently from separate clients to one server

Run the script as

`./test_client_server.sh [server port]`

If you get a permissions error, run `chmod 744 test_client_server.sh` to give the script execute privileges.

The test script will print "SUCCESS" if the message is sent and received correctly. Otherwise it will print a diff of the sent and received message if the diff output is human-readable, i.e., just for tests 1 and 4.

### Debugging hints

Here are some debugging tips. If you are still having trouble, ask a question on Piazza or see an instructor during office hours.

* There are defined buffer size constants in the scaffolding code. Use them. If you are not using one of them, either you have hard-coded a value, which is bad style, or you are very likely doing something wrong.
* There are multiple ways to read and write from stdin/stdout in Go. Any method is acceptable as long as it does not read an unbounded amount into memory at once and does not modify the message.
* If you are using buffered I/O to write to stdout, make sure to call `flush` or the end of a long message may not write.
* Remember to close the socket at the end of the client program.
* When testing, make sure you are using `127.0.0.1` as the server IP argument to the client and the same server port for both client and server programs.
* If you get "address already in use" errors, make sure you don't already have a server running. Otherwise, restart your ssh session with the command `logout` followed by `vagrant ssh`.
* If you are getting other connection errors, try a different port between 10000 and 60000.

### Q&A

* **I'm getting an error when I run the command `vagrant up`. What do I do?**
Many errors/warnings are not a problem and do not need to be addressed, such a `==> default: stdin: is not a tty`. Usually, errors starting with `==> default` should not be worried about, but others should, in particular if they cause the process to be aborted. Use `vagrant status` to see if the VM is running after `vagrant up`; if it is not, then there is a real problem. Here are some known errors and how to fix them:
    * **"A Vagrant environment or target machine is required to run this command..."**:
      you must run `vagrant up` from a subdirectory of the directory containing the Vagrantfile (in the case, `assignments`).
    * **"Vagrant cannot forward the specified ports on this VM, since they would collide with some other application that is already listening on these ports..."**:
      perhaps you cloned the repository twice and the VM is already running on one of them. Since they both use the same port, they cannot run at the same time. You may also have some other application using port 8888. To help find what is using it, follow
      [these](http://osxdaily.com/2014/05/20/port-scanner-mac-network-utility/) instructions for macOS,
      [these](https://techtalk.gfi.com/scan-open-ports-in-windows-a-quick-guide/) for Windows and
      [these](https://wiki.archlinux.org/index.php/Nmap#Port_scan) for Linux (you may have to install `nmap`).
      Use 127.0.0.1 as the IP and 8888-8888 as the port range in your port scan.
    * **"VBoxManage.exe: error: Details: code E_FAIL (0x80004005), component MachineWrap, interface IMachine"**:
      Try rebooting your host machine (i.e. your laptop). If this fails to fix the issue, you may have other virtualization software installed that conflicts with Vagrant or VirtualBox, which you will need to disable.
    * **"Windows detected. Remember to set DISPLAY"**:
      This relates to the X11 server you installed in a previous step. We are not using the X11 server for this assignment, so you can ignore this message for now.

  If this did not help you fix the problem, please ask on Piazza or at office hours.
* **I'm using Cmder on Windows, and the terminal window is flickering rapidly when I'm connected to the virtual machine**
  This is a [known issue](https://github.com/Maximus5/ConEmu/issues/1625) that happens especially when trying to use a terminal application like Emacs or vim. The root cause is an incompatibility between the Windows OpenSSH implementation (which is likely preinstalled on your machine) and the ConEmu software underlying Cmder. You will have to uninstall Windows' version of OpenSSH by following the directions [here](https://github.com/Maximus5/ConEmu/issues/1625#issuecomment-413459750). This is no great loss, as you've already installed a more reliable SSH client as part of the VM setup steps above.

* **Do I need to handle signals such as SIGINT to clean up the server process when the user presses `ctrl-c`?**
  No, it is not necessary in this assignment. The default response to signals is good enough.
  Keep in mind it would be good practice to do so in general, however.
* **Should I use stream (TCP) or datagram (UDP) sockets?**
  Please use stream (TCP) sockets, to ensure that the exact message is delivered.
  Streams ensure reliable, in-order packet transmission, whereas datagram packets are not guaranteed to be delivered.
* **Should the client wait to receive a reply from the server?**
  No, in this assignment it should exit immediately after sending all the data.
* **Should the server handle client connections concurrently (in separate processes)?**
  No, as stated in the client specification, this is not required in this assignment.
  So no need to use `syscall.ForkExec()` or any other system-level `fork()`!

### Submission and grading

To submit your client and server, `git commit` the changes to your code and
`git push` them to your GitHub classroom repository.

See the [GitHub classroom README] ***TODO: Add link once directory structure finalized*** for more detailed instructions.

We will grade your submissions by compiling your client and server, and then
sending messages back and forth between each of your submitted programs and a
reference server or client, as appropriate. Your code will be scored based on
how many different kinds of messages are transmitted correctly, and how well
your implementations adhere to various aspects of the specification. Within a
couple minutes of submitting your assignment, the GitHub autograder will add a
comment to your most recent commit on GitHub, indicating your test results.

You may submit in this way and receive feedback as many times as you like,
whenever you like, but a lateness penalty will be applied to submissions
received after the deadline. 

Code that does not compile is graded harshly; if you want partial credit on code that doesn't compile, comment it out and make sure your file compiles!

Remember that you should submit two files: a server file and a client file, both written in Go!
