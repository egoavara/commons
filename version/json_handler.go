package version

import (
	"github.com/pkg/errors"
)

func (s *Version) MarshalJSON() ([]byte, error) {
	return []byte(`"` + Marshal(*s) + `"`), nil
}

func (s *Version) UnmarshalJSON(data []byte) error {
	if len(data) < 2{
		return errors.New("Version must be string")
	}
	*s = Unmarshal(string(data[1:len(data) - 1]))
	return nil
}
