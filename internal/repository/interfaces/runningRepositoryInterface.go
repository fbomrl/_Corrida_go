package interfaces

import (
	"github.com/fbomrl/_Corrida_go/internal/model"
)

type RunningRepositoryInterface interface {
	CreateRunning(running model.Running) error
}
