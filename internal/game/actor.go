package game

import (
	"github.com/DmitriiTrifonov/blackhat/internal/models"
)

type Actor interface {
	Capture(node *models.Node) bool
	ChooseNext(node *models.Node) *models.Node
}
