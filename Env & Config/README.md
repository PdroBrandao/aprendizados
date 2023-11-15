## .env & config/app

### Porque estou comparando aparentemente coisas que não devem ser comparadas?

Precisava trazer uma url pra dentro de um template de email, e esta não deveria ser estática. 
Para trazer uma váriavel "global" me veio a mente o arquivo .env, porém descobri que preciso instância-lo tbm nas config do app.


---
### Pra que serve o .env?

Antes de falar no arquivo .env, devemos falar das variaveis de ambiente. São váriaveis do SO ou do definidas no código, no caso no arquivo .env.
Essas váriaveis são responsáveis por informações sobre o ambiente no quão o programa está sendo executado, como chaves de apis, senhas, urls.

Normalmente ele não é versionado.

---
### Pra que serve os configs do app?

São configs que normalmente são versionadas. Lidam com configurações globais de app como: Configurações de log, portas http, configurações de cache, filas, mensagens, etc...

---
### Qual a diferença entre eles? 

O .env é utilizado para armazenar váriaveis de ambiente locais e indicado para informações mais sensíveis, já o config é utilizado para 

---
### Devo referênciar uma variável de qual dos dois?

Como a leitura das váriaveis do .env são feitas diretamente pelo Sistema Operacional ou por alguma lib especifica, no caso do go a lib "os". 
Acredito que acessar uma variável via config é uma melhor prática, pra não precisar usar mais libs que o necessário. 



