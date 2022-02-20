def high(x):
    arr,numArr= x.split(' '),[]
    for val in arr:
        intNum = 0
        letters = list(val)
        for word in letters:
            intNum += ord(word) - 96
        numArr.append(intNum)
    return arr[numArr.index(max(numArr))]
        

print(high('what time are we climbing up the volcano'))