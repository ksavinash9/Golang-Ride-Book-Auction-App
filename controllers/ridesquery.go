package controllers
import (
    "net/http"
    "fmt"
    "time"
    "encoding/json"
    "../models"
    "../utils"
)

/* 
    Rides APIs Controller

    Rides Query - Handles the CRUD APIs for Rides Model
*/
func RidesQuery(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            ID := r.URL.Query().Get("id")
            if ID != "" {
                res, err := models.GetRide(ID)
                if err != nil {
                    fmt.Fprintln(w, err);
                } else {
                    fmt.Fprintln(w, utils.GetJson(res));
                }
            } else {
                res, err := models.AllRides()
                if err != nil {
                    fmt.Fprintln(w, err);
                } else {
                    fmt.Fprintln(w, utils.GetJson(res));
                }
            }
        case "POST":
            decoder := json.NewDecoder(r.Body)
            var ride models.Ride
            err := decoder.Decode(&ride)
            if err != nil {
                fmt.Println(err)
            }
            guid := utils.GetGuid()
            ride.AuctionUUID = guid
            ride.RideTime = time.Now()
            _, err = models.InsertRide(ride)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                var auctionUUID string = guid
                var auction_max_price string = ride.RidePrice
                var auction_inital_price string = utils.GetInitialPrice(ride.RidePrice)
                var auction_steps = ride.AuctionSteps
                if 1 > ride.AuctionSteps {
                    auction_steps = 1
                } 
                var auction_duration = ride.AuctionDuration
                if 1 > ride.AuctionDuration {
                    auction_duration = 60
                } 
                _, err = models.InsertAuction(models.Auction{AuctionUUID: auctionUUID, AuctionMaxPrice: auction_max_price, AuctionStartTime: time.Now(), AuctionInitialPrice: auction_inital_price, AuctionSetPrice: auction_inital_price, AuctionSteps: auction_steps, AuctionDuration: auction_duration})
                if err != nil {
                    fmt.Fprintln(w, err);
                }
                var rideOutput models.RideOutput
                rideOutput.AuctionUUID = guid
                rideOutput.AuctionStartTime = time.Now()
                rideOutput.AuctionEndTime = utils.GetEndTime(ride.AuctionDuration)
                rideOutput.AuctionInitialPrice = auction_inital_price
                rideOutput.AuctionFinalPrice = ride.RidePrice
                fmt.Fprintln(w, utils.GetJson(rideOutput));
            }
        case "PUT":
            ID := r.URL.Query().Get("id")
            decoder := json.NewDecoder(r.Body)
            var ride models.Ride   
            err := decoder.Decode(&ride)
            if err != nil {
                fmt.Println(err)
            }
            ride.RideTime = time.Now()
            res, err := models.UpdateRide(ride, ID)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Rows Updated `, res);
            }
        case "DELETE":
            ID := r.URL.Query().Get("id")
            res, err := models.RemoveRide(ID)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Rows Deleted `, res);
            }
        default:
    }
}

