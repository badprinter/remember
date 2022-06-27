package models

import "time"

type Sessions struct {
	User_id    uint
	Login_time time.Time
	Alive      time.Time
	JWT        string
}

func (s *Sessions) TimeRemaining() time.Duration {
	return time.Now().Sub(s.Alive)
}
