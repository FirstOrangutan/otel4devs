package tailtracer

import (
	"fmt"
	"time"

)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
   config.ReceiverSettings `mapstructure:",squash"`
   Interval       string `mapstructure:"interval"`
   NumberOfTraces int `mapstructure:"number_of_traces"`
}


