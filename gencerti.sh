openssl pkcs12 -export \
  -out localhost.pfx \
  -inkey localhost.key \
  -in localhost.crt \
  -certfile rootCA.pem