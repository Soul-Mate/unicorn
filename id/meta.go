package id

const (
	defaultMachineBit   = 8
	defaultSeqBit       = 12
	defaultTimestampBit = 41
	defaultTypeBit      = 1
	defaultReservedBit  = 1
)

type meta struct {
	machine   uint8 // 机器
	seq       uint8 // 序列号
	timestamp uint8 // 时间戳
	typ       uint8 // 类型
	reserved  uint8 // 保留字段
}

func newMeta() *meta {
	return &meta{
		machine:   defaultMachineBit,
		seq:       defaultSeqBit,
		timestamp: defaultTimestampBit,
		typ:       defaultTypeBit,
		reserved:  defaultReservedBit,
	}
}

// getSeqLeftShift 获取seq需要左移的长度
func (m *meta) getSeqLeftShift() uint64 {
	return uint64(m.machine)
}

// getTimestampLeftShift 获取timestamp需要左移的长度
func (m *meta) getTimestampLeftShift() uint64 {
	return uint64(m.machine + m.seq)
}

// getTypeLeftShift 获取type需要左移的长度
func (m *meta) getTypeLeftShift() uint64 {
	return uint64(m.machine + m.seq + m.timestamp)
}

// getReservedLeftShift 获取reserved 需要左移的长度
func (m *meta) getReservedLeftShift() uint64 {
	return uint64(m.machine + m.seq + m.timestamp + m.reserved)
}

// getMaxMachine 获取machine最大值
func (m *meta) getMaxMachine() int64 {
	return -1 ^ - 1<<m.machine
}

// getMaxSequence 获取sequence最大值
func (m *meta) getMaxSequence() int64 {
	return -1 ^ -1<<m.seq
}

// getMaxTimestamp 获取timestamp最大值
func (m *meta) getMaxTimestamp() int64 {
	return -1 ^ -1<<m.timestamp
}

// getMaxIdType 获取type最大值
func (m *meta) getMaxIdType() int64 {
	return -1 ^ -1<<m.typ
}

// getMaxReserved 获取 reserved最大值
func (m *meta) getMaxReserved() int64 {
	return -1 ^ -1<<m.reserved
}

