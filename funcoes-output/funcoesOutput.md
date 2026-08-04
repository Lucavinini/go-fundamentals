# Funções de Saída em GO
Go possui três funções para gerar saída de dados no console: `fmt.Print()`, `fmt.Println()` e `fmt.Printf()`. Essas funções são fornecidas pelo pacote `fmt`, que é parte da biblioteca padrão do Go.

## 1. `fmt.Print()`
A função `fmt.Print()` é usada para imprimir valores no console sem adicionar uma nova linha ao final. Ela aceita múltiplos argumentos e os imprime na ordem em que são fornecidos.

Exemplo:
```go
	var i,j string = "Hello","world"

	fmt.Print(i)
	fmt.Print(j)
```
Saída:
```
Helloworld
```

## 2. `fmt.Println()`
A função `fmt.Println()` é semelhante à `fmt.Print()`, mas adiciona uma nova linha ao final da saída. Ela também aceita múltiplos argumentos e os imprime na ordem em que são fornecidos.

Exemplo:
```go
  var i,j string = "Hello","world"

  fmt.Println(i)
  fmt.Println(j)
```
Saída:
```
Hello
world
```

## 3. `fmt.Printf()`
A `Printf()função` primeiro formata seus argumentos com base no verbo de formatação fornecido e, em seguida, os imprime.

Aqui usaremos dois verbos de formatação:

`%v`  é usado para imprimir o value dos argumentos.
`%T` é usado para imprimir o type dos argumentos.

Ex:
```go

	var nomeCurso string = "Computação"
	var tempoConclusao int = 4

	fmt.Printf("O nome do curso é %v e o seu tipo é %T\n", nomeCurso, nomeCurso)
}
```
Saída: 
```
O nome do curso é Computação e o seu tipo é string
```
