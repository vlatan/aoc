#! /usr/bin/env python3


def main():
    print(f"Part 1: {number('2020/15/day15.txt', 2020)}")
    print(f"Part 2: {number('2020/15/day15.txt', 30000000)}")


def number(filename, n):
    with open(filename, "r") as file:
        data = list(map(int, file.read().split(",")))
    turns = {data[i]: i + 1 for i in range(len(data))}

    num = data[-1]
    for i in range(len(data), n):
        try:
            last_seen, turns[num] = turns[num], i
            num = i - last_seen
        except KeyError:
            turns[num] = i
            num = 0
    return num


if __name__ == "__main__":
    main()
