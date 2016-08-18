package main

import (
    "fmt"
    "time"
    "math"
    "./models"
    "./utils"
    "github.com/jasonlvhit/gocron"
)

func taskUpdatePrice() {
    auctions, err := models.AllAuctions()
    if err != nil {
        fmt.Println(err);
    } else {
        for i, _ := range auctions {
            setPrice, _ := utils.CalculateStepPrice(auctions[i].AuctionInitialPrice, auctions[i].AuctionFinalPrice, auctions[i].AuctionSteps, auctions[i].AuctionDuration, auctions[i].AuctionStartTime)
            _, err := models.SetAuctionPriceByUUID(auctions[i].AuctionUUID, setPrice)
            if err != nil {
                fmt.Println(err);
            } else {
                fmt.Println("Updated Set Price of UUID ", auctions[i].AuctionUUID)
            }
        }
    }
}

func taskSetFinalPrice() {
    fmt.Println("Auction FINAL PRICE CRON")
    auctions := []models.Auction{}
    auctions, err := models.AllAuctions()
    if err != nil {
        fmt.Println(err);
    } else {
        for i, _ := range auctions {
            duration := time.Now().Sub(auctions[i].AuctionStartTime).Minutes()
            if int(math.Abs(float64(duration))) >= auctions[i].AuctionDuration {
                _, err := models.SetAuctionFinalPrice(auctions[i].AuctionUUID, auctions[i].AuctionSetPrice)
                if err != nil {
                    fmt.Println(err);
                } else {
                    fmt.Println("Updated Final Price of UUID ", auctions[i].AuctionUUID)
                }
            }
        }
    }
}

func main() {
    /*
        Two Schedule running Concurrently
        to pickup taskUpdatePrice & taskSetFinalPrice tasks
     */
    fmt.Println("CRON JOBS")

    updateSetPrice := gocron.NewScheduler()
    updateSetPrice.Every(1).Minute().Do(taskUpdatePrice)
    <- updateSetPrice.Start()

    setFinalPrice := gocron.NewScheduler()
    setFinalPrice.Every(1).Minute().Do(taskSetFinalPrice)
    <- setFinalPrice.Start()

}