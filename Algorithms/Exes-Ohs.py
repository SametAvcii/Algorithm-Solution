def xo(s):
    letters= list(s)
    x_count = 0
    o_count = 0

    for i in letters:
        if i.lower() == 'x':
            x_count += 1
        elif i.lower() == 'o':
            o_count += 1

    return x_count == o_count