package cache

type Stat struct {
	Count     int64
	KeySlice  int64
	ValueSize int64
}

func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySlice += int64(len(k))
	s.ValueSize -= int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySlice -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
