#Fibonacci Web API

---

Did you know that the Fibonacci Sequence can be found in nature? 
Flower petals grow in a manner that is consistent with the Fibonacci Sequence. This pattern can be seen in lilies (three petals), buttercups (five petals), delphiniums (8 petals) and many more!
<p align="right"><img src="https://i0.wp.com/eminimind.com/wp-content/uploads/2018/06/Fibonacci-Nature.jpg?fit=1024%2C768&ssl=1" width="250" height="180"/></p>

---

##Installation
Clone the Fibonacci Web API repository to your local system:
```sh 
git clone https://github.com/ecklenke9/fibonacci-web-api.git
```


##Running the Application
There are two ways to run this application:

Docker:
* Run the following cmd at the root level of the application:
```sh 
make docker
```

Locally:
* Run the following cmd at the root level of the application:
```sh 
make local
```

##Calling the Endpoints
The following Fibonacci related data can be retrieved from the Fibonacci Web API endpoints: 

Fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144)
```http request
GET http://localhost:8080/api/fibonacci/{ordinal}
```
Output:
```json
{
    "fibonacciNumber": 2
}
```
Fetch all of the Fibonacci numbers in data store
```http request
GET http://localhost:8080/api/fibonacci/all
```
Output: 
```json
{
    "allFibonacciResults": [
        {
            "ordinal": 0,
            "fibNum": 1
        },
        {
            "ordinal": 1,
            "fibNum": 1
        },
        {
            "ordinal": 2,
            "fibNum": 1
        },
        {
            "ordinal": 3,
            "fibNum": 2
        }
    ]
}
```
Fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120)
```http request
GET http://localhost:8080/api/fibonacci/memoizedResults/{value}
```
Output:
```json
{
    "memoizedResults": 4
}
```
Clear the data store
```http request
DELETE http://localhost:8080/api/fibonacci/clear
```
Output:
```json
{
    "message": "Database cleared"
}
```

---
Languages and Tools Used for this Application:
<p align="left"> <a href="https://www.docker.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.postgresql.org" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a> </p>
