package id

import (
	"sync/atomic"
	"unsafe"
)

type data struct {
	sequence      uint64
	lastTimestamp uint64
}

type AtomicGenerator struct {
	id   *Id
	data *data
	addr *unsafe.Pointer
}

func NewAtomicGenerator(machineId int, epoch uint64, serviceType int) (*AtomicGenerator, error) {
	cfg := &config{
		machineId:   machineId,
		epoch:       epoch,
		serviceType: serviceType,
	}

	id, err := NewId(cfg)
	if err != nil {
		return nil, err
	}

	gen := &AtomicGenerator{
		id:  id,
		data: &data{0, 0},
	}

	gen.addr = (*unsafe.Pointer)(unsafe.Pointer(gen.data))
	atomic.StorePointer(gen.addr, unsafe.Pointer(gen.data))
	return gen, nil
}



func (gen *AtomicGenerator) Make() (uint64, error) {
	var sequence, timestamp uint64
	for ; ; {
		oldDataPointer := atomic.LoadPointer(gen.addr)
		oldData := (*data)(oldDataPointer)

		// 旧的sequence
		sequence = oldData.sequence

		// 计算当前的时间戳
		timestamp = genTimestamp(gen.id.cfg.epoch)

		// 如果时间戳相同，去进行sequence的对比
		if timestamp == oldData.lastTimestamp {
			// 计算当前的sequence是否为空，如果为空需要等待下一个时钟
			if sequence = (sequence + 1) & uint64(gen.id.meta.getMaxSequence()); sequence == 0 {
				timestamp = waitNextClock(gen.id.cfg.epoch, oldData.lastTimestamp)
			}

			// 时间戳不同，将sequence置为0
		} else {
			sequence = 0
		}

		// 生成新的数据
		newData := &data{
			sequence:      sequence,
			lastTimestamp: timestamp,
		}

		// CAS设置成功就可以进行新的id计算
		if atomic.CompareAndSwapPointer(gen.addr, oldDataPointer, unsafe.Pointer(newData)) {
			uuid := gen.id.calculate(sequence, timestamp)
			return uuid, nil
		}
	}
}

func (gen *AtomicGenerator) Extract(uuid uint64) (*ExtractData) {
	return gen.id.transfer(uuid)
}
