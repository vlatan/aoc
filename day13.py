#! /usr/bin/env python3

import time


def main():
    print(f"Part 1: {earliest_bus('day13.txt')}")
    print('Getting the Part 2...')
    print(f"Part 2: {earliest_timestamp('day13.txt')}")


def earliest_bus(filename):
    with open(filename, 'r') as file:
        time = int(file.readline())
        ids = [int(i) for i in file.readline().split(',') if i.isdigit()]
    mins = [num - time % num if time % num != 0 else 0 for num in ids]
    smallest_diff = min(mins)
    bus_id = ids[mins.index(smallest_diff)]
    return smallest_diff * bus_id


def earliest_timestamp(filename):
    with open(filename, 'r') as file:
        output = file.readlines()[1].strip().split(',')
    buses = [[int(i)] for i in output if i.isdigit()]

    count, k = 1, 0
    for i in range(1, len(output)):
        if output[i].isalpha():
            count += 1
        else:
            buses[k].append(count)
            k += 1
            count = 1
    buses[-1].append(0)

    i, incr = 0, buses[0][0]
    while True:
        local, count = i, 0
        for j in range(len(buses)):
            if local % buses[j][0] != 0:
                break
            count += 1
            local += buses[j][1]
        if count == len(buses):
            return i
        i += incr


if __name__ == '__main__':
    main()
