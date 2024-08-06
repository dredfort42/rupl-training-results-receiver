package structs

import (
	"github.com/lib/pq"
)

// SessionGeneralData is a struct for JSON
type SessionGeneralData struct {
	Email      string `json:"email"`
	DeviceName string `json:"device_name"`
	StartTime  int64  `json:"session_start_time"`
	EndTime    int64  `json:"session_end_time"`
}

// SessionRouteData is a struct for JSON and database
type SessionRouteData struct {
	Timestamp          int64   `json:"timestamp"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	Altitude           float64 `json:"altitude"`
	VerticalAccuracy   float64 `json:"vertical_accuracy"`
	Speed              float64 `json:"speed"`
	SpeedAccuracy      float64 `json:"speed_accuracy"`
	Course             float64 `json:"course"`
	CourseAccuracy     float64 `json:"course_accuracy"`
}

// SessionQuantityData is a struct for JSON
type SessionQuantityData struct {
	Timestamp int64  `json:"timestamp"`
	Quantity  string `json:"quantity"`
}

// TrainingSession is a struct for JSON
type TrainingSession struct {
	Session             SessionGeneralData    `json:"session"`
	RouteData           []SessionRouteData    `json:"route_data"`
	StepCount           []SessionQuantityData `json:"step_count"`
	RunningPower        []SessionQuantityData `json:"running_power"`
	VerticalOscillation []SessionQuantityData `json:"vertical_oscillation"`
	EnergyBurned        []SessionQuantityData `json:"energy_burned"`
	HeartRate           []SessionQuantityData `json:"heart_rate"`
	StrideLength        []SessionQuantityData `json:"stride_length"`
	GroundContactTime   []SessionQuantityData `json:"ground_contact_time"`
	Speed               []SessionQuantityData `json:"speed"`
	Distance            []SessionQuantityData `json:"distance"`
	VO2Max              []SessionQuantityData `json:"vo2_max"`
}

// DBSessionDataInt is a struct for database
type DBSessionDataInt struct {
	Timestamp int64
	Data      int
}

// DBSessionDataFloat32 is a struct for database
type DBSessionDataFloat32 struct {
	Timestamp int64
	Data      float32
}

// DBTrainingSession represents a session entry in the database
type DBTrainingSession struct {
	SessionUUID              string
	SessionStartTime         int64
	SessionEndTime           int64
	Email                    string
	DeviceName               string
	RouteData                pq.StringArray
	StepCount                pq.StringArray
	RunningPower             pq.StringArray
	VerticalOscillation      pq.StringArray
	EnergyBurned             pq.StringArray
	HeartRate                pq.StringArray
	StrideLength             pq.StringArray
	GroundContactTime        pq.StringArray
	Speed                    pq.StringArray
	Distance                 pq.StringArray
	VO2MaxMLPerMinPerKg      float32
	AvrSpeedMPerS            float32
	AvrHeartRateCountPerS    float32
	AvrPowerW                float32
	AvrVerticalOscillationCm float32
	AvrGroundContactTimeMs   float32
	AvrStrideLengthM         float32
	TotalDistanceM           float32
	TotalStepsCount          int
	TotalEnergyBurnedKcal    int
}
