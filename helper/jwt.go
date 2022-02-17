package helper

import "time"

type JwtMap struct {
	Id  int64         `json:"id"`
	Exp time.Duration `json:"exp"`
}
