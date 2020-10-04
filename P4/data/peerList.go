package data

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type peerList struct {
	Peers []string
}

func(p *peerList) ToString() string{
	str := ""
	for _, peer := range p.Peers {
		str += peer+ ",\t"
	}
	return str
}


func (p *peerList) ToJson() (string, error) {
	value, err := json.Marshal(p)
	if err != nil {
		return "", errors.New("Cannot encode Question to Json")
	}
	return string(value), nil
}

func PeerListFromJson(inputJson string) (peerList, error) {
	peers := peerList{}
	err := json.Unmarshal([]byte(inputJson), &peers)
	if err != nil {
		return peers, errors.New("Cannot decode Json to Question")
	}
	return peers, nil
}