// Um arquivo .go consiste nas seguintes partes:

package main //declaração de pacote
import ("fmt") //Importação de pacotes a serem usados

func main(){ //Função principal
	fmt.Println("Oláaaa") //Declarações e expressões

	/* Declaração tradicional var + nome + tipo
	
	var MinhaIdade int = 10
	var MeuPeso float32 = 68.02
	var MeuNome string = "Lucas Vinícius"
	var StatusCorteDeCabelo bool = false
	
	fmt.Println(MinhaIdade, MeuPeso, MeuNome, StatusCorteDeCabelo)
	*/
	
	// Declaração inferida - tipo da variável conforme o valor
	// variablename := value
	// var = value

	/*
	x := 40
	var student2 = "Jane"
	fmt.Println(student2, x)
	*/

	var studentname string
	studentname = "Lucas"
	fmt.Println(studentname)
	
}