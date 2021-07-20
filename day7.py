#! /usr/bin/env python3

import re


def get_data(filename):
    with open(filename, 'r') as file:
        data = [re.sub(r' bag(s)|\.|', '', line) for line in file]
        data = [item.strip().split(' contain ') for item in data]
        data = [[item[0], item[1].split(', ')] for item in data]
    return data


# part 1
def color_count(color):
    """ Returns a set of unique bag colors that
        ultimately contain the given bag color. """
    result = set()
    for item in data:
        for col in item[1]:
            if color in col:
                result.add(item[0])
                result |= color_count(item[0])
    return result


# part 2
def nested_count(color):
    result = 0
    for item in data:
        if item[0] == color:
            for i in range(len(item[1])):
                if item[1][i][:2] != 'no':
                    bag_num = int(item[1][i][0])
                    bag_color = item[1][i][2:]
                    result += bag_num + bag_num * nested_count(bag_color)
    return result


if __name__ == '__main__':
    data = get_data('day7.txt')
    print(f"Part 1: {len(color_count('shiny gold'))}")
    print(f"Part 2: {nested_count('shiny gold')}")
