# Interfaces in Go

As interfaces em Go permitem definir um conjunto de assinaturas de métodos que qualquer tipo pode implementar. 
Isto fornece uma maneira flexivel de escrever código.

## O que são?

Por definição, interface é um tipo que define um conjunto de assinaturas de métodos.
Se um tipo implementa todos os métodos contidos na assinatura da interface, este tipo implementa (pode usar) a interface. 
Além disso a implementação do método deve ser identica do ponto de vista dos argumentos recebidos e do retorno esperado.
No exemplo abaixo, a interface `Shape` define as assinaturas de método: `Area`, `Perimeter` e `Circumference`.

```
package interfaces


type Shape interface {
	Area(n ...int) float64
	Perimeter(n ...int) float64
	Circumference(n ...int) float64
}
```

Qualquer tipo que possua as assinaturas de métodos acima pode satisfazer esta interface.
(ps: A definição da interface e das assinaturas com letras maiusculas significam que elas serão exportadas, ou seja, poderão ser utilizadas em outros pacotes)

## Como implementa-las?

No exemplo abaixo, o tipo `Rectangle` implementa a interface `Shape`.

```
package shapes


type Rectangle struct{}
```

Agora precisamos definir os métodos que a interface `Shape` espera. Definindo assim, o comportamento do tipo Rectangle para calcular a area, por exemplo:

```
func (r Rectangle) Area(length, width int) float64 {
	return float64(length * width)
}
```

Desta forma, o tipo `Rectangle` poderá implementar a interface `Shape` e consequentemente seu método que calcula a área. Tá, mas o que ganhamos com isso?

Essa abordagem proporciona flexibilidade ao lidar com diferentes formas geométricas, permitindo calcular a area de um circulo, quadrado, etc... Isso resulta em um código mais flexível e fácil de estender. 

## Como usa-las?

No exemplo abaixo nós:
- Importamos a biblioteca shapes
- Implementamos a interface Shape para um Retangulo
- Chamamos o método Area do tipo rectangle para calcular a área do Retangulo

Importante observarmos que tanto a interface quanto seus métodos são exportaveis (iniciando com letras maiusculas), porque, neste caso, precisamos deles em outros pacotes.

```
package main

import (
	"shapes"
)

func main() {
    ...
    rectangle := shapes.Rectangle{} //Rectangle é uma implementação da interface Shape para um Retangulo.
    rectangle.Area(10, 10) //Area é o método que define o comportamento do tipo Rectangle para calcular a área.
}
```


## Sobre o projeto Geometry.

Este tópico contempla um projeto que visa utilizar as interfaces para calcular a área de diversas formas geométricas.
Tem como objetivo, único e exclusivo, o aprendizado na aplicação de interfaces em Go.



