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

