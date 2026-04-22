package rabbit

type Event struct {
	EventID string      `json:"event_id"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
}
