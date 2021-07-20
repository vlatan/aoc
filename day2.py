#! /usr/bin/env python3

import re

with open('day2.txt', 'r') as file:
    data = [re.split('-| |: ', line.strip()) for line in file]

# part 1
total = 0
for item in data:
    if int(item[0]) <= item[3].count(item[2]) <= int(item[1]):
        total += 1
print(f'Part 1: {total}')

# part 2
valid = 0
for item in data:
    if item[3][int(item[0]) - 1] == item[2] and item[3][int(item[1]) - 1] != item[2]:
        valid += 1
    elif item[3][int(item[0]) - 1] != item[2] and item[3][int(item[1]) - 1] == item[2]:
        valid += 1
print(f'Part 2: {valid}')
