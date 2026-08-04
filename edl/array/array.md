# Arrays em Go
São usados para armazenar uma coleção de elementos do mesmo tipo. Um array tem um tamanho fixo, que é definido no momento da sua criação.

## Declaração de Arrays
Em Go, existem duas maneiras de declarar um array:

1. Com o var
```go
var nomeArray[5]int{1, 3, 4 , 5, 6} //Com a definição do tamanho

//ou

var nomeArray2 = []int{1, 3, 4 , 5, 6} //Sem a definição do tamanho
```

2. Com o operador de atribuição curta
```go
nomeArray := [5]int{1, 3, 4 , 5, 6} //Com a definição do tamanho
//ou
nomeArray2 := []int{1, 3, 4 , 5, 6} //Sem a definição do tamanho
```

[DICA]: Você pode acessar um elemento específico da matriz fazendo referência ao número do índice. Em Go, os índices de array começam em 0. Isso significa que [0] é o primeiro elemento, [1] é o segundo elemento, etc.]