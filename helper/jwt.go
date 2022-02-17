package helper

import "time"

type JwtMap struct {
	Id  uint64        `json:"id"`
	Exp time.Duration `json:"exp"`
}
