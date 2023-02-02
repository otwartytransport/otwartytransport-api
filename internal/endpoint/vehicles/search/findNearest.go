// package search

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/otwartytransport/otwartytransport-api/internal/parsers"
// 	"github.com/otwartytransport/otwartytransport-api/internal/vehicles/context/"
// 	"github.com/paulmach/orb"
// )

// type Nearest struct {
// 	Lat float64 `query:"lat"`
// 	Lon float64 `query:"lon"`
// }

// func FindNearestVehicle(ctx *fiber.Ctx) error {
// 	n := new(Nearest)
// 	parsers.QueryParser(ctx, n)

// 	return ctx.JSON(context.GetQuadtree(ctx).Find(orb.Point{n.Lon, n.Lat}))
// }