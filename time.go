package doze

import (
	"encoding/json"
	"strconv"
	"time"
)

// Used to represent a time point
type Time time.Time

// Serializes the time into a unix timestamp
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// Deserializes an int64 a Time point
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return nil
	}
	*(*time.Time)(t) = time.Unix(value, 0)
	return nil
}

// Returns the unix timestamp as an int64
func (t *Time) Unix() int64 {
	return ((*time.Time)(t)).Unix()
}
