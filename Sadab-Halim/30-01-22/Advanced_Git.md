# Learnt [Advanced Git](https://www.youtube.com/watch?v=qsTthZi23VE)

## Interactive Rebase
It is a tool for optimizing and cleaning up commit history

- change a commit's message
- delete commits
- reorder commits
- combine multiple commits into one
- edit/split an existing commit into multiple new ones

**NOTE:**
- Do NOT use Interactive Rebase on commits that you've already pushed/shared on a remote repo. <br/>
  Instead, use it for cleaning up your local commit history before mergin it into a shared team branch.
  
### Step By Step 
1. How far back do you want to go? <br/>
   What should be the "base" commit?
2. $ git rebase -i HEAD~3
3. In the editor, only determine which actions you want to perform. <br/>
   Don't change commit data in this step, yet!
