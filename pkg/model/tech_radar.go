package model

type TechRadarTopic struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Ring string `json:"ring"`
}

type AdapterTechRadar struct {
	Data    []*TechRadarTopic `json:"data"`
	Message string            `json:"message"`
}
