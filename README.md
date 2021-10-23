# AskTick
Ask for comparison and evaluate every vote tick

## Run locally
```sh
docker-compose up --build
```


### postgres commandline

#### connect to postgres
```sh
psql -U <username>
```

#### show all databases
```sh
\l
```

#### show all databases
```sh
GRANT CONNECT ON DATABASE <database> TO <username>;
```


#### connect to DB
```sh
\c <database>
```

#### show tables
```sh
\dt
```