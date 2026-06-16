package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	s gocron.Scheduler
}

func New() (*Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{s: s}, nil
}

func (s *Scheduler) Add(
	name string,
	cronExpr string,
	task func() error,
) error {
	job, err := s.s.NewJob(
		gocron.CronJob(cronExpr, false),
		gocron.NewTask(task),
	)
	if err != nil {
		return err
	}

	next, _ := job.NextRun()
	fmt.Printf("[%s] next run: %s\n",
		name,
		next.Format(time.RFC822),
	)

	return nil
}

func (s *Scheduler) AddAndRunNow(
	name string,
	cronExpr string,
	task func() error,
	runNow bool,
) error {
	if err := s.Add(name, cronExpr, task); err != nil {
		return err
	}

	if runNow {
		go func() {
			if err := task(); err != nil {
				fmt.Printf("[%s] %v\n", name, err)
			}
		}()
	}

	return nil
}

func (s *Scheduler) Start() {
	s.s.Start()
}

func (s *Scheduler) Shutdown() error {
	return s.s.Shutdown()
}
