package Stratum

type setVersionMask struct {
	Notification
}

func (au *setVersionMask) Params() uint32 {}

func NewSetVersionMask(RequestID, uint32) (*setVersionMask, error) {}
