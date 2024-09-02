#! /usr/bin/env python3

import time


def main():
    print(f"Part 1: {earliest_bus('2020/13/day13.txt')}")


def earliest_bus(filename):
    with open(filename, "r") as file:
        time = int(file.readline())
        ids = [int(i) for i in file.readline().split(",") if i.isdigit()]
    mins = [num - time % num if time % num != 0 else 0 for num in ids]
    smallest_diff = min(mins)
    bus_id = ids[mins.index(smallest_diff)]
    return smallest_diff * bus_id


if __name__ == "__main__":
    main()
