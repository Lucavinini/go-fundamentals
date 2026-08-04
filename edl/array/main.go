package main

import ("fmt")

func main(){
	var notasPeriodo = [5]int{4,6,7,8,0}
	notasPeriodo2 := []float32{4.6,7.5,3.4}
	disciplinaCurso := [4]string{"BD", "Árvores", "Integração"}

	fmt.Println(notasPeriodo)
	fmt.Println(notasPeriodo2)
	fmt.Println(disciplinaCurso[2])
}