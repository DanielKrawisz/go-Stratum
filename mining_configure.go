package Stratum

// https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki

type ConfigureParams struct {
	Supported  []string
	Parameters map[string]interface{}
}

type ConfigureResult map[string]interface{}

type VersionRollingConfigurationRequest struct {
	Mask        uint32
	MinBitCount byte
}

func (p *VersionRollingConfigurationRequest) Read(*ConfigureParams) error {}

type VersionRollingConfigurationResult struct {
	Mask uint32
}

func (p *VersionRollingConfigurationResult) Read(*ConfigureResult) error {}

type MinimumDifficultyConfigurationRequest struct {
	Difficulty Difficulty
}

func (p *MinimumDifficultyConfigurationRequest) Read(*ConfigureParams) error {}

type MinimumDifficultyConfigurationResult struct {
	Accepted bool
}

func (p *MinimumDifficultyConfigurationResult) Read(*ConfigureResult) error {}

type SubscribeExtranonceConfigurationRequest struct{}

func (p *SubscribeExtranonceConfigurationRequest) Read(*ConfigureParams) error {}

type SubscribeExtranonceConfigurationResult struct {
	Accepted bool
}

func (p *SubscribeExtranonceConfigurationResult) Read(*ConfigureResult) error {}

type InfoConfigurationRequest struct {
	ConnectionURL *string
	HWVersion     *string
	SWVersion     *string
	HWID          *string
}

func (p *InfoConfigurationRequest) Read(*ConfigureParams) error {}

type InfoConfigurationResult struct {
	Accepted bool
}

func (p *InfoConfigurationResult) Read(*ConfigureResult) error {}

func (p *ConfigureParams) Add(interface{}) error {}

func (p *ConfigureResult) Add(interface{}) error {}

func ConfigureRequest(MessageID, ConfigureParams) request {}

func ConfigureResponse(MessageID, ConfigureResult) response {}
