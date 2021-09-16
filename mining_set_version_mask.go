package Stratum

type setVersionMask struct {
	Notification
}

func (au *setVersionMask) Params() (uint32, error) {}

func NewSetVersionMask(uint32) *setVersionMask {}
