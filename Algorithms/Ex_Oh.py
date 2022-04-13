def ExOh(strParam):

  list=[]
  for i in strParam:
    if i == 'x' or 'o':
      list.append(i)

  count_x=0
  count_o=0
  for i in list:
    if i == 'x':
      count_x+=1
    elif i=='o':
      count_o+=1    

  if count_o==count_x:
    return True
  else:
    return False  
  


# keep this function call here 
print ExOh(raw_input())