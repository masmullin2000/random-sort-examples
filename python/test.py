import random

randomlist = []

for i in range(0,1000000):
    x = random.randint(0,100) 
    randomlist.append(x)

for i in range(0,1000):
    my_list = randomlist.copy()
    my_list.sort()

#print(randomlist)
