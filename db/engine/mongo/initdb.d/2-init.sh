
mongoimport -u user -p user --db app --collection bans --file /docker-entrypoint-initdb.d/bans.json --jsonArray
# mongoimport -u user -p user --db app --collection patties --file /docker-entrypoint-initdb.d/patties.json --jsonArray
# mongoimport -u user -p user --db app --collection vegetables --file /docker-entrypoint-initdb.d/vegetables.json --jsonArray
# mongoimport -u user -p user --db app --collection ingredients --file /docker-entrypoint-initdb.d/ingredients.json --jsonArray
