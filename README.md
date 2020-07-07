
### Usage

##### 1. Pull
```bash
git pull git@github.com:HayataSato/go_psql_test.git
```

##### 2. Create DB
```bash
sudo /etc/init.d/postgresql start
qsql -U DB_USER -p
create database DB_NAME
```

##### 3. Modify .env
* Rename ".env.example" file to ".env"
* Add environment-variables to .env

##### 4. build & run
```bash
# build  (※ $GO111MODULE = on)
go build

# run
psql_test
```

##### 5. Check DB & index-page
```sql
select * from users;
```
http://localhost:8080/


