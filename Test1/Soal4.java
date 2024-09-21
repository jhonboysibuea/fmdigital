import java.time.LocalDate;
import java.time.temporal.ChronoUnit;

public class Soal4 {

    public static boolean canTakeCutiPribadi(int cutiBersama, LocalDate tanggalJoin, LocalDate tanggalRencanaCuti, int durasiCuti) {
        int cutiKantor = 14;  // Cuti kantor tahunan
        int maxCutiBerturut = 3;  // Maksimal cuti berturut-turut adalah 3 hari

        // Hitung jumlah cuti pribadi
        int cutiPribadi = cutiKantor - cutiBersama;

        // Hitung tanggal setelah 180 hari sejak tanggal join
        LocalDate tanggalBolehCuti = tanggalJoin.plusDays(180);

        // Jika karyawan meminta cuti sebelum 180 hari, langsung return false
        if (tanggalRencanaCuti.isBefore(tanggalBolehCuti)) {
            System.out.println("False: Karyawan belum mencapai 180 hari kerja.");
            return false;
        }

        // Cek apakah durasi cuti melebihi batas maksimal cuti berturut-turut
        if (durasiCuti > maxCutiBerturut) {
            System.out.println("False: Cuti pribadi tidak boleh lebih dari 3 hari berturut-turut.");
            return false;
        }

        // Hitung jumlah hari dari tanggalBolehCuti sampai akhir tahun
        LocalDate akhirTahun = LocalDate.of(tanggalRencanaCuti.getYear(), 12, 31);
        long sisaHariTahun = ChronoUnit.DAYS.between(tanggalBolehCuti, akhirTahun);

        // Hitung jatah cuti pribadi berdasarkan sisa hari hingga akhir tahun
        int jatahCutiPribadi = (int) (sisaHariTahun * cutiPribadi / 365.0); // Pembulatan ke bawah

        // Jika jatah cuti pribadi kurang dari durasi yang diajukan, return false
        if (jatahCutiPribadi < durasiCuti) {
            System.out.println("False: Jatah cuti pribadi tidak mencukupi.");
            return false;
        }

        // Jika semua validasi lolos, karyawan boleh mengambil cuti
        System.out.println("True: Karyawan boleh mengambil cuti pribadi.");
        return true;
    }

    public static void main(String[] args) {
        // Contoh penggunaan
        int cutiBersama = 7;
        LocalDate tanggalJoin = LocalDate.of(2021, 5, 1);
        LocalDate tanggalRencanaCuti = LocalDate.of(2021, 11, 10);
        int durasiCuti = 1;

        boolean hasil = canTakeCutiPribadi(cutiBersama, tanggalJoin, tanggalRencanaCuti, durasiCuti);
        System.out.println(hasil);
    }
}
