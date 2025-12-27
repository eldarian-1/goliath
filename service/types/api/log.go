package api

type Log struct {
	Level   string  `json:"level"`
	Message string  `json:"message"`
	Broker  *string `json:"broker"`
}
