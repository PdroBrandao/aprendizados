## ID x UUID

### Qual a diferença entre os dois?

Ambos são identificadores, usados pra identificar recursos de forma exclusiva. O ID - normalmente - respeita o formato e ordem dos números inteiros de forma descrente, enquanto o UUID já tem um formato estabelecido, são 32 caracteres, separados 8-4-4-4-12, por exemplo:

```
f47ac10b-58cc-4372-a567-0e02b2c3d479
```


### Quando devo utilizar ID?

ID são mais simples de gerar e ocupam menos espaço, são indicados pra situações onde a garantia de unicidade dentro do sistema é suficiente.

### Quando devo utilizar UUID?

UUID são recomendados para sistemas distribuidos, bancos de dados descentralizados.
Ocupam mais espaços e são mais complexos de gerar.

