// kalkulator dian
package main
import ( 
"fmt"
"os"
)
func main() {
	var bil1, bil2 float64
	var operator string
	// Membaca bilangan pertama
	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scanln(&bil1)
	// Membaca operator matematika
	fmt.Print("Masukan operator matematika (+, - , *, /): ")
	fmt.Scanln(&operator)
	//Masukan bilangan kedua
	fmt.Print("Masukan bilangan kedua: ")
	fmt.Scanln(&bil2)
	
	var hasil float64
	switch operator {
	case "+":
		hasil = bil1 + bil2
	case "-":
		hasil = bil1 - bil2
	case "*":
		hasil = bil1 * bil2
	case "/":
		if bil2 !=0 {
			hasil = bil1 / bil2
		} else {
				fmt.Println("Error: Pembagian dengan nol tidal diizinkan.")
				os.Exit(1)
		}
		default:
			fmt.Println("Operator matematika tidak valid.")
			os.Exit(1)
		}
		// Menampilkan hasil
		fmt.Printf("Hasil dari %.2f %s %.2f = %.2f\n", bil1, operator, bil2, hasil)
}