def Fibonacci(n):
    if n<0:
        print("Geçersiz Değer Girdiniz.")   
    elif n==1 or n==2:
        return 1
    
    elif n==0:
        return 0
    else:
        return Fibonacci(n-1)+Fibonacci(n-2)
    
n=int(input("Bulmak istediğiniz fibonacci değerini giriniz:"))
print(Fibonacci(n))
    
        