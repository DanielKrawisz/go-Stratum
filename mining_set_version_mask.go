package Stratum

type SetVersionMaskParams struct {
	Mask uint32
}

func (p *SetVersionMaskParams) read(n *notification) error {}

func SetVersionMask(uint32) *notification {}
