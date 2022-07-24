#!/bin/bash
echo "What is your name?"
read USER_NAME
echo "Hello $USER_NAME"
echo "Creating a  ${USER_NAME}_file"
touch "${USER_NAME}_file"