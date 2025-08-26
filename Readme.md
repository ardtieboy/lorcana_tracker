# Lorcana tracker

Tool to track my lorcana collection

## Start the service

    go run main.go

or

    go run main.go --initDB true 

to initialize the database

## Implemented queries with HTTP pie

Get all cards

    http 'localhost:8080/card'

    or

    http 'localhost:8080/card/TFC-1'

Get stats of an (owned) card

     http 'localhost:8080/card_in_collection/TFC-1'

Update the status of an owned card

    echo '{"card_id":"TFC-191", "owned_normal_copies": 9, "owned_foil_copies":2, "whish_list": true}' | http PUT http://localhost:8080/card_in_collection

## Example HTTP requests

### Count the number of cards in a set

    http GET http://localhost:8080/card | jq '[.[] | select(.set_id == "TFC")] | length'

### Check if a specific card is in your collection

    http GET http://localhost:8080/card_in_collection/TFC-191          

### Update the card count of a specific card

    echo '{"card_id": "TFC-191", "owned_foil_copies": 3, "owned_normal_copies": 0, "whish_list": true}' | http PUT http://localhost:8080/card_in_collection

## API Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) to generate OpenAPI documentation from the Go code.

### Generating Documentation

The documentation is generated based on the annotations in `controller/controller.go`. To update the documentation after making changes to the comments, run the following command:

```bash
swag init
```

This will update the files in the `docs/` directory.

### Viewing Documentation

When the application is running, the interactive Swagger UI can be accessed at:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
