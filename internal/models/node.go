package models

type Node struct {
	Level          int
	Owner          Owner
	Login          string
	Password       string
	IPAddress      string
	ConnectedNodes []*Node
}
