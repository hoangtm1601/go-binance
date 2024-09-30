package models

type CandleInterval string

const (
	OneMin           CandleInterval = "1min"
	FiveMin          CandleInterval = "5min"
	FifteenMin       CandleInterval = "15min"
	ThirtyMin        CandleInterval = "30min"
	SixtyMin         CandleInterval = "60min"
	TwoFortyMin      CandleInterval = "240min"
	SevenTwentyMin   CandleInterval = "720min"
	FourteenFortyMin CandleInterval = "1440min"
)
