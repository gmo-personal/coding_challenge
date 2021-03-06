# Prototype Admin Dashboard coding_challenge
Prototype admin dashboard that collects and displays usage metrics submitted by a fakeIOT client.

### Installation:

```
go get github.com/gmo-personal/coding_challenge
cd ~/go/src/github.com/gmo-personal/coding_challenge
```

### Setup database: 

Assuming mysql is installed, 

Use the following command in mysql to create the schema.

`create schema <schema name>`

### Setup server:

Run the following command to spin up the docker container. The site can be accessed on https://localhost:443

replace \<username> and \<password> with local mysql database username and password respectively.

`srv_port=443 srv_cert_path=certs/server-cert.pem srv_key_path=certs/server-key.pem db_user=<username> db_pass=<password> db_schema=<schema name> db_host=host.docker.internal db_port=3306 docker-compose up --build`

### Run tests:

Use the following command to run the tests. 

`docker exec -it gmo-fullstack bash`

`./test.sh`

### Site:

Login to dashboard: `https://localhost:443` by default.

The default login for the base user `testacct-0000-0000-0000-000000000000`:

username: `t@gmail.com`

password: `t`

### Running with FakeIOT

The following is the hard coded bearer token. It is a randomly generated
len 32 string.

`0b2a4a30602b7f6d1079593cabfdb9386eff2041b863a9c678915158eb60fdac`

Sample full command:

`fakeiot --token=0b2a4a30602b7f6d1079593cabfdb9386eff2041b863a9c678915158eb60fdac --url="https://127.0.0.1:443" --ca-cert=./fixtures/ca-cert.pem run --period=9s --freq=0.1s --users=90 --account-id=testacct-0000-0000-0000-000000000000`