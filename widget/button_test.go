package widget

import (
	"testing"

	internalevent "github.com/blizzy78/ebitenui/internal/event"
	"github.com/matryer/is"
)

func TestButton_PressedEvent_User(t *testing.T) {
	is := is.New(t)

	var eventArgs *ButtonPressedEventArgs

	b := newButton(t,
		ButtonOpts.WithPressedHandler(func(args *ButtonPressedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonPress(b, t)

	is.True(eventArgs != nil)
}

func TestButton_ReleasedEvent_User(t *testing.T) {
	is := is.New(t)

	var eventArgs *ButtonReleasedEventArgs

	b := newButton(t,
		ButtonOpts.WithReleasedHandler(func(args *ButtonReleasedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonRelease(b, t)

	is.True(eventArgs != nil)
}

func TestButton_ClickedEvent_User(t *testing.T) {
	is := is.New(t)

	var eventArgs *ButtonClickedEventArgs

	b := newButton(t,
		ButtonOpts.WithClickedHandler(func(args *ButtonClickedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonClick(b, t)

	is.True(eventArgs != nil)
}

func newButton(t *testing.T, opts ...ButtonOpt) *Button {
	t.Helper()

	b := NewButton(opts...)
	internalevent.ExecuteDeferredActions()
	render(b, t)
	return b
}
