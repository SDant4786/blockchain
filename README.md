# blockchain
Based off of this tutorial:https://jeiwan.net/posts/building-blockchain-in-go-part-7/



Export NODE_ID=3000
Go build .
./src createwallet
    15pD2r9KvFtbRQiTd6fUK4WVr13AXno1pk
./src createblockchain -address 15pD2r9KvFtbRQiTd6fUK4WVr13AXno1pk
cp blockchain_3000.db blockchain_genesis.db

