using System;

namespace Algorithms
{
    class Program
    {
        static void Main(string[] args)
        {

            //1 ile 100 Arasında girilen sayıların 3 veya 5e bölünebilenlerini ekrana yazan kod
            int x, y, z;
            x = 0;
            y = 0;
            z = 0;

            for (int i = 1; i < 100; i++)
            {
                if (i % 3 == 0)
                {
                    Console.WriteLine(i);
                    y++;
                }
                if (i % 5 == 0)
                {
                    Console.WriteLine(i);
                    z++;
                }
            }
            Console.WriteLine("3'e bölünebilen sayi miktarı :" + y);
            Console.WriteLine("5'e bölünebilen sayi miktarı :" + z);



        }
    }
}
