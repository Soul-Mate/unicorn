package id

import (
	"fmt"
)

type Generator interface {
	Make() uint64
	UnMake() ExtractData

}

type ExtractData struct {
	MachineId   int		// 机器ID
	Sequence    uint64 	// 序列号
	Timestamp   uint64 	// 时间戳
	ServiceType int 	// 服务类型
	Reserved    int 	//  保留字段
}

type config struct {
	epoch       uint64
	machineId   int
	serviceType int
	reserved    int
}

type Id struct {
	cfg       *config
	meta      *meta
}

func NewId(cfg *config) (*Id, error) {
	if err := checkConfig(cfg); err != nil {
		return nil, err
	}

	id := &Id{
		cfg: cfg,
		meta:newMeta(),
	}

	return id, nil
}

func checkConfig(cfg *config) error {
	if cfg.epoch == 0 {
		return fmt.Errorf("epoch cannot be empty, the id type supports: : \n\t%d: max peak type\n\t%d: min granularity type\n",
			SecondIdType, MilliSecondIdType)
	}

	if cfg.machineId < 1 || cfg.machineId > 256 {
		return fmt.Errorf("machine id is not in range, machine id range: %d ~ %d\n",
			1, 256)
	}

	//if cfg.Version != UnavailableVersion && cfg.Version != NormalVersion {
	//	return fmt.Errorf("version is unsupported value, the version supports: : \n\t%d: unavailable version\n\t%d: normal version\n",
	//		UnavailableVersion, NormalVersion)
	//}

	return nil
}

func (id *Id) calculate(sequence, lastTimestamp uint64) uint64 {
	var uuid uint64
	uuid |= uint64(id.cfg.machineId)
	uuid |= uint64(sequence << id.meta.getSeqLeftShift())
	uuid |= uint64(lastTimestamp << id.meta.getTimestampLeftShift())
	uuid |= uint64(id.cfg.serviceType << id.meta.getTypeLeftShift())
	uuid |= uint64(id.cfg.reserved << id.meta.getReservedLeftShift())
	return uuid
}

func (id *Id) transfer(uuid uint64) *ExtractData {
	data := &ExtractData{}
	data.MachineId = int(uuid & uint64(id.meta.getMaxMachine()))
	data.Sequence = (uuid >> id.meta.getSeqLeftShift()) & uint64(id.meta.getMaxSequence())
	data.Timestamp = convertTimestamp(id.cfg.epoch, uuid >> id.meta.getTimestampLeftShift() & uint64(id.meta.getMaxTimestamp()))
	data.Reserved = int((uuid >> id.meta.getReservedLeftShift()) & uint64(id.meta.getMaxReserved()))
	data.ServiceType = int((uuid >> id.meta.getTypeLeftShift()) & uint64(id.meta.getMaxIdType()))
	return data
}
