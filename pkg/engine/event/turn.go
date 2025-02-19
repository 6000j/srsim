package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type TurnTargetsAddedEventHandler = handler.EventHandler[TurnTargetsAdded]
type TurnTargetsAdded struct {
	Targets   []key.TargetID `json:"targets"`
	TurnOrder []TurnStatus   `json:"turn_order"`
}

type TurnResetEventHandler = handler.EventHandler[TurnReset]
type TurnReset struct {
	ResetTarget key.TargetID `json:"reset_target"`
	GaugeCost   int64        `json:"gauge_cost"`
	TurnOrder   []TurnStatus `json:"turn_order"`
}

type GaugeChangeEventHandler = handler.EventHandler[GaugeChange]
type GaugeChange struct {
	Key       key.Reason   `json:"key"`
	Target    key.TargetID `json:"target"`
	Source    key.TargetID `json:"source"`
	OldGauge  int64        `json:"old_gauge"`
	NewGauge  int64        `json:"new_gauge"`
	TurnOrder []TurnStatus `json:"turn_order"`
}

type CurrentGaugeCostChangeEventHandler = handler.EventHandler[CurrentGaugeCostChange]
type CurrentGaugeCostChange struct {
	Key     key.Reason   `json:"key"`
	Source  key.TargetID `json:"source"`
	OldCost int64        `json:"old_cost"`
	NewCost int64        `json:"new_cost"`
}

type TurnStatus struct {
	ID    key.TargetID `json:"id"`
	Gauge int64        `json:"gauge"`
	AV    float64      `json:"av"`
}
