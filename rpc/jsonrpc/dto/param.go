package dto

import "time"

type SumParam struct {
	A int
	B int
}

type SumRes struct {
	Sum  int
	Time time.Time
}
