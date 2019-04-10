#!/usr/bin/python3

with open("input.txt") as r:
    count = int(r.readline().strip())

r = open("tmp.txt", mode="r")
w = open("output.txt", mode="w")

last = -1
for l in r:
    ac = l.split(":")
    while int(ac[0]) != (last + 1):
        w.write("-1\n")
        last += 1
    w.write("{}".format(ac[1]))

r.close()
w.close()

i = 0
with open("output.txt") as r:
    for s in r:
        i += 1

with open("output.txt", mode="a") as r:
    m = count - i
    for i in range(m):
        r.write("-1\n")
