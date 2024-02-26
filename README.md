Overview
-------------------

`gitpot` is an extensible git implementation library written in **pure Go**. The project is using 
[`Cobra CLI`](https://github.com/spf13/cobra) to get functional CLI capabilties likes parsing commands, flags and parameters for the `gitpot` cli app.



Comparison with git
-------------------

**gitpot** is a toy implementation of git written in **golang**. In it's current state it does not ensure reverse compatibility with git since it doesn't use the exact same compression and storage techniques inside the *objects* directory of the git. 


*git* is a humongous project with years of development by thousands of contributors, making it challenging for **gitpot** to implement all the features. However you can still use basic commands like `init`, `add`, `rm`, `commit`, `branch`, `checkout` and `status`. 

It uses roughly the same file structure inside of the `.git` or `.gitpot` folder and uses the same schema for vital objects like `bolb`, `tree` and `commit`. However they are stored in the plain text format for easy understanding and debugging and can be switched to any compression format whenever needed.


Installation
------------

You can run the below command to get a binary and install *gitpot* as it is from the current active branch:
```
go build
```
After building the binary you can use it in the same way like git.

Examples
---------

This will initialize an empty gitpot project by creating `.gitpot directory`
```
./gitpot init
```
___

This will add the given files to staging area
```
./gitpot add <files or folders>
```
___

This will give information about all the tracked/untracked and staged/unstaged files and folders in the working directroy
```
./gitpot status
```
___

This will commit all the new staged changes in the index in the form of a new commit on the current active branch
```
./gitpot commit --m "commit message"
```
___


This will create a new branch and show all branchnames in case `branchname` is not passed
```
./gitpot branch <branchname>
```
___

This will checkout to the given branchname and make it active branch
```
./gitpot checkout <branchname>
```
___


Contribute
----------

Contributions are more than welcome, if you are interested please take a look at the project codebase. A lot of improvements can be done by parallelising a bunch of stuff an optimising disk read-writes.
