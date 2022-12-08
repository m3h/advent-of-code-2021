#!/usr/bin/env python3
import re
import numpy as np
import itertools

def read_beacons(path):

    beacons = {}
    scanner = None
    with open(path) as f:
        for line in f:
            if line.strip() == "":
                continue

            m = re.match(r'--- scanner (.*) ---', line)
            if m is not None:
                scanner = m.group(1)
                beacons[scanner] = []
            else:
                beacon = [int(x) for x in line.split(',')]
                beacons[scanner] += [beacon]
    
    beacons = {k: np.array(v) for k,v in beacons.items()}
    return beacons
    
            

def combinations(a, r):
    ret = list()

    for i in range(len(a)):
        ai = a[i]
        if r == 1:
            ret += [[ai]]
        else:
            for cm in combinations(a, r-1):
                ret += [[ai, *cm]]
    return ret

def matrix_permutations(a: np.ndarray):
    r = list()

    base = list(range(a.shape[-1]))
    for shuffle in itertools.permutations(base):
        for sign in combinations([1, -1], r=len(base)):
            b = a[..., shuffle]

            for i in range(len(base)):
                b[..., i] *= sign[i]
            r += [b]
    return r

def merge_2(a, b, threshold):
    aps = matrix_permutations(a)
    bps = matrix_permutations(b)

    for ap in aps:
        for bp in bps:
            d = ap - bp
            d = list(d.flatten())

            for v in set(d):
                print(v)




    print(a)
    print(b)
def merge_all(beacons, threshold):
    for si in beacons:
        for sj in beacons:
            if sj == si:
                continue
            
            merged = merge_2(beacons[si], beacons[sj], threshold)
                
            

def main():
    fpath = "./aoc/adventofcode/day19/day19_input_test_small_1.txt"
    threshold = 3
    beacons = read_beacons(fpath)

    merge_all(beacons, threshold)

if __name__ == "__main__":
    main()
