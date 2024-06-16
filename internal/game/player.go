package game

import (
	"fmt"
	"github.com/DmitriiTrifonov/blackhat/internal/models"
)

type Player struct {
}

func (p *Player) Capture(node *models.Node) bool {
	return node.Capture(models.OwnerPlayer)
}

func (p *Player) ChooseNext(node *models.Node) *models.Node {
	for i, c := range node.ConnectedNodes {
		fmt.Printf("%d) node ip:%s, lvl:%d, owner:%s \n", i+1, c.IPAddress, c.Level, c.Owner)
	}

	fmt.Print("enter node number to select next target: ")

	var nextNodeNum int
	_, err := fmt.Scan(&nextNodeNum)
	if err != nil {
		fmt.Println("error during node number parsing")

		return p.ChooseNext(node)
	}

	if nextNodeNum > len(node.ConnectedNodes) || nextNodeNum < 1 {
		fmt.Println("error during node number parsing")

		return p.ChooseNext(node)
	}

	return node.ConnectedNodes[nextNodeNum-1]
}

func (p *Player) Traverse(current, final *models.Node) {
	if current == final {
		return
	}

	for i, node := range current.ConnectedNodes {
		fmt.Printf("%d) node ip:%s, lvl:%d, owner:%s \n", i+1, node.IPAddress, node.Level, node.Owner)

		p.Traverse(node, final)
	}
}
