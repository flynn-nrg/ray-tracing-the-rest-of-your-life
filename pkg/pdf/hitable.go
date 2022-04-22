package pdf

import (
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/hitabletarget"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/vec3"
)

// Ensure interface compliance.
var _ PDF = (*Hitable)(nil)

// Hitable represents a hitable PDF.
type Hitable struct {
	o       *vec3.Vec3Impl
	hitable hitabletarget.HitableTarget
}

// NewHitable returns an instance of a hitable PDF.
func NewHitable(p hitabletarget.HitableTarget, origin *vec3.Vec3Impl) *Hitable {
	return &Hitable{
		o:       origin,
		hitable: p,
	}
}

func (h *Hitable) Value(direction *vec3.Vec3Impl) float64 {
	return h.hitable.PDFValue(h.o, direction)
}

func (h *Hitable) Generate() *vec3.Vec3Impl {
	return h.hitable.Random(h.o)
}
