package main
import "fmt"

// Fonction pour calculer la somme de deux entiers
func add(a int, b int) int {
    return a + b
}
func main(){
  // Appel de la fonction avec des valeurs
    result := add(3, 5)
    fmt.Println("La somme est :", result) // Affiche "La somme est : 8"
}
