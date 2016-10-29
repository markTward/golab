package main

const (
	configFile = "config.yaml"
)

type Config struct {
	Config struct {
		JobQueue struct {
			Number int
		}

		WorkerQueue struct {
			Number int
		}

		Worker struct {
			Number int
		}

		StatChan struct {
			Number int
		}

		Jobs struct {
			Number int
		}
		Testing struct {
			Duration int
		}
	}
}
