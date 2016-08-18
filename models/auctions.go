package models
import (
    "fmt"
    _"database/sql"
    "time"
)

func GetAuction(auctionID string) (Auction, error) {
    //Retrieve
    res := Auction{}

    var auctionUUID             string
    var auction_max_price       string
    var auction_start_time      time.Time
    var auction_inital_price    string
    var auction_set_price       string
    var auction_final_price     string
    var auction_steps           int
    var auction_duration        int

    err := db.QueryRow(`SELECT auction_uuid, auction_max_price, auction_start_time, auction_initial_price, auction_set_price, auction_final_price, auction_steps, auction_duration from blacklane_auctions where id=$1`, auctionID).Scan(&auctionUUID, &auction_max_price, &auction_start_time, &auction_inital_price, &auction_set_price, &auction_final_price, &auction_steps, &auction_duration)
    if err == nil {
        res = Auction{AuctionUUID: auctionUUID, AuctionMaxPrice: auction_max_price, AuctionStartTime: auction_start_time, AuctionInitialPrice: auction_inital_price, AuctionSetPrice: auction_set_price, AuctionFinalPrice: auction_final_price, AuctionSteps: auction_steps, AuctionDuration: auction_duration}
    }

    return res, err
}

func AllAuctions() ([]Auction, error) {
    //Retrieve
    res := []Auction{}

    rows, err := db.Query(`SELECT auction_uuid, auction_max_price, auction_start_time, auction_initial_price, auction_set_price, auction_final_price, auction_steps, auction_duration from blacklane_auctions`)
    // defer rows.Close()
    if err == nil {
        for rows.Next() {
            fmt.Println(rows)
            var auctionUUID             string
            var auction_max_price       string
            var auction_start_time      time.Time
            var auction_inital_price    string
            var auction_set_price       string
            var auction_final_price     string
            var auction_steps           int
            var auction_duration        int

            err = rows.Scan(&auctionUUID, &auction_max_price, &auction_start_time, &auction_inital_price, &auction_set_price, &auction_final_price, &auction_steps, &auction_duration)
            if err == nil {
                currentAuction := Auction{AuctionUUID: auctionUUID, AuctionMaxPrice: auction_max_price, AuctionStartTime: auction_start_time, AuctionInitialPrice: auction_inital_price, AuctionSetPrice: auction_set_price, AuctionFinalPrice: auction_final_price, AuctionSteps: auction_steps, AuctionDuration: auction_duration}
                res = append(res, currentAuction)
            } else {
                return res, err
            }
        }
    } else {
        return res, err
    }

    return res, err
}

func InsertAuction(auction Auction) (int, error) {
    //Create
    var auctionID int

    err := db.QueryRow(`INSERT INTO blacklane_auctions(auction_uuid, auction_max_price, auction_start_time, auction_initial_price, auction_set_price, auction_final_price, auction_steps, auction_duration) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, auction.AuctionUUID, auction.AuctionMaxPrice, auction.AuctionStartTime, auction.AuctionInitialPrice, auction.AuctionSetPrice, auction.AuctionFinalPrice, auction.AuctionSteps, auction.AuctionDuration).Scan(&auctionID)

    if err != nil {
        return 0, err
    }

    fmt.Printf("Auction inserted ID %d", auctionID)
    return auctionID, err
}

func UpdateAuction(auction Auction, auctionID string) (int, error) {
    //Update

    res, err := db.Exec(`UPDATE blacklane_auctions set auction_uuid=$1, auction_max_price=$2, auction_initial_price=$3, auction_set_price=$4, auction_final_price=$5, auction_steps=$6, auction_duration=$7 where id=$8`, auction.AuctionUUID, auction.AuctionMaxPrice, auction.AuctionInitialPrice, auction.AuctionSetPrice, auction.AuctionFinalPrice, auction.AuctionSteps, auction.AuctionDuration, auctionID)
    if err != nil {
        return 0, err
    }

    rowsUpdated, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsUpdated), err
}

func SetAuctionPriceByID(auctionID , auctionSetPrice string) (int, error) {
    //Set Auction Price

    res, err := db.Exec(`UPDATE blacklane_auctions set auction_set_price=$1 where id=$2`, auctionSetPrice, auctionID)
    if err != nil {
        return 0, err
    }

    rowsUpdated, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsUpdated), err
}

func SetAuctionPriceByUUID(auctionUUID , auctionSetPrice string) (int, error) {
    //Set Auction Price

    res, err := db.Exec(`UPDATE blacklane_auctions set auction_set_price=$1 where auction_uuid=$2`, auctionSetPrice, auctionUUID)
    if err != nil {
        return 0, err
    }

    rowsUpdated, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsUpdated), err
}

func SetAuctionFinalPrice(auctionUUID , auctionFinalPrice string) (int, error) {
    //Set Auction Final Price

    res, err := db.Exec(`UPDATE blacklane_auctions set auction_final_price=$1 where auction_uuid=$2`, auctionFinalPrice, auctionUUID)
    if err != nil {
        return 0, err
    }

    rowsUpdated, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsUpdated), err
}

func RemoveAuction(auctionID string) (int, error) {
    //Delete

    res, err := db.Exec(`delete from blacklane_auctions where id = $1`, auctionID)
    if err != nil {
        return 0, err
    }

    rowsDeleted, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsDeleted), nil
}