package widget

import (
	"testing"

	internalevent "github.com/blizzy78/ebitenui/internal/event"
	"github.com/matryer/is"
)

func TestGraphic_PreferredSize(t *testing.T) {
	is := is.New(t)

	i := newImageEmptyWithSize(47, 11, t)
	g := newGraphic(t,
		GraphicOpts.WithImage(i))
	w, h := g.PreferredSize()
	is.Equal(w, i.Bounds().Dx())
	is.Equal(h, i.Bounds().Dy())
}

func newGraphic(t *testing.T, opts ...GraphicOpt) *Graphic {
	t.Helper()

	g := NewGraphic(opts...)
	internalevent.ExecuteDeferredActions()
	render(g, t)
	return g
}
