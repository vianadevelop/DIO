import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Digite um número inteiro para verificar se é primo: ");
        int numero = scanner.nextInt();

        if (ehprimo(numero)) {
            System.out.println(numero + " é um número primo");
        } else {
            System.out.println(numero + " não é um número primo");
        }
        scanner.close();
    }

    public static boolean ehprimo(int numero) {
        if (numero <= 1) {
            return false;
        }

        for (int i = 2; i < numero; i++) {
            if (numero % i == 0) {
                return false;
            }
        }

        return true;
    }


}