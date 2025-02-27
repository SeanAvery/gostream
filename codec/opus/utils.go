package opus

import (
	"time"

	"github.com/edaniels/golog"

	"github.com/viamrobotics/gostream"
	"github.com/viamrobotics/gostream/codec"
)

// DefaultStreamConfig configures Opus as the audio encoder for a stream.
var DefaultStreamConfig gostream.StreamConfig

func init() {
	DefaultStreamConfig.AudioEncoderFactory = NewEncoderFactory()
}

// NewEncoderFactory returns an Opus audio encoder factory.
func NewEncoderFactory() codec.AudioEncoderFactory {
	return &factory{}
}

type factory struct{}

func (f *factory) New(sampleRate, channelCount int, latency time.Duration, logger golog.Logger) (codec.AudioEncoder, error) {
	return NewEncoder(sampleRate, channelCount, latency, logger)
}

func (f *factory) MIMEType() string {
	return "audio/opus"
}
