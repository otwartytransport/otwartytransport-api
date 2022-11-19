package dynamicData

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	// "github.com/paulmach/orb/quadtree"
)

type UploadData struct {
	Pointer    orb.Point `json:"point"`
	Properties geojson.Properties
	ID         int `json:"id"`
}

func ToDynamicFeature(f *UploadData) DynamicFeature {
	return DynamicFeature{
		Pointer:    f.Pointer,
		Properties: f.Properties,
		ID:         f.ID,
	}
}

func uploadVehicle(ctx *fiber.Ctx) error {
	u := new(UploadData)

	if err := ctx.BodyParser(u); err != nil {
		return err
	}

	UpdateVehicle(ctx, u)

	return ctx.JSON("ds")
}
