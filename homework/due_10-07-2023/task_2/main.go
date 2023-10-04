package main

import "fmt"

type AutomobileSpecification struct {
	Name               string
	NDoors             int
	RimCenterBore      float32
	RimBoltPatter      string
	Petrol             string
	EngineVolumeLiters float32
	EngineName         string
	HorsePower         float32
	WeightKilogram     int
	HasSunroof         bool
	HasRoof            bool
	NSportFeatures     int
}

func isSportCar(spec *AutomobileSpecification) bool {
	temp_sport_features := spec.NSportFeatures
	if spec.NDoors < 4 {
		temp_sport_features += 1
	}

	if spec.HorsePower > 150 {
		temp_sport_features += 1
	}

	if spec.WeightKilogram < 1050 {
		temp_sport_features += 1
	}

	return temp_sport_features > 2
}

func main() {
	car1 := AutomobileSpecification{
		Name:               "ВАЗ2106",
		NDoors:             4,
		RimCenterBore:      58.5,
		RimBoltPatter:      "4x98",
		Petrol:             "A92",
		EngineVolumeLiters: 1.57,
		EngineName:         "2106",
		HorsePower:         71.5,
		WeightKilogram:     1035,
		HasSunroof:         false,
		HasRoof:            true,
	}

	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
	// did something change?
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
	// did something change?
	fmt.Printf("is %s a sport cat? %v\n", car1.Name, isSportCar(&car1))
}
