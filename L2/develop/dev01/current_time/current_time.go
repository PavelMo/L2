package current_time

import (
	"github.com/beevik/ntp"
	"time"
)

// GetCurrTime Получаем время с NTP сервера и в случае ошибки возвращаем её
func GetCurrTime() (time.Time, error) {
	currTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return currTime, nil
}
