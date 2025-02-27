package gostream

import (
	"github.com/edaniels/golog"

	"github.com/viamrobotics/gostream/codec"
)

// A StreamConfig describes how a Stream should be managed.
type StreamConfig struct {
	Name                string
	VideoEncoderFactory codec.VideoEncoderFactory
	AudioEncoderFactory codec.AudioEncoderFactory

	// TargetFrameRate will hint to the stream to try to maintain this frame rate.
	TargetFrameRate int

	Logger golog.Logger
}
