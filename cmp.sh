#!/bin/bash

timeout 320 ./cenzori <./input.txt >tmp
sort -g  <tmp >tmp.txt
./cmp.py
