#!/bin/bash

timeout 260 ./cenzori <./input.txt >tmp
sort -g  <tmp >tmp.txt
./cmp.py
