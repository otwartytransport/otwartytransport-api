package dynamicData

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
)

type Nearest struct {
	Lat float64 `query:"lat"`
	Lon float64 `query:"lon"`
}

func FindNearestVehicle(ctx *fiber.Ctx) error {
	n := new(Nearest)
	QueryParser(ctx, n)

	return ctx.JSON(GetQuadtree(ctx).Find(orb.Point{n.Lon, n.Lat}))
}
