GREP is a powerful pattern searching tool in unix based system. 

 -------------------------------
 | command1 | xargs | command2 |
 -------------------------------

- there are many shell commands that allow you to take input 
- ls | grep "^d" -> all the files/directories that start will d.
- ls | grep "d$" -> all the files/directories that end with d. 
- seq 5 will give you number from 1 to 5. 
- Now if we want to echo then we can't because echo doesn't take standard input. 
- If you will not give command2 then the default is echo
