package models

import "time"

type Node struct {
	Level          int
	Owner          Owner
	Login          string
	Password       string
	IPAddress      string
	ConnectedNodes []*Node
}

func (n *Node) Capture(attacker Owner) bool {
	t := time.Duration(n.Level) * time.Second
	<-time.After(t)
	n.Owner = attacker
	if n.Owner == OwnerFirewall {
		return true
	}
	for _, c := range n.ConnectedNodes {
		if c.Owner == OwnerFirewall {
			return true
		}
	}

	return false
}
