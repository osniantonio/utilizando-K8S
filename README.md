## Utilizando K8S

## 1) Servidor Web - Nginx
Utilize a imagem base do Nginx Alpine
Disponibilize 3 réplicas
Quando alguém acessar o IP externo do LoadBalancer do serviço criado, ou em caso de utilização do Minikube usando "minikube service nome-do-servico", deve ser exibido no browser: Code.education Rocks.<br />
URL: <a href="http://34.71.234.201/">http://34.71.234.201/</a>

## 2) Configuração do MySQL
Faça o processo de configuração de um servidor de banco de dados MySQL
Utilize secret em conjunto com as variáveis de ambiente
Utilize disco persistente para gravar as informações dos dados.
Obs.: Execute o seguinte comando para criar a secret: kubectl create secret generic mysql-pass --from-literal=password='SENHA'

## 3) Desafio Go!
Crie um aplicativo Go que disponibilize um servidor web na porta 8000 que quando acessado seja exibido em HTML (em negrito) Code.education Rocks!
A exibição dessa string deve ser baseada no retorno de uma função chamada "greeting". Essa função receberá a string como parâmetro e a retornará entre as tags <b></b>.
Como ótimo desenvolvedor(a), você deverá criar o teste dessa função.
Ative o processo de CI no Google Cloud Build para garantir que a cada PR criada faça com que os testes sejam executados.
Gere a imagem desse aplicativo de forma otimizada e publique-a no Docker Hub
Utilizando o Kubernetes, disponibilize o serviço do tipo Load Balancer que quando acessado pelo browser acesse a aplicação criada em Go.
URL: <a href="https://hub.docker.com/r/osniantonio/utilizando-k8s">https://hub.docker.com/r/osniantonio/utilizando-k8s</a>