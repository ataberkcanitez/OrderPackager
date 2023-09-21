# Order Packager
## RE Partners Coding challenge

This application calculates the number of packs needed to fulfill customer orders based on specified pack sizes


## How to Run
- Build the Docker image:
 ```
$ docker build -t pack-calculator .
```
- Run the Docker Container
```
$ docker run -p 8080:8080 pack-calculator
```

### API Endpoint

- `POST` - `/calculate-packs`: Calculate the number of packs needed for an order
#### Request Example
Body:
```json
{
  "items_to_ship": 12001
}
```
Response:
```json
{
  "pack_counts": {
    "5000": 2,
    "1000": 1,
    "250": 1
  }
}
```