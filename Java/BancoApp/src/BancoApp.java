import java.text.DecimalFormat;
import java.util.Scanner;

public class BancoApp {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        double valorInicial = scanner.nextDouble();

        double taxaJuros = scanner.nextDouble();

        int periodo = scanner.nextInt();

        double valorFinal = valorInicial;


            for (int i = 0; i < periodo; i++) {
            valorFinal *= (1 + taxaJuros);
        }

        DecimalFormat df = new DecimalFormat("0.00");

        System.out.println("Valor final do investimento: R$ " +  df.format(valorFinal));

        scanner.close();
    }
}