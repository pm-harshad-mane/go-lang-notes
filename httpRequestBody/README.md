# Problem:
Once you read the httpRequest.Body it becomes empty so you get a blank value when you access it again.

# Setup
How to run? ` go run main.go `

# Details
- When you call the old API with following curl call, we get empty value on second attemp on console
```
curl --location --request POST 'http://localhost:99/wakandaOld' \
	--header 'Content-Type: text/plain' \
	--data-raw '{ "name": "Harshad", "age": 34 }'
```
- We have modified the similar code in another API where this issue is not observed
```
curl --location --request POST 'http://localhost:99/wakandaNew' \
	--header 'Content-Type: text/plain' \
	--data-raw '{ "name": "Harshad", "age": 34 }'
```

