## Adding a new exoplanet.

curl --location 'http://{host}/exoplanets' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "Test Planet",
    "Description": "test descpricption",
    "Type": "Terrestrial",
    "Mass": 1.0,
    "Radius": 1.0,
    "Distance": 100
}'

## Listing all exoplanets.

curl --location 'http://{host}/exoplanets'

## Retrieving an exoplanet by its ID.

curl --location 'http://{host}/exoplanets/{id}'

## Updating an exoplanet by its ID.

curl --location --request PUT 'http://{host}/exoplanets/{id}' \
--header 'Content-Type: application/json' \
--data '{
    "Distance": 42
}'

## Deleting an exoplanet by its ID.

curl --location --request DELETE 'http://{host}/exoplanets/812050'

## Estimating the fuel required for a voyage based on exoplanet data.

curl --location 'http://{host}/exoplanets/{id}/fuel-estimation?crew={crew-capacity}'