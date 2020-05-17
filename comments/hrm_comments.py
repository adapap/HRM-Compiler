from __future__ import print_function
import numpy as np
import turtle
from argparse import ArgumentParser
from base64 import decodestring
from zlib import decompress


def parse_images(filepath):
    with open(filepath) as f:
        raw_data = bytes(f.read().rstrip(';'), encoding='utf-8')
        # add base64 padding
        if len(raw_data) % 4 != 0:
            raw_data += b'=' * (2 - (len(raw_data) % 2))
        print(raw_data)
        # decode base64 -> decode zlib -> convert to byte array
        data = np.fromstring(decompress(decodestring(raw_data)), dtype=np.uint8)
        assert data.shape == (1028,)
        path_len, = data[:4].view(np.uint32)
        path = data[4:4 + 4 * path_len].view(np.uint16).reshape((-1, 2))
        yield path

    
def main():
    ap = ArgumentParser()
    ap.add_argument('--speed', type=int, default=10,
                    help='Number 1-10 for drawing speed, or 0 for no added delay')
    ap.add_argument('program')
    args = ap.parse_args()

    for path in parse_images(args.program):
        if not path.size:
            continue
        pen_up = (path == 0).all(axis=1)
        # convert from path (0 to 65536) to turtle coords (0 to 655.36)
        path = path / 100.
        turtle.speed(args.speed)
        turtle.setworldcoordinates(0, 655.36, 655.36, 0)
        turtle.pen(shown=False, pendown=False, pensize=10)
        for i, pos in enumerate(path):
            if pen_up[i]:
                turtle.penup()
            else:
                turtle.setpos(pos)
                turtle.pendown()
                turtle.dot(size=10)
        input('Press enter to continue')
        turtle.clear()
    turtle.bye()

main()
