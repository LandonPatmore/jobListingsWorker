package types

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
	AmountOfReadings int
}

func (w WaterTemperatureStation) LogStruct() WaterTemperatureStation {
	w.AmountOfReadings = len(w.Data)
	w.Data = nil
	return w
}
