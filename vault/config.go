package vault

type VaultBridgeConfig struct {
	CustomerRefIdPrefix string
	BridgeResources     map[string]map[uint8]*ResourceChain `mapstructure:"bridgeResources"`
}

type ResourceChain struct {
	ChainName   string `mapstructure:"chainName"`
	CoinKey     string `mapstructure:"coinKey"`
	CoinName    string `mapstructure:"coinName"`
	CoinDecimal int    `mapstructure:"coinDecimal"`
	CoinType    string `mapstructure:"coinType"`
}
