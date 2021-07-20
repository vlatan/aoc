#! /usr/bin/env python3

import re

with open('day4.txt', 'r') as file:
    content = file.read().replace('\n', ' ').split('  ')
    content = [re.split(':| ', item) for item in content]
    for i in range(len(content)):
        keys = [content[i][j] for j in range(0, len(content[i]), 2)]
        values = [content[i][k] for k in range(1, len(content[i]), 2)]
        content[i] = dict(zip(keys, values))

# part 1
fields = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}
present = [p for p in content if not fields - {k for k in p}]
print(f'Part 1: {len(present)}')

# part 2
valid = 0
for psp in present:
    if not 1920 <= int(psp['byr']) <= 2002:
        continue
    if not 2010 <= int(psp['iyr']) <= 2020:
        continue
    if not 2020 <= int(psp['eyr']) <= 2030:
        continue
    p1 = re.compile(r'^1([5-8]\d|9[0-3])(cm)$')
    p2 = re.compile(r'^59|6\d|7[0-6](in)$')
    if not (p1.match(psp['hgt']) or p2.match(psp['hgt'])):
        continue
    p = re.compile(r'^#([a-f0-9]{6})$')
    if not re.compile(r'^#([a-f0-9]{6})$').match(psp['hcl']):
        continue
    p = re.compile(r'^(amb|blu|brn|gry|grn|hzl|oth)$')
    if not re.compile(r'^(amb|blu|brn|gry|grn|hzl|oth)$').match(psp['ecl']):
        continue
    p = re.compile(r'^\d{9}$')
    if not re.compile(r'^\d{9}$').match(psp['pid']):
        continue
    valid += 1
print(f'Part 2: {valid}')
