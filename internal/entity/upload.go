package entity

type UploadData struct {
	Service string            `json:"service"`
	Data    map[string]string `json:"data"`
}
