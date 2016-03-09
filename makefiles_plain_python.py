#!/usr/bin/python3

def make_a_file(i):
    f = "out_python/{:05d}".format(i)
    try:
        with open(f, "w") as fp:
            fp.write("hello")
        # print("wrote", f)
    except:
        print("failed on", f)


def make_a_lot_of_files(how_many):
    for i in range(how_many):
        make_a_file(i)
    return None

print(make_a_lot_of_files(100000))
