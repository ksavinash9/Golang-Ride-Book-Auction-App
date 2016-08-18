package utils
import (
    "time"
    "math"
    "strconv"
)

func GetInitialPrice(ridePrice string) string {
    intPrice := ToInt(ridePrice)
    val := math.Ceil(float64(intPrice)/float64(10))
    return ToString(val)
}

func GetEndTime(duration int) time.Time {
    now := time.Now()
    then := now.Add(time.Duration(duration) * time.Minute)
    return then
}

func CalculateStepPrice(auctionInitialPrice, auctionFinalPrice string, auctionSteps, auctionDuration int, auctionStartTime time.Time) (string, time.Time) {
    initialPrice,_ := strconv.ParseFloat(auctionInitialPrice, 64)
    finalPrice,_ := strconv.ParseFloat(auctionFinalPrice, 64)
    duration := time.Now().Sub(auctionStartTime).Minutes()

    if int(duration) >= auctionDuration {

        return auctionFinalPrice, time.Now().Add(time.Duration(60) * time.Minute)

    } else {

        stepPrice := math.Floor(initialPrice + (float64(duration)/float64(auctionDuration))*(finalPrice - initialPrice))

        return ToString(stepPrice), time.Now().Add(time.Duration(5) * time.Minute)
    }
}