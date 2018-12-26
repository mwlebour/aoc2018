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
        carts[(x,y)] = (char, x, y, 0)

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

def new_pos(c, m):
    do_map = {
        '/<': lambda i: 'v',
        '/>': lambda i: '^',
        '/^': lambda i: '>',
        '/v': lambda i: '<',
        '\\<': lambda i: '^',
        '\\>': lambda i: 'v',
        '\\^': lambda i: '<',
        '\\v': lambda i: '>',
        '+<' : lambda i: {0: 'v', 1: '<', 2: '^'}[i],
        '+>' : lambda i: {0: '^', 1: '>', 2: 'v'}[i],
        '+^' : lambda i: {0: '<', 1: '^', 2: '>'}[i],
        '+v' : lambda i: {0: '>', 1: 'v', 2: '<'}[i],
    }
    if c[0] == '<': nx, ny = c[1],c[2]-1
    if c[0] == '>': nx, ny = c[1],c[2]+1
    if c[0] == '^': nx, ny = c[1]-1,c[2]
    if c[0] == 'v': nx, ny = c[1]+1,c[2]
    l = '%s%s'%(m[nx][ny],c[0])
    n = do_map[l](c[3]) if l in do_map else c[0]
    return (n, nx, ny, (c[3]+1)%3 if m[nx][ny] == '+' else c[3])


def move(carts,m):
    new_carts = defaultdict(list)
    to_delete = {}
    for pos, c in sorted(carts.items()):
        if pos in to_delete: continue
        del carts[pos]
        new_c = new_pos(c,m)
        nx, ny = new_c[1], new_c[2]
        new_carts[(nx,ny)].append(new_c)
        if len(new_carts[(nx,ny)]) == 2:
            del new_carts[(nx,ny)]
        if (nx,ny) in carts:
            del new_carts[(nx,ny)]
            to_delete[(nx,ny)] = 1
        if len(new_carts) + len(carts) - len(to_delete) == 1:
            print(carts, new_carts)

    return {k:v[0] for k,v in new_carts.items()}

def print_map(m, c):
    os.system('clear')
    for i,y in enumerate(m):
        for j,v in enumerate(y):
            if (i,j) in c:
                sys.stdout.write(c[(i,j)][0])
            else:
                sys.stdout.write(v)
        print()
    sys.stdout.flush()


def main1(the_map ,carts):
    i = 0
    while True :
        # print_map(the_map, carts)
        # os.system('sleep .1')
        carts = move(carts, the_map)
        if len(carts) == 1:
            break
        i += 1
        if i % 100 == 0:
            print(i, len(carts))
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
