- Git log
    
    By default means without any flags git log command will show will SHA, author, date and commit message. 
    
    Some flags related to git log : 
    
    - `git log  --oneline` will show you one commit per line, the seven characters of SHA, the commit message.
    - `git log --stat` will show you:
        - the files that were modified in each commit
        - the number of lines added or removed
        - a summary line with the total number of files and lines changed
    - `git log --oneline --graph` works best if you want to have a look at the project in a brief way. It is not good for big and large codebase. It becomes chaos.
    

---

- Git Stash
    
    When I started with git then I was confused about the usage of this command. This functionality is useful when you’ve made changes to a branch that you aren’t ready to commit, but you need to switch to another branch.
    

---