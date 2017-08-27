db:
	docker run --rm --name tidb-server -d -p 4000:4000 pingcap/tidb:latest
db-connect:
	mysql -h 127.0.0.1 -P 4000 -u root -D test --prompt="tdib> "
db-seed:
	go run main.go seed
