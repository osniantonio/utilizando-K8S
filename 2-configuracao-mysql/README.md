## Utilizando K8S

## 2) Configuração do MySQL
Faça o processo de configuração de um servidor de banco de dados MySQL
Utilize secret em conjunto com as variáveis de ambiente
Utilize disco persistente para gravar as informações dos dados.
Obs.: Execute o seguinte comando para criar a secret: kubectl create secret generic mysql-pass --from-literal=password='SENHA'