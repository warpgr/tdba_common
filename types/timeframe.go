package types

import "time"

type TimeFrame string

const (
	OneWeek        TimeFrame = "1W"
	ThreeDay       TimeFrame = "3D"
	OneDay         TimeFrame = "1D"
	EightHour      TimeFrame = "8h"
	FourHour       TimeFrame = "4h"
	OneHour        TimeFrame = "1h"
	ThirtyMinutes  TimeFrame = "30m"
	FifteenMinutes TimeFrame = "15m"
	FiveMinutes    TimeFrame = "5m"
	ThreeMinutes   TimeFrame = "3m"
	OneMinutes     TimeFrame = "1m"
)

func GetTimeFrameInMilliseconds(timeFrame TimeFrame) int64 {
	return GetTimeFrameInSeconds(timeFrame) * 1000
}

func GetTimeFrameInSeconds(timeFrame TimeFrame) int64 {
	if OneWeek == timeFrame {
		return 604800
	} else if ThreeDay == timeFrame {
		return 259200
	} else if OneDay == timeFrame {
		return 86400
	} else if EightHour == timeFrame {
		return 28800
	} else if FourHour == timeFrame {
		return 14400
	} else if OneHour == timeFrame {
		return 3600
	} else if ThirtyMinutes == timeFrame {
		return 1800
	} else if FifteenMinutes == timeFrame {
		return 900
	} else if FiveMinutes == timeFrame {
		return 300
	} else if ThreeMinutes == timeFrame {
		return 180
	} else if OneMinutes == timeFrame {
		return 60
	} else {
		return 0
	}
}

// USe unix milliseconds.
func IsExpired(timePoint int64) bool {
	return time.Now().UnixMilli() > timePoint
}

func IsExpiredForDuration(timePoint int64, timeFrame TimeFrame) bool {
	return (time.Now().UnixMilli() + GetTimeFrameInMilliseconds(timeFrame)) > timePoint
}
