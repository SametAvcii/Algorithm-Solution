
def find_next_square(sq):
   import math
   if math.sqrt(sq).is_integer():
    n=sq+1
    while(math.sqrt(n).is_integer() == False):
     n+=1
    return n  
   else:
    return -1