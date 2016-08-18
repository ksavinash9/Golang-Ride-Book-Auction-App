package models
import (
    "fmt"
    _"database/sql"
    "time"
    _"github.com/lib/pq"
)

func GetRide(rideID string) (Ride, error) {
    //Get Ride
    res := Ride{}


    var userID int
    var driverID int
    var ride_price string
    var ride_time time.Time
    var auctionUUID string

    err := db.QueryRow(`SELECT user_id, driver_id, ride_price, ride_time, auction_uuid from blacklane_rides where id = $1`, rideID).Scan(&userID, &driverID, &ride_price, &ride_time, &auctionUUID)
    if err == nil {
        res = Ride{UserID: userID, DriverID: driverID, RidePrice: ride_price, RideTime: ride_time, AuctionUUID: auctionUUID}
    }

    return res, err
}

func AllRides() ([]Ride, error) {
    //Get All Rides
    res := []Ride{}

    rows, err := db.Query(`SELECT user_id, driver_id, ride_price, ride_time, auction_uuid from blacklane_rides`)
    if err == nil {
        for rows.Next() {
            var userID int
            var driverID int
            var ride_price string
            var ride_time time.Time
            var auctionUUID string

            err = rows.Scan(&userID, &driverID, &ride_price, &ride_time, &auctionUUID)
            if err == nil {
                currentRide := Ride{UserID: userID, DriverID: driverID, RidePrice: ride_price, RideTime: ride_time, AuctionUUID: auctionUUID}
                res = append(res, currentRide)
            } else {
                return res, err
            }
        }
    } else {
        return res, err
    }

    return res, err
}

func InsertRide(ride Ride) (int, error) {
    //Create Ride
    var rideID int

    err := db.QueryRow(`INSERT INTO blacklane_rides(user_id, driver_id, ride_price, ride_time, auction_uuid) VALUES($1, $2, $3, $4, $5) RETURNING id`, ride.UserID, ride.DriverID, ride.RidePrice, ride.RideTime, ride.AuctionUUID).Scan(&rideID)

    if err != nil {
        return 0, err
    }

    fmt.Printf("Ride inserted ID %d", rideID)
    return rideID, err
}

func UpdateRide(ride Ride, rideID string) (int, error) {
    //Update Ride

    res, err := db.Exec(`UPDATE blacklane_rides set user_id=$1, driver_id=$2, ride_price=$3, ride_time=$4, auction_uuid=$5 where id=$6 RETURNING id`, ride.UserID, ride.DriverID, ride.RidePrice, ride.RideTime, ride.AuctionUUID, rideID)
    if err != nil {
        return 0, err
    }

    rowsUpdated, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsUpdated), err
}

func RemoveRide(rideID string) (int, error) {
    //Delete Ride
    res, err := db.Exec(`delete from blacklane_rides where id = $1`, rideID)
    if err != nil {
        return 0, err
    }

    rowsDeleted, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(rowsDeleted), nil
}