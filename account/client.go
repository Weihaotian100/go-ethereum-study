package account

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ClientTest() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
