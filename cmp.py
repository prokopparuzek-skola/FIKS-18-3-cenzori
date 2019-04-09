#!/usr/bin/python3

i = 0
with open("output.txt") as r:
    for s in r:
        i += 1

with open("input.txt") as r:
    count = int(r.readline().strip())

with open("output.txt", mode="a") as r:
    m = count - i
    for i in range(m):
        r.write("-1\n")
