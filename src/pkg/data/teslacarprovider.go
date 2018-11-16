package data

type TeslaCarProvider struct {
	// External parameters
	VIN         string // API interactions needs to be tied
	Credentials UsernamePasswordCredential

	// Internal state
	initialized   bool
	hasDriveState bool
	apiEndpoint   string
}

func NewTeslaCarProvider(vin string, username string, password string) *TeslaCarProvider {
	teslaCarProvider := &TeslaCarProvider{VIN: vin, Credentials: UsernamePasswordCredential{Username: username, Password: password}}
	return teslaCarProvider
}

func (tesla *TeslaCarProvider) GetLocation() (LatLong, error) {
	return LatLong{}, nil
}

func (tesla *TeslaCarProvider) GetState() (LockState, error) {
	tesla.initialize()
	// TODO: Implement read lock state
	return UNKNOWN, nil
}

func (tesla *TeslaCarProvider) SetState(state LockState) (LockState, error) {
	tesla.initialize()
	// TODO: Implement write lock state
	return UNKNOWN, nil
}

// Lazy initialization function.
func (tesla *TeslaCarProvider) initialize() error {
	if tesla.initialized {
		return nil
	}

	// Get Vehicle

	tesla.initialized = true
	return nil
}

type UsernamePasswordCredential struct {
	Username string
	Password string
}
