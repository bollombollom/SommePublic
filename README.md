# SommePublic

Un module public, avec ses tests. 
Il vit dans un dépôt github public. 


## déroulé
```
 2009  cd SommePublic/
 2010  go mod init
 2012  go mod init github.com/bollombollom/SommePublic
 2013  touch somme.go
 2014  touch somme_test.go
 2015  code .
 2017  go test -v
 2020  git add .
 2021  git commit -m "Test module public"
 2022  git push
 2023  git tag v1.0.0
 2025  git push origin v1.0.0
```

## tester
```
go test -v
``` 

## intégrer
