# GO-PSQL (name subject to change)

### How to run (in docker) assuming you're in `jigdra/go-psql` directory
- `docker-compose up -d postgres go-psql`
- `cd sql/ && ./init.sh`
- ??????
- profit

### How to run outside of docker
###### not tested, don't care
- start your postgres server
- `psql -f init.sql`
- `sudo -u postgres psql`
  ###### in shell
> `psql> CREATE USER admin WITH PASSWORD 'password'`;
> 
> `psql> GRANT ALL PRIVILEGES ON jigdra.* TO admin;`
> 
> `psql> \q`
- should work

