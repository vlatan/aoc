#! /usr/bin/env python3
# https://adventofcode.com/2020/day/23

def main():
    file = 'test.txt'
    data = get_data(file)
    print(data)
    # print(data)
    print(solve1(data, 10))
    # print(f'Part 1: {solve_1(file)}')
    # print(f'Part 2: {solve_2(file)}')

    # lst = CircularList()
    # for i in data:
    #     lst.add(i)
    # print(lst())


def get_data(filename):
    with open(filename, 'r') as f:
        return list(map(int, f.read().strip()))


class Node:
    def __init__(self, value=None):
        self.value = value
        self.next = None


class CircularList:
    def __init__(self):
        self.head = self.tail = Node()
        self.head.next = self.tail
        self.tail.next = self.head

    def add(self, value):
        """ Add node at the end of the list """
        new_node = Node(value)
        # if list is empty
        if self.head.value is None:
            self.head = new_node
            self.tail = new_node
            new_node.next = self.head
        else:
            self.tail.next = new_node
            self.tail = new_node
            new_node.next = self.head

    def __call__(self):
        """ Return the list """
        current, lst = self.head, []
        if current.value is None:
            return lst
        lst.append(current.value)
        while current.next != self.head:
            current = current.next
            lst.append(current.value)
        return lst


def solve1(data, n):
    for i in range(n):
        # print(f'cups = {data}')
        new, x = [0] * len(data), i % 9
        # print(f'current_index = {x}')
        # current_cup = data[x]
        new[x] = data[x]
        # print(f'current_cup = {current_cup}')
        # three_cups = [data[wrap_right(data, j)] for j in range(x + 1, x + 4)]
        three_cups = []
        for _ in range(3):
            try:
                three_cups.append(data.pop(x + 1))
            except IndexError:
                three_cups.append(data.pop(x + 1 - len(data)))

        # print(f'nex_three_cups = {three_cups}')
        # print('*' * 40)
        next_cup = wrap_left(new[x] - 1)
        while next_cup in three_cups:
            next_cup = wrap_left(new[x] - 1)

        for j in range(x, data:
            if


        new[x], indx=current_cup, x + 1
        for k in range(len(three_cups)):
            try:
                new[indx]=three_cups[k]
            except IndexError:
                new[indx - len(data)]=three_cups[k]
            finally:
                indx += 1

        print(new)

    #     while data.index(next_cup) < x:
    #         data.append(data.pop(0))
    #     j = data.index(next_cup)
    #     for k in range(len(data)):
    #         if data[k] in three_cups:
    #             pass
    #     data = data[:j + 1] + three_cups + data[j + 1:]
    # return data


def wrap_left(num):
    if num == 0:
        return 9
    return num


def wrap_right(data, cup_index):
    if cup_index == len(data):
        return 0
    return cup_index


if __name__ == '__main__':
    main()
