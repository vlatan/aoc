#! /usr/bin/env python3


def main():
    part1, part2 = solve("2020/18/day18.txt")
    print(f"Part 1: {part1}")
    print(f"Part 2: {part2}")


def solve(filename):
    with open(filename, "r") as file:
        part1 = part2 = 0
        for ln in file:
            ln = ln.strip().replace(" ", "")
            result1, result2 = solve_line(ln)
            part1 += result1
            part2 += result2
        return part1, part2


def solve_line(line):
    if line[0].isdigit():
        result1, result2, i = int(line[0]), int(line[0]), 1
    else:
        start, stop, i = subline(0, line)
        result1, result2 = solve_line(line[start:stop])

    products = []
    while i < len(line) - 1:
        if line[i] == "*" and line[i + 1].isdigit():
            result1 *= int(line[i + 1])
            products.append(result2)
            result2 = int(line[i + 1])
            i += 2
        elif line[i] == "+" and line[i + 1].isdigit():
            result1 += int(line[i + 1])
            result2 += int(line[i + 1])
            i += 2
        else:
            start, stop, next_i = subline(i + 1, line)
            res1, res2 = solve_line(line[start:stop])
            if line[i] == "*":
                result1 *= res1
                products.append(result2)
                result2 = res2
            else:
                result1 += res1
                result2 += res2
            i = next_i
    if products:
        prod = 1
        for num in products:
            prod *= num
        result2 *= prod
    return result1, result2


def subline(indx, line):
    opening, closing, start, stop = 1, 0, indx + 1, 0
    for i in range(start, len(line)):
        if line[i] == "(":
            opening += 1
        elif line[i] == ")":
            closing += 1
            stop = i
            if opening == closing:
                return start, stop, i + 1


if __name__ == "__main__":
    main()
