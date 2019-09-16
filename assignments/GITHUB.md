# General Assignment Instructions

This course uses GitHub Classroom for submitting and grading assignments. This
guide details the steps you will need to take to be ready to use GitHub to
submit your assignments.

## Do these once at the beginning of the course

### Create a GitHub Account

If you already have a GitHub account you are willing to use for this course,
then you're done - proceed to the next step.

Otherwise, visit the [GitHub registration page](https://github.com/join) and
follow the onscreen instructions to register for an account. You will want to
select "Free" as your subscription option, and ensure the "Help me set up an
organization" box is unchecked.

## Do these for the first assignment only

### Accept the GitHub assignment invitation

Before you start working on each new assignment, you will need to accept the
corresponding assignment invitation.

The first time you do so, you will be prompted to provide GitHub Classroom
permission to access various parts of your GitHub account. This is required
in order to allow instructors to view your assignment files, and to allow the
autograder to find your submissions.

![GitHub Authorize](https://github.com/cos316/COS316-Public/blob/master/assignments/images/github-authorize.png)

### Associate your GitHub ID with your NetID (One-time only)

After granting this permission, you will also be prompted to select your NetID
from the class roster, so that your GitHub account ID can be mapped to your NetID.
**Do not skip this step or it will be impossible to give you credit for the assignments.***


## Do these once per new assignment

### Accept the GitHub assignment invitation

Every assignment has its own invitation link, which you will need to click on
to get started with that assignment. Be sure you're logged into the GitHub
account that you will be using for this class.

### Creating or Joining a Team

For assignments where teams are allowed, the next step is to join a team. Only
one member of your group should create a team, and then all other members of your
group should join that existing team. Be sure that you are joining the
correct team, and that your teammate(s) are expecting you to join - *Do not join
a group without the permission of the other members.*

Be careful when creating or joining a team, as there is no way for you to change
teams on your own once you have selected one.

### Getting access to assignment code

Once you have joined a team, you will be provided with a link to your team's
personal assignment repository, including the starter code for that assignment.

Your URL should be of the form:
`https://github.com/cos316/<assignment#>-<assignment-name>-<team-name>`

#### Cloning (downloading) your repository

Once you have access to this repository, you will want to *clone* a copy of the
repository to your own machine so that you can work on it.

There are two ways to accomplish this:

1.  Via command prompt. Open a terminal window using the Git-capable terminal
    from the first assignment setup. Navigate to the directory you are using for
    this course.

    Then enter the following command:

    `git clone <URL of your assignment repo>`

    For example:

    `git clone https://github.com/cos316/assignment1-socket-programming-team-2`

    You will need to enter your GitHub credentials to authenticate yourself.

2.  Via GitHub Desktop. GitHub Desktop is a graphical Git/GitHub interface for
    macOS and Windows. You may find it easier to use, especially if you have
    never used Git or similar version control software before.

    Download GitHub Desktop [here](https://desktop.github.com/). Log in with
    your GitHub account. (If you are not prompted, you can do this manually
    from `File -> Options`).

    Then choose `File -> Clone repository` and select the repo from the list,
    or enter the URL manually.

#### Syncing your repository with Vagrant

Once you have the assignment code downloaded to your machine, you will need to
make it accessible from within your Vagrant VM.

To do this, you will need to uncomment a line in your Vagrantfile. Find the
section in your Vagrantfile marked *Synced Assignment Directories*, and replace
the `<path ...>` argument with the actual path from your current working directory
to your cloned assignment repository.

If you cloned using GitHub Desktop, the default path will be
`~/Documents/GitHub/<repo name>`.

If you cloned using `git clone` from the command-line, the default path will
be a subdirectory of whatever your working directory was what you ran the command.

You should feel free to add, remove, or restructure synced folders as desired to
suit your desired workflow.

The next time you start your VM, your assignment files will be accessible in
`/assignmentN`, where N is the number of the assignment.

### Pushing (submitting) to get online feedback

To receive feedback on your submissions, you will need to push them to GitHub.
If you already know how to do this, you may skip this section.

Editing and saving files in your repository locally (that is, on your laptop or
from within Vagrant) *does not* automatically save these changes to GitHub.

Saving changes to GitHub is a two-step process. First, you will have to *commit*
your changes to the local copy of the repository you downloaded, and then you
have to *push* these commits to the remote copy of the repository, hosted on
GitHub.

As before, there are two ways to do this:

1.  Via command line. To commit your changes, execute the following command
    from within your cloned local copy of the assignment repository:

    `git commit -a -m "<brief summary of changes>"`

    Then, to push your changes to the remote repository hosted on GitHub:

    `git push`

2.  Via GitHub Desktop. Select the changes you would like to commit using the
    change list on the left, and enter a description of the changes in the text
    box on the bottom left. Then click *Commit to master*.

    To push your changes to the remote GitHub repository, click the *Push* button
    in the top toolbar, next to the *Current Branch* dropdown. If someone else
    has made changes to the repository since the last time you pushed, you may
    first need to *fetch* or *pull* these changes, which is done using the same
    button.

### Viewing test results

Every time you or a teammate pushes a change to your GitHub repository, your
submission will be graded as it is at that moment in time, and a snapshot of
your results will be saved on GitHub as a comment on that commit.

To view these comments (for either your most recent submission or any past
submission), click the *X Commits* button at the top of your repository's *Code*
view, where X is the number of commits made so far. Then, click on the comment
button for the commit you would like to inspect, and you will see your results.
