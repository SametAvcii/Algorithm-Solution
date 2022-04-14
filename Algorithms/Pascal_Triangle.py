def GetPascal(n):
  list=[]
  list.append(1)
  if n==0:
    return list
  prev=GetPascal(n-1)

  for i in range(1,len(prev)):
    a=prev[i-1]+prev[i]
    list.append(a)
  list.append(1)  
  return list


def PascalsTriangle(arr):
  arr1=[]
  for i in range(1,7):
    arr1=GetPascal(i)
    if arr[1]==arr1[1]:      
      break
  a=len(arr)  
  return arr1[a]

  


# keep this function call here 
print PascalsTriangle(raw_input())