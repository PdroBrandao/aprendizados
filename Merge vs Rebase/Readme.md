## Merge vs Rebase

Apresentei uma certa dificuldade nestes conceitos. Ao me deparar com a situação de colaborar com um projeto onde o fluxo de atualizações é grande, por várias pessoas, é mais que necessário dominar os comandos que "juntam" o código. 

### Merge?

O Merge mescla alterações e os históricos de 2 branchs diferentes. O comando "git merge branchA" mesclará as alterações feitas na "branchA" na branch atual. 
Todo Merge cria um novo commit. O Merge deixa um histórico "claro" das mudanças realizadas, porém, se não souber usar deixará um histórico de commits bagunçado.

### Rebase?

É o nosso organizador de históricos dos commits. O comando "git rebase branchA" reescreverá o histórico de commits da minha branch atual para a "branchA" e deixará as duas branchs com o histórico semelhante. É possivel tambem, alterar a ordem dos commits, o que pode ser feito de forma simplificada quando acionamos o parametro "-i" para tornar o rebase iterativo, ou seja, abrirá uma editor para que possamos "amassar", "deletar" e alterar a ordem dos commits. 

Não é aconselhavel utiliza-lo em branchs compartilhadas, porque pode gerar confusões.

Ele pode evitar merges desnecessários e ele deixará o histórico de commits mais limpo.

### Merge vs Rebase

Como o Merge mantém o histórico original é aconselhavél utiliza-lo quando queremos evidenciar que houve uma mesclagem. Quando visualizarmos em um "git log --graph" por exemplo, branchs diferentes terão linhas diferentes e se unirão de forma explicita no merge. Merge também é aconselhavel no contexto de branchs públicas e compartilhadas pois evitam problemas referentes a reescrita do histórico.

Já o Rebase é aconselhavel para branchs de desenvolvimento, já que pode organizar e limpar o historico de commits desta respectiva branch. 

### Configurar o pull.rebase false ou true, para um projeto o atual que estou atuando?

O pull irá puxar a versão que está no remoto. Então a questão aqui é, se você deseja que os commits do remoto sobescrevem os commits do local e coloque seus commits locais no topo dessas alterações faça o --rebase. Ele não gera nenhum novo commit e deixará o histórico mais limpo.

Se você deseja preservar a estrutura original do histórico, use o --no-rebase.

Eu decido usar o --rebase para projetos que tem mais ramificações e usar o --no-rebase para projetos mais simples.

