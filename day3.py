#! /usr/bin/env python3

with open('day3.txt', 'r') as file:
    lines = [line.strip() for line in file]
    lines = [line * len(lines) for line in lines]

 # part 1
index, count = 0, 0
for line in lines:
    if line[index] == '#':
        count += 1
    index += 3
print(f'Part 1: {count}')

# part 2
total, steps = 1, [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
for step in steps:
    count, i, index = 0, 0, 0
    while i < len(lines):
        if lines[i][index] == '#':
            count += 1
        index += step[0]
        i += step[1]
    total *= count
print(f'Part 2: {total}')
