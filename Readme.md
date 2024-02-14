# Lorcana tracker

Tool to track my lorcana collection

## Implemented queries with HTTP pie

Get all cards

    http 'localhost:8080/cards'

    or

    http 'localhost:8080/cards?collection=all'

Get all owned cards

    http 'localhost:8080/cards?collection=owned'

Get all missing cards

    http 'localhost:8080/cards?collection=missing'

get all cards from the ROF set

    http 'localhost:8080/cards?set=ROF'

get all cards from the ROF and TFC set

    http 'localhost:8080/cards?set=ROF&set=TFC'

get all cards from all sets

    http 'localhost:8080/cards'

get all sets in the database

    http 'localhost:8080/sets'