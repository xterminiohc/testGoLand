# testGoLand


Dependencies: install Go and libraries for sqlite and gorm

To start run main



Test de RestFul services:

Get Token
  curl --location --request GET 'http://localhost:9091/token'


Get Greeting
  curl --location --request GET 'http://localhost:9091/greet' \ --header 'Token: {token} '



  
