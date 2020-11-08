## To create your PostgresQL DB follow the following steps ##

Access Postgres shell like so:
```
sudo -u postgres psql
```

Create Postgres DB like so:
```
create database studs;
create user studUser with encrypted password 'studPassword';
grant all privileges on database studs to studUser;
```

Fill out the following .env to conform to your database
```
DB_HOST='localhost'
PGHOST='localhost'

DB_PORT='5432'
PGPORT='5432'

DB_NAME='studs'
PGDATABASE='studs'

DB_USER='studuser'
PGUSER='studuser'

DB_PASSWORD='studPassword'
PGPASSWORD='studPassword'

```

Here's an .env script. Named `.envrc`
```
#!/usr/bin/env zsh

for x in `cat .env`; do
  todo=$(echo "export $x")
  echo "$todo"
  eval "$todo"
done

```