package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"nb_client/models"
	"nb_client/pkg/util"
	"strings"
)

func CreateScheduler(config models.AutoConfig, function interface{}) (s *gocron.Scheduler) {

	interval, periodo, err := util.GetIntervalePeriodo(config.IntervalTime)

	if err != nil {
		panic(interval)
	}
	s = gocron.NewScheduler()

	switch strings.ToLower(periodo) {
	case "m":
		newSchedulerMinutes(interval, s, function, config)
	case "s":
		newSchedulerSecond(interval, s, function, config)
	case "h":
		newSchedulerHours(interval, s, function, config)
	default:
		panic("Tempo não mapeado")
	}

	return
}

func newSchedulerSecond(interval uint64, s *gocron.Scheduler, function interface{}, args models.AutoConfig) {
	err := s.Every(interval).Second().Do(function, args)
	if err != nil {
		return
	}
}

func newSchedulerMinutes(interval uint64, s *gocron.Scheduler, function interface{}, args models.AutoConfig) {

	err := s.Every(interval).Minutes().Do(function, args)
	if err != nil {
		return
	}

}

func newSchedulerHours(interval uint64, s *gocron.Scheduler, function interface{}, args models.AutoConfig) {

	err := s.Every(interval).Hours().Do(function, args)
	if err != nil {
		return
	}

}
