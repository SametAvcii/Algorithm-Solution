from random import randint
rand=randint(1,100)
sayac=0

while True:
    sayac+=1
    sayi=int(input('1 ile 100 Arasında Bir Sayı Giriniz:'))
    if(sayi==0):
        print('Yanlış Değer Girdiniz')
        break
    elif sayi<rand:
        print("Daha Yüksek Bir Değer Giriniz")
        continue
    elif sayi>rand:
        print("Daha Düşük Bir Sayı Giriniz")
        continue
    else:
        print('Rastgele Seçilen Sayı {0}!'.format(rand))
        print('Tahmin Sayınız {0}'.format(sayac))
    