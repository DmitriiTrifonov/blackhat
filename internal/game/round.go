package game

import (
	"encoding/binary"
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

	r.NodesMap = make(map[string]*models.Node)

	playerNodeStart := &models.Node{
		Level:          0,
		Owner:          models.OwnerPlayer,
		Login:          "",
		Password:       "",
		IPAddress:      generateIPAddr(),
		ConnectedNodes: nil,
	}

	r.NodesMap[playerNodeStart.IPAddress] = playerNodeStart

	r.PlayerNodeStart = playerNodeStart

	return r
}

func generateIPAddr() string {
	buf := make([]byte, 4)
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
