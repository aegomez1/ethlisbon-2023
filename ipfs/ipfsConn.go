package ipfsConn

import (
	"bytes"
	"encoding/json"
	"ethlisbon-2023/configs"
	"fmt"

	ipfs "github.com/ipfs/go-ipfs-api"
)

// Connects to ipfs, send json payload, and return cid(hash)
// Makes a call to IPFS
// Sends CID, Returns ipfsBody
func IpfsCall(cid string) (ipfsBody map[string]interface{}) {
	// Where local node is running - Connect to the IPFS API
	IPFSConn := configs.GoDotEnvVariable("IPFSConnection")
	shell := ipfs.NewShell(IPFSConn)
	data, err := shell.Cat(cid)
	if err != nil {
		fmt.Println(err)
	}

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	// https://golangcode.com/convert-io-readcloser-to-a-string/
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()

	ipfsBody = map[string]interface{}{}
	json.Unmarshal([]byte(newStr), &ipfsBody)
	fmt.Println(ipfsBody)
	return ipfsBody
}
