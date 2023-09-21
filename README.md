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

### API Endpoint

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