package dinopay

import "encoding/json"

type EventEnvelope struct {
    Id   string          `json:"id"`
    Type string          `json:"type"`
    Time string          `json:"time"`
    Data json.RawMessage `json:"data"`
}
