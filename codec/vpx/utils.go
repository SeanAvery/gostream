package vpx

import (
	"fmt"

	"github.com/edaniels/golog"

	"github.com/viamrobotics/gostream"
	"github.com/viamrobotics/gostream/codec"
)

// DefaultStreamConfig configures vpx as the encoder for a stream.
var DefaultStreamConfig gostream.StreamConfig

func init() {
	DefaultStreamConfig.VideoEncoderFactory = NewEncoderFactory(Version8)
}

// NewEncoderFactory returns a vpx factory for the given vpx codec.
func NewEncoderFactory(codecVersion Version) codec.VideoEncoderFactory {
	return &factory{codecVersion}
}

type factory struct {
	codecVersion Version
}

func (f *factory) New(width, height, keyFrameInterval int, logger golog.Logger) (codec.VideoEncoder, error) {
	return NewEncoder(f.codecVersion, width, height, keyFrameInterval, logger)
}

func (f *factory) MIMEType() string {
	switch f.codecVersion {
	case Version8:
		return "video/vp8"
	case Version9:
		return "video/vp9"
	default:
		panic(fmt.Errorf("unknown codec version %q", f.codecVersion))
	}
}
