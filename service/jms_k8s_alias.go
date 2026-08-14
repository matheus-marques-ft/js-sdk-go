package service

// GetK8sShellAliases fetches the real JumpServer user's saved kubectl shell
// aliases for the given connect token, so Koko can inject them into the
// generated .bashrc when opening a Kubernetes shell session.
func (s *JMService) GetK8sShellAliases(tokenId string) (aliases map[string]string, err error) {
	data := map[string]string{
		"id": tokenId,
	}
	var resp struct {
		Aliases map[string]string `json:"aliases"`
	}
	_, err = s.authClient.Post(SuperConnectTokenK8sAliasesURL, data, &resp)
	return resp.Aliases, err
}

// SetK8sShellAlias persists a single kubectl shell alias for the real
// JumpServer user behind the given connect token.
func (s *JMService) SetK8sShellAlias(tokenId, name, command string) (err error) {
	data := map[string]string{
		"id":      tokenId,
		"name":    name,
		"command": command,
	}
	_, err = s.authClient.Post(SuperConnectTokenK8sAliasesSetURL, data, nil)
	return
}
