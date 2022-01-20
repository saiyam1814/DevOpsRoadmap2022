## Wildcards and Matching File Names

|Wildcard|Result|
|:-:|:-:|
|?|Matches and single character|
|*|Matches any string of characters|
|[set]|Matchess any character in the set of characters, of example [adf] will match any occurence of a, d, or f|
|[!set]|Matches any character not in the set of characters|

To search for files using the ? wildcard, replace a each unknown character with ?.

## find
It is an extremely useful and often-used utility program in the daily life of a Linux system administrator. It recurses down the filesystem tree from any particular directory (or set of directories) and locates files that match specified conditions. the default pathanme is always the present working directory.

![find](https://courses.edx.org/assets/courseware/v1/102046563ac484a6047300c801886837/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/findubuntu.png)

Searching for file in other directories

![Using find](https://courses.edx.org/assets/courseware/v1/ea8161eed2e8b061792778df2dec70d7/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/findrhel7.png)

Another good use of find is being able to run commands on the files that match your search criteria. The -exec option is used for this purpose

The {} (squiggly brackets) is a placeholder that will be filled with all the file names that result from the find expression, and the preceding command will be run on each one individually.

![Finding and Removing Files that end with .swp](https://courses.edx.org/assets/courseware/v1/cbdf6dc606a39eace7d669077837e628/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch06_screen41.jpg)

## Finding Files Based on Time and Size

**find / -ctime 3**
-ctime - last changed
-atime - accessed/last read
-mtime - modified written
(n) is the number of days
+n greater than n
-n less than n

**find / -size 0**
(c) - bytes
(k) - kilobytes
(M) - megabytes
(G) - gigabytes

![Finding Files Based on Time and Size](https://courses.edx.org/assets/courseware/v1/007f36e6f54ef7e1547682492e8a9b93/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/findsizerhel7.png)

## Package Mangers: Two Levels

Both package management systems operate on two distinct levels: a low-level tool (such as dpkg or rpm) takes care of the details of unpacking individual packages, running scripts, getting the software installed correctly, while a high level tool (such as apt-get, dnf, yum, or zypper) works with groups of packages, downloads packages from the vendor, and figures out dependencies.

![Package Mangers: Two Levels](https://courses.edx.org/assets/courseware/v1/b2cfd35138881e077bfc97915aed86b8/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/Package_Managers.png)

## Working With Different Package Management Systems

![Working with Different package Management Systems](https://courses.edx.org/assets/courseware/v1/272cd4906572ff37d0352281abe81dbe/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/Different_Package_Mmanagement_Tools.png)

|Operation|rpm|deb|
|:-:|:-:|:-:|
|Install package|rpm -i foo.rpm|dpkg --install foo.deb|
|Install package, dependencies|dnf install foo|apt-get install foo|
|Remove package|rpm -e foo.rpm|dpkg --remove foo.deb|
|Remove package, dependencies|dnf remove foo|apt-get autoremove foo|
|Update package|rpm -U foo.rpm|dpkg --install foo.deb|
|Update package, dependencies|dnf update foo|apt-get install foo|
|Update entire system|dnf update|apt-get dist-upgrade|
|Show all installed packages|rpm -qa or dnf list instaled|dpkg -- list|
|Ger information on package|rpm -qil foo| dpkg --listfiles foo|
|Show packages named foo|dnf list "foo"|apt-cache search foo|
|Show all available packages|dnf list|apt-cache dumpavail foo|
|Show all available|dnf list|apt-cache dumpavail|
|What package is file part of?|rpm -qf file|dpkg --search file|