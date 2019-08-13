package id

import (
	"time"
)

func genTimestamp(epoch  uint64) uint64 {
	return uint64(time.Now().UnixNano()/1000000) - epoch
}

func waitNextClock(epoch, lastTimestamp uint64) uint64 {
	ts := uint64(time.Now().UnixNano() / 1000000)
	for {
		if ts <= lastTimestamp {
			ts = genTimestamp(epoch)
		} else {
			break
		}
	}
	return ts
}

func convertTimestamp(epoch, idTs uint64) uint64  {
	return (idTs + epoch) / 1000
}

//type TimerUtil struct {
//	epoch  uint64
//}
//func (t *TimerUtil) Timestamp() uint64 {
//
//	//// 最大峰值类型, 使用秒级时间戳
//	//if t.idType == SecondIdType {
//	//	return (uint64(time.Now().UnixNano()/1000000) - t.epoch) / 1000
//	//}
//	//
//	//// 最小粒度类型, 使用毫秒级时间戳
//	//if t.idType == MilliSecondIdType {
//	//	return uint64(time.Now().UnixNano()/1000000) - t.epoch
//	//}
//	//
//	//return (uint64(time.Now().UnixNano()/1000000) - t.epoch) / 1000
//}
//
//func (t *TimerUtil) WaitNextClock(lastTimestamp uint64) uint64 {
//	ts := uint64(time.Now().UnixNano() / 1000000)
//	for {
//		if ts <= lastTimestamp {
//			ts = t.Timestamp()
//		} else {
//			break
//		}
//	}
//	return ts
//}
//func (t *TimerUtil) ConvertTimestamp(uuidTimestamp uint64) uint64 {
//	return (uuidTimestamp + t.epoch) / 1000
//	//switch t.idType {
//	//case SecondIdType:
//	//	return (uuidTimestamp*1000 + t.epoch) / 1000
//	//case MilliSecondIdType:
//	//	return (uuidTimestamp + t.epoch) / 1000
//	//default:
//	//	return (uuidTimestamp*1000 + t.epoch) / 1000
//	//}
//}
