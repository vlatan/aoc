#! /usr/bin/env python3

import time


def main():
    rules, lines = get_data('test.txt')
    # print(rules, end='\n\n')
    # print(lines, end='\n\n')
    # res = build_lines(rules, 0)
    # print(res, end='\n\n')

    # print(build_lines(rules, lines, 0))
    start = time.time()
    print('The process of getting the solution for Part 1 has started...')
    print('This naive implementation will take around 3.2 minutes to fetch the result...')
    print(f"Part 1: {solve_1(rules, lines)}")
    print(f'It took {round((time.time() - start) / 60, 2)} minutes.')


def get_data(filename):
    with open(filename, 'r') as file:
        data = [item.split('\n') for item in file.read().strip().split('\n\n')]
        rules, lines = [s.split(': ') for s in data[0]], data[1]
        rules = {int(s[0]): s[1].replace('"', '').split() for s in rules}
        for key, value in rules.items():
            if '|' in value:
                split = value.index('|')
                rules[key] = (value[:split], value[split + 1:])
            elif value[0] in ['a', 'b']:
                rules[key] = value[0]
    return rules, set(lines)


def build_lines(rules, lines, i):
    results = [rules[0]]

    def recurse(rules, i):
        tmp = []
        for j in range(len(results)):
            tmp1, tmp2 = [], []
            for char in results[j]:
                if char.isdigit() and int(char) == i:
                    if isinstance(rules[i], tuple):
                        tmp1 += rules[i][0]
                        tmp2 += rules[i][1]
                    elif isinstance(rules[i], list):
                        tmp1 += rules[i]
                        tmp2 += rules[i]
                    else:
                        tmp1.append(rules[i])
                        tmp2.append(rules[i])
                else:
                    tmp1.append(char)
                    tmp2.append(char)
            results[j] = tmp1
            if len(tmp2) == len(tmp1) and tmp2 != tmp1:
                tmp.append(tmp2)
        results.extend(tmp)

        if isinstance(rules[i], tuple):
            for group in rules[i]:
                for char in group:
                    recurse(rules, int(char))
        elif isinstance(rules[i], list):
            for char in rules[i]:
                recurse(rules, int(char))

    recurse(rules, i)
    results = {''.join(res) for res in results}
    return len(lines & results)


def solve_1(rules, lines):
    results = [rules[0]]

    print(results)

    while True:
        tmp = []
        for item in results:
            tmp1, tmp2 = [], []
            for i in range(len(item)):
                if item[i].isdigit():
                    rule = rules[int(item[i])]
                    if isinstance(rule, tuple):
                        tmp1 += rule[0] + item[i + 1:]
                        tmp2 += rule[1] + item[i + 1:]
                        break
                    elif isinstance(rule, list):
                        tmp1 += rule
                        tmp2 += rule
                    else:
                        tmp1.append(rule)
                        tmp2.append(rule)
                else:
                    tmp1.append(item[i])
                    tmp2.append(item[i])
            tmp.append(tmp1)
            if tmp2 != tmp1:
                tmp.append(tmp2)

        results, done = tmp, True
        for item in results:
            for char in item:
                if char.isdigit():
                    done = False
                    break
            if not done:
                break

        print(results, end="\n\n\n")

        if done:
            results = {''.join(res) for res in results}
            return len(lines & results)


if __name__ == '__main__':
    main()
