#!/usr/bin/env python3

import math

def padded_bin(h: str) -> str:
    assert len(h) == 1
    n = int(h, 16)
    b = bin(n)[2:]
    b_padded = b.rjust(4, '0')
    return b_padded


def hex_to_binary(h: str) -> str:
    return ''.join([
        padded_bin(x) for x in h
    ])

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

    # bin_len = 3+3+5*nibbles 
    # pad_len = math.ceil(bin_len/4)*4 - bin_len

    # padding = pc.get(pad_len)
    # assert padding == "0"*pad_len

    val = int2(val_bin)

    print("literal value:", val)
    return val


class StringConsumer:
    def __init__(self, s):
        self.s = s
    def empty(self):
        return len(self.s) == 0
    def get(self, n):
        val, self.s = general_get(n, self.s)
        return val

def decode_len_0(pc):

    version_sum = 0

    subpacket_len = int2(pc.get(15))

    subpacket_bin = pc.get(subpacket_len)
    # create new pc
    pcs = StringConsumer(subpacket_bin)
    while not pcs.empty():
        sub_version_sum = decode_packet(pcs)
        version_sum += sub_version_sum
    
    return version_sum

def decode_len_1(pc):
    version_sum = 0

    subpacket_count = int2(pc.get(11))
    for i in range(subpacket_count):
        sub_version_sum = decode_packet(pc)
        version_sum += sub_version_sum
    return version_sum


def decode_packet(pc):


    version = int2(pc.get(3))
    typeid = int2(pc.get(3))

    version_sum = version
    if typeid == 4:
        decode_literal(pc)
        return version_sum
    
    # anything else implies operator
    len_type = pc.get(1)
    if len_type == "0":
        version_sum += decode_len_0(pc)
    elif len_type == "1":
        version_sum += decode_len_1(pc)
    else:
        assert False, f"Unexpected len_type: {len_type}"
    
    return version_sum


def main(ph):
    pb = hex_to_binary(ph)

    pc = StringConsumer(pb)

    print("RESULT", decode_packet(pc))

if __name__ == "__main__":
    ph = "D2FE28"
    ph = "38006F45291200"
    ph = "EE00D40C823060"
    ph = "8A004A801A8002F478" # vs sum of 16
    ph = "620080001611562C8802118E34" # vs sum of 12
    ph = "C0015000016115A2E0802F182340" # vs_sum = 23
    ph = "A0016C880162017C3686B18A3D4780" # vs_sum = 31

    ph = "020D74FCE27E600A78020200DC298F1070401C8EF1F21A4D6394F9F48F4C1C00E3003500C74602F0080B1720298C400B7002540095003DC00F601B98806351003D004F66011148039450025C00B2007024717AFB5FBC11A7E73AF60F660094E5793A4E811C0123CECED79104ECED791380069D2522B96A53A81286B18263F75A300526246F60094A6651429ADB3B0068937BCF31A009ADB4C289C9C66526014CB33CB81CB3649B849911803B2EB1327F3CFC60094B01CBB4B80351E66E26B2DD0530070401C82D182080803D1C627C330004320C43789C40192D002F93566A9AFE5967372B378001F525DDDCF0C010A00D440010E84D10A2D0803D1761045C9EA9D9802FE00ACF1448844E9C30078723101912594FEE9C9A548D57A5B8B04012F6002092845284D3301A8951C8C008973D30046136001B705A79BD400B9ECCFD30E3004E62BD56B004E465D911C8CBB2258B06009D802C00087C628C71C4001088C113E27C6B10064C01E86F042181002131EE26C5D20043E34C798246009E80293F9E530052A4910A7E87240195CC7C6340129A967EF9352CFDF0802059210972C977094281007664E206CD57292201349AA4943554D91C9CCBADB80232C6927DE5E92D7A10463005A4657D4597002BC9AF51A24A54B7B33A73E2CE005CBFB3B4A30052801F69DB4B08F3B6961024AD4B43E6B319AA020020F15E4B46E40282CCDBF8CA56802600084C788CB088401A8911C20ECC436C2401CED0048325CC7A7F8CAA912AC72B7024007F24B1F789C0F9EC8810090D801AB8803D11E34C3B00043E27C6989B2C52A01348E24B53531291C4FF4884C9C2C10401B8C9D2D875A0072E6FB75E92AC205CA0154CE7398FB0053DAC3F43295519C9AE080250E657410600BC9EAD9CA56001BF3CEF07A5194C013E00542462332DA4295680"

    main(ph)