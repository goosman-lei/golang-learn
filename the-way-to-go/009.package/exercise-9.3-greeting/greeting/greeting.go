package greeting

import "time"

func IsMorning() bool {
    nowHour := time.Now().Hour()
    return nowHour >= 6 && nowHour <= 11
}
func IsAfternoon() bool {
    nowHour := time.Now().Hour()
    return nowHour >= 12 && nowHour <= 19
}
func IsEvening() bool {
    nowHour := time.Now().Hour()
    return nowHour >= 20 || nowHour <= 5
}
