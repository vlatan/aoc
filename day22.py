#! /usr/bin/env python3
# https://adventofcode.com/2020/day/22

def main():
    file = 'day22.txt'
    print(f'Part 1: {solve_1(file)}')
    print(f'Part 2: {solve_2(file)}')


def get_data(filename):
    with open(filename, 'r') as f:
        data = f.read().strip().split('\n\n')
        return [list(map(int, i.split('\n')[1:])) for i in data]


def solve_1(filename):
    d = get_data(filename)
    while d[0] and d[1]:
        if d[0][0] > d[1][0]:
            d[0] += [d[0].pop(0), d[1].pop(0)]
        else:
            d[1] += [d[1].pop(0), d[0].pop(0)]
    winner = max(d[0], d[1], key=lambda x: len(x))
    return sum(x[0] * x[1] for x in list(enumerate(reversed(winner), 1)))


def solve_2(filename):
    def play(d):
        seen0, seen1 = set(), set()
        while True:
            p0, p1 = tuple(d[0]), tuple(d[1])
            if p0 in seen0 or p1 in seen1:
                return 0
            seen0.add(p0)
            seen1.add(p1)

            if d[0][0] <= len(d[0][1:]) and d[1][0] <= len(d[1][1:]):
                w = play([d[0][1:d[0][0] + 1], d[1][1:d[1][0] + 1]])
                d[w] += [d[w].pop(0), d[abs(w - 1)].pop(0)]
            elif d[0][0] > d[1][0]:
                d[0] += [d[0].pop(0), d[1].pop(0)]
            else:
                d[1] += [d[1].pop(0), d[0].pop(0)]

            if not d[1]:
                return 0
            elif not d[0]:
                return 1
    d = get_data(filename)
    winner = d[play(d)]
    return sum(x[0] * x[1] for x in list(enumerate(reversed(winner), 1)))


if __name__ == '__main__':
    main()
