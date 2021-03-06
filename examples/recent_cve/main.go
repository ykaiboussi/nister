package main

import (
	"fmt"

	"github.com/ykaiboussi/nister"
)

func main() {
	data := nister.HighCVE()
	for _, v := range data {
		fmt.Println("ID: ", v.CVE.MetaData.ID)
	}
}
