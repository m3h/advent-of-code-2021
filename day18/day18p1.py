#!/usr/bin/env python3
import math


def addv(arr, i, v):
    if type(arr[i]) == int:
        arr[i] += v
    else:
        addv(arr[i], i, v)


def explode_helper(arr, depth=0):
    if type(arr) == int:
        return arr, False, None, None
    elif depth >= 4 and type(arr[0]) == int and type(arr[1]) == int:
        return 0, True, arr[0], arr[1]
    else:
        # left
        arr[0], exploded, l, r = explode_helper(arr[0], depth+1)
        if exploded:
            if r is not None:
                addv(arr, 1, r)
                r = None
            return arr, True, l, r
        # right
        arr[1], exploded, l, r = explode_helper(arr[1], depth+1)
        if exploded:
            if l is not None:
                addv(arr, 0, l)
                l = None
            return arr, True, l, r
        
        return arr, False, None, None


def explode(arr):
    arr, exploded, _, _ = explode_helper(arr)

    return arr, exploded

            
def split(arr):
    for i in range(2):
        if type(arr[i]) == int:
            if arr[i] >= 10:
                arr[i] = [int(math.floor(arr[i]/2)), int(math.ceil(arr[i]/2))]
                return arr, True
        else:
            arr[i], modified = split(arr[i])
            if modified:
                return arr, True 
    return arr, False


def reduce(arr):

    print("reduce", arr)
    for fn in [explode, split]:
        arr, reduced = fn(arr)

        if reduced:
            return reduce(arr)
    return arr


def add(l, r):

    r = [l, r]

    return reduce(r)

if __name__ == "__main__":
    # l, r = [1,2], [[3,4],5]
    # print(add(l, r))

    a = [[[[[9,8],1],2],3],4]
    a = [7,[6,[5,[4,[3,2]]]]]
    a = [[6,[5,[4,[3,2]]]],1]
    a = [[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]
    a = [10, 11]

    a = [[[[4,3],4],4],[7,[[8,4],9]]]
    b = [1,1]
    print(a, "+", b)
    print(add(a, b))