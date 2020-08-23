package main

import (
	"blockchain/utils"
	"blockchain/wallet"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
)

const templateDirectory = "wallet_server/templates"

type WalletServer struct {
	domain   string
	port     uint16
	rootPath string
	gateway  string
}

func NewWalletServer(domain string, port uint16, rootPath string, gateway string) *WalletServer {
	return &WalletServer{domain: domain, port: port, rootPath: rootPath, gateway: gateway}
}

func (s *WalletServer) Port() uint16 {
	return s.port
}

func (s *WalletServer) Gateway() string {
	return s.gateway
}

func (s *WalletServer) Run() {
	// Publish resources
	//fileServer := http.FileServer(http.Dir("./resources"))
	//strip := http.StripPrefix("/resources/", fileServer)
	//http.Handle(path.Join(s.rootPath, "resources"), strip)

	http.HandleFunc(path.Join(s.rootPath, "index.html"), s.index)
	http.HandleFunc(path.Join(s.rootPath, "wallet"), s.wallet)
	portStr := strconv.Itoa(int(s.port))
	utils.Logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.domain, portStr), nil))
}

func (s *WalletServer) index(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(templateDirectory, "index.html"))
		_ = t.Execute(writer, "")
	default:
		utils.Logger.Error("Invalid HTTP method")
	}
}

func (s *WalletServer) wallet(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		writer.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		_, _ = io.WriteString(writer, string(m[:]))
	default:
		writer.WriteHeader(http.StatusBadRequest)
		utils.Logger.Error("Invalid HTTP method")
	}
}
