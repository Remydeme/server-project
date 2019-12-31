package models

type Crs struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Position struct {
	Type              string  `json:"type"`
	Coordinates       Point   `json:"coordinate"`
	LocationAccuracy  float32 `json:"locationAccuracy"`
	LocationSource    string  `json:"locationSource"`
	Crs               Crs     `json:"crs"`
	Altitude          float32 `json:"altitude"`
	Speed             float32 `json:"speed"`
	DeviceOrientation int32   `json:"deviceOrientation"`
	GPSOrientation    int32   `json:"gpsOrientation"`
}
