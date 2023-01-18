package model

type HiringPosition struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AdapterHiringPosition struct {
	Data    []*HiringPosition `json:"data"`
	Message string            `json:"message"`
}
