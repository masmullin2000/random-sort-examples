import random

# randomlist = []
# for i in range(0, 1000000):
    # x = random.randint(0, 99) 
    # randomlist.append(x)

def do_it():
    randomlist = [0] * 1_000_000
    for i in range(0, 1_000_000):
        randomlist[i] = random.randint(0,99)

    for i in range(0, 250):
        my_list = randomlist.copy()
        my_list.sort()

if __name__ == "__main__":
    do_it()
