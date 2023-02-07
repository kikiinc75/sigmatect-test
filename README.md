# sigmatech-test
Create restful-api for solution on PT xyz using golang 

## Set up project
1. clone this repository to your computer
   - git clone https://github.com/kikiinc75/sigmatect-test
2. restore sql file and setup your own environment on database\mysql.go
3. Running the API with go run main.go
   

## Unit testing
For unit testing, I tested starting from fetching the url with the method until the return the JSON results, matched using the assert method.
Because of I'm testing the gin.Engine, they give me the GIN-DEBUG that makes the result a bit messy in my oppinion, but we still got the PASS or FAIL result.
For the testing path, I put it on /test folder, so to check it you have to type cd test and run go test -v, or just test it one by one

## Docker
For building the image type on cmd 
`docker build -t sigmatech-test:1.0 .`
For the docker conatiner, type on cmd 
`docker compose up`
