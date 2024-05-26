package game

import (
	"encoding/binary"
	"fmt"
	"github.com/DmitriiTrifonov/blackhat/internal/models"
	"math/rand/v2"
	"net"
	"time"
)

type Round struct {
	PlayerNodeStart     *models.Node
	FirewallNodeStart   *models.Node
	Target              *models.Node
	CurrentPlayerNode   *models.Node
	CurrentFirewallNode *models.Node
	NodesMap            map[string]*models.Node
	CaptureTime         time.Time
	BreachDetected      bool
}

func NewRound() *Round {
	r := &Round{}

	nodesMap := make(map[string]*models.Node)

	playerNodeStart := &models.Node{
		Level:          0,
		Owner:          models.OwnerPlayer,
		Login:          "",
		Password:       "",
		IPAddress:      generateIPAddr(),
		ConnectedNodes: nil,
	}

	nodesMap[playerNodeStart.IPAddress] = playerNodeStart

	return r
}

func generateIPAddr() string {
	buf := make([]byte, 0, 4)
	num := rand.Uint32()
	binary.LittleEndian.PutUint32(buf, num)

	return net.IP(buf).String()
}

func (r *Round) CheckPlayerWin() bool {
	return r.Target.Owner == models.OwnerPlayer || r.FirewallNodeStart.IPAddress == models.OwnerPlayer
}

func (r *Round) CheckPlayerLoose() bool {
	return r.PlayerNodeStart.Owner == models.OwnerFirewall
}

func (r *Round) CheckConnected(node *models.Node) {
	for _, c := range node.ConnectedNodes {
		fmt.Printf("node ip:%s, lvl:%d, owner:%s \n", c.IPAddress, c.Level, c.Owner)
	}
}

func (r *Round) Capture(node *models.Node) {
	t := time.Duration(node.Level) * time.Second
	<-time.After(t)
	node.Owner = models.OwnerPlayer
	if node.Owner == models.OwnerFirewall {
		r.BreachDetected = true

		return
	}
	for _, c := range node.ConnectedNodes {
		if c.Owner == models.OwnerFirewall {
			r.BreachDetected = true

			return
		}
	}
}
