# jumia_test_go


## Comments

I made the two services, calling each other through http service. Maybe if I chose to use Redis or another local service could be better.

If these services were in production, the best answer is to use AWS SNS or AWS SQS to avoid this delay.

Unfortunately, I couldn’t solve the giant JSON return from Service B. Maybe I could use pagination, but I don't know if it would be allowed.

The function GetCountryByPhone (Service A) was created using the simple switch case because the given regex doesn't match with the phone list in CSV file

## Usage

```
docker-compose build  
docker-compose run
```

## Services
### Service A
| URL  |  METHOD | BODY  | FIELD / DESC  |
|---|---|---|---|
|  http://localhost:8062/upload | POST  | form-data  | file : with csv file  |

			

### Service B
| URL                               | METHOD | BODY  | FIELD / DESC                       |
|-----------------------------------|--------|-------|------------------------------------|
| http://localhost:8063/order/new   | POST   | raw   | json with order                    |
| http://localhost:8063/cargo/daily | GET    | empty | get json with daily cargo manifest |


## Tests
###service_a/use_cases
#### Countries
```go test -v countries_test.go countries.go```
#### Processor
```go test -v processor_test.go processor.go csv_importer.go countries.go```
