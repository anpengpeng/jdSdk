package jdSdk

type JdBaseApiRequest interface {
	GetConfig() *JdBaseConfig
	SetConfig(*JdBaseConfig)
	GetParamsObject() interface{}
	GetMethodName() string
}

type BaseJdApiRequest struct {
	config *JdBaseConfig
	client *JdClient
}

func (b *BaseJdApiRequest) GetConfig() *JdBaseConfig {
	return b.config
}

func (b *BaseJdApiRequest) SetConfig(config *JdBaseConfig) {
	b.config = config
}

func (b *BaseJdApiRequest) GetClient() *JdClient {
	return b.client
}

func (b *BaseJdApiRequest) SetClient(client *JdClient) {
	b.client = client
}
