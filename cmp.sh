#!/bin/bash

timeout 260 ./cenzori <./input.txt >./output.txt
./cmp.py
