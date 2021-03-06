package material

import (
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/hitrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/pdf"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/ray"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/scatterrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/texture"
)

// Ensure interface compliance.
var _ Material = (*Isotropic)(nil)

// Isotropic represents an isotropic material.
type Isotropic struct {
	nonEmitter
	albedo texture.Texture
}

// NewIsotropic returns a new instances of the isotropic material.
func NewIsotropic(albedo texture.Texture) *Isotropic {
	return &Isotropic{
		albedo: albedo,
	}
}

// Scatter computes how the ray bounces off the surface of a diffuse material.
func (i *Isotropic) Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *scatterrecord.ScatterRecord, bool) {
	scattered := ray.New(hr.P(), randomInUnitSphere(), r.Time())
	attenuation := i.albedo.Value(hr.U(), hr.V(), hr.P())
	pdf := pdf.NewCosine(hr.Normal())
	scatterRecord := scatterrecord.New(nil, false, attenuation, pdf)
	return scattered, scatterRecord, true
}

// ScatteringPDF implements the probability distribution function for isotropic materials.
func (i *Isotropic) ScatteringPDF(r ray.Ray, hr *hitrecord.HitRecord, scattered ray.Ray) float64 {
	return 0
}
