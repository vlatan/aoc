#! /usr/bin/env python3


def main():
    data = get_data("2020/10/day10.txt")
    print(f"Part 1: {count_differences(data)}")


def get_data(filename):
    with open(filename, "r") as file:
        data = sorted([int(line.strip()) for line in file])
        return [0] + data + [data[-1] + 3]


def count_differences(data):
    ones, threes = 0, 0
    for i in range(len(data) - 1):
        if data[i + 1] - data[i] == 1:
            ones += 1
        elif data[i + 1] - data[i] == 3:
            threes += 1
    return ones * threes


if __name__ == "__main__":
    main()
