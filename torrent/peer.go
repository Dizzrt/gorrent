package torrent

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Peer struct {
	Ip   net.IP
	Port uint16
}

const (
	IpLen   int = 4
	PortLen int = 2
	PeerLen int = IpLen + PortLen
)

func buildPeers(data []byte) []Peer {
	peerCnt := len(data) / PeerLen
	if len(data)%PeerLen != 0 {
		fmt.Println("received malformed peers")
		return nil
	}

	peers := make([]Peer, peerCnt)
	for i := 0; i < peerCnt; i++ {
		offset := i * PeerLen
		peers[i].Ip = net.IP(data[offset : offset+IpLen])
		peers[i].Port = binary.BigEndian.Uint16(data[offset+IpLen : offset+PeerLen])
	}

	return peers
}

// TODO
// func GetPeersByHttp() []Peer{}
// func GetPeersByUdp() []Peer{}

func GetPeers(bt *Torrent) []Peer {
	// TODO
	return nil
}
