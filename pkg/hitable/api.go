// Package hitable implements the methods used to compute intersections between a ray and geometry.
package hitable

import (
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/aabb"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/hitrecord"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/material"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/ray"
	"github.com/flynn-nrg/ray-tracing-the-rest-of-your-life/pkg/vec3"
)

// Hitable defines the methods to compute ray/geometry operations.
type Hitable interface {
	Hit(r ray.Ray, tMin float64, tMax float64) (*hitrecord.HitRecord, material.Material, bool)
	BoundingBox(time0 float64, time1 float64) (*aabb.AABB, bool)
	PDFValue(o *vec3.Vec3Impl, v *vec3.Vec3Impl) float64
	Random(o *vec3.Vec3Impl) *vec3.Vec3Impl
}
