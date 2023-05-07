package models

type Str struct {
	Value string
	Sum   int32
}

func (s *Str) Lesser(other *Str) bool {
	if s.Sum < other.Sum {
		return true
	}
	return false
}

func (s *Str) Equal(other *Str) bool {
	if s.Sum == other.Sum {
		return true
	}
	return false
}

func (s *Str) Greater(other *Str) bool {
	if s.Sum > other.Sum {
		return true
	}
	return false
}

func (s *Str) Swap(other *Str) {
	s.Value, other.Value = other.Value, s.Value
	s.Sum, other.Sum = other.Sum, s.Sum
}
