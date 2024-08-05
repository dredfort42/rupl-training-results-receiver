package structs

import (
	"github.com/lib/pq"
)

// LastTrainingResult is a struct for JSON
type LastTrainingResult struct {
	Email      string `json:"email"`
	DeviceUUID string `json:"device_uuid"`
	StartTime  int64  `json:"session_start_time"`
	EndTime    int64  `json:"session_end_time"`
}

// RouteData is a struct for JSON and database
type RouteData struct {
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

// JSONLastResultTypeData is a struct for JSON
type JSONLastResultTypeData struct {
	Timestamp int64  `json:"timestamp"`
	Quantity  string `json:"quantity"`
}

// JSONLastTrainingResult is a struct for JSON
type JSONLastTrainingResult struct {
	Session             LastTrainingResult       `json:"session"`
	RouteData           []RouteData              `json:"route_data"`
	StepCount           []JSONLastResultTypeData `json:"step_count"`
	RunningPower        []JSONLastResultTypeData `json:"running_power"`
	VerticalOscillation []JSONLastResultTypeData `json:"vertical_oscillation"`
	EnergyBurned        []JSONLastResultTypeData `json:"energy_burned"`
	HeartRate           []JSONLastResultTypeData `json:"heart_rate"`
	StrideLength        []JSONLastResultTypeData `json:"stride_length"`
	GroundContactTime   []JSONLastResultTypeData `json:"ground_contact_time"`
	Speed               []JSONLastResultTypeData `json:"speed"`
	Distance            []JSONLastResultTypeData `json:"distance"`
	VO2Max              []JSONLastResultTypeData `json:"vo2_max"`
}

// DBResultDataInt is a struct for database
type DBResultDataInt struct {
	Timestamp int64
	Data      int
}

// DBResultDataFloat32 is a struct for database
type DBResultDataFloat32 struct {
	Timestamp int64
	Data      float32
}

// DBTrainingResult represents a session entry in the database
type DBTrainingResult struct {
	SessionUUID              string
	SessionStartTime         int64
	SessionEndTime           int64
	Email                    string
	DeviceUUID               string
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
