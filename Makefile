build:
	geth init --datadir ./blockchain genesis.json
	geth --datadir ./blockchain --syncmode "full" --port 30304 --http --http.addr "localhost" --http.port 8552 --http.api "txpool,eth,net,web3,personal,admin,miner" --http.corsdomain="*" --networkid 13 --allow-insecure-unlock  --ipcdisable --authrpc.port 8553 console
	
run:
	go run main.go