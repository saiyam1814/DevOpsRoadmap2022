#! /bin/bash 
# The first indicates what interpreter to use to run the file (in this case, bash)
  
# This is the basic bash script  
# Add Shebang  on top 
# Make this script executable by chmod +x  hello.sh 
# Running the script as bash hello.sh, the first line is ignored.
 
echo " Hello World! Welcome to Bash"  

echo 
echo 
echo -n  "This is new line"
echo "Todays date is: $(date)"
echo "The users name is: $USER"
echo -e "This is a first line.\n and this is a second line"
echo "Welcome to bash scripting"
echo -e "\033[1;31m Colorfull Echo"
echo -e "\033[0;34m Blue Color Echo "
echo -e "\033[1;35m Purple Bold Color Echo "
echo "Abdul      Kalam"       # This is a comment, too!
echo "Mohan Das"
echo "JL * Nehru"
echo Bhagat * Singh
echo Mohan      Bhagwat
echo "Suchitra" Kriplani
echo Taylor "     " Reed
echo "Lala "*" Lajpat"
echo \n
echo `Hargobind` Khorana
echo 'Salman' Khan