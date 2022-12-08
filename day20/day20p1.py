#!/usr/bin/env python3
import copy

def binarize(arr):
    for i in range(len(arr)):
        if type(arr[i]) == list:
            arr[i] = binarize(arr[i])
        else:
            arr[i] = {
                '#': 1,
                '.': 0,
            }[arr[i]]
    return arr


def read_input(path):
    with open(path) as f:
        iea = list(f.readline().strip())

        assert f.readline().strip() == ""
        img = list()
        for line in f:
            if line.strip() == "":
                break
            img += [list(line.strip())]
    return binarize(iea), binarize(img)


def pad_img(img, pad, v):
    p = list()
    for _ in range(pad):
        p += [[v]*(len(img[0])+2*pad)]
    for ln in img:
        p += [([v]*pad) + ln + ([0]*pad)]
    for _ in range(pad):
        p += [[v]*(len(img[0])+2*pad)]
    return p

def at(arr, i, j, padv):
    if i < 0 or j < 0 or i >= len(arr) or j >= len(arr[i]):
        return padv
    else:
        return arr[i][j]


def get_block(arr, i, j, s, v):
    ret = "0b"
    for x in range(i - s//2, i + s//2+1):
        for y in range(j - s//2, j + s//2+1):
            ret += str(at(arr, x, y, v))
    return int(ret, 2)

def enhance(img_i, iea, padv):
    in_s = 3
    pad = in_s-1
    img_i_p = pad_img(img_i, pad, padv)

    print("post pad")
    pimg(img_i_p)

    img_o = [[0]*len(img_i_p[i]) for i in range(len(img_i_p))]
    
    for i in range(len(img_i_p)):
        for j in range(len(img_i_p[i])):
            idx = get_block(img_i_p, i, j, in_s, padv)

            img_o[i][j] = iea[idx]
    print("enhanced")
    pimg(img_o)
    return img_o


def count(img):
    s = 0
    for ln in img:
        s += sum(ln)
    return s

def pimg(img):
    print()
    for ln in img:
        print(ln)
    print()

def main():
    # path = "./aoc/adventofcode/day20/day20_input_test.txt"
    # path = "./aoc/adventofcode/day20/day20_input_one.txt"
    path = "./aoc/adventofcode/day20/day20_input.txt"

    iea, img = read_input(path)
    print("iea")
    print(iea)
    print("original img")
    pimg(img)

    img_o1 = enhance(img, iea, 0)
    img_o2 = enhance(img_o1, iea, img_o1[0][0])

    print("img enhanced - step 1")
    pimg(img_o1)


    print("img enhanced - step 2")
    pimg(img_o2)

    c = count(img_o2)
    print("COUNT", c)

if __name__ == "__main__":
    main()

