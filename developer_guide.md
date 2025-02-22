SIGN UP 
```aiignore
 curl -X POST http://localhost:8080/signup   -H "Content-Type: application/json"   -d '{"username": "your_username", "password": "your_password" , "name":"rm", "email": "ritu251297@gmail.com"}'
```

LOGIN
```aiignore
curl -X POST http://localhost:8080/login   -H "Content-Type: application/json"   -d '{"email": "ritu251297@gmail.com", "password": "your_password"}
```
response
```aiignore
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZjAzNTYyMjEtMjAzMC00ZjllLWE5NzMtOGM1MGM2MzE1YWMzIiwiZXhwIjoxNzQwMjgxNTM1fQ.4t14l_UeeVbNwI5yWKdJ7aHt5SvukS4m6PFaJT7bhKo","userId":"f0356221-2030-4f9e-a973-8c50c6315ac3"}
```


GET User profiles 

```aiignore
 curl -X GET "http://localhost:8080/user/profile/f0356221-2030-4f9e-a973-8c50c6315ac3"      -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZjAzNTYyMjEtMjAzMC00ZjllLWE5NzMtOGM1MGM2MzE1YWMzIiwiZXhwIjoxNzQwMjgxNTM1fQ.4t14l_UeeVbNwI5yWKdJ7aHt5SvukS4m6PFaJT7bhKo"      -H "Content-Type: application/json"

```
response
```aiignore
{"id":"f0356221-2030-4f9e-a973-8c50c6315ac3","name":"rm","email":"ritu251297@gmail.com","password":""}
```