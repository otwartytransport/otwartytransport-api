package context

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/quadtree"
)

var DynamicDataQuadtreeName = "dynamicVehicles"

func GetQuadtree(ctx *fiber.Ctx) *quadtree.Quadtree {
	qt := ctx.Locals(DynamicDataQuadtreeName)
	return qt.(*quadtree.Quadtree)
}

func InitQuadTree(ctx *fiber.Ctx, features []DynamicFeature) {
	qt := quadtree.New(orb.Bound{Min: orb.Point{0, 0}, Max: orb.Point{90, 90}})

	for _, feature := range features {
		qt.Add(feature)
	}

	ctx.Locals(DynamicDataQuadtreeName, qt)
}

func UpdateVehicle(ctx *fiber.Ctx, feature *UploadData) {
	qt := GetQuadtree(ctx)

	qt.Remove(orb.Point{}, func(p orb.Pointer) bool {
		f := p.(DynamicFeature)
		return f.ID == feature.ID
	})

	f := ToDynamicFeature(feature)

	qt.Add(f)
}
