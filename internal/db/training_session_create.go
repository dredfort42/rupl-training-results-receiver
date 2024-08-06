package db

import (
	"errors"
	s "training_sessions_receiver/internal/structs"
)

// TrainingSessionCreate creates a new session in the database
func TrainingSessionCreate(session s.DBTrainingSession) (err error) {
	if TrainingSessionExistsCheckByTime(session.SessionStartTime, session.SessionEndTime, session.Email) {
		return errors.New("session already exists")
	}

	query := `
		INSERT INTO ` + DB.TableTrainingSessions + ` (
				session_uuid,
				session_start_time,
				session_end_time,
				email,
				device_name,
				route_data,
				step_count,
				running_power,
				vertical_oscillation,
				energy_burned,
				heart_rate,
				stride_length,
				ground_contact_time,
				speed,
				distance,
				vo2max_mL_per_min_per_kg
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		)`

	_, err = DB.Database.Exec(query,
		session.SessionUUID,
		session.SessionStartTime,
		session.SessionEndTime,
		session.Email,
		session.DeviceName,
		session.RouteData,
		session.StepCount,
		session.RunningPower,
		session.VerticalOscillation,
		session.EnergyBurned,
		session.HeartRate,
		session.StrideLength,
		session.GroundContactTime,
		session.Speed,
		session.Distance,
		session.VO2MaxMLPerMinPerKg)
	if err != nil {
		return
	}

	err = calculateExtraData(session.SessionUUID)

	return
}
