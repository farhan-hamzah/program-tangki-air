package main
import "fmt"

func tangkiAir(n, m *int){
	var b bool
	if *n > 0{
		*n -= *m
	}
	if *n <= 0 {
		b = true
	}else {
		b = false
	}
	fmt.Print(b)
}
func main(){
	var n, m int
	fmt.Scan(&n, &m)
	tangkiAir(&n, &m)
	for n > 0{
		fmt.Scan(&m)
		tangkiAir(&n, &m)
	}
}