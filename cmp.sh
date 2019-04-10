#!/bin/bash

timeout 260 ./cenzori <./input.txt | sort > tmp.txt
./cmp.py
