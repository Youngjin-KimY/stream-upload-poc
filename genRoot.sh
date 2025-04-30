openssl genrsa -out rootCA.key 4096

openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 3650 \
  -out rootCA.pem -config localhost-openssl.cnf -extensions v3_ca