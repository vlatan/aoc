#! /usr/bin/env python3

def get_data(filename):
    with open(filename, 'r') as file:
        return [line.strip() for line in file]


def terminate_loop(i=0):
    if not data[i]:
        return 0
    acc, cmd, data[i] = int(data[i][4:]), data[i][:3], None
    if cmd == 'nop':
        return terminate_loop(i + 1)
    elif cmd == 'acc':
        return acc + terminate_loop(i + 1)
    elif cmd == 'jmp':
        return terminate_loop(i + acc)


def fix_loop(i=0):
    to_change = [i for i in range(len(data)) if data[i][:3] in ['nop', 'jmp']]
    i, total, visited = 0, 0, []
    while i < len(data):
        cmd, value = data[i][:3], int(data[i][4:])
        if i in visited:
            i, total, visited = 0, 0, []
            to_change.pop(0)
        else:
            visited.append(i)
            if cmd == 'acc':
                total += value
                i += 1
            elif i == to_change[0]:
                if cmd == 'nop':
                    i += value
                else:
                    i += 1
            elif cmd == 'nop':
                i += 1
            else:
                i += value
    return total


if __name__ == '__main__':
    data = get_data('day8.txt')
    print(f"Part 1: {terminate_loop()}")
    data = get_data('day8.txt')
    print(f"Part 2: {fix_loop()}")
