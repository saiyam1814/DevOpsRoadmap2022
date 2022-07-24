#!/bin/bash
# Note that there must be no spaces around the "=" sign: VAR=value Is correct; VAR = value doesn't work
STRING="HELLO WORLD!!!" # STRING is not reserved work 
platform="Platform Number is 1"
message="The Train has left"
readonly message
# message="The Train has arrived"  this will throw error as message is read only
unset platform 
num=12
floating_num=12.44
bool=true

#--------------------------------------
echo STRING
echo $STRING
echo $platform # does not print platform as variable is unset
echo $message
echo "Number  : $num"
echo "Decimal : $floating_num"
echo "Boolean :$bool"

echo "File Name: $0"
echo "________________________________________"
sampledir=/var
ls $sampledir
echo "________________________________________"
#---------------------------------------
# In Bash scripting, a global variable is a variable that can be used anywhere inside the script. 
# A local variable will only be used within the function that it is declared in.

sum="global variable"

function localFn {
# Define bash local variable
# This variable is local to bash function only
local sum="local variable" # "local" is bash reserved word
echo $sum
}

echo $sum
localFn # call function



echo $sum # Note the bash global variable did not change

