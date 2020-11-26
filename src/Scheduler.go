package main

import (
	"context"
)

// Scheduler takes a plan and it executes it.
type Scheduler struct {
	// stepCounter keep track of the number of steps exectued by the scheduler.
	// It is used for debug and logged out at the end of every execution.
	stepCounter int
}

// Execute accept an plan as input and it execute it until there are not anymore
// procedures to do
func (s *Scheduler) Execute(ctx context.Context, p Plan) error {
	// uuidGenerator := uuid.New()

	// logger := s.logger.With(zap.String("execution_id", uuidGenerator.String()))

	// start := time.Now()
	// if loggableP, ok := p.(Loggable); ok {
	// 	loggableP.WithLogger(logger)
	// }
	// logger.Info("Started execution plan " + p.Name())
	println("Starting execution of plan " + p.Name())

	s.stepCounter = 0
	for {
		steps, err := p.Create(ctx)
		if err != nil {
			// logger.Error(err.Error())
			println(err.Error())
			return err
		}
		if len(steps) == 0 {
			break
		}
		err = s.react(ctx, steps)
		if err != nil {
			// logger.Error(err.Error(), zap.String("execution_time", time.Since(start).String()), zap.Int("step_executed", s.stepCounter))
			println(err.Error())
			return err
		}
	}
	// logger.Info("Plan executed without errors.", zap.String("execution_time", time.Since(start).String()), zap.Int("step_executed", s.stepCounter))
	println("Plan executed with no errors!")
	return nil
}

// react is a recursive function that goes over all the steps and the one
// returned by previous steps until the plan does not return anymore steps
func (s *Scheduler) react(ctx context.Context, steps []Procedure) error {
	for _, step := range steps {
		s.stepCounter = s.stepCounter + 1
		// if loggableS, ok := step.(Loggable); ok {
		// 	loggableS.WithLogger(logger)
		// }
		innerSteps, err := step.Do(ctx)
		if err != nil {
			// logger.Error("Step failed.", zap.String("step", step.Name()), zap.Error(err))
			println("Step failed " + step.Name() + " " + err.Error())
			return err
		}
		if len(innerSteps) > 0 {
			if err := s.react(ctx, innerSteps); err != nil {
				return err
			}
		}
	}
	return nil
}
