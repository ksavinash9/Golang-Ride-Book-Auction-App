package utils
import (
    "encoding/base64"
    "encoding/hex"
    "io"
    "crypto/rand"
    "crypto/md5"
    "strconv"
    "encoding/json"
)

//md5 Util function - get md5sum of string
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

// GetGuid util for getting uuids
func GetGuid() string {
    b := make([]byte, 48)

    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
// Basic String Util for interfacing strings
func ToString(args ...interface{}) string {
    result := ""
    for _, arg := range args {
        switch val := arg.(type) {
        case int:
            result += strconv.Itoa(val)
        case string:
            result += val
        case float64:
            result += strconv.FormatFloat(val, 'f', 0, 64)
        }
    }
    return result
}

func ToInt(input string) int {
    i, _ := strconv.Atoi(input)
    return i
}

func GetJson(input interface{}) string {
    b, _ := json.Marshal(input)
    return string(b)
}
