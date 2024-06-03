package validator

import (
	"exoplanet/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidExoplanetName(t *testing.T) {
	exoplanet := models.Exoplanet{
		Type:     "Terrestrial",
		Mass:     1.0,
		Radius:   1.0,
		Distance: 100,
	}
	err := ValidateExoplanet(exoplanet)
	fmt.Println(err.Error())
	assert.Error(t, err)
}
func TestInvalidRadius(t *testing.T) {
	exoplanet := models.Exoplanet{
		Name:        "Kepler-22b",
		Description: "First exoplanet found in the habitable zone",
		Type:        "Terrestrial",
		Mass:        2.4,
		Radius:      10.1,
		Distance:    600,
	}
	err := ValidateExoplanet(exoplanet)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestInvalidDistance(t *testing.T) {
	exoplanet := models.Exoplanet{
		Name:        "Proxima b",
		Description: "Closest known exoplanet",
		Type:        "Terrestrial",
		Mass:        1.3,
		Radius:      1.1,
		Distance:    1001,
	}
	err := ValidateExoplanet(exoplanet)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestInvalidMass(t *testing.T) {
	exoplanet:= models.Exoplanet{
		Name:        "Kepler-20e",
		Description: "An Earth-size exoplanet",
		Type:        "Terrestrial",
		Mass:        10.1,
		Radius:      1.0,
		Distance:    300,
	}
	err := ValidateExoplanet(exoplanet)
	fmt.Println(err.Error())
	assert.Error(t, err)
}


func TestInvalidType(t *testing.T) {
	exoplanet:= models.Exoplanet{
		Name:        "Unknown",
		Description: "An unknown type of planet",
		Type:        "Unknown",
		Mass:        1.0,
		Radius:      1.0,
		Distance:    100,
	}
	err := ValidateExoplanet(exoplanet)
	fmt.Println(err.Error())
	assert.Error(t, err)
}


