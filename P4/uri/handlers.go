package uri

import (
	"../data"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)



var SELF_ADDRESS string
var sbc = data.NewSBC()
var peerList []string
var PortID string
var address = "http://localhost:"
var PList string


func InitSelfAddress(port string) {
	SELF_ADDRESS = "http://localhost:" + port
	PortID = port

	go KeepGeneratingBlocks()

}
func Upload(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data.EncodesToJSON(sbc.Bc.Chain)))

}
func Download(){
	response, _ := http.Get("http://localhost:6689/upload")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println("error")
	}

	data.DecodeFromJSON(string(body))

}
func ShowHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sbc.Bc.Show()))
}
func Register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := readRequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	peerList, err := data.PeerListFromJson(reqBody)
	PList = peerList.ToString()

	fmt.Println(peerList.ToString())


	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(reqBody))
}

func readRequestBody(r *http.Request) (string, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("cannot read request body")
	}
	defer r.Body.Close()
	return string(reqBody), nil
}


func KeepGeneratingBlocks() {
	for {
		if sbc.Bc.Length == 0 {
			for1stBlock()
		} else {
			forAfter1stBlock()
		}

	}
}


func for1stBlock() {
	bloc := data.NewBlock()
	if sbc.Bc.Length == 0 {
		height := 1
		timeStamp := ""
		parentHash := "genesis"
		size := 0
		headerData := data.NewHeader(int32(height), timeStamp, "", parentHash, int32(size),
			0)

		bloc.HeaderData = headerData
		bloc.Value = "ok ok"

		jsonBA, _ := json.Marshal(bloc)

		hash := sha256.Sum256(jsonBA)

		hashstr := hex.EncodeToString(hash[:])
		bloc.HeaderData.Hash = hashstr

		nonce := data.SimplePOW(bloc, 3)
		bloc.HeaderData.Nonce = nonce


		sbc.Bc.Insert(bloc)
		sbc.Bc.Length++
	}
}


func forAfter1stBlock() {
	bloc := data.NewBlock()
		//fmt.Println("bc length is 0")
		height := sbc.Bc.Length+1
		timeStamp := ""//time.Now().String()
		parentHash := sbc.Bc.GetLatestBlock()[0].HeaderData.Hash
		size := 0
		headerData := data.NewHeader(int32(height), timeStamp, "", parentHash, int32(size),
			0)

		//bloc := data.NewBlock()
		bloc.HeaderData = headerData
		bloc.Value = "ok ok"

		jsonBA, _ := json.Marshal(bloc)

		hash := sha256.Sum256(jsonBA)

		hashstr := hex.EncodeToString(hash[:])
		bloc.HeaderData.Hash = hashstr
		//block.Hash = hash

		nonce := data.SimplePOW(bloc, 3)
		bloc.HeaderData.Nonce = nonce
		//fmt.Println("block :\n" + data.EncodeToJSON(bloc))

		sbc.Bc.Insert(bloc)
		SendHeartBeat(bloc)
		sbc.Bc.Length++
		//fmt.Println("bc.Length",bc.Length)
		//fmt.Println("Showing Blockchain")
		//fmt.Println(bc.Show())

}

func SendHeartBeat(b data.Block) {
	hb := data.HeartBeat{PortID,address,data.EncodeToJSON(b),"i am peerlist"}
	hbJson, _ := hb.HBDEncodeToJson()
	for _, peer := range peerList{
		uri := peer+"/heartbeat"
		http.Post(uri, "application/json", bytes.NewBuffer([]byte(hbJson)))
	}

}

func ReceiveHeartBeat(w http.ResponseWriter, r *http.Request){
	reqBody, err := readRequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	hb := data.HBDecodeFromJson(reqBody)
	bloc := data.DecodeFromJSON(hb.BlocJson)
	sbc.Bc.Insert(bloc)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(reqBody))
}

