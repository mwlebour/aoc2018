#!/usr/bin/env python3

import os, sys
from functools import partial
from collections import defaultdict

def parse(lines):

    # the_map = [[' '] * len(lines) for i in range(len(lines[0])-1)]
    the_map = [[' '] * (len(lines[0])-1) for i in range(len(lines))]
    def set_map(char, x, y):
        the_map[x][y] = char

    carts = {}
    def make_cart(char, x, y):
        carts[(x,y)] = (char, x, y)

    do_map = {
        ' ': [],
        '\n': [],
        '|' : [partial(set_map,'|')],
        '^' : [partial(set_map,'|'), partial(make_cart,'^')],
        'v' : [partial(set_map,'|'), partial(make_cart,'v')],
        '-' : [partial(set_map,'-')],
        '>' : [partial(set_map,'-'), partial(make_cart,'>')],
        '<' : [partial(set_map,'-'), partial(make_cart,'<')],
        '/' : [partial(set_map,'/')],
        '\\': [partial(set_map,'\\')],
        '+' : [partial(set_map,'+')],
    }

    for i,line in enumerate(lines):
        for j, c in enumerate(line):
            for func in do_map[c]:
                func(i,j)

    return the_map, carts

def move(carts,m):
    do_map = {
        '/<': lambda i: 'v',
        '/>': lambda i: '^',
        '/^': lambda i: '>',
        '/v': lambda i: '<',
        '\\<': lambda i: '^',
        '\\>': lambda i: 'v',
        '\\^': lambda i: '<',
        '\\v': lambda i: '>',
                          l,      s,      r
        '+<' : lambda i: {0: '<', 1: '<', 2: '<'}[i],
        '+>' : lambda i: {0: '<', 1: '<', 2: '<'}[i],
        '+^' : lambda i: {0: '<', 1: '<', 2: '<'}[i],
        '+v' : lambda i: {0: '<', 1: '<', 2: '<'}[i],
    }
    for pos, c in carts.items():
        if c[0] == '<': carts[pos] = (c[0], c[1] - 1, c[2])
        if c[0] == '>': carts[pos] = (c[0], c[1] + 1, c[2])
        if c[0] == '^': carts[pos] = (c[0], c[1], c[2] - 1)
        if c[0] == 'v': carts[pos] = (c[0], c[1], c[2] + 1)





def print_map(m, c):
    for i,y in enumerate(m):
        for j,v in enumerate(y):
            if (i,j) in c:
                sys.stdout.write(c[(i,j)][0])
            else:
                sys.stdout.write(v)
        print()


def main1(the_map ,carts):
    print_map(the_map, carts)
    pass

def main2(the_map ,carts):
    pass


if __name__ == '__main__':
    fn = 'unit.out'
    if 'FULL' in os.environ:
        fn = 'input.out'
    with open(fn) as f:
        lines = f.readlines()
    the_map, carts = parse(lines)
    main1(the_map, carts)
    main2(the_map, carts)
