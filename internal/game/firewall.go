package game

import (
	"fmt"
	"github.com/DmitriiTrifonov/blackhat/internal/models"
)

type Firewall struct {
	a Actor
}

func (f *Firewall) Capture(node *models.Node) bool {
	return node.Capture(models.OwnerFirewall)
}

func (f *Firewall) ChooseNext(node *models.Node) *models.Node {
	for i, c := range node.ConnectedNodes {
		fmt.Printf("%d) node ip:%s, lvl:%d, owner:%s \n", i+1, c.IPAddress, c.Level, c.Owner)
	}

	fmt.Print("enter node number to select next target: ")

	var nextNodeNum int
	_, err := fmt.Scan(&nextNodeNum)
	if err != nil {
		fmt.Println("error during node number parsing")

		return f.ChooseNext(node)
	}

	if nextNodeNum > len(node.ConnectedNodes) || nextNodeNum < 1 {
		fmt.Println("error during node number parsing")

		return f.ChooseNext(node)
	}

	return node.ConnectedNodes[nextNodeNum-1]
}
