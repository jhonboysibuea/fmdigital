import java.util.Stack;

public class Soal3 {

    public static boolean isValidString(String s) {
        // Cek panjang string sesuai dengan batasan yang diberikan
        if (s.length() < 1 || s.length() > 4096) {
            return false;
        }

        // Stack untuk menyimpan karakter pembuka
        Stack<Character> stack = new Stack<>();

        // Loop melalui setiap karakter dalam string
        for (char c : s.toCharArray()) {
            // Jika karakter adalah pembuka, masukkan ke stack
            if (c == '<' || c == '{' || c == '[') {
                stack.push(c);
            } 
            // Jika karakter adalah penutup
            else if (c == '>' || c == '}' || c == ']') {
                // Jika tidak ada pasangan pembuka di stack atau tidak cocok
                if (stack.isEmpty() || !isMatchingPair(stack.pop(), c)) {
                    return false;
                }
            }
        }

        // String valid jika semua pembuka telah ditutup (stack harus kosong)
        return stack.isEmpty();
    }

    // Fungsi untuk memeriksa apakah karakter pembuka dan penutup cocok
    private static boolean isMatchingPair(char open, char close) {
        return (open == '<' && close == '>') ||
               (open == '{' && close == '}') ||
               (open == '[' && close == ']');
    }

    public static void main(String[] args) {
        // String panjang yang diberikan
        String string1 = "{{[<>[{{}}]]}}"
                + "{<{[[{{[]<{{[{[]<>}]}}<>>}}]]}>}"
                + "{{[{<[[{<{<<<[{{{[]{<{[<[[<{{[[[[[<{[{<[<<[[<<{[[{[<<"
                + "<<<<<[{[{[{{<{[[<{<<<{<{[<>]}>}>>[]>}>]]}>}}]}]}]>>>>>><"
                + ">]}]]}>>]]>>]>}]}>]]]]]}}>]]>]}>}}}}]>>>}>}]]>}]}]"
                + "[<{<{[{[{}[[<[<{{[<[<[[[<{{[<<<[[[<[<{{[<<{{<{<{<[<{["
                + "{{[{{{{[<<{{{<{[{[[[{<<<[{[<{<<<>>>}>]}]>>>}]]]}]}>}}"
                + "}>>]}}}}]}}]}>]>}>}>}}>>]}}>]>]]]>>>]}}>]]]>]>]}}>]>]"
                + "]]}]}>}>]"
                + "[[{[[<{{{{[[[<[{[[<{{{{[{[{[[[<<{<{[{<<<[[<[{[<{{["
                + "{[<[[<<[{<<[[[{<[{[[{{<<>[<<{{<<{[[[<{}{[{{{[[{{[[<[{}]"
                + ">]]}}]]}}}]}>]]]}>>}}>>]>}}]]}]>}]]]>>}]>>]]>]}]}}>]}]>"
                + "]]>>>}]}>}>>]]]}]}]}}}}>]]}]>]]]}}}}>]]}]]"
                + "[{}<>]";

        // Menampilkan hasil validasi
        System.out.println(isValidString(string1)); 
    }
}
