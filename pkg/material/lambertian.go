package material

import (
	"math"

	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/hitrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/onb"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/pdf"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/ray"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/scatterrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/texture"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/vec3"
)

// Ensure interface compliance.
var _ Material = (*Lambertian)(nil)

// Lambertian represents a diffuse material.
type Lambertian struct {
	albedo texture.Texture
}

// NewLambertian returns an instance of the Lambert material.
func NewLambertian(albedo texture.Texture) *Lambertian {
	return &Lambertian{
		albedo: albedo,
	}
}

// Scatter computes how the ray bounces off the surface of a diffuse material.
func (l *Lambertian) Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *scatterrecord.ScatterRecord, bool) {
	uvw := onb.New()
	uvw.BuildFromW(hr.Normal())
	direction := uvw.Local(vec3.RandomCosineDirection())
	scattered := ray.New(hr.P(), vec3.UnitVector(direction), r.Time())
	albedo := l.albedo.Value(hr.U(), hr.V(), hr.P())
	pdf := pdf.NewCosine(hr.Normal())
	scatterRecord := scatterrecord.New(nil, false, albedo, pdf)
	return scattered, scatterRecord, true
}

// Emitted returns black for Lambertian materials.
func (l *Lambertian) Emitted(_ ray.Ray, _ *hitrecord.HitRecord, _ float64, _ float64, _ *vec3.Vec3Impl) *vec3.Vec3Impl {
	return &vec3.Vec3Impl{}
}

// ScatteringPDF implements the probability distribution function for diffuse materials.
func (l *Lambertian) ScatteringPDF(r ray.Ray, hr *hitrecord.HitRecord, scattered ray.Ray) float64 {
	cosine := vec3.Dot(hr.Normal(), vec3.UnitVector(scattered.Direction()))
	if cosine < 0 {
		cosine = 0
	}

	return cosine / math.Pi
}
