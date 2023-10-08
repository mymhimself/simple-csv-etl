package source

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionWithHost(t *testing.T) {
	s := &rmqSource{}
	err := OptionWithHost("any")(s)
	require.NoError(t, err)
	require.Equal(t, "any", s.config.host)
}

func TestOptionWithQueueName(t *testing.T) {
	s := &rmqSource{}
	err := OptionWithQueueName("any")(s)
	require.NoError(t, err)
	require.Equal(t, "any", s.config.queueName)
}

func TestOptionWithConsumerName(t *testing.T) {
	s := &rmqSource{}
	err := OptionWithConsumerName("any")(s)
	require.NoError(t, err)
	require.Equal(t, "any", s.config.name)
}

func TestOptionWithArgs(t *testing.T) {
	s := &rmqSource{}
	err := OptionWithArgs(map[string]interface{}{
		"Key1": "Val1",
		"Key2": 2,
	})(s)
	require.NoError(t, err)
	require.NotEmpty(t, s.config.args)
}
