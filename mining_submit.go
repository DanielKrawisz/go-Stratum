package Stratum

import "encoding/hex"

type SubmitParams Share

func Submit(id MessageID, share SubmitParams) request {
	var sx []interface{}
	if share.GeneralPurposeBits != nil {
		sx = make([]interface{}, 6)
		sx[5] = encodeLittleEndian(*share.GeneralPurposeBits)
	} else {
		sx = make([]interface{}, 5)
	}

	sx[0] = string(share.Name)
	sx[1] = encodeID(share.JobID)
	sx[2] = hex.EncodeToString(share.ExtraNonce2)
	sx[3] = encodeBigEndian(share.Time)
	sx[4] = encodeBigEndian(share.Nonce)

	return Request(id, MiningSubmit, sx)
}

type SubmitResult BooleanResult

func SubmitResponse(id MessageID, b bool) response {
	return BooleanResponse(id, b)
}
