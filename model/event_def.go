package model

type (
	EventType int
	Level     int
)

const (
	Checkpoint EventType = 1
)

type Event struct {
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Lang           string         `json:"lang"`
	CheckpointData CheckpointData `json:"checkpoint_data"`
}

type CheckpointData struct {
	Text   []string `json:"text"`
	Points []Point  `json:"points"`
}

type Point struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Events map[EventType]map[Level]Event

func (e EventType) String() string {
	switch e {
	case Checkpoint:
		return "Checkpoint"
	default:
		return "Unknown"
	}
}

type Summaries map[EventType]map[Level]string
