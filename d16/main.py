#!/usr/bin/env python3

from collections import defaultdict

opcodes = {
    'addr': lambda i, o: o[i[1]] + o[i[2]]       , #stores into register C the result of adding register A and register B.
    'addi': lambda i, o: o[i[1]] + i[2]          , #stores into register C the result of adding register A and value B.
    'mulr': lambda i, o: o[i[1]] * o[i[2]]       , #stores into register C the result of multiplying register A and register B.
    'muli': lambda i, o: o[i[1]] * i[2]          , #stores into register C the result of multiplying register A and value B.
    'banr': lambda i, o: o[i[1]] & o[i[2]]       , #stores into register C the result of the bitwise AND of register A and register B.
    'bani': lambda i, o: o[i[1]] & i[2]          , #stores into register C the result of the bitwise AND of register A and value B.
    'borr': lambda i, o: o[i[1]] | o[i[2]]       , #stores into register C the result of the bitwise OR of register A and register B.
    'bori': lambda i, o: o[i[1]] | i[2]          , #stores into register C the result of the bitwise OR of register A and value B.
    'setr': lambda i, o: o[i[1]]                 , #copies the contents of register A into register C. (Input B is ignored.)
    'seti': lambda i, o: i[1]                    , #stores value A into register C. (Input B is ignored.)
    'gtir': lambda i, o: int(i[1] > o[i[2]])     , #sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
    'gtri': lambda i, o: int(o[i[1]] > i[2])     , #sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
    'gtrr': lambda i, o: int(o[i[1]] > o[i[2]])  , #sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
    'eqir': lambda i, o: int(i[1] == o[i[2]])    , #sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
    'eqri': lambda i, o: int(o[i[1]] == i[2])    , #sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
    'eqrr': lambda i, o: int(o[i[1]] == o[i[2]]) , #sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
}

lines = []
with open('input.out') as f:
    lines = f.readlines()

def to_list(s):
    return [int(i) for i in s.split('[')[1][:-1].split(',')]


finder = {x:defaultdict(int) for x in opcodes.keys()}
while len(lines) > 0:
    if not lines[0].strip():
        lines = lines[1:]
        continue
    if not lines[0].startswith('Before'): break
    test = lines[0:3]
    lines = lines[3:]
    begin, instr, end = [l.strip() for l in test]
    begin = to_list(begin)
    end = to_list(end)
    instr = [int(i) for i in instr.split()]
    n = 0
    for k,op in opcodes.items():
        t = begin[:]
        t[instr[-1]] = op(instr,begin)
        if end == t:
            finder[k][instr[0]] += 1

def pop_smallest(f):
    small = 0
    small_key = None
    for k, v in f.items():
        if small_key is None or len(v) < small:
            small = len(v)
            small_key = k

    assert len(f[small_key]) == 1, "must be one"
    determined = (small_key,list(f.pop(small_key).keys()).pop())
    return f, determined

def remove_all(f,d):
    for k,v in f.items():
        v.pop(d,None)

    return f

new_opcodes = {}
while len(finder) > 0:
    finder, determined = pop_smallest(finder)
    finder = remove_all(finder,determined[1])
    new_opcodes[determined[1]] = opcodes[determined[0]]

assert len(new_opcodes) == len(opcodes), "better be same length"

reg = [0,0,0,0]
for line in lines:
    instr = [int(i) for i in line.strip().split()]
    reg[instr[-1]] = new_opcodes[instr[0]](instr,reg)
    print(reg)
