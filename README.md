# stream-upload-poc


The script is for certificate of local dev and test. It is not for prd use
---
Reason Why I should set 'ssl'?<br>
`Web Stream API` is supported by http/2 perfectly, and when http/2 spec like `Web Stream API` is used , Browser forces http/s to use ssl
---

how to run this project


```bash
# in the root location
tsc # build ts file, move into static folder
go run main.go # run backend
```
