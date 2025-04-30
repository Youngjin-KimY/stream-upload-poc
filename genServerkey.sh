openssl genrsa -out localhost.key 2048

openssl req -new -key localhost.key -out localhost.csr -subj "/CN=localhost"