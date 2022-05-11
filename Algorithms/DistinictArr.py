def DistinctList(arr):

  arrCount=[]
  count=0
  for i in arr:
    arrCount.append(count)
    count=0
    for x in arr:
      if i==x:
        count+=1

  arrCount.sort
  if arrCount[-1]<=1:
    return 0
  return arrCount[-1]

# keep this function call here 
print(DistinctList(input()))