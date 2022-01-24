# Linux Commands

### `touch`

Allows you to create a blank file.

### `cat`

Create a new file, or show the content of the file

```shell
cat > filename
# create a new file and enter into insertion mode

cat filename
# to display the contents of the file

cat file1 file2 > file3
# joins the contents of file1 and file2 and stores them in file3

```

### `ls`

List out all the files and directorires available in a directory.

There some options
```shell
ls -a : shows the hidden files
ls -R : list out all the files ain the sub-directories.
ls -l : will list the files and directories with 
        detailed information like the permissions size, owner, etc. 
```

![](https://i.imgur.com/SddjpCr.png)


1.  First Column: represents file type and permission given on the file. Below is the description of all type of files.
2. Second Column: represents the number of memory blocks taken by the file or directory.
3. Third Column: represents owner of the file. This is the Unix user who created this file.
4. Fourth Column: represents group of the owner.
5. Fifth Column: represents file size in bytes.
6. Sixth Column: represents date and time when this file was created or modified last time.
7. Seventh Column: represents file or directory name.

```shell
 ❯ ls -l
total 52
drwxr-xr-x 2 siddhantprateek siddhantprateek  4096 Jan 14 18:02 Desktop
 |-|  <- owner permission

drwxr-xr-x 2 siddhantprateek siddhantprateek  4096 Jan 16 16:51 Documents
    |-| <- group permission

drwxr-xr-x 4 siddhantprateek siddhantprateek  4096 Jan 16 20:35 Downloads
       |-|  <- other permission

w: file can be altered
r: file can be readable
x: file can be executable
```

| Number | Octal Permission Representation | Ref |
| -------- | -------- | -------- |
| 0     | No Permission     | `---`     |
| 1     | Execute Permission     | `--x`     |
| 2     | Write Permission     | `-w-`     |
| 3     | Execute and write Permission: 1(execute) + 2(write)     | `-wx`     |
| 4     | Read Permission     | `r--`     |
| 5     | Read and Execute Permission: 4(read) + 1(execute)     | `r-x`     |
| 6     | Read and Write Permission: 4(read) + 2(write)     | `rw-`     |
| 7     | All Permission: 4(read) + 2(write) + 1(execute)     | `rwx`     |
### `chmod`

This commands is used to change file or directory permission.

There are two do ways to change a file permission:
- symbolic mode
- absolute mode

Types of permission a file have :
- [`u`]Owner permissions: The owner's permissions determine what actions the owner of the file can perform on the file.
- [`g`]Group permissions: The group's permissions determine what actions a user, who is a member of the group that a file belongs to, can perform on the file.
- [`o`]Other (world) permissions: The permissions for others indicate what action all other users can perform on the file.

| Chmod Operator | Description | 
| -------- | -------- |
|  `+` | Adds the designated permission(s) to a file or directory. |
|  `-` |  Removes the designated permission(s) from a file or directory. |
|  `=` | Set the designated permission(s) |
#### Symbolic Mode

For understanding symbolic mode, lets take a example
```shell
# creates a file name 'text'
$ touch text

# to see what permission does the file have
$ ls -l text 
-rwxrwxrwx 1 mavrick mavrick 0 Jan 23 17:51 text

# to remove excute permission from Owner 
$ chmod u-x text

$ ls -l text
-rw-rwxrwx 1 mavrick mavrick 0 Jan 23 18:02 text

```

#### Absolute mode

```shell
chmod 777 text
# gives all the permission to `text` file
# u = 7 (rwx)
# g = 7 (rwx)
# o = 7 (rwx)

chmod 742 text
# gives all the permission to `text` file
# u = 7 (rwx)
# g = 4 (r--)
# o = 2 (-w-)
```


### `systemctl`

### `systemd`


### `vim`

General vim commands and shortcuts

| Keys | Description | 
| -------- | -------- |
| `i` | insert mode | 
| `esc` | to escape the insertion mode |
| `:q` | to exit vim editor |
| `:wq` | to write and exit the vim editor  | 


### `grep`

The grep command can also search the contents of files. Here we’re searching for the word `cool` in all text files in the current directory.
```shell
grep cool *.txt
```

### `echo`

Echo command prints (echos) of a string of text to the terminal.

```shell
echo "Hello Linux" 
```

### `history`
History command shows all the commands that you have used in the past for the current terminal session.

```shell
history
```

This can help you refer to the old commands you have entered and re-used them in your operations again.

### `man`

To show the mannuak instrution of any commands

```shell
man command_name

#example
$ man ls
$ man git
$ man cd

```

### `rm`

To delete an existing file we use `rm` command

```shell
rm filename
```

### `pwd`

Displays the path of the present working directory

### `mv`
 
 To change the file name

```shell
mv old_file_name new_file_name
```

### `cp`

To make a copy of a file

```shell
cp source_file destination
```

### `diff`

Compares the content of two file line by line

```shell
diff file1 file2
```
### `wc`

To count total number of lines, words contained in a file.

```shell
wc filename
```

#### Some short Commands

- `whoami` While you're logged in to the system, you might be willing to know.
- `who` who is logged in to the computer at the same time.