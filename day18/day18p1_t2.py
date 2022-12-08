#!/usr/bin/env python3
import math
import copy
import tqdm

class Node:
    def __init__(self, l, r, p):
        self.l = l
        self.r = r
        self.p = p

    def __getitem__(self, key):
        if type(key) != int:
            raise TypeError
        elif not (0 <= key <= 1):
            raise IndexError

        if key == 0:
            return self.l
        elif key == 1:
            return self.r

    def __int__(self):
        if self.is_leaf():
            return self.l
        else:
            return 3*int(self.l) + 2*int(self.r)    

    def __setitem__(self, key, value):
        # key check
        self[key]
        # check value type
        if type(value) != Node:
            raise TypeError

        if key == 0:
            self.l = value
        elif key == 1:
            self.r = value

    def __add__(self, other):
        lc = copy.deepcopy(self)
        rc = copy.deepcopy(other)

        n = Node(lc, rc, None)
        lc.p, rc.p = n, n
        n.reduce()
        return n


    def add_to(self, d, v):
        c = self

        while c.p[d] == c:
            c = c.p
            if c.p is None:
                return False
        c = c.p[d]
        od = 1 if d == 0 else 0

        while not c.is_leaf():
            c = c[od]
        c.l += v.l
        return True


    def is_leaf(self):
        return self.r is None

    def reduce(self):
        if self.explode() or self.split():
            self.reduce()
            return True
        else:
            return False

    def split(self):
        if self.is_leaf():
            if self.l >= 10:
                lv, rv = int(math.floor(self.l/2)), int(math.ceil(self.l/2))

                self.l = Node(lv, None, self)
                self.r = Node(rv, None, self)
                return True
            else:
                return False
        
        for i in [0, 1]:
            if self[i].split():
                return True
            
        return False

    def explode(self, depth=0):
        if self.is_leaf():
            return False
        elif depth >= 4:
            if self.l.is_leaf() and self.r.is_leaf():
                l, r = self.l, self.r

                self.l = 0
                self.r = None

                self.add_to(0, l)
                self.add_to(1, r)

                return True

        if self.l.explode(depth+1):
            return True
        if self.r.explode(depth+1):
            return True
        return False


    def __repr__(self):
        return self.__str__()

    def __str__(self):
        if self.r is None:
            return str(self.l)
        else:
            return f"[{self.l}, {self.r}]"


def list_to_tree(lst, parent=None):

    if type(lst) == int:
        return Node(lst, None, parent)
    else:
        n = Node(None, None, parent)
        l = list_to_tree(lst[0], parent=n)
        r = list_to_tree(lst[1], parent=n)
        n.l, n.r = l, r
        return n


def read_file_list(path):
    input_arrs = list()
    with open(path) as f:
        for line in f:
            input_arrs.append(eval(line))
    # convert to nodes
    for i in range(len(input_arrs)):
        input_arrs[i] = list_to_tree(input_arrs[i])
    return input_arrs


def read_file(path):
    input_arrs = read_file_list(path)

    n = input_arrs[0]
    for i in range(1, len(input_arrs)):
        b = input_arrs[i]
        n = n + b
    return n

if __name__ == "__main__":
    # ab = [[[[[9,8],1],2],3],4]
    # ab = [7,[6,[5,[4,[3,2]]]]]
    # ab = [[6,[5,[4,[3,2]]]],1]
    # ab = [[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]] 

    # a = [[[[4,3],4],4],[7,[[8,4],9]]]
    # b = [1,1]

    # a = list_to_tree(a)
    # b = list_to_tree(b)
    # ab = a+b



    # print(f"ab:", ab)
    # tree = ab
    # print(tree)
    # print(f"tree:", tree)

    # print("reduced", tree.reduce())
    # print(f"tree:", tree)

    # print(read_file("./aoc/adventofcode/day18/day18_input_test_small_1.txt"))
    # print(read_file("./aoc/adventofcode/day18/day18_input_test_medium_1.txt"))

    # tree = list_to_tree([[1,2],[[3,4],5]])

    # tree = list_to_tree([[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]])
    # print(tree)
    # print(int(tree))

    # n = read_file("./aoc/adventofcode/day18/day18_input_test.txt")
    # print(n)
    # print(int(n))


    # n = read_file("./aoc/adventofcode/day18/day18_input.txt")
    # print(n)
    # print(int(n))


    arr = read_file_list("./aoc/adventofcode/day18/day18_input_test.txt")
    arr = read_file_list("./aoc/adventofcode/day18/day18_input.txt")

    msum = -1
    for i in tqdm.tqdm(range(len(arr))):
        for j in range(len(arr)):
            a = arr[i]
            b = arr[j]

            s = int(a+b)
            msum = max(msum, s)
    print(msum)

