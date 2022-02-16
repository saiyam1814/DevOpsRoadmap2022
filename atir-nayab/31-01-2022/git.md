# Git

## commit
- add the right changes
- compose a good commit message
- add specific changes of a file to staging area
`git add -p <filename>`
- commit message
  - subject = conise summary of what happened
  - body = more detailed explanation
    - what is now different then before?
    - what's the reason for the change?
    - Is there anything to watch out for / anything particular remarkable?  
  - `git commit` - git will open editor for commit message

## branching strategies 
Agree on a branching worflow in your team
- Git allows you to create branches - but it doesn't tell you how to use them!
- you need a written best practice of how work is ideally structured in your team - to avoid mistakes & collisions.
- if highly depends on your team/team size, on your project, and how you handle releases.
- It helps to onboard new team memebers ("This is how we work here"). 

## integrating changes & structuring Releases
- Mainline development ("Always be integrating")
- State, Release, and Feature Branches
- few branches
- relatively small commits
- high quality testing & QA standards 

## state, Release, and feature branches
branches enhance structure & workflows
- different  types of branches
- ...fulfill different types of jobs

## long running & short lived branches
two main types of branches
- long running
  - exist through the complete lifetime of the project
  - often, they mirror "stages" in your dev life cycle
  - commong convention: no direct commits!
- short lived branches
  - for new feature, bug fixes, refacrotings, experiments...
  - will be deleted after integration (merge/rebase)

## branching strategies
- github flow
  - very simple, very lean: only one long-running. branch ("main") + feature branches
- gitflow
  - more structure, more rules
  - long-running: "main"+"develop"
  - short-lived: features, releases, hotfixes

**The best branching model**
- cosider your project, release cycle, and team
- take inspiration from existing models (like "Gitflow" or "Github Flow")
- ...and create your own model!

## pull request
communicating about and reviewing code
- without a pull request you would jump right to merging, your code

## merge conflicts
- when they might occur
- what they actually are
- how to solve them

## how to undo a conflict and start over
- `git merge --abort`
- `git rebase --abort`

## How to solve a conflict
simply clean up the file

## merge vs rebase
- rebase rewrite commit history
- mrege don't rewrite commit history

**warning**
- do not use rebase on commits that you've already pushed/shared on a remote repository
- instead, use it for cleaning up your local commit history before merging it into a shared team branch.