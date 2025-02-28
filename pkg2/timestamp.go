package pkg2

import (
	"fmt"
	"strconv"
	"strings"
)

type Interval struct {
	Begin *Timestamp
	End   *Timestamp
	isMax bool
}

func MaxInterval() *Interval {
	return &Interval{
		Begin: &Timestamp{isInf: false},
		End:   &Timestamp{isInf: true},
		isMax: true,
	}
}

func ParseInterval(str string) (*Interval, error) {
	if str == "" {
		return MaxInterval(), nil
	}

	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("interval doesn't have correct format, expected '<left>-<right>' but didn't, found %v '-'\n\t'%v'", len(parts)-1, str)
	}

	beginStr := parts[0]
	endStr := parts[1]

	begin, err := ParseTimestamp(beginStr)
	if err != nil {
		return nil, err
	}

	end, err := ParseTimestamp(endStr)
	if err != nil {
		return nil, err
	}

	return &Interval{
		Begin: begin,
		End:   end,
	}, nil
}

func (i *Interval) String() string {
	return i.Begin.String() + "-" + i.End.String()
}

func (i *Interval) PathFormat() string {
	return i.Begin.PathFormat() + "_" + i.End.PathFormat()
}

// TODO: I don't really like this but meh :))
func (i *Interval) IsMax() bool {
	return i.isMax
}

func (i *Interval) Duration() *Timestamp {
	return TimestampFromSeconds(i.End.Seconds() - i.Begin.Seconds())
}

type Timestamp struct {
	Hour   int
	Minute int
	Second float64
	isInf  bool
}

func ParseTimestamp(timestampStr string) (*Timestamp, error) {
	var err error = nil
	var minutes, hours int64 = 0, 0
	var seconds float64
	index := 0

	if timestampStr == "inf" {
		return &Timestamp{
			isInf: true,
		}, nil
	}

	switch comps := strings.Split(timestampStr, ":"); len(comps) {
	case 3:
		hours, err = strconv.ParseInt(comps[index], 10, 64)
		if err != nil {
			return nil, err
		}
		index += 1
		fallthrough
	case 2:
		minutes, err = strconv.ParseInt(comps[index], 10, 64)
		if err != nil {
			return nil, err
		}
		index += 1
		fallthrough
	case 1:
		seconds, err = strconv.ParseFloat(comps[index], 64)
		if err != nil {
			return nil, err
		}

		return &Timestamp{
			Hour:   int(hours),
			Minute: int(minutes),
			Second: seconds,
		}, nil
	default:
		return nil, fmt.Errorf("unexpected amount of components in the timestamp (%d) wanted between(1, 3)\n\t%s", len(comps), timestampStr)
	}
}

func TimestampFromSeconds(seconds float64) *Timestamp {
	hours := int(seconds / (60 * 60))
	seconds -= float64(hours) * (60 * 60)

	minutes := int(seconds / 60)
	seconds -= float64(minutes * 60)

	return &Timestamp{
		Hour:   hours,
		Minute: minutes,
		Second: seconds,
		isInf:  false,
	}
}

func (ts *Timestamp) String() string {
	if ts.isInf {
		return "inf"
	}

	return fmt.Sprintf("%02d:%02d:%0.2f", ts.Hour, ts.Minute, ts.Second)
}

func (ts *Timestamp) PathFormat() string {
	if ts.isInf {
		return "inf"
	}

	// TODO: do we want to exclude the miliseconds? Maybe yes?
	return fmt.Sprintf("%02d-%02d-%02.0f", ts.Hour, ts.Minute, ts.Second)
}

func (ts *Timestamp) Seconds() float64 {
	return float64(ts.Hour*60*60) + float64(ts.Minute*60) + ts.Second
}
