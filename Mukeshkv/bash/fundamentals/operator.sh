#!/bin/sh

val=`expr 2 + 2`
echo "Total value : $val"

a=10
b=20
echo "Sum of $a & $b = `expr $a + $b`"
echo "Difference of $a & $b = `expr $a - $b`"
echo "Division of $a & $b = `expr $a / $b`"
echo "Multiplication  of $a & $b = `expr $a * $b`"
echo "% of $a & $b = `expr $b % $a`"
echo "is =  $a & $b = [$a == $b]"
echo "is!= [$a != $b]" [$a != $b]
