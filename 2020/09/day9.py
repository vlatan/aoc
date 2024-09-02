#! /usr/bin/env python3


def main():
    data = get_data("2020/09/day9.txt")
    print(f"Part 1: {invalid_number(data)}")
    print(f"Part 2: {weakness(data)}")


def get_data(filename):
    with open(filename, "r") as file:
        return [int(line.strip()) for line in file]


def invalid_number(data):
    preamble, new_data = data[:25], data[25:]
    while new_data:
        if correct_sum(preamble, new_data[0]):
            preamble.pop(0)
            preamble.append(new_data[0])
            new_data.pop(0)
        else:
            return new_data[0]


def weakness(data):
    num = invalid_number(data)
    for i in range(len(data)):
        current_list = [data[i]]
        for j in range(i + 1, len(data)):
            current_list.append(data[j])
            if sum(current_list) == num:
                return min(current_list) + max(current_list)
            elif sum(current_list) > num:
                break


def correct_sum(dataset, num):
    for i in range(len(dataset)):
        for j in range(i + 1, len(dataset)):
            if dataset[i] + dataset[j] == num:
                return True
    return False


if __name__ == "__main__":
    main()
