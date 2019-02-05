package types

import "fmt"

type WaterTemperatureStation struct {
	Metadata struct {
		Id   string
		Name string
		Lat  string
		Lon  string
	}
	Data [] struct {
		T string
		V string
		F string
	}
}

func (w WaterTemperatureStation) Format() {
	fmt.Printf("Id: %s\n"+
		"Name: %s\n"+
		"Lat: %s\n"+
		"Lon: %s\n\n", w.Metadata.Id, w.Metadata.Name, w.Metadata.Lat, w.Metadata.Lon)
}
