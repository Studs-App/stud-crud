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