def Consecutive(arr):

  list=[]
  for i in arr:
    list.append(i)
  list.sort()
  list2=[]

  for i in range(list[0],list[-1]+1):
    list2.append(i)
    
  return len(list2)-len(list)  
# keep this function call here 
print Consecutive(raw_input())