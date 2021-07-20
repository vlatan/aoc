#! /usr/bin/env python3


def main():
    data = get_data('test.txt')
    print(data)
    print(f"Part 1: {count_differences(data)}")
    print(f"Part 2: {non_simultaneous_changes(data)}")


def get_data(filename):
    with open(filename, 'r') as file:
        data = sorted([int(line.strip()) for line in file])
        return [0] + data + [data[-1] + 3]


def count_differences(data):
    ones, threes = 0, 0
    for i in range(len(data) - 1):
        if data[i + 1] - data[i] == 1:
            ones += 1
        elif data[i + 1] - data[i] == 3:
            threes += 1
    print(ones, threes)
    return ones * threes


# def count_combinations(dataset):
#     i, data, count = 1, dataset.copy(), {tuple(dataset)}
#     while i < len(data) - 1:
#         count |= continuous_change(data)
#         count |= individual_change(data)
#         if data[i + 1] - data[i - 1] <= 3:
#             tmp = data.pop(i)
#             count |= continuous_change(data)
#             count |= individual_change(data)
#             data.insert(i, tmp)
#         i += 1
#     return len(count)


# def count_combinations(dataset):
#     i, data, count = 1, dataset.copy(), {tuple(dataset)}
#     while i < len(data) - 1:
#         if data[i + 1] - data[i - 1] <= 3:
#             j, tmp = 1, data.pop(i)
#             new_data = data.copy()
#             count.add(tuple(new_data))
#             while j < len(new_data) - 1:
#                 if new_data[j + 1] - new_data[j - 1] <= 3:
#                     new_tmp = new_data.pop(j)
#                     count.add(tuple(new_data))
#                     data.insert(j, new_tmp)
#                 j += 1
#             data.insert(i, tmp)
#         i += 1
#     return len(count)


def non_simultaneous_changes(dataset):
    count = 0
    indexes = distinct_changes(dataset)
    print(indexes)
    for i in range(len(indexes) - 1):
        if sum(indexes[i:i + 3]) / 3 == indexes[i + 1]:
            print(indexes[i:i + 3])
            count += 1
    if count:
        return int((len(indexes) - 1)**count / 2)
    else:
        return int((len(indexes) - 1) / 2)


def distinct_changes(dataset):
    changes, i, data, indexes = 0, 1, dataset.copy(), []
    while i < len(data) - 1:
        if data[i + 1] - data[i - 1] <= 3:
            tmp = data.pop(i)
            # print(data)
            changes += 1
            indexes.append(tmp)
            data.insert(i, tmp)
        i += 1
    # print(indexes)
    return indexes


def continuous_change(dataset, i=1):
    count, data = {tuple(dataset)}, dataset.copy()
    while i < len(data) - 1:
        if data[i + 1] - data[i - 1] <= 3:
            data.remove(data[i])
            count.add(tuple(data))
        else:
            i += 1
    return count


def individual_change(dataset, i=1):
    count, data = {tuple(dataset)}, dataset.copy()
    while i < len(data) - 1:
        if data[i + 1] - data[i - 1] <= 3:
            tmp = data.pop(i)
            count.add(tuple(data))
            data.insert(i, tmp)
        i += 1
    return count


# (0), 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 6, 7, 10, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 5, 7, 10, 12, 15, 16, 19, (22)
# (0), 1, 4, 6, 7, 10, 11, 12, 15, 16, 19, (22)
(0), 1, 4, 6, 7, 10, 12, 15, 16, 19, (22)
# (0), 1, 4, 7, 10, 11, 12, 15, 16, 19, (22)
# (0), 1, 4, 7, 10, 12, 15, 16, 19, (22)

# {(0, 1, 4, 7, 10, 11, 12, 15, 16, 19, 22),
# (0, 1, 4, 7, 10, 12, 15, 16, 19, 22),
# (0, 1, 4, 6, 7, 10, 11, 12, 15, 16, 19, 22),
# (0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22)}


if __name__ == '__main__':
    main()
