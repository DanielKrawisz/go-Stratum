package Stratum

import "errors"

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

func (p *ConfigureResult) supports(extension string) bool {
	for supported := range *p {
		if supported == extension {
			return true
		}
	}

	return false
}

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

func (p *ConfigureParams) addVersionRolling(x VersionRollingConfigurationRequest) error {
	if !p.supports("version_rolling") {
		return errors.New("request already contains version_rolling")
	}

	p.Supported = append(p.Supported, "version_rolling")

	p.Parameters["version_rolling.mask"] = encodeLittleEndian(x.Mask)
	p.Parameters["version-rolling.min-bit-count"] = x.MinBitCount

	return nil
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

	if !accepted {
		return &VersionRollingConfigurationResult{
			Accepted: false,
			Mask:     0,
		}
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
		Accepted: true,
		Mask:     mask,
	}
}

func (p *ConfigureResult) addVersionRolling(x VersionRollingConfigurationResult) error {
	if !p.supports("version_rolling") {
		return errors.New("result already contains version_rolling")
	}

	map[string]interface{}(*p)["version_rolling"] = x.Accepted
	if x.Accepted {
		map[string]interface{}(*p)["version_rolling.mask"] = encodeLittleEndian(x.Mask)
	}

	return nil
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

func (p *ConfigureParams) addMinimumDifficulty(x MinimumDifficultyConfigurationRequest) error {
	if !p.supports("minimum_difficulty") {
		return errors.New("request already contains minimum_difficulty")
	}

	p.Supported = append(p.Supported, "minimum_difficulty")

	p.Parameters["minimum_difficulty.value"] = x.Difficulty

	return nil
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

func (p *ConfigureResult) addMinimumDifficulty(x MinimumDifficultyConfigurationResult) error {
	if !p.supports("minimum_difficulty") {
		return errors.New("result already contains minimum_difficulty")
	}

	map[string]interface{}(*p)["minimum_difficulty"] = x.Accepted

	return nil
}

type SubscribeExtranonceConfigurationRequest struct{}

func (p *ConfigureParams) ReadSubscribeExtranonce() *SubscribeExtranonceConfigurationRequest {
	if !p.supports("subscribe_extranonce") {
		return nil
	}

	return &SubscribeExtranonceConfigurationRequest{}
}

func (p *ConfigureParams) addSubscribeExtranonce(x SubscribeExtranonceConfigurationRequest) error {
	if !p.supports("subscribe_extranonce") {
		return errors.New("request already contains subscribe_extranonce")
	}

	p.Supported = append(p.Supported, "subscribe_extranonce")

	return nil
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

func (p *ConfigureResult) addSubscribeExtranonce(x SubscribeExtranonceConfigurationResult) error {
	if !p.supports("subscribe_extranonce") {
		return errors.New("result already contains subscribe_extranonce")
	}

	map[string]interface{}(*p)["subscribe_extranonce"] = x.Accepted

	return nil
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

func (p *ConfigureParams) addInfo(x InfoConfigurationRequest) error {
	if !p.supports("info") {
		return errors.New("request already contains info")
	}

	p.Supported = append(p.Supported, "info")

	p.Parameters["info.connection-url"] = x.ConnectionURL
	p.Parameters["info.hw-version"] = x.HWVersion
	p.Parameters["info.sw-version"] = x.SWVersion
	p.Parameters["info.hw-id"] = x.HWID

	return nil
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

func (p *ConfigureResult) addInfo(x InfoConfigurationResult) error {
	if !p.supports("info") {
		return errors.New("result already contains info")
	}

	map[string]interface{}(*p)["info"] = x.Accepted

	return nil
}

func (p *ConfigureParams) Add(z interface{}) error {
	switch x := z.(type) {
	case VersionRollingConfigurationRequest:
		return p.addVersionRolling(x)
	case MinimumDifficultyConfigurationRequest:
		return p.addMinimumDifficulty(x)
	case SubscribeExtranonceConfigurationRequest:
		return p.addSubscribeExtranonce(x)
	case InfoConfigurationRequest:
		return p.addInfo(x)
	default:
		return errors.New("Unrecognized extension request")
	}
}

func (p *ConfigureResult) Add(z interface{}) error {
	switch x := z.(type) {
	case VersionRollingConfigurationResult:
		return p.addVersionRolling(x)
	case MinimumDifficultyConfigurationResult:
		return p.addMinimumDifficulty(x)
	case SubscribeExtranonceConfigurationResult:
		return p.addSubscribeExtranonce(x)
	case InfoConfigurationResult:
		return p.addInfo(x)
	default:
		return errors.New("Unrecognized extension request")
	}
}

func ConfigureRequest(id MessageID, p ConfigureParams) request {
	params := make([]interface{}, 2)
	params[0] = p.Supported
	params[1] = p.Parameters
	return Request(id, MiningConfigure, params)
}

func ConfigureResponse(id MessageID, r ConfigureResult) response {
	return Response(id, r)
}
