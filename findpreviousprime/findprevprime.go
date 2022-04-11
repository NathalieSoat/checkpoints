package checkpoints

//import "fmt"

func FindPrevPrime(nb int) int {
	result := nb
	if nb < 2 {
		return 2
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return FindPrevPrime((nb - 1))
		}
	}
	return result
}

//func main() {
//	fmt.Println(FindPrevPrime(10))
//	fmt.Println(FindPrevPrime(4))
//}
