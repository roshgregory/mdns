package main

import (
	"context"

	host "github.com/libp2p/go-libp2p/core/host"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

type THS struct {
	Id      peer.ID
	Moniker string
	Round   int
}

type Message struct {
	From         P2P
	To           peer.ID
	Type         int
	Payload_name string
	Payload      string
	Status       int
}

type P2P struct {

	// Represents the libp2p host
	Host             host.Host
	Host_ip          string
	Ctx              context.Context
	Peers            []THS
	Connectedparties int
	ThisParty        int

	//Used for indexing peers
	Sorted_Peers []THS
	My_Index     int

	//Used for comms

}
