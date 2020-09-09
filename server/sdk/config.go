package sdk

import (
	"go.uber.org/zap"
)

type Endpoint struct {
	Target    string
	Authority string
}

// export Config just for reference
type Config struct {
	// local service address
	Bind string
	// log configuration
	Log *zap.Config
	// api endpoints
	DBV1 *Endpoint
	GWV1 *Endpoint
}
