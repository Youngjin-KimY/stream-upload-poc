openssl x509 -req -in localhost.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial \
  -out localhost.crt -days 825 -sha256 -extfile localhost-openssl.cnf -extensions v3_ca