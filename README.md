# blockchain
Based off of this tutorial:https://jeiwan.net/posts/building-blockchain-in-go-part-7/

# unit tests
from root:
go test -v ./...

# CLI manual testing
cd src
go build -o blockchain_go

3 terminals
T1
    export NODE_ID=3000
    ./blockchain_go createwallet (store wallet address somewhere "T1_WALLET_1")
    ./blockchain_go createblockchain -address (T1_WALLET_1)
    cp blockchain_3000.db blockchain_genesis.db 
T2
    export NODE_ID=3001
    ./blockchain_go createwallet (store wallet address somewhere "T2_WALLET_1")
    ./blockchain_go createwallet (store wallet address somewhere "T2_WALLET_2")
    ./blockchain_go createwallet (store wallet address somewhere "T2_WALLET_3")
T1 
    ./blockchain_go send -from T1_WALLET_1 -to T2_WALLET_1 -amount 10 -mine
    ./blockchain_go send -from T1_WALLET_1 -to T2_WALLET_2 -amount 10 -mine
    ./blockchain_go startnode
T2
    cp blockchain_genesis.db blockchain_3001.db
    ./blockchain_go startnode
    CTRL+C
    ./blockchain_go getbalance -address T2_WALLET_1 (should be 10)
    ./blockchain_go getbalance -address T2_WALLET_2 (should be 10)
T3
    export NODE_ID=3002
    cp blockchain_genesis.db blockchain_3002.db
    ./blockchain_go createwallet (store wallet address somewhere "T3_WALLET_1")
    ./blockchain_go startnode -miner T3_WALLET_1
T2
    ./blockchain_go send -from T2_WALLET_1 to T2_WALLET_3 -amount 1
    CTRL+C
    ./blockchain_go startnode (to download new blocks)
    ./blockchain_go getbalance -address T2_WALLET_1 (should be 9)
    ./blockchain_go getbalance -address T2_WALLET_2 (should be 10)
    ./blockchain_go getbalance -address T2_WALLET_3 (should be 1)


# todo
unit tests 
metrics
docker container
metric and dashboard containers
