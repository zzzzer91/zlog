package trace

import "github.com/sirupsen/logrus"

type Config struct {
	EnableLevels   []logrus.Level
	ErrorSpanLevel logrus.Level
}

type Option func(s *Config)

func WithEnableLevels(levels []logrus.Level) Option {
	return func(s *Config) {
		s.EnableLevels = levels
	}
}

func WithErrorSpanLevel(l logrus.Level) Option {
	return func(s *Config) {
		s.ErrorSpanLevel = l
	}
}
