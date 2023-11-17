## Editor-Config

### Qual a utilidade do Editor-Config?

Em um projeto onde existem vários colaboradores, utilizando várias IDEs e Editores diferentes, padronizar a formatação é uma boa prática.

O [editor-config](https://editorconfig.org/) permite esta padronização. Atravez de um arquivo de configuração ele define uma séria de regras sobre a formatação do código. Como por exemplo: Espaços para identação, Tamanho da fonte, Espaços, etc..

Estas configurações podem ser aplicadas em arquivos diferentes. Por exemplo, com as chaves [] você poderá definir quais arquivos serão afetados por uma regra especifica:

```
[{Makefile,go.mod,go.sum,*.go,.gitmodules}]

indent_style = tab
indent_size = 4
```

Os arquivos que estão dentro das chaves, serão afetados pelas configurações que estão logo abaixo, como o style da identação e o tamanho da mesma. 

---

### Como instalar o Editor-Config?

editor-config é uma extenção e deve ser instalada no seu editor. Grande parte dos editores do mercado o aceitam.

1º - Instale a extenção no seu editor
2º - Crie um arquivo na raiz do seu projeto chamado ".editorconfig"
3º - Configure as regras do arquivo, conforme falamos anteriormente. 

---

### Porque uma quebra de linha no final do arquivo faz diferença?

Existe uma diferença em como diferentes sistemas operacionais lidam com a quebra de linha.
Por exemplo, o Windows usa 2 caracteres enquanto o Linux/Unix utiliza apenas 1 caractere.

A falta de quebra de linha no final do arquivo pode ocasionar, em alguns casos, a não visualização da última linha do arquivo e até problemas com mesclagem (merge) de arquivos durante controle de versão.

Além da questão entre os sistemas operacionais, algumas ferramentas tem como padrão esperar a quebra da última linha, como os linters (Ferramentas de verificação de código) que pode emitir avisos ou erros se não tiver a quebra de linha.

---