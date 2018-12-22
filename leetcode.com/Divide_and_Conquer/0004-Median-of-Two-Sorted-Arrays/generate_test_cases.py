import random, json, sys


test_cases = []


def lp(title, list):
    print( "{}_keys: [{}]".format(title, ",".join(["{:>3}".format(i) for i in range(len(list))])))
    print( "{}_vals: [{}]".format(title, ",".join(["{:>3}".format(i) for i in list])))

def median(a, b):
    c = sorted(a+b)

    values = []

    if len(c)%2 == 0:
        values.append(c[(len(c)//2)-1])

    values.append( c[len(c)//2] )


    # print("Total {} / I is {}".format(len(c), sum(values) / len(values)), values)

    # lp('c', c)


    return sum(values) / len(values)

def medians(a, b):
    medians_list = []

    c = sorted(a+b)
    if len(c)%2 == 0:
        medians_list.append((len(c)//2)-1)

    i = (len(c)//2)
    medians_list.append(i)
    return medians_list


for i in range (777):

    a,b=[],[]
    top_a = random.randint(2, 2**3)
    top_b = random.randint(2, 2**3)

    for i in range(top_a):
        a.append(random.randint(0, top_a*5))

    for i in range(random.randint(2, 2**4)):
        b.append(random.randint(0, top_b*5))

    a, b = sorted(a), sorted(b)
    c = sorted(a+b)

    test_case = {
        'a' : a,
        'b' : b,
        'c' : c,
        'mi' : medians(a, b),
        'm' : median(a, b) * 1.0
    }


    test_cases.append(test_case)

# // (\n\s{12})
with open('test-cases.json', 'w') as f:
    json.dump(test_cases, fp=f, indent=4)


# print(test_cases)




# def search(n, a, b):

#     bg, sm = (a, b) if len(a) > len(b) else (b, a)

#     lp('bg', bg)
#     lp('sm', sm)

#     clen = len(a) + len(b)

#     if bg[len(bg)-1] >= sm[0]:
#         return sm[n] if n < len(sm) else bg[n-len(sm)]
#     if sm[len(sm)-1] >= bg[0]:
#         return bg[n] if n < len(bg) else sm[n-len(bg)]



#     print(a,b, clen)



# print(search(6, [1, 2, 5], [1, 5, 6, 8]))


# print()
# lp('a', a)
# print("-"*40)
# lp('b', b)
# print("-"*40)
# expected = test(a, b)
# print(expected)

