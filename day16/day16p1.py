#!/usr/bin/env python3

import math

def hex_to_binary(h: str) -> str:
    return bin(int(h, 16))[2:]

def int2(b: str) -> int:
    return int(b, 2)

def general_get(n, pb):
    val = pb[0:n]
    pb = pb[n:]
    return val, pb

def decode_literal(pc):
    val_bin = ""

    nibbles = 0
    last = False
    while not last:
        nibbles += 1
        b = pc.get(5)
        last = b[0] == "0"

        val_bin += b[1:]

    bin_len = 3+3+5*nibbles 
    pad_len = math.ceil(bin_len/4)*4 - bin_len

    padding = pc.get(pad_len)
    assert padding == "0"*pad_len

    val = int2(val_bin)

    print("literal value:", val)


class StringConsumer:
    def __init__(self, s):
        self.s = s
    def empty(self):
        return len(self.s) > 0
    def get(self, n):
        val, self.s = general_get(n, self.s)
        return val


def decode_packet(pc):


    version = int2(pc.get(3))
    typeid = int2(pc.get(3))

    if typeid == 4:
        return decode_literal(pc, version)
    
    # anything else implies operator

    version_sum = 0
    len_type = pc.get(1)
    if len_type == "0":

        subpacket_len = int2(pc.get(15))

        subpacket_bin = pc.get(subpacket_len)
        # create new pc
        pcs = StringConsumer(subpacket_bin)
        while not pcs.empty():
            sub_version_sum = decode_packet(pcs)
            version_sum += sub_version_sum
    elif len_type == "1":
        subpacket_count = int2(pc.get(11))
        for i in range(subpacket_count):
            sub_version_sum = decode_packet(pc)
            version_sum += sub_version_sum
    else:
        assert False, f"Unexpected len_type: {len_type}"
    
    return version_sum
    
    return version_sum



def main(ph):
    pb = hex_to_binary(ph)

    pc = StringConsumer(pb)

    print("RESULT", decode_packet(pc))

if __name__ == "__main__":
    ph = "D2FE28"
    main(ph)