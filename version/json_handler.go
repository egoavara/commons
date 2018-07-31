package version

func (s *Version) MarshalJSON() ([]byte, error) {
	return []byte(Marshal(*s)), nil
}

func (s *Version) UnmarshalJSON(data []byte) error {
	*s = Unmarshal(string(data))
	return nil
}