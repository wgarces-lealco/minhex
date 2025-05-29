package secrets

type AWSSecretsManager struct{}

func NewAWSSecretsManager() *AWSSecretsManager {
	return &AWSSecretsManager{}
}

func (s *AWSSecretsManager) GetSecret(name string) (string, error) {
	return "", nil
}
