#!/usr/bin/python3
import asyncio


@asyncio.coroutine
def make_a_file(i):
    f = "out_python/{:05d}".format(i)
    try:
        with open(f, "w") as fp:
            fp.write("hello")
        # print("wrote", f)
    except:
        print("failed on", f)


def make_a_lot_of_files(how_many):
    loop = asyncio.get_event_loop()
    tasks = [asyncio.ensure_future(make_a_file(i)) for i in range(how_many)]
    loop.run_until_complete(asyncio.wait(tasks))


if __name__ == "__main__":
    from sys import argv
    how_many = int(argv[1]) if len(argv) == 2 else 100000
    make_a_lot_of_files(how_many)
