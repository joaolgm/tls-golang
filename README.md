# tls-golang
Exemplo de uso simplificado de TLS. O cliente já possui o certificado do server e utiliza ele para fazer chamada HTTP.

## Gerando certificado
Altere a ultima linha do arquivo __san.cnf__ com o DNS desejado
```
DNS.1 = <seu_DNS>
```

A partir daqui os arquivos __server.csr__, __server.key__, __server.crt__, que estão como exemplo, serão sobrescritos.

Rode o __opensssl genrsa__ a seguir para criar uma chave para o server:
```
openssl genrsa -out server.key 2048
```

Rode o __openssl req__ para gerar uma requisição de criação de certificado. Lembre novamente de alterar o caminho com seu DNS:
```
openssl req -new -key server.key -out server.csr -subj "/CN=<seu_DNS>" -config server.cnf
```

Finalmente rode o __openssl x509__ para gerar o certificado:
```
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt -extensions v3_req -extfile server.cnf
```

Com o certificado gerado, faça uma copia dele para o client(pulando alguns passos do TLS).
