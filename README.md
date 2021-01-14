# Especificações

- Docker
- GO 1.14
- MySQL 8

#### Baixar o projeto utilizando o git:
 ``` 
    git clone https://github.com/Phaael/phaael-playground.git
 ```

 1. Executar o comando para fazer o build e aguardar a execução:
    ``` 
       docker-compose build 
    ```

 2. Depois de executado o comando anterior, executar o seguinte:
    ``` 
       docker-compose up -d 
    ```
 3. Para verificar se o processo ocorreu tudo certo, e se a aplicação subiu, rodar o seguindo comando:
    ```
       docker-compose ps
    ```

1. Exemplos de chamadas a api:
   - criando uma conta:
    ```
   curl -X POST http://localhost:8080/accounts -d '{"document_number": "12345678900"}'
    ```

   - Criando uma transação:
   ```
   curl -X POST http://localhost:8080/transactions -d '{"account_id": 4,"operation_type_id": 1,"amount": 123.45}'
   ```

   - Consultando uma conta:
   ```
    curl -X GET http://localhost:8080/accounts/1234
   ```
1. Rodar os testes:
    ```
    go test -coverprofile=coverage.out ./...

   ```


- Obs: Após o comando  docker-compose up -d é necessário em média uns 30s para a aplicação subir e estar disponivel.

   
 

   

    


   
  