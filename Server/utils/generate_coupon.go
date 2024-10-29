package response


import (
    "math/rand"
    "strconv"
    "strings"
    "time"
)

func GenerateCouponCode() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomInt := r.Int63()
    base36String := strconv.FormatInt(randomInt, 36)
    if len(base36String) > 8 {
        base36String = base36String[2:8]
    } else if len(base36String) > 2 {
        base36String = base36String[2:]
    } else {
        base36String = "default"
    }
    return strings.ToUpper(base36String)
}
