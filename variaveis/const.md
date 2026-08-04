## Constantes em GO
Se uma variável deve ter um valor fixo que não pode ser alterado, você pode usar a palavra chave `const` para declará-la.

A const declara a variável como "constante", o que significa que ela é imutável e somente leitura .

### Exemplo de declaração de constante
```go
const NOTA int = 8
```
### Aviso: A constante deve ser inicializada no momento da declaração, caso contrário, ocorrerá um erro de compilação.

### Regras para declaração de constantes
- Os nomes das constantes geralmente são escritos em letras maiúsculas (para facilitar a identificação e a diferenciação das variáveis).
- As constantes podem ser declaradas tanto dentro quanto fora de uma função.

### Tipos de constantes
- Constante tipada: Quando você declara uma constante com um tipo específico, como `int`, `float64`, `string`, etc.

Exemplo:
```go
const PI float64 = 3.145445 // o programador atribuiu um tipo específico à constante
```

- Constante não tipada: Quando você declara uma constante sem especificar um tipo, o Go infere o tipo com base no valor atribuído.

Exemplo:
```go
const PI = 3.145445 // o Go infere o tipo da constante com base no valor atribuído
```