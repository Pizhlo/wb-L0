package data

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789"

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomString(n int) string {
	b := make([]byte, n)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func randomUIID() uuid.UUID {
	return uuid.New()
}

func randomTime() time.Time {
	return time.Now()
}

func randomTimeISO() int64 {
	return time.Now().Unix()
}

func randomChoise(slice []string) string {
	source := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	r1 := rand.New(source)

	n := r1.Int() % len(slice)

	return slice[n]
}

func randomPhone() string {
	phoneStr := "+7"

	phone := randomInt(9000000000, 9999999999)

	phoneStr += strconv.Itoa(phone)

	return phoneStr
}

func randomEmail(n int) string {
	login := randomString(n)

	login += "@mail.ru"

	return login
}
