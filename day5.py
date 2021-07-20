#! /usr/bin/env python3

table = {'F': '0', 'B': '1', 'L': '0', 'R': '1'}
with open('day5.txt', 'r') as f:
    seats = {int(''.join(table[c] for c in line.strip()), 2) for line in f}
    my_seat = list(set(range(min(seats), max(seats) + 1)) - seats).pop()
    print(f'Part 1: {max(seats)}')  # part 1
    print(f'Part 2: {my_seat}')     # part 2
