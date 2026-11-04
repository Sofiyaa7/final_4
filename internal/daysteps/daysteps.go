package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	pars := strings.Split(data, ",")
	if len(pars) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}
	steps, err := strconv.Atoi(pars[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("колличество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(pars[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceM := float64(steps) * stepLength
	distanceKM := distanceM / mInKm

	calories, err := spentcalories.WalkingSpentCalories(
		steps, weight, height, duration,
	)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf(
		"Количесто шагов: %d. \n Дистанция составила%.2f км.\n Вы сожгли %.2f ккал.",
		steps,
		distanceKM,
		calories,
	)
}
