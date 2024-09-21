import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Soal2 {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Total belanja seorang customer: Rp ");
        long totalBelanja = scanner.nextLong();

        System.out.print("Pembeli membayar: Rp ");
        long uangDibayar = scanner.nextLong();

        Object result = hitungKembalian(totalBelanja, uangDibayar);
        
        if (result instanceof String) {
            System.out.println(result);
        } else {
            Map<Long, Long> pecahan = (Map<Long, Long>) result;
            long kembalian = calculateKembalian(totalBelanja, uangDibayar);
            System.out.println("Kembalian yang harus diberikan kasir: " + (uangDibayar - totalBelanja) + ", dibulatkan menjadi " + kembalian);
            System.out.println("Pecahan uang:");
            for (Map.Entry<Long, Long> entry : pecahan.entrySet()) {
                long nilai = entry.getKey();
                long jumlah = entry.getValue();
                if (nilai >= 1000) {
                    System.out.println(jumlah + " lembar " + nilai);
                } else {
                    System.out.println(jumlah + " koin " + nilai);
                }
            }
        }
    }

    public static Object hitungKembalian(long totalBelanja, long uangDibayar) {
        if (uangDibayar < totalBelanja) {
            return "Jumlah uang yang dibayarkan kurang dari total belanja.";
        }

        long kembalian = uangDibayar - totalBelanja;
        kembalian = (kembalian / 100) * 100; // dibulatkan ke bawah Rp.100

        long[] pecahanUang = {100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100};
        Map<Long, Long> hasilPecahan = new HashMap<>();

        for (long nilai : pecahanUang) {
            if (kembalian >= nilai) {
                long jumlah = kembalian / nilai;
                hasilPecahan.put(nilai, jumlah);
                kembalian -= jumlah * nilai;
            }
        }

        return hasilPecahan;
    }

    public static long calculateKembalian(long totalBelanja, long uangDibayar) {
        long kembalian = uangDibayar - totalBelanja;
        return (kembalian / 100) * 100; // dibulatkan ke bawah Rp.100
    }
}
