#!/bin/bash

timeout 260 ./cenzori <./input.txt >tmp
sort <tmp >tmp.txt
./cmp.py
