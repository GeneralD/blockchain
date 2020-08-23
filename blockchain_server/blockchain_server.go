package main

import (
	"blockchain/blockchain"
	"blockchain/utils"
	"blockchain/wallet"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

type BlockchainServer struct {
	domain   string
	port     uint16
	rootPath string
}

func NewBlockchainServer(domain string, port uint16, rootPath string) *BlockchainServer {
	return &BlockchainServer{domain: domain, port: port, rootPath: rootPath}
}

var (
	cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)
)

func (s *BlockchainServer) Port() uint16 {
	return s.port
}

func (s *BlockchainServer) Run() {
	http.HandleFunc(path.Join(s.rootPath, "chain"), s.GetChain)
	portStr := strconv.Itoa(int(s.port))
	utils.Logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.domain, portStr), nil))
}

func (s *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minerWallet.Address(), s.port)
		cache["blockchain"] = bc
		utils.Logger.Debugf("private_key %v", minerWallet.PrivateKeyString())
		utils.Logger.Debugf("public_key  %v", minerWallet.PublicKeyString())
		utils.Logger.Debugf("address_key %v", minerWallet.Address())
	}
	return bc
}

func (s *BlockchainServer) GetChain(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		bc := s.GetBlockchain()
		m, _ := bc.MarshalJSON()
		_, _ = io.WriteString(writer, string(m[:]))
	default:
		utils.Logger.Error("Invalid HTTP method")
	}
}
