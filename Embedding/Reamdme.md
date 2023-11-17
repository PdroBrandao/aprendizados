## Embedding

### O que é o Embedding?

Do inglẽs embutimento, significa incorporar/embutir. No contexto de estruturas de dados (structs), quando utilizamos o embedding estamos trazendo um tipo para dentro de um tipo, ou seja, incorporando um tipo dentro de outro tipo. Com o embedding uma struct pode ter outras structs incorporadas, isso faz com que você consiga acessar todos métodos e campos do tipo incorporado.

O Uso do Embedding facilita a reutilização do código e promove um design flexivel. O uso de forma demasiada pode causar confusões, uma vez que se utilizar muitos niveis de embedding pode gerar dificuldade em identicar métodos, por exemplo.

---

### Composição vs Herança

Herança existe em linguagens orientadas a objetos. Nada mais é que a possibilidade de uma classe ou tipo, herdar métodos e campos de outra.
Como o Go não é orientada a objetos, podemos utilizar o embedding para chegar nos "mesmos" resultados. 

Incorporando um tipo dentro de outro, conseguimos acessar seus respectivos campos e métodos. 

