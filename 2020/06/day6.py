#! /usr/bin/env python3

with open("2020/06/day6.txt", "r") as file:
    data = ["" if line == "\n" else line.strip() for line in file] + [""]

# part 1
group, total = set(), 0
for line in data:
    if line:
        group |= {char for char in line}
    elif group:
        total += len(group)
        group = set()
print(f"Part 1: {total}")

# part 2
group, total = [], 0
for i in range(len(data)):
    if not data[i - 1] and not data[i + 1]:
        total += len(data[i])
    elif data[i]:
        group.append({char for char in data[i]})
    elif group:
        total += len(set.intersection(*group))
        group = []
print(f"Part 2: {total}")
