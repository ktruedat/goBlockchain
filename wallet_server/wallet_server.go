package main

import (
	"encoding/json"
	"github.com/ktruedat/goBlockchain/utils"
	"github.com/ktruedat/goBlockchain/wallet"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

const templateDir = "wallet_server/templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port: port, gateway: gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(templateDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Println("Invalid HTTP method")
	}
}

func (ws *WalletServer) CreateWallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		createdWallet := wallet.NewWallet()
		m, _ := createdWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid HTTP method")
	}

}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		if err := decoder.Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("JSON decoding unsuccessful: %v", err)
			return
		}

		if !t.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("Bad Request")
			return
		}
		publicKey := utils.PublicKeyFromString(t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(t.Value, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Parse error")
			return
		}
		value32 := float32(value)
		w.Header().Add("Content-Type", "application/json")

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid HTTP method")
		return
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.CreateWallet)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
