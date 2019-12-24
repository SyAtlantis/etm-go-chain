package utils

import (
	"time"
)

type slots struct {
	Interval    int // 出块时间
	Delegates   int // 代理数量（不为101时需要同步修改gegesisBlock.json）
	RoundBlocks int // 每轮区块的数量
	Leading     int // 前导位数
	PowTimeOut  int // pow超时时间（单位s）
}

func NewSlots() slots {
	return slots{
		Interval:    3,
		Delegates:   101,
		RoundBlocks: 101,
		Leading:     7,
		PowTimeOut:  2,
	}
}

// 定义起始时间
func (s *slots) BeginEpochTime() time.Time {
	d := time.Date(2018, time.October, 12, 12, 0, 0, 0, time.UTC)
	return d
}

// 获取相对于起始时间的时间戳
func (s *slots) GetTime(timestamps ...int64) int64 {
	var timestamp int64
	if len(timestamps) == 0 {
		timestamp = time.Now().Unix()
	} else {
		timestamp = timestamps[0]
	}
	d := s.BeginEpochTime()
	t0 := d.Unix()
	return timestamp - t0
}

// 获取当前时间戳
func (s *slots) GetRealTime(epochTimes ...int64) int64 {
	var epochTime int64
	if len(epochTimes) == 0 {
		epochTime = s.GetTime()
	} else {
		epochTime = epochTimes[0]
	}
	d := s.BeginEpochTime()
	t0 := d.Unix()
	return (t0 + epochTime)*1000
}

// 获取slot
func (s *slots) GetSlotNumber(epochTimes ...int64) int64 {
	var epochTime int64
	if len(epochTimes) == 0 {
		epochTime = s.GetTime()
	} else {
		epochTime = epochTimes[0]
	}
	return epochTime / int64(s.Interval)
}

// slot转换为时间戳
func (s *slots) GetSlotTime(slot int64) int64 {
	return slot * int64(s.Interval)
}

// 获取下一个slot
func (s *slots) GetNextSlot() int64 {
	slot := s.GetSlotNumber()
	return slot + 1
}

// 获取加一轮后的slot
func (s *slots) GetLastSlot(nextSlot int64) int64 {
	return nextSlot + int64(s.Delegates)
}

// 获取当前data的时间戳
func (s *slots) RoundTime(data time.Time) int64 {
	return data.Unix()
}

// 获取一天能出块的数量
func (s *slots) GetHeightPerDay() int64 {
	return 24 * 60 * 60 / int64(s.Interval)
}
