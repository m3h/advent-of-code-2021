#!/usr/bin/env python3


def fire(yv, xv, tyl, tyh, txl, txh):

    py, px = 0, 0

    max_height = py

    positions = []
    # Assume we're going from top left to bottom right
    xvd = (xv > 0) * -1 + (xv < 0) * 1

    while not ((py < tyh and py < tyl) or (px > txl and px > txh)):
        if tyl <= py <= tyh and txl <= px <= txh:
            return True, positions, max_height

        px += xv
        py += yv
        max_height = max(py, max_height)

        positions.append((py, px))

        if xv != 0:
            xv += xvd
        yv -= 1

    return False, positions, max_height


def graph_fire(yv, xv, tyl, tyh, txl, txh):
    print(f"({xv, yv}, T: x=({txl}, {txh}), y=({tyl}, {tyh})")

    hit, positions, max_height = fire(yv, xv, tyl, tyh, txl, txh)

    positions_y = [p[0] for p in positions]
    positions_x = [p[1] for p in positions]
    Sx, Sy = 0, 0
    S = (Sy, Sx)
    all_points_x = [Sx] + positions_x + [txl, txh]
    all_points_y = [Sy] + positions_y + [tyl, tyh]

    minx, maxx = min(all_points_x), max(all_points_x)
    miny, maxy = min(all_points_y), max(all_points_y)

    for y in range(maxy, miny-1, -1):
        for x in range(minx, maxx+1):
            p = (y, x)

            if p == S:
                print('S', end='')
            elif p in positions:
                print('#', end='')
            elif tyl <= p[0] <= tyh and txl <= p[1] <= txh:
                print('T', end='')
            else:
                print('.', end='')
        print()

    if not hit:
        if len(positions) <= 1:
            print("TARGET MISSED IMMEDIATELY! ({len(positions)} steps)")
        else:
            print(f"TARGET MISSED! ({len(positions)} steps)")
    else:
        print(f"TARGET HIT! ({len(positions)} steps)")
    return hit, positions, max_height


    # graph_fire(2, 7, tyl, tyh, txl, txh)
    # graph_fire(3, 6, tyl, tyh, txl, txh)
    # graph_fire(0, 9, tyl, tyh, txl, txh)
    # graph_fire(-4, 17, tyl, tyh, txl, txh)
if __name__ == "__main__":

    tyl = -10
    tyh = -5
    txl = 20
    txh = 30

    yv = 0
    yv_d = 1
    xv = 1
    xv_d = 1

    while True:

        while True:
            # hit, positions, max_height = graph_fire(yv, xv, tyl, tyh, txl, txh)
            hit, positions, max_height = fire(yv, xv, tyl, tyh, txl, txh)

            yv += yv_d

            lx = positions[-1][1]
            if lx < txl and lx < txh:
                break
            if yv_d == 1:
                ly = positions[-1][0]
                if ly > tyl and ly > tyh:
                    yv = 0
                    yv_d = -1
            else:
                fy = positions[0][0]
                if fy < tyl and fy < tyh:
                    break
        
        xv += xv_d

        if xv > txl and xv > txh:
            break