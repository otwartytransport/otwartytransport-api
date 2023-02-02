// package search

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/otwartytransport/otwartytransport-api/internal/parsers"
// 	"github.com/otwartytransport/otwartytransport-api/internal/vehicles/context/"
// 	"github.com/paulmach/orb"
// )

// type KNearest struct {
// 	Lat float64 `query:"lat"`
// 	Lon float64 `query:"lon"`
// 	K 	int		`query:"k"`
// }

// func FindKNearestVehicle(ctx *fiber.Ctx) error {
// 	n := new(KNearest)
// 	parsers.QueryParser(ctx, n)

// 	return ctx.JSON(context.GetQuadtree(ctx).KNearest(nil, orb.Point{n.Lon, n.Lat}, n.K))
// }