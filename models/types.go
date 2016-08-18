package models

import "time"


//Ride represents a Ride object
type Ride struct {
	UserID          	int
	DriverID        	int
	RidePrice       	string
	RideTime 			time.Time
	AuctionUUID       	string
	AuctionSteps 		int
	AuctionDuration 	int
}

// Auction represent a Auction object
type Auction struct {
	AuctionUUID			string
	AuctionMaxPrice 	string
	AuctionStartTime    time.Time
	AuctionInitialPrice string
	AuctionSetPrice		string
	AuctionFinalPrice 	string
	AuctionSteps 		int
	AuctionDuration 	int
}

type RideOutput struct {
	AuctionUUID			string
	AuctionStartTime	time.Time
	AuctionEndTime		time.Time
	AuctionInitialPrice string
	AuctionFinalPrice	string
}

type AuctionSetPrice struct {
	AuctionUUID			string
	AuctionStepPrice	string
	AuctionStepExpiry	time.Time
}