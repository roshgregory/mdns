package main

func main() {
	var party P2P

	//Init P2P values to Party
	p2p_ch := make(chan P2P)
	go func(p2p_ch chan P2P) {
		P2p_init(p2p_ch)
	}(p2p_ch)

	party = <-p2p_ch
	Create_peer(&party)
}
