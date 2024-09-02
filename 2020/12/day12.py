#! /usr/bin/env python3


def main():
    data = get_data("2020/12/day12.txt")
    print(f"Part 1: {manhattan_distance_v1(data)}")
    print(f"Part 2: {manhattan_distance_v2(data)}")


def get_data(filename):
    with open(filename, "r") as file:
        return [line.strip() for line in file]


def manhattan_distance_v1(dataset):
    facing, x, y = 0, 0, 0
    for item in dataset:
        if item[0] == "N":
            y += int(item[1:])
        elif item[0] == "S":
            y -= int(item[1:])
        elif item[0] == "E":
            x += int(item[1:])
        elif item[0] == "W":
            x -= int(item[1:])
        elif item[0] == "L":
            facing = (facing + int(item[1:])) % 360
        elif item[0] == "R":
            facing = (facing - int(item[1:])) % 360
        elif item[0] == "F":
            if facing == 0:
                x += int(item[1:])
            elif facing == 90:
                y += int(item[1:])
            elif facing == 180:
                x -= int(item[1:])
            elif facing == 270:
                y -= int(item[1:])
    return abs(x) + abs(y)


def manhattan_distance_v2(dataset):
    x, y, xw, yw = 0, 0, 10, 1
    for item in dataset:
        if item[0] == "N":
            yw += int(item[1:])
        elif item[0] == "S":
            yw -= int(item[1:])
        elif item[0] == "E":
            xw += int(item[1:])
        elif item[0] == "W":
            xw -= int(item[1:])
        elif item[0] == "F":
            x += int(item[1:]) * xw
            y += int(item[1:]) * yw
        elif item[0] == "L":
            if int(item[1:]) == 90:
                xw, yw = -yw, xw
            elif int(item[1:]) == 180:
                xw, yw = -xw, -yw
            elif int(item[1:]) == 270:
                xw, yw = yw, -xw
        elif item[0] == "R":
            if int(item[1:]) == 90:
                xw, yw = yw, -xw
            elif int(item[1:]) == 180:
                xw, yw = -xw, -yw
            elif int(item[1:]) == 270:
                xw, yw = -yw, xw
    return abs(x) + abs(y)


if __name__ == "__main__":
    main()
