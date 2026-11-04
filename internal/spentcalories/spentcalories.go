package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	pars := strings.Split(data, ",")
	if len(pars) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных")
	}
	steps, err := strconv.Atoi(pars[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("колличество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(pars[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, pars[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	hours := duration.Hours()
	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var distanceKm, speed, calories float64
	switch activity {
	case "бег":
		distanceKm = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "ходьба":
		distanceKm = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция:%.2f км.\n Скорость:%.2f км/ч \n Сожгли калорий: %.2f",
		activity,
		duration.Hours(),
		distanceKm,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неверно введены данные")
	}
	meanS := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanS * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("неверно введены данные")
	}
	meanS := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanS * durationInMinutes / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
