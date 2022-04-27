def PrimeMover(num):
  a=[]
  b=2
  c=100
  x=0
  control=True
  while len(a)!=num-1:
    control=True
    b+=1
    for i in range(2, b):
      if (b % i) == 0:
        control = False    
        break
    if control==True:    
     a.append(b)  
  return a[-1]


# keep this function call here 
print(PrimeMover(input()))