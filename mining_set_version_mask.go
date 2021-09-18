package Stratum

import "errors"

type SetVersionMaskParams struct {
	Mask uint32
}

func (p *SetVersionMaskParams) read(n *notification) error {
	if len(n.params) != 1 {
		return errors.New("invalid format")
	}

	mask, ok := n.params[0].(string)
	if !ok {
		return errors.New("invalid format")
	}

	var err error
	p.Mask, err = decodeLittleEndian(mask)
	if err != nil {
		return err
	}

	return nil
}

func SetVersionMask(u uint32) notification {
	return Notification(MiningSetVersionMask, []interface{}{encodeLittleEndian(u)})
}
