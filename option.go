package auth

type Option func(*Manager) error

func OptConfig(config Config) Option {
	return func(m *Manager) error {
		m.Config = config
		return nil
	}
}

func OptPersist(persist Persist) Option {
	return func(m *Manager) error {
		m.Persist = persist
		return nil
	}
}

func OptLogin(login LoginMethod) Option {
	return func(m *Manager) error {
		m.LoginMethod = login
		return nil
	}
}

func OptAuthenticator(auth Authenticator) Option {
	return func(m *Manager) error {
		m.Auth = auth
		return nil
	}
}
