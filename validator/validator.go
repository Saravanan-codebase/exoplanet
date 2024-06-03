package validator

import (
	"fmt"

	"exoplanet/models"
)

func ValidateExoplanet(exoplanet models.Exoplanet) error {
	if exoplanet.Name == "" || exoplanet.Description == "" {
		return fmt.Errorf("name and description are required")
	}
	if exoplanet.Distance < 10 || exoplanet.Distance > 1000 {
		return fmt.Errorf("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return fmt.Errorf("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type == "Terrestrial" {
		if exoplanet.Mass < 0.1 || exoplanet.Mass > 10 {
			return fmt.Errorf("mass must be between 0.1 and 10 Earth-mass units for terrestrial planets")
		}
	} else if exoplanet.Type != "GasGiant" {
		return fmt.Errorf("type must be either GasGiant or Terrestrial")
	}

	return nil
}
