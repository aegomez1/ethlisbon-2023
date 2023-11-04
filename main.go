package main

import (
	//"encoding/json"

	"encoding/json"
	ipfsConn "ethlisbon-2023/ipfs"
	parsing "ethlisbon-2023/parsing"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("audit-request/{address}", getAudit).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func getAudit(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	contractAddress := (params["address"])
	audit, _ := recursive(contractAddress)
	res := map[string]interface{}{"data": audit}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(res); err != nil {
		panic(err)
	}
}

func recursive(contractAddress string) (audit map[string]interface{}, frontendLink string) {
	cid := web3Call(contractAddress)          // Smart address sent to web3, returns cid
	contractDetails := ipfsConn.IpfsCall(cid) // CID send to ipfs, pulls down contractDetails
	cidTarget := "glossifi-audit"
	auditCid := parsing.JsonParsing(cidTarget, contractDetails) // Finds cid of audit
	frontendTarget := "glossifi-frontend"
	frontendLink = parsing.JsonParsing(frontendTarget, contractDetails) // Finds front-end link, how to display on front-end**
	audit = ipfsConn.IpfsCall(auditCid)                                 // Cid of audit is sent to ipfs, pulls down audit

	return audit, frontendLink
}

// Makes a call to smart contract index to find CID based on address
// Sends address, Returns CID
// Need to have a parser or send it thru the json parser
func web3Call(contractAddress string) (cid string) {

	return cid
}

// Parse thru contractDetails.json to find Audit CID and Vue App link
// Makes second call to ipfs(maybe have as one ipfs call function)
// Sends audit CID, returns Audit
// Audit is displayed on front end with link
