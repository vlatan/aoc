#! /usr/bin/env python3


def main():
    data = get_data("2020/16/day16.txt")
    error, departure = solution(data)
    print(f"Part 1: {error}")
    print(f"Part 2: {departure}")


def get_data(filename):
    with open(filename, "r") as file:
        data = [item.split("\n") for item in file.read().strip().split("\n\n")]

        fields = [item.split(": ") for item in data[0]]
        fields = {item[0]: item[1].split(" or ") for item in fields}
        for key, value in fields.items():
            field = []
            for i in range(2):
                tmp = value[i].split("-")
                field += list(range(int(tmp[0]), int(tmp[1]) + 1))
            fields[key] = tuple(field)

        ticket = tuple(map(int, data[1][1].split(",")))
        tickets = set(tuple(map(int, item.split(","))) for item in data[2][1:])

    return fields, ticket, tickets


def solution(data):
    fields, my_ticket, tickets = data
    error, invalid_tickets = 0, set()
    for ticket in tickets:
        for value in ticket:
            if value not in set([i for k in fields.values() for i in k]):
                error += value
                invalid_tickets.add(ticket)
    valid_tickets = list(tickets - invalid_tickets)

    corr_fields = {}
    for j in range(len(my_ticket)):
        all_common_fields = []
        for i in range(len(valid_tickets)):
            current_common_fields = set()
            for key, value in fields.items():
                if valid_tickets[i][j] in value:
                    current_common_fields.add(key)
            all_common_fields.append(current_common_fields)
        corr_fields[j] = list(set.intersection(*all_common_fields))

    fields_to_remove = set()
    while sum([len(v) for v in corr_fields.values()]) != len(my_ticket):
        for key, value in corr_fields.items():
            if len(value) == 1:
                fields_to_remove.add(value[0])
            else:
                for field in fields_to_remove:
                    try:
                        corr_fields[key].remove(field)
                    except ValueError:
                        pass

    departure_values = 1
    for key, value in corr_fields.items():
        if "departure" in value[0]:
            departure_values *= my_ticket[key]

    return error, departure_values


if __name__ == "__main__":
    main()
