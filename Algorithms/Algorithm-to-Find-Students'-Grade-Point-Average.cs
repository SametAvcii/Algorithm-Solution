using System;

namespace Algorithms
{
    class Program
    {
        static void Main(string[] args)
        {
            //Notu girilen 10 öğrencinin notlarının ortalaması 70'den büuyüykse iyi değilse kötü yazdıran algoritması
            int x = 0;
            int toplam = 0;
            int ortalama = 0;
            for (int i = 1; i < 11; i++)
            {
               
                Console.WriteLine(i+". öğrencinin notunu giriniz:");
                x =Convert.ToInt32(Console.ReadLine());
                toplam += x;

            }
            ortalama = toplam / 10;
            if (ortalama>=70)
            {
                Console.WriteLine("Öğrencileirn not ortalaması İyi");
            }
            else
            {
                Console.WriteLine("Öğrencileirn not ortalaması Kötü");
            }



        }
    }
}
