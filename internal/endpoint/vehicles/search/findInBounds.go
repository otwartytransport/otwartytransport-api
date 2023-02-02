// package search

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/otwartytransport/otwartytransport-api/internal/parsers"
// 	"github.com/otwartytransport/otwartytransport-api/internal/vehicles/context/"
// 	"github.com/paulmach/orb"
// )


// type InBounds struct {
// 	MinLat float64 `query:"minLat"`
// 	MinLon float64 `query:"minLon"`
// 	MaxLat float64 `query:"maxLat"`
// 	MaxLon float64 `query:"maxLon"`
// }

// // FindVehicleInBounds godoc
// // @Summary Return API dynamicVehicles
// // @Produce json
// // @Success 200 {object} []orb.Pointer
// // @Router  /findInBounds [get]

// func FindVehicleInBounds(ctx *fiber.Ctx) error {
// 	bbox := new(InBounds)
// 	parsers.QueryParser(ctx, bbox)

// 	bound := orb.Bound{Min: orb.Point{bbox.MinLon, bbox.MinLat}, Max: orb.Point{bbox.MaxLon, bbox.MaxLat}}
// 	return ctx.JSON(context.GetQuadtree(ctx).InBound(nil, bound))
// }
