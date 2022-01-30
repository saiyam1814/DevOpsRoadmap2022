- Welcome to the Devops World 
- We are targeting to learn devops skills in the timeframe of six months. Well it's a growing field so we cannot learn everything out there and neither anyone expects to know everything from us. 
- The whole syllabus is mentioned in the repository. 

In this guide I will show you step by step how you can contribute your learnings here.

### Steps Involved
- You need a github account and you need the fork the this [repository](https://github.com/saiyam1814/DevOpsRoadmap2022). 

- After your have forked the repo you will see `<username>/DevOpsRoadmap2022` under your account.
  
- Now you need to clone the repo into your local system. For that open the terminal and clone. 
```
git clone https://github.com/<username>/DevOpsRoadmap2022
```
- then you need go inside the directory 
```cd DevOpsRoadmap2022 ```
- You need to set upstream url, run this command to do so:
```

git remote add upstream https://github.com/saiyam1814/DevOpsRoadmap2022.git
```
> For learning more about upstream or origin and github, Watch [Kunal Kushwaha's](https://www.youtube.com/watch?v=apGV9Kg7ics) or [freecodecamp](https://www.youtube.com/watch?v=RGOj5yH7evk) course on github.

- Run these commands to add a different branch:
   - You can have anything as <branch_name> , but by convention it should indicate what you are working on
      -example : <branch_name> can be **add-folder**, when you are adding your folder. 
```
git branch <branch_name>

git checkout <branch_name>
``` 
- Make a directory with your name
```mkdir <Your_name>```
   - For adding spaces between name, type as mkdir firstName\ lastName .
      -example : mkdir Saiyam\ pathak
- To change directory into the folder, run ```cd <Your_name>
- Add a readme.md file , to tell about yourself and show your progress.

- To push your changes to github, Run the following commands.
```
git add .

git commit -m "<username> added Readme file"

git push origin <branch_name>
```

- Go to your github forked repo, You will see an option to "Compare and Pull request".
- Click on that and Then You will see an option to "Create pull request". Click on that.
- that's it you have made your pull request.
- After your request is accepted, You will see the folder of your name on the repository.
 
- You can create day-wise folder or topic-wise for writing down your daily learnings. Checkout Samples in the repository. 
- Learn something new each day and write down your notes. 
- After you have written down your notes create a pull request. We are waiting for your pull request.

- And always remember, Learning should be your goal. Learn at your own pace. 
- You can always ask your doubts on the [Discord Channel](https://saiyampathak.com/discord)
  
 
