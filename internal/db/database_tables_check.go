package db

// trainingSessionsTableCheck checks if the training results table exists, if not, it creates it
func trainingSessionsTableCheck() {
	query := `
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'route_data'
				) THEN 
					CREATE TYPE route_data AS (
						timestamp BIGINT,
						latitude DOUBLE PRECISION,
						longitude DOUBLE PRECISION,
						horizontal_accuracy DOUBLE PRECISION,
						altitude DOUBLE PRECISION,
						vertical_accuracy DOUBLE PRECISION,
						speed DOUBLE PRECISION,
						speed_accuracy DOUBLE PRECISION,
						course DOUBLE PRECISION,
						course_accuracy DOUBLE PRECISION
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'step_count'
				) THEN 
					CREATE TYPE step_count AS (
						timestamp BIGINT,
						steps_count INT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'running_power'
				) THEN 
					CREATE TYPE running_power AS (
						timestamp BIGINT,
						power_W INT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'vertical_oscillation'
				) THEN 
					CREATE TYPE vertical_oscillation AS (
						timestamp BIGINT,
						vertical_oscillation_cm FLOAT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'energy_burned'
				) THEN 
					CREATE TYPE energy_burned AS (
						timestamp BIGINT,
						energy_burned_kcal FLOAT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'heart_rate'
				) THEN 
					CREATE TYPE heart_rate AS (
						timestamp BIGINT,
						heart_rate_count_per_s FLOAT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'stride_length'
				) THEN 
					CREATE TYPE stride_length AS (
						timestamp BIGINT,
						stride_length_m FLOAT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'ground_contact_time'
				) THEN 
					CREATE TYPE ground_contact_time AS (
						timestamp BIGINT,
						ground_contact_time_ms INT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'speed'
				) THEN 
					CREATE TYPE speed AS (
						timestamp BIGINT,
						speed_m_per_s FLOAT
					);
				END IF;
				IF NOT EXISTS (
					SELECT 1 
					FROM pg_type 
					WHERE typname = 'distance'
				) THEN 
					CREATE TYPE distance AS (
						timestamp BIGINT,
						distance_m FLOAT
					);
				END IF;
			END $$;

			CREATE TABLE IF NOT EXISTS ` + DB.TableTrainingSessions + ` (
				session_uuid VARCHAR(255) NOT NULL PRIMARY KEY,
				session_start_time BIGINT NOT NULL,
				session_end_time BIGINT NOT NULL,
				email VARCHAR(255) NOT NULL,
				device_name VARCHAR(255) DEFAULT 'unknown',
				route_data route_data[] DEFAULT '{}'::route_data[] ,
				step_count step_count[] DEFAULT '{}'::step_count[] ,
				running_power running_power[] DEFAULT '{}'::running_power[] ,
				vertical_oscillation vertical_oscillation[] DEFAULT '{}'::vertical_oscillation[] ,
				energy_burned energy_burned[] DEFAULT '{}'::energy_burned[] ,
				heart_rate heart_rate[] DEFAULT '{}'::heart_rate[] ,
				stride_length stride_length[] DEFAULT '{}'::stride_length[] ,
				ground_contact_time ground_contact_time[] DEFAULT '{}'::ground_contact_time[] ,
				speed speed[] DEFAULT '{}'::speed[] ,
				distance distance[] DEFAULT '{}'::distance[] ,
				vo2max_mL_per_min_per_kg FLOAT ,
				avr_speed_m_per_s FLOAT ,
				avr_heart_rate_count_per_s FLOAT ,
				avr_power_w FLOAT ,
				avr_vertical_oscillation_cm FLOAT ,
				avr_ground_contact_time_ms FLOAT ,
				avr_stride_length_m FLOAT ,
				total_distance_m FLOAT ,
				total_steps_count INT ,
				total_energy_burned_kcal INT ,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
		`

	_, err := DB.Database.Exec(query)
	if err != nil {
		panic(err)
	}
}

// CheckTables checks if the tables exists, if not, it creates it
func tablesCheck() {
	trainingSessionsTableCheck()
}
