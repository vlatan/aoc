#! /usr/bin/env python3

with open('day1.txt', 'r') as file:
    nums = [int(line.strip()) for line in file]

# part 1
abort, pairs = False, set()
for i in range(len(nums)):
    for j in range(i + 1, len(nums)):
        if nums[i] + nums[j] == 2020:
            print(f"Part 1: {nums[i] * nums[j]}")
        elif nums[i] + nums[j] < 2020:
            pairs.add((nums[i], nums[j]))

# part 2
abort = False
for num in nums:
    for pair in pairs:
        if num not in pairs and sum(pair) + num == 2020:
            print(f"Part 2: {num * pair[0] * pair[1]}")
            abort = True
            break
    if abort:
        break
