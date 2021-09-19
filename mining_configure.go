package Stratum

// https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki

type ConfigureParams struct {
	Supported  []string
	Parameters map[string]interface{}
}

func (p *ConfigureParams) supports(extension string) bool {
	for _, supported := range p.Supported {
		if supported == extension {
			return true
		}
	}
	return false
}

type ConfigureResult map[string]interface{}

type VersionRollingConfigurationRequest struct {
	Mask        uint32
	MinBitCount byte
}

func (p *ConfigureParams) ReadVersionRolling() *VersionRollingConfigurationRequest {
	if !p.supports("version_rolling") {
		return nil
	}

	m, ok := p.Parameters["version_rolling.mask"]
	if !ok {
		return nil
	}

	maskstr, ok := m.(string)
	if !ok {
		return nil
	}

	mask, err := decodeLittleEndian(maskstr)
	if err != nil {
		return nil
	}

	b, ok := p.Parameters["version-rolling.min-bit-count"]
	if !ok {
		return nil
	}

	mbc, ok := b.(uint64)
	if !ok {
		return nil
	}

	if mbc > 255 {
		return nil
	}

	return &VersionRollingConfigurationRequest{
		Mask:        mask,
		MinBitCount: byte(mbc),
	}
}

type VersionRollingConfigurationResult struct {
	Accepted bool
	Mask     uint32
}

func (p *ConfigureResult) ReadVersionRolling() *VersionRollingConfigurationResult {
	v, ok := (*p)["version_rolling"]
	if !ok {
		return nil
	}

	accepted, ok := v.(bool)
	if !ok {
		return nil
	}

	m, ok := (*p)["version_rolling.mask"]
	if !ok {
		return nil
	}

	maskstr, ok := m.(string)
	if !ok {
		return nil
	}

	mask, err := decodeLittleEndian(maskstr)
	if err != nil {
		return nil
	}

	return &VersionRollingConfigurationResult{
		Accepted: accepted,
		Mask:     mask,
	}
}

type MinimumDifficultyConfigurationRequest struct {
	Difficulty Difficulty
}

func (p *ConfigureParams) ReadMinimumDifficulty() *MinimumDifficultyConfigurationRequest {
	if !p.supports("minimum_difficulty") {
		return nil
	}

	d, ok := p.Parameters["minimum_difficulty.value"]
	if !ok || !ValidDifficulty(d) {
		return nil
	}

	return &MinimumDifficultyConfigurationRequest{
		Difficulty: d,
	}
}

type MinimumDifficultyConfigurationResult struct {
	Accepted bool
}

func (p *ConfigureResult) ReadMinimumDifficulty() *MinimumDifficultyConfigurationResult {
	v, ok := (*p)["minimum_difficulty"]
	if !ok {
		return nil
	}

	accepted, ok := v.(bool)
	if !ok {
		return nil
	}

	return &MinimumDifficultyConfigurationResult{
		Accepted: accepted,
	}
}

type SubscribeExtranonceConfigurationRequest struct{}

func (p *ConfigureParams) ReadSubscribeExtranonce() *SubscribeExtranonceConfigurationRequest {
	if !p.supports("subscribe_extranonce") {
		return nil
	}

	return &SubscribeExtranonceConfigurationRequest{}
}

type SubscribeExtranonceConfigurationResult struct {
	Accepted bool
}

func (p *ConfigureResult) ReadSubscribeExtranonce() *SubscribeExtranonceConfigurationResult {
	v, ok := (*p)["subscribe_extranonce"]
	if !ok {
		return nil
	}

	accepted, ok := v.(bool)
	if !ok {
		return nil
	}

	return &SubscribeExtranonceConfigurationResult{
		Accepted: accepted,
	}
}

type InfoConfigurationRequest struct {
	ConnectionURL *string
	HWVersion     *string
	SWVersion     *string
	HWID          *string
}

func (p *ConfigureParams) ReadInfo() *InfoConfigurationRequest {
	if !p.supports("info") {
		return nil
	}

	var info InfoConfigurationRequest
	var connectionUrl string
	var hwVersion string
	var swVersion string
	var hwID string

	if x, ok := p.Parameters["info.connection-url"]; ok {
		connectionUrl, ok = x.(string)
		if !ok {
			return nil
		}

		info.ConnectionURL = &connectionUrl
	}

	if x, ok := p.Parameters["info.hw-version"]; ok {
		hwVersion, ok = x.(string)
		if !ok {
			return nil
		}

		info.HWVersion = &hwVersion
	}

	if x, ok := p.Parameters["info.sw-version"]; ok {
		swVersion, ok = x.(string)
		if !ok {
			return nil
		}

		info.SWVersion = &swVersion
	}

	if x, ok := p.Parameters["info.hw-id"]; ok {
		hwID, ok = x.(string)
		if !ok {
			return nil
		}

		info.HWID = &hwID
	}

	return &info
}

type InfoConfigurationResult struct {
	Accepted bool
}

func (p *ConfigureResult) ReadInfo() *InfoConfigurationResult {
	v, ok := (*p)["info"]
	if !ok {
		return nil
	}

	accepted, ok := v.(bool)
	if !ok {
		return nil
	}

	return &InfoConfigurationResult{
		Accepted: accepted,
	}
}

func (p *ConfigureParams) Add(interface{}) error {}

func (p *ConfigureResult) Add(interface{}) error {}

func ConfigureRequest(MessageID, ConfigureParams) request {}

func ConfigureResponse(MessageID, ConfigureResult) response {}
