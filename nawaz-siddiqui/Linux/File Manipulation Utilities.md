# File Manipulation Utilities

Linux provides various file manipulation utilities to make sure you work with text file seamlessly. HEre are some of the file utilities:

## Sort

As the name suggests, Sort is used to rearrange the lines of a text file in ascending or descending order. Some examples of the Sort command are:

| Syntax | Usage |
| --- | --- |
| sort filename | Sort the lines in the specified file, according to the characters at the beginning of each line |
| cat file1 file2 | sort | Combine the two files, then sort the lines and display the output on the terminal |
| sort -r filename | Sort the lines in reverse order |
| sort -k 3 filename | Sort the lines by the 3rd field on each line instead of the beginning |

## Uniq

Uniq removes duplicate consecutive lines in a text file. It is often advised to use uniq with sort command in order to get the duplicate text removed. Examples:

**sort file1 file2 | uniq > file3**

Above command sort file1 and file2 and pipes the unique output to file3.

**uniq -c filename**

The above command is used to count the duplicate entries in the text file.

## Paste

Used to combine to file in a good-looking manner. Suppose you have two file, one with the student names and the other with their roll no, the above command can paste them in a single file without much effort. For example:

**$ paste children.txt rollno.txt**

If you want to separate consecutive values in parallel, use -d (delimiters). Some common delimiters are 'space', 'tab', '|', 'comma', etc. Example:

**$ paste -d, file1 file2**

The above file uses comma(,) as a delimiter to separate the values of file1 and file2.

## Join

Join is the enhanced version of paste. Suppose in a file you have roll no. and name of the student, on the other file you have roll no. and their test scores, you can combine both of them using roll no. as a common reference. For example, the below  command joins two file:

**join file1 file2**

## Split

The command breaks up large text files. The default breakup value is 1000 lines segment. Example:

**split file1**

## Regular Expressions and Search Patterns

These are text strings, used while searching a specific pattern. Some of the symbols and their uses are:

| Search Patterns | Usage |
| --- | --- |
| .(dot) | Match any single character |
| a|z | Match a or z |
| $ | Match end of a line |
| ^ | Match beginning of a line |
| * | Match preceding item 0 or more times |

## Grep

Grep is used as a text searching tool. For example:

| Pattern | Usage |
| --- | --- |
| grep green file1 | Matches and prints the line containing “green” in file1 |
| grep -n green file1 | Prints the line number containing the word “green” in file1 |
| grep -nC 2 green file1 | prints the context of 2 lines above and below containing the word “green” in file1 |
| grep -r “green” | recursively search for the word “green”  in the current directory and nested subdirectory. |
| grep -v [pattern] filename | Print all lines that do not match the pattern |
