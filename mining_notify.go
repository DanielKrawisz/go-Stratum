package Stratum

type NotifyParams struct {
	JobID         JobID
	Digest        [32]byte
	GenerationTX1 []byte
	GenerationTX2 []byte
	Path          [][32]byte
	Version       uint32
	Target        []byte
	Timestamp     uint32
}

func (s *NotifyParams) MarshallJSON() ([]byte, error) {}

func (s *NotifyParams) UnmarshallJSON([]byte) error {}

type notify struct {
	Notification
}

func (n *notify) Params() (NotifyParams, error) {}

func NewNotify(MethodID, NotifyParams) *notify {}
