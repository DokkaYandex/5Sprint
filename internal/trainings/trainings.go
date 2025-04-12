package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrinvalidFormat = errors.New("Invalid Format")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("Incorrect amount of data", ErrinvalidFormat)
	}
	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("Incorrect number of steps", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("Negative number of steps")
	}
	t.TrainingType = slice[1]

	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("Invalid walk duration format", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("negative duration of the walk", ErrinvalidFormat)
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distanc := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	calorie := 0.0
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calorie, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("Calorie calculation error")
		}
	case "Бег":
		calorie, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("Calorie calculation error")
		}
	default:
		return "неизвестный тип тренировки", ErrinvalidFormat
	}
	str := fmt.Sprintf("Тип тренеровки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distanc, meanSpeed, calorie)
	return str, nil
}
