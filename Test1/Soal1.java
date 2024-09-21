import java.util.Scanner;

public class Soal1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        System.out.print("Masukkan jumlah string: ");
        int N = scanner.nextInt();
        scanner.nextLine(); // untuk membersihkan newline
        
        String[] strings = new String[N];
        
        for (int i = 0; i < N; i++) {
            System.out.print("Masukkan string " + (i + 1) + ": ");
            strings[i] = scanner.nextLine();
        }
        
        String result = findMatchingStrings(N, strings);
        System.out.println(result);
    }

    public static String findMatchingStrings(int N, String[] strings) {
        StringBuilder output = new StringBuilder();
        boolean foundMatch = false;

        for (int i = 0; i < N; i++) {
            for (int j = i + 1; j < N; j++) {
                if (strings[i].equalsIgnoreCase(strings[j])) {
                    if (!foundMatch) {
                        output.append((i + 1)).append(" ");
                        foundMatch = true;
                    }
                    output.append((j + 1)).append(" ");
                }
            }
        }
        
        return output.length() > 0 ? output.toString().trim() : "False";
    }
}
