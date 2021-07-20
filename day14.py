#! /usr/bin/env python3

from itertools import product


def main():
    # data = get_data('day14.txt')
    print(f"Part 1: {bitmask('day14.txt')}")
    print(f"Part 2: {decoder('day14.txt')}")


def bitmask(filename):
    mem = {}
    with open(filename, 'r') as file:
        for line in file:
            if line[0:4] == 'mask':
                m = line.strip()[7:]
            else:
                data = line.strip().replace('mem[', '').split('] = ')
                bn = [s for s in f'{int(data[1]):036b}']
                bn = [m[i] if m[i].isdigit() else bn[i] for i in range(len(m))]
                mem[data[0]] = int(''.join(bn), 2)
    return sum([value for value in mem.values()])


def decoder(filename):
    mem = {}
    with open(filename, 'r') as file:
        for line in file:
            if line[0:4] == 'mask':
                mask = line.strip()[7:]
                x = [i for i in range(len(mask)) if mask[i] == 'X']
            else:
                data = line.strip().replace('mem[', '').split('] = ')
                m, v = [s for s in f'{int(data[0]):036b}'], int(data[1])
                m = ['1' if mask[i] == '1' else m[i] for i in range(len(m))]
                prod = product(['0', '1'], repeat=len(x))
                for p in prod:
                    tmp = m
                    for i in range(len(p)):
                        tmp[x[i]] = p[i]
                    mem[''.join(tmp)] = v
    return sum([value for value in mem.values()])


if __name__ == '__main__':
    main()
