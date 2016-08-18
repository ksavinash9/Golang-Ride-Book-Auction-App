package controllers
import (
    "net/http"
    "fmt"
    "encoding/json"
    "../models"
    "../utils"
)

/* 
Auction APIs Controller

AuctionQuery - Handles the CRUD APIs for Auction Model
*/
func AuctionsQuery(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            ID := r.URL.Query().Get("id")
            if ID != "" {
                res, err := models.GetAuction(ID)
                if err != nil {
                    fmt.Fprintln(w, err);
                } else {
                    fmt.Fprintln(w, utils.GetJson(res));
                }
            } else {
                res, err := models.AllAuctions()
                if err != nil {
                    fmt.Fprintln(w, err);
                } else {
                    fmt.Fprintln(w, utils.GetJson(res));
                }
            }
        case "POST":
            decoder := json.NewDecoder(r.Body)
            var auction models.Auction   
            err := decoder.Decode(&auction)
            if err != nil {
                fmt.Println(err)
            }
            res, err := models.InsertAuction(auction)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Row Created with ID `, res);
            }
        case "PUT":
            ID := r.URL.Query().Get("id")
            decoder := json.NewDecoder(r.Body)
            var auction models.Auction   
            err := decoder.Decode(&auction)
            if err != nil {
                fmt.Println(err)
            }
            res, err := models.UpdateAuction(auction, ID)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Rows Updated `, res);
            }
        case "DELETE":
            ID := r.URL.Query().Get("id")
            res, err := models.RemoveAuction(ID)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Rows Deleted `, res);
            }
        default:
    }
}

func AuctionUpdateSetPrice(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "PUT":
            ID := r.URL.Query().Get("id")
            setPrice := r.URL.Query().Get("price")
            decoder := json.NewDecoder(r.Body)
            fmt.Println(decoder)
            var auction models.Auction   
            err := decoder.Decode(&auction)
            if err != nil {
                fmt.Println(err)
            }
            res, err := models.SetAuctionPriceByID(ID, setPrice)
            if err != nil {
                fmt.Fprintln(w, err);
            } else {
                fmt.Fprintln(w, `Rows Updated `, res);
            }
        default:
    }
}

func AuctionGetSetPrice(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            ID := r.URL.Query().Get("id")
            if ID != "" {
                res, err := models.GetAuction(ID)
                if err != nil {
                    fmt.Fprintln(w, err);
                } else {
                    setPrice, stepExpiry := utils.CalculateStepPrice(res.AuctionInitialPrice, res.AuctionFinalPrice, res.AuctionSteps, res.AuctionDuration, res.AuctionStartTime)
                    auctionSetPrice := models.AuctionSetPrice{AuctionUUID: res.AuctionUUID, AuctionStepPrice: setPrice, AuctionStepExpiry: stepExpiry}
                    fmt.Fprintln(w, utils.GetJson(auctionSetPrice));
                }
            }
        default:
    }
}


