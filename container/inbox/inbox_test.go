package inbox_test

import (
	"doko/container/inbox"
	"fmt"
	"testing"
)

type Processor struct{}

func (slf *Processor) Invoke(envelopes []string) {
	for _, envelope := range envelopes {
		println(envelope)
	}
}

func TestInbox(t *testing.T) {
	box := inbox.NewInbox[string](inbox.WithInitialSize(21))
	defer box.Stop(true)
	box.Start(&Processor{})
	for i := 0; i < 100; i++ {
		if err := box.Send(fmt.Sprintf("envelope %d", i)); err != nil {
			break
		}
	}
}
