package messages

import (
	"encoding/json"
)

type Video struct {
	VideoId  int64  `json:"video_id"`
	FileName string `json:"file_name"`
}

func (_ Video) GetTopic() string {
	return "video-processing"
}

func (v Video) ToBytes() ([]byte, error) {
	return json.Marshal(v)
}
