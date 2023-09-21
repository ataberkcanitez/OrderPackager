# Order Packager
## RE Partners Coding challenge

This application calculates the number of packs needed to fulfill customer orders based on specified pack sizes


## How to Run
- Build the Docker image:
 ```
$ docker build -t order-packager .
```
- Run the Docker Container
```
$ docker run -p 8080:8080 order-packager
```

### API Endpoints

---
- `POST` - `/calculate-packs`: Calculate the number of packs needed for an order
#### Request Example
Body:
```json
{
  "itemsToShip": 12001
}
```
Response:
```json
{
  "packCounts": [
    {
      "Pack": {
        "ID": 5,
        "Size": 5000
      },
      "Amount": 2
    },
    {
      "Pack": {
        "ID": 4,
        "Size": 2000
      },
      "Amount": 1
    }
  ]
}
```
---
- `POST` - `/packs`: Create a new pack with size
#### Request Example
Body:
```json
{
  "id": "10",
  "amount": 15000
}
```
if id is not used before:
```json
{
    "pack": {
        "ID": "10",
        "Size": 15000
    },
    "success": true
}
```
if id is used before:
```json
{
  "details": "Pack already exists",
  "error": "Internal Server Error",
  "success": false
}
```
---
- GET - `packs`: Get all packs
#### Request Example
Response:
```json
{
    "packs": [
        {
            "ID": "2",
            "Size": 500
        },
        {
            "ID": "3",
            "Size": 1000
        },
        {
            "ID": "4",
            "Size": 2000
        },
        {
            "ID": "5",
            "Size": 5000
        },
        {
            "ID": "10",
            "Size": 15000
        },
        {
            "ID": "1",
            "Size": 250
        }
    ],
    "success": true
}
```
---
- GET - `/packs/:id`: Get a pack by id
Reponse:
```json
  {
    "pack": {
        "ID": "2",
        "Size": 500
    },
    "success": true
}
```
---
To simplify the api calls, I provided a postman collection, it's under the /docs/postman folder
please use it to send requests.
Also please make sure that there is no other process runs on 8080 port. if so, it will cause a conflict.
You can adjust your host as `localhost:<port:8080>` you can change the port in the Dockerfile as well if you want.