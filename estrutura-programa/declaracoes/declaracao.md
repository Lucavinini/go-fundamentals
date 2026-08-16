2.2

Uma declaração dá nome a uma entidade do programa e especica algumas ou todas as suas propriedades. Há quatro tipos principais de declaração: var, const, type e func.

Um programa Go é armazenado em um ou mais arquivos cujos nomes terminam com .go. Cada arquivo começa com uma declaração package que informa o pacote do qual o arquivo faz parte. A declaração package é seguida de qualquer declaração import e, em seguida, por uma sequência de declarações de tipos, variáveis, constantes e funções no nível de pacote (package-level), em qualquer ordem. Por exemplo, o programa a seguir declara uma constante, uma função e duas variáveis: