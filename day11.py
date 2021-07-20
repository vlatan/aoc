#! /usr/bin/env python3

from itertools import product


def main():
    print(f"Part 1: {occupied_seats_v1()}")
    print(f"Part 2: {occupied_seats_v2()}")


def get_data(filename):
    with open(filename, 'r') as file:
        return [[i for i in line.strip()] for line in file]


def occupied_seats_v1():
    data, changes = get_data('day11.txt'), {}
    coordinates = list(product([-1, 0, 1], repeat=2))
    coordinates.remove((0, 0))
    while True:
        for i in range(len(data)):
            for j in range(len(data[i])):
                if data[i][j] == '.':
                    continue
                neighbors = []
                for c in coordinates:
                    if i + c[0] == -1 or j + c[1] == -1:
                        continue
                    try:
                        neighbors.append(data[i + c[0]][j + c[1]])
                    except IndexError:
                        continue
                if data[i][j] == 'L' and neighbors.count('#') == 0:
                    changes[(i, j)] = '#'
                elif data[i][j] == '#' and neighbors.count('#') >= 4:
                    changes[(i, j)] = 'L'
        if not changes:
            return sum([item.count('#') for item in data])
        for key, value in changes.items():
            data[key[0]][key[1]] = value
        changes = {}


def occupied_seats_v2():
    data, changes = get_data('day11.txt'), {}
    while True:
        for i in range(len(data)):
            for j in range(len(data[i])):
                if data[i][j] == '.':
                    continue
                elif data[i][j] == 'L' and count_visible(data, (i, j)) == 0:
                    changes[(i, j)] = '#'
                elif data[i][j] == '#' and count_visible(data, (i, j)) >= 5:
                    changes[(i, j)] = 'L'
        if not changes:
            return sum([item.count('#') for item in data])
        for key, value in changes.items():
            data[key[0]][key[1]] = value
        changes = {}


def count_visible(data, point):
    count, k, i, j = 0, 1, point[0], point[1]
    status = {'top_left': False, 'top': False, 'top_right': False,
              'left': False, 'right': False, 'bottom_left': False,
              'bottom': False, 'bottom_right': False}
    while k < len(data):
        if not status['top_left'] and i - k >= 0 and j - k >= 0:
            if data[i - k][j - k] == '#':
                count += 1
            if data[i - k][j - k] != '.':
                status['top_left'] = True

        if not status['top'] and i - k >= 0:
            if data[i - k][j] == '#':
                count += 1
            if data[i - k][j] != '.':
                status['top'] = True

        if not status['top_right'] and i - k >= 0 and j + k < len(data[0]):
            if data[i - k][j + k] == '#':
                count += 1
            if data[i - k][j + k] != '.':
                status['top_right'] = True

        if not status['left'] and j - k >= 0:
            if data[i][j - k] == '#':
                count += 1
            if data[i][j - k] != '.':
                status['left'] = True

        if not status['right'] and j + k < len(data[0]):
            if data[i][j + k] == '#':
                count += 1
            if data[i][j + k] != '.':
                status['right'] = True

        if not status['bottom_left'] and i + k < len(data) and j - k >= 0:
            if data[i + k][j - k] == '#':
                count += 1
            if data[i + k][j - k] != '.':
                status['bottom_left'] = True

        if not status['bottom'] and i + k < len(data):
            if data[i + k][j] == '#':
                count += 1
            if data[i + k][j] != '.':
                status['bottom'] = True

        if not status['bottom_right'] and i + k < len(data) and j + k < len(data[0]):
            if data[i + k][j + k] == '#':
                count += 1
            if data[i + k][j + k] != '.':
                status['bottom_right'] = True
        k += 1
    return count


if __name__ == '__main__':
    main()
