package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type DynamicFeature struct {
	orb.Pointer
	Properties geojson.Properties
	ID         interface{}
}

func map2[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func LoadQuadtreeToContext(ctx *fiber.Ctx) error {
	var arr []geojson.Feature

	jsonFile, err := os.Open("./sample-data/dynamic.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &arr); err != nil {
		log.Fatal(err)
	}

	features := map2(arr, func(f geojson.Feature) DynamicFeature {
		feature := DynamicFeature{
			Pointer:    f.Geometry.(orb.Point),
			Properties: f.Properties,
			ID:         f.ID,
		}
		return feature
	})

	InitQuadTree(ctx, features)
	return ctx.Next()
}
