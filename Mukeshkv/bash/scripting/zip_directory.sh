#!/bin/bash
#tar -czf testdir.tar.gz  testdir

# Zip with variable file name
fileName=testdir$(date +%Y%m%d).tar.gz
tar -czf $fileName testdir 