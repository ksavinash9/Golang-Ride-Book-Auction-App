A Simple Ride Booking Auction Backend Layer
===========================================

## Challenges

- [x] booking of rides
- [x] auction price start with lowest price
- [x] linear progression of price with auction duration and steps
- [x] configurable peroid of time and steps
- [x] implemented auction pool, lowest offering price and configuring auction
- [x] persistence of data in postgres
- [x] cron jobs for handling background jobs
- [x] unit test for APIs

## Usage

### Installation
-----------
```sh
$ go install
```

### Running Server
-----------
```sh
$ go run main.go
```

### Running Cron Jobs
-----------
```sh
$ go run cronjobs.go
```

### Testing
-----------
```sh
$ go test test/
```

![testing-image](https://raw.githubusercontent.com/swarnavinash/Golang-Ride-Book-Auction-App/master/Testing.png)


### EndPoints
-------------

GET - http://localhost:9000/ride
```json
[
  {
    "UserID": 101,
    "DriverID": 1001,
    "RidePrice": "1024",
    "RideTime": "2016-07-21T14:30:25.077289Z",
    "AuctionUUID": "e399e1011b836296342e04c3d985a21b",
    "AuctionSteps": 0,
    "AuctionDuration": 0
  },
  {
    "UserID": 202,
    "DriverID": 2002,
    "RidePrice": "2048",
    "RideTime": "2016-07-21T14:31:53.625309Z",
    "AuctionUUID": "fb203a2bcac46e9c190a7effbca1f9a1",
    "AuctionSteps": 0,
    "AuctionDuration": 0
  }
]
```
GET - http://localhost:9000/ride?id=1
```json
{
  "UserID": 101,
  "DriverID": 1001,
  "RidePrice": "1024",
  "RideTime": "2016-07-21T14:30:25.077289Z",
  "AuctionUUID": "e399e1011b836296342e04c3d985a21b",
  "AuctionSteps": 0,
  "AuctionDuration": 0
}
```
POST - http://localhost:9000/ride
```json
Request Body:
{
    "UserID" : 101,
    "DriverID": 1001,
    "RidePrice": "1024"
}
```
```json
Response:
{
  "AuctionUUID": "e399e1011b836296342e04c3d985a21b",
  "AuctionStartTime": "2016-07-21T14:30:25.125096105+08:00",
  "AuctionEndTime": "2016-07-21T14:30:25.12509649+08:00",
  "AuctionInitialPrice": "103",
  "AuctionFinalPrice": "1024"
}
```
PUT - http://localhost:9000/ride?id=:id
```json
Request Body
{
    "UserID" : 101,
    "DriverID": 10001,
    "RidePrice": "1024"
}
```
```json
Response 
Rows Updated  1
```
DELETE - http://localhost:9000/ride?id=:id

```json
Response 
Rows Deleted  1
```
GET - http://localhost:9000/auction

```json
Response
[
  {
    "AuctionUUID": "e399e1011b836296342e04c3d985a21b",
    "AuctionMaxPrice": "1024",
    "AuctionStartTime": "2016-07-21T14:30:25.107045Z",
    "AuctionInitialPrice": "103",
    "AuctionSetPrice": "103",
    "AuctionFinalPrice": "",
    "AuctionSteps": 1,
    "AuctionDuration": 60
  },
  {
    "AuctionUUID": "fb203a2bcac46e9c190a7effbca1f9a1",
    "AuctionMaxPrice": "2048",
    "AuctionStartTime": "2016-07-21T14:31:53.637673Z",
    "AuctionInitialPrice": "205",
    "AuctionSetPrice": "205",
    "AuctionFinalPrice": "",
    "AuctionSteps": 1,
    "AuctionDuration": 60
  }
]
```
GET - http://localhost:9000/auction?id=:id
```json
{
	"AuctionUUID": "fb203a2bcac46e9c190a7effbca1f9a1",
	"AuctionMaxPrice": "2048",
	"AuctionStartTime": "2016-07-21T14:31:53.637673Z",
	"AuctionInitialPrice": "205",
	"AuctionSetPrice": "205",
	"AuctionFinalPrice": "",
	"AuctionSteps": 1,
	"AuctionDuration": 60
}
```
POST http://localhost:9000/auction
```json
{
	"AuctionUUID": "fb203a2bcac46e9c190a7effbca1f9a1",
	"AuctionMaxPrice": "2048",
	"AuctionSteps": 1,
	"AuctionDuration": 60
}
```
```json
Response
Auction Inserted ID 949
```
PUT - http://localhost:9000/auction?id=:id

```json
Request Body
{
	"AuctionUUID": "fb203a2bcac46e9c190a7effbca1f9a1",
	"AuctionMaxPrice": "2048",
	"AuctionSteps": 1,
	"AuctionDuration": 60
}
```
```json
Response 
Rows Updated  1
```
DELETE - http://localhost:9000/ride?id=:id

```json
Response 
Rows Deleted  1
```
GET - http://localhost:9000/auction/getprice

```json
Response 
{
  "AuctionUUID": "f1349385af4ca4612c16557bc65a9535",
  "AuctionStepPrice": "93",
  "AuctionStepExpiry": "2016-07-21T14:46:38.446518804+08:00"
}
```
PUT - http://localhost:9000/auction/setPrice?id=:id

```json
Response
Rows Updated 1
```