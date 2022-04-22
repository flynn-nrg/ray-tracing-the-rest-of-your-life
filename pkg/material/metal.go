package material

import (
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/hitrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/ray"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/scatterrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/vec3"
)

// Ensure interface compliance.
var _ Material = (*Metal)(nil)

// Metal represents metallic materials.
type Metal struct {
	albedo *vec3.Vec3Impl
	fuzz   float64
}

// NewMetal returns an instance of the metal material.
func NewMetal(albedo *vec3.Vec3Impl, fuzz float64) *Metal {
	return &Metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

// Scatter computes how the ray bounces off the surface of a metallic object.
func (m *Metal) Scatter(r ray.Ray, hr *hitrecord.HitRecord) (*ray.RayImpl, *scatterrecord.ScatterRecord, bool) {
	reflected := reflect(vec3.UnitVector(r.Direction()), hr.Normal())
	specular := ray.New(hr.P(), vec3.Add(reflected, vec3.ScalarMul(randomInUnitSphere(), m.fuzz)), r.Time())
	attenuation := m.albedo
	scatterRecord := scatterrecord.New(specular, true, attenuation, nil)
	return nil, scatterRecord, true
}

// Emitted returns black for metallic materials.
func (m *Metal) Emitted(_ ray.Ray, _ *hitrecord.HitRecord, _ float64, _ float64, _ *vec3.Vec3Impl) *vec3.Vec3Impl {
	return &vec3.Vec3Impl{}
}

// ScatteringPDF implements the probability distribution function for metals.
func (m *Metal) ScatteringPDF(r ray.Ray, hr *hitrecord.HitRecord, scattered ray.Ray) float64 {
	return 0
}
