package data

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type Block struct {

	HeaderData Header
	Value 	   string
}

type HeartBeat struct {

	Id 			string
	Address 	string
	BlocJson    string
	PeerMapJson string
}

func (hb *HeartBeat) HBDEncodeToJson() (string, error) {
	//var hb HeartBeat
	value, err := json.Marshal(hb)
	if err != nil {
		return "", errors.New("Cannot encode Question to Json")
	}
	return string(value), nil
}

func HBDecodeFromJson(hbjson string) HeartBeat {
	var hb HeartBeat
	json.Unmarshal([]byte(hbjson), &hb)
	return hb
}

/*func(h *HeartBeat) ToString() string{
	str := ""
	for _, peer := range h.Peers {
		str += peer+ ",\t"
	}
	return str
}
*/

func NewBlock() Block {
	return Block{
		HeaderData: NewEmptyHeader(),
		Value:      "",
	}
}

type Header struct {
	Height     int32
	TimeStamp  string
	Hash       string
	ParentHash string
	Size	   int32
	Nonce	   int64
}

func NewEmptyHeader() Header {
	return Header{
		Height:     0,
		TimeStamp:  "",
		Hash:       "",
		ParentHash: "",
		Size:       0,
		Nonce:      0,
	}
}

func NewHeader(height int32, timeStamp string, hash string, parentHash string, size int32, nonce int64) Header {
	return Header{Height: height, TimeStamp: timeStamp, Hash: hash, ParentHash: parentHash, Size: size, Nonce: nonce}
}

//func NewHeader() Header {
//	return Header{
//		Height:     0,
//		timeStamp:  "",
//		Hash:       "",
//		ParentHash: "",
//		size:       0,
//		nonce:      0,
//	}
//}
func SimplePOW (block Block, difficulty int) int64{
	//fmt.Println("In POW")
	x := rand.Intn(99999999)
	for{
		y := sha256.Sum256([]byte(block.Value + string(x) + block.HeaderData.ParentHash))
		StringOfZeros := strings.Repeat("0", difficulty)
		hashstr:= hex.EncodeToString(y[:])
		//fmt.Println("hashstr : "+ hashstr)
		if strings.HasPrefix( hashstr , StringOfZeros) {
			break
		} else{
			x = rand.Intn(99999999)
			//fmt.Println("rand : ", x)
		}
	}

	return int64(x)
}
func (block Block) Initial(bc Blockchain, Value string, Height int32, timeStamp string, ParentHash string, size int32) {

	for {
		HeaderData := Header{Height, timeStamp, "", ParentHash, size,
			0}
		block = Block{HeaderData, Value}

		jsonBA, _ := json.Marshal(block)

		hash := sha256.Sum256(jsonBA)

		hashstr := hex.EncodeToString(hash[:])
		block.HeaderData.Hash = hashstr
		//block.Hash = hash

		Nonce := SimplePOW(block, 10)
		block.HeaderData.Nonce = Nonce

		bc.Insert(block)
		//bc.Show()
	}

}

func DecodeFromJSON(JSONstring string) Block{
	block := Block{}
	HeaderData := Header{}
	replacer := strings.NewReplacer(",", "", ".", "", ";", "", ":", " ", "\"", "")
	JSONstring = replacer.Replace(JSONstring)
	values := strings.Fields(JSONstring)

	for index := range values {
		if strings.Compare(values[index], "height") == 0 {

			height, err := strconv.ParseInt(values[index + 1], 10, 32)
			if err == nil{

			}
			HeaderData.Height = int32(height)
		}
		if strings.Compare(values[index], "timeStamp") == 0 {

			//timeStamp, err := strconv.ParseInt(values[index + 1], 10, 64)
			//if err == nil{
			//
			//}
			HeaderData.TimeStamp = values[index]
		}
		if strings.Compare(values[index], "size") == 0 {

			size, err := strconv.ParseInt(values[index + 1], 10, 32)
			if err == nil{

			}
			HeaderData.Size = int32(size)
		}
		if strings.Compare(values[index], "hash") == 0 {


			HeaderData.Hash = values[index + 1]
		}
		if strings.Compare(values[index], "parentHash") == 0 {

			HeaderData.ParentHash = values[index + 1]
		}

		if strings.Compare(values[index], "value") == 0{

			block.Value = values[index+1]
		}
		if strings.Compare(values[index], "nonce") == 0{

			nonce, err := strconv.ParseInt(values[index + 1], 10, 64)
			if err == nil{

			}
			HeaderData.Nonce = nonce
		}


	}

	block.HeaderData = HeaderData

	return block

}

func EncodeToJSON(block Block) string{
	JSONstring := "{\n\t\"height\":"+strconv.Itoa(int(block.HeaderData.Height)) + ",\n"
	JSONstring = JSONstring + "\t\"timeStamp\":"+ (block.HeaderData.TimeStamp) + ",\n"
	JSONstring = JSONstring + "\t\"hash\":"+ "\"" + block.HeaderData.Hash + "\"" + ",\n"
	JSONstring = JSONstring + "\t\"parentHash\":"+ "\"" + block.HeaderData.ParentHash + "\"" + ",\n"
	JSONstring = JSONstring + "\t\"size\":"+ strconv.Itoa(int(block.HeaderData.Size)) + ",\n"
	JSONstring = JSONstring + "\t\"value\":" + block.Value + ",\n"
	JSONstring = JSONstring + "'\t\"nonce\":" + strconv.Itoa(int(block.HeaderData.Nonce)) + "\n}\n"



	return JSONstring

}
