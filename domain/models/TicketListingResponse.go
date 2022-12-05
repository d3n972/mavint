package models

import "time"

type TicketListingResponse struct {
	IsOfDetailedSearch bool `json:"isOfDetailedSearch"`
	Route              []struct {
		SameOfferId    int    `json:"sameOfferId"`
		OfferGroupCode string `json:"offerGroupCode"`
		OfferIdentity  string `json:"offerIdentity"`
		RouteServices  []struct {
			ListOrder                   string `json:"listOrder"`
			Description                 string `json:"description"`
			RestrictiveStartStationCode string `json:"restrictiveStartStationCode"`
			RestrictiveEndStationCode   string `json:"restrictiveEndStationCode"`
			Sign                        struct {
				FontName  string `json:"fontName"`
				Character string `json:"character"`
			} `json:"sign"`
			TrainStopKind interface{} `json:"trainStopKind"`
		} `json:"routeServices"`
		TransfersCount    int         `json:"transfersCount"`
		TravelRouteLength interface{} `json:"travelRouteLength"`
		Departure         struct {
			Time         time.Time   `json:"time"`
			TimeExpected time.Time   `json:"timeExpected"`
			TimeFact     time.Time   `json:"timeFact"`
			DelayMin     int         `json:"delayMin"`
			TimeZone     interface{} `json:"timeZone"`
		} `json:"departure"`
		Arrival struct {
			Time         time.Time   `json:"time"`
			TimeExpected time.Time   `json:"timeExpected"`
			TimeFact     time.Time   `json:"timeFact"`
			DelayMin     int         `json:"delayMin"`
			TimeZone     interface{} `json:"timeZone"`
		} `json:"arrival"`
		TravelTimeMin  string `json:"travelTimeMin"`
		Name           string `json:"name"`
		LastStation    string `json:"lastStation"`
		DepartureTrack struct {
			Name             interface{} `json:"name"`
			ChangedTrackName interface{} `json:"changedTrackName"`
		} `json:"departureTrack"`
		ArrivalTrack struct {
			Name             interface{} `json:"name"`
			ChangedTrackName interface{} `json:"changedTrackName"`
		} `json:"arrivalTrack"`
		Services []struct {
			ListOrder                   string `json:"listOrder"`
			Description                 string `json:"description"`
			RestrictiveStartStationCode string `json:"restrictiveStartStationCode"`
			RestrictiveEndStationCode   string `json:"restrictiveEndStationCode"`
			Sign                        struct {
				FontName  string `json:"fontName"`
				Character string `json:"character"`
			} `json:"sign"`
			TrainStopKind interface{} `json:"trainStopKind"`
		} `json:"services"`
		TravelClasses []struct {
			Name     string `json:"name"`
			Fullness int    `json:"fullness"`
			Price    struct {
				Amount                  float64 `json:"amount"`
				AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
				Currency                struct {
					Name    string `json:"name"`
					UicCode string `json:"uicCode"`
				} `json:"currency"`
			} `json:"price"`
		} `json:"travelClasses"`
		Details struct {
			Distance      float64     `json:"distance"`
			TrainFullName string      `json:"trainFullName"`
			Days          interface{} `json:"days"`
			Tickets       []struct {
				OfferId                  string        `json:"offerId"`
				ServiceId                string        `json:"serviceId"`
				ServerServiceInformation string        `json:"serverServiceInformation"`
				Name                     string        `json:"name"`
				Price1StClass            float64       `json:"price1stClass"`
				Price2NdClass            float64       `json:"price2ndClass"`
				Fullness                 int           `json:"fullness"`
				PassengerId              int           `json:"passengerId"`
				TakeOverModes            []string      `json:"takeOverModes"`
				OfferValidFrom           time.Time     `json:"offerValidFrom"`
				NamedAdditionals         []interface{} `json:"namedAdditionals"`
				OfferValidTo             time.Time     `json:"offerValidTo"`
				ClientDiscounts          []struct {
					DiscountScale float64 `json:"discountScale"`
					GrossPrice    struct {
						AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
						Amount                  float64 `json:"amount"`
						Currency                struct {
							Key       string `json:"key"`
							Name      string `json:"name"`
							IsDefault bool   `json:"isDefault"`
						} `json:"currency"`
					} `json:"grossPrice"`
					NetPrice struct {
						AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
						Amount                  float64 `json:"amount"`
						Currency                struct {
							Key       string `json:"key"`
							Name      string `json:"name"`
							IsDefault bool   `json:"isDefault"`
						} `json:"currency"`
					} `json:"netPrice"`
					TakeoverMode   *int        `json:"takeoverMode"`
					PaymentMode    interface{} `json:"paymentMode"`
					DiscountReason string      `json:"discountReason"`
				} `json:"clientDiscounts"`
				CustomerCountDiscountName string `json:"customerCountDiscountName"`
				CarClassNumber            string `json:"carClassNumber"`
				CarClassIndependent       bool   `json:"carClassIndependent"`
				PlaceReservationNeeded    bool   `json:"placeReservationNeeded"`
				QuotaReservationNeeded    bool   `json:"quotaReservationNeeded"`
				AllowedInvoiceKind        int    `json:"allowedInvoiceKind"`
				DiscountedGrossPrice      struct {
					AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
					Amount                  float64 `json:"amount"`
					Currency                struct {
						Key       string `json:"key"`
						Name      string `json:"name"`
						IsDefault bool   `json:"isDefault"`
					} `json:"currency"`
				} `json:"discountedGrossPrice"`
				GrossPrice struct {
					AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
					Amount                  float64 `json:"amount"`
					Currency                struct {
						Key       string `json:"key"`
						Name      string `json:"name"`
						IsDefault bool   `json:"isDefault"`
					} `json:"currency"`
				} `json:"grossPrice"`
				GrossUnitPrice struct {
					AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
					Amount                  float64 `json:"amount"`
					Currency                struct {
						Key       string `json:"key"`
						Name      string `json:"name"`
						IsDefault bool   `json:"isDefault"`
					} `json:"currency"`
				} `json:"grossUnitPrice"`
				GrossPriceExchanged struct {
					AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
					Amount                  float64 `json:"amount"`
					Currency                struct {
						Key       string      `json:"key"`
						Name      string      `json:"name"`
						IsDefault interface{} `json:"isDefault"`
					} `json:"currency"`
				} `json:"grossPriceExchanged"`
				GrossUnitPriceExchanged struct {
					AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
					Amount                  float64 `json:"amount"`
					Currency                struct {
						Key       string      `json:"key"`
						Name      string      `json:"name"`
						IsDefault interface{} `json:"isDefault"`
					} `json:"currency"`
				} `json:"grossUnitPriceExchanged"`
				IsGroup                  bool   `json:"isGroup"`
				Direction                int    `json:"direction"`
				DirectionDescription     string `json:"directionDescription"`
				CustomerTypeDiscountName string `json:"customerTypeDiscountName"`
				DiscountName             string `json:"discountName"`
				TrainDependent           bool   `json:"trainDependent"`
				Refundable               bool   `json:"refundable"`
				CacheRenewParams         struct {
					Reserved            bool      `json:"reserved"`
					ReservationFixed    bool      `json:"reservationFixed"`
					ReservationCode     string    `json:"reservationCode"`
					AllowedOvertime     int       `json:"allowedOvertime"`
					LastReservationTime time.Time `json:"lastReservationTime"`
					PaymentDeadline     time.Time `json:"paymentDeadline"`
					ReservationID       string    `json:"reservationID"`
				} `json:"cacheRenewParams"`
				ClientDiscount []struct {
					DiscountScale float64 `json:"discountScale"`
					GrossPrice    struct {
						AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
						Amount                  float64 `json:"amount"`
						Currency                struct {
							Key       string `json:"key"`
							Name      string `json:"name"`
							IsDefault bool   `json:"isDefault"`
						} `json:"currency"`
					} `json:"grossPrice"`
					NetPrice struct {
						AmountInDefaultCurrency float64 `json:"amountInDefaultCurrency"`
						Amount                  float64 `json:"amount"`
						Currency                struct {
							Key       string `json:"key"`
							Name      string `json:"name"`
							IsDefault bool   `json:"isDefault"`
						} `json:"currency"`
					} `json:"netPrice"`
					TakeoverMode   *int        `json:"takeoverMode"`
					PaymentMode    interface{} `json:"paymentMode"`
					DiscountReason string      `json:"discountReason"`
				} `json:"clientDiscount"`
				IsInternational bool `json:"isInternational"`
				PlaceTicket     struct {
					TicketId                       *int        `json:"ticketId"`
					CertificateID                  interface{} `json:"certificateID"`
					PlaceReservationServices       interface{} `json:"placeReservationServices"`
					NearCarriageNumber             interface{} `json:"nearCarriageNumber"`
					NearSeatPosition               interface{} `json:"nearSeatPosition"`
					ReservationType                int         `json:"reservationType"`
					ReservationKind                *string     `json:"reservationKind"`
					PlaceReservationServicesFilled bool        `json:"placeReservationServicesFilled"`
					UicReservationCode             interface{} `json:"uicReservationCode"`
					CarriageNumber                 interface{} `json:"carriageNumber"`
					SeatPosition                   interface{} `json:"seatPosition"`
					Location                       interface{} `json:"location"`
					CustomerTypeDiscountName       *string     `json:"customerTypeDiscountName"`
					IsGroup                        bool        `json:"isGroup"`
					ForGrouppedFareTicket          bool        `json:"forGrouppedFareTicket"`
					Train                          *struct {
						AggregatedServiceIds        []string    `json:"aggregatedServiceIds"`
						Name                        string      `json:"name"`
						SeatReservationCode         string      `json:"seatReservationCode"`
						Code                        string      `json:"code"`
						CompanyCode                 string      `json:"companyCode"`
						Route                       interface{} `json:"route"`
						StartStationReservationCode string      `json:"startStationReservationCode"`
						EndStationReservationCode   string      `json:"endStationReservationCode"`
						StartStation                struct {
							Id                            int    `json:"id"`
							IsAlias                       bool   `json:"isAlias"`
							Name                          string `json:"name"`
							Code                          string `json:"code"`
							BaseCode                      string `json:"baseCode"`
							IsInternational               bool   `json:"isInternational"`
							CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
							Country                       string `json:"country"`
							CoutryIso                     string `json:"coutryIso"`
							IsIn1081                      bool   `json:"isIn108_1"`
						} `json:"startStation"`
						EndStation struct {
							Id                            int    `json:"id"`
							IsAlias                       bool   `json:"isAlias"`
							Name                          string `json:"name"`
							Code                          string `json:"code"`
							BaseCode                      string `json:"baseCode"`
							IsInternational               bool   `json:"isInternational"`
							CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
							Country                       string `json:"country"`
							CoutryIso                     string `json:"coutryIso"`
							IsIn1081                      bool   `json:"isIn108_1"`
						} `json:"endStation"`
						StartDate        time.Time `json:"startDate"`
						OrigStartStation struct {
							Id                            int    `json:"id"`
							IsAlias                       bool   `json:"isAlias"`
							Name                          string `json:"name"`
							Code                          string `json:"code"`
							BaseCode                      string `json:"baseCode"`
							IsInternational               bool   `json:"isInternational"`
							CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
							Country                       string `json:"country"`
							CoutryIso                     string `json:"coutryIso"`
							IsIn1081                      bool   `json:"isIn108_1"`
						} `json:"origStartStation"`
						OrigEndStation struct {
							Id                            int    `json:"id"`
							IsAlias                       bool   `json:"isAlias"`
							Name                          string `json:"name"`
							Code                          string `json:"code"`
							BaseCode                      string `json:"baseCode"`
							IsInternational               bool   `json:"isInternational"`
							CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
							Country                       string `json:"country"`
							CoutryIso                     string `json:"coutryIso"`
							IsIn1081                      bool   `json:"isIn108_1"`
						} `json:"origEndStation"`
						Start           time.Time `json:"start"`
						VirtualStart    bool      `json:"virtualStart"`
						Arrive          time.Time `json:"arrive"`
						VirtualArrive   bool      `json:"virtualArrive"`
						Distance        float64   `json:"distance"`
						ClosedTrackway  bool      `json:"closedTrackway"`
						FullName        string    `json:"fullName"`
						FullNameAndType string    `json:"fullNameAndType"`
						Kinds           []struct {
							Name                string `json:"name"`
							SortName            string `json:"sortName"`
							Code                string `json:"code"`
							Priority            int    `json:"priority"`
							BackgrouColorCode   string `json:"backgrouColorCode"`
							ForegroundColorCode string `json:"foregroundColorCode"`
							Sign                struct {
								FontName  string `json:"fontName"`
								Character string `json:"character"`
							} `json:"sign"`
							StartStation struct {
								Id                            int         `json:"id"`
								IsAlias                       bool        `json:"isAlias"`
								Name                          interface{} `json:"name"`
								Code                          interface{} `json:"code"`
								BaseCode                      interface{} `json:"baseCode"`
								IsInternational               bool        `json:"isInternational"`
								CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
								CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
								Country                       interface{} `json:"country"`
								CoutryIso                     interface{} `json:"coutryIso"`
								IsIn1081                      bool        `json:"isIn108_1"`
							} `json:"startStation"`
							EndStation struct {
								Id                            int         `json:"id"`
								IsAlias                       bool        `json:"isAlias"`
								Name                          interface{} `json:"name"`
								Code                          interface{} `json:"code"`
								BaseCode                      interface{} `json:"baseCode"`
								IsInternational               bool        `json:"isInternational"`
								CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
								CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
								Country                       interface{} `json:"country"`
								CoutryIso                     interface{} `json:"coutryIso"`
								IsIn1081                      bool        `json:"isIn108_1"`
							} `json:"endStation"`
						} `json:"kinds"`
						KindsToDisplay interface{} `json:"kindsToDisplay"`
						Kind           struct {
							Name                string `json:"name"`
							SortName            string `json:"sortName"`
							Code                string `json:"code"`
							Priority            int    `json:"priority"`
							BackgrouColorCode   string `json:"backgrouColorCode"`
							ForegroundColorCode string `json:"foregroundColorCode"`
							Sign                struct {
								FontName  string `json:"fontName"`
								Character string `json:"character"`
							} `json:"sign"`
							StartStation struct {
								Id                            int         `json:"id"`
								IsAlias                       bool        `json:"isAlias"`
								Name                          interface{} `json:"name"`
								Code                          interface{} `json:"code"`
								BaseCode                      interface{} `json:"baseCode"`
								IsInternational               bool        `json:"isInternational"`
								CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
								CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
								Country                       interface{} `json:"country"`
								CoutryIso                     interface{} `json:"coutryIso"`
								IsIn1081                      bool        `json:"isIn108_1"`
							} `json:"startStation"`
							EndStation struct {
								Id                            int         `json:"id"`
								IsAlias                       bool        `json:"isAlias"`
								Name                          interface{} `json:"name"`
								Code                          interface{} `json:"code"`
								BaseCode                      interface{} `json:"baseCode"`
								IsInternational               bool        `json:"isInternational"`
								CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
								CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
								Country                       interface{} `json:"country"`
								CoutryIso                     interface{} `json:"coutryIso"`
								IsIn1081                      bool        `json:"isIn108_1"`
							} `json:"endStation"`
						} `json:"kind"`
						Services []struct {
							ListOrder                   string `json:"listOrder"`
							Description                 string `json:"description"`
							RestrictiveStartStationCode string `json:"restrictiveStartStationCode"`
							RestrictiveEndStationCode   string `json:"restrictiveEndStationCode"`
							Sign                        struct {
								FontName  string `json:"fontName"`
								Character string `json:"character"`
							} `json:"sign"`
							TrainStopKind interface{} `json:"trainStopKind"`
						} `json:"services"`
						ActualOrEstimatedStart  interface{} `json:"actualOrEstimatedStart"`
						ActualOrEstimatedArrive interface{} `json:"actualOrEstimatedArrive"`
						HavarianInfok           struct {
							AktualisKeses float64     `json:"aktualisKeses"`
							KesesiOk      interface{} `json:"kesesiOk"`
							HavariaInfo   interface{} `json:"havariaInfo"`
							UzletiInfo    interface{} `json:"uzletiInfo"`
							KesesInfo     string      `json:"kesesInfo"`
						} `json:"havarianInfok"`
						DirectTrains         interface{} `json:"directTrains"`
						CarrierTrains        interface{} `json:"carrierTrains"`
						StartTrack           interface{} `json:"startTrack"`
						EndTrack             interface{} `json:"endTrack"`
						JeEszkozAlapId       float64     `json:"jeEszkozAlapId"`
						FullType             string      `json:"fullType"`
						FullShortType        string      `json:"fullShortType"`
						FullNameAndPiktogram struct {
							Collection string `json:"(Collection)"`
						} `json:"fullNameAndPiktogram"`
						Footer           interface{}    `json:"footer"`
						ViszonylatiJel   ViszonylatiJel `json:"viszonylatiJel"`
						ViszonylatObject struct {
							StartStationCode  string      `json:"startStationCode"`
							StartTime         time.Time   `json:"startTime"`
							StartTimeZone     string      `json:"startTimeZone"`
							EndStationCode    string      `json:"endStationCode"`
							EndTime           time.Time   `json:"endTime"`
							EndTimeZone       string      `json:"endTimeZone"`
							TravelTime        float64     `json:"travelTime"`
							StartTrack        interface{} `json:"startTrack"`
							EndTrack          interface{} `json:"endTrack"`
							InnerStationCodes []string    `json:"innerStationCodes"`
						} `json:"viszonylatObject"`
						Description    interface{} `json:"description"`
						SameCar        bool        `json:"sameCar"`
						StartTimeZone  interface{} `json:"startTimeZone"`
						ArriveTimeZone interface{} `json:"arriveTimeZone"`
						TrainId        string      `json:"trainId"`
					} `json:"train"`
					TrainCode                      *string     `json:"trainCode"`
					StartStationCode               *string     `json:"startStationCode"`
					EndStationCode                 *string     `json:"endStationCode"`
					StartTime                      interface{} `json:"startTime"`
					ArriveTime                     interface{} `json:"arriveTime"`
					SpecialisHelyAdatok            interface{} `json:"specialisHelyAdatok"`
					ServiceCode                    *string     `json:"serviceCode"`
					TrainCategory                  interface{} `json:"trainCategory"`
					VerificationNumber             interface{} `json:"verificationNumber"`
					HkGuid                         string      `json:"hkGuid"`
					Tarifa                         interface{} `json:"tarifa"`
					IsInternational                bool        `json:"isInternational"`
					KapcsoltHelyfoglalasID         interface{} `json:"kapcsoltHelyfoglalasID"`
					KapcsoltAjanlat                interface{} `json:"kapcsoltAjanlat"`
					GlobalDij                      bool        `json:"globalDij"`
					VacantPlaceCount               *int        `json:"vacantPlaceCount"`
					VacantPlaceCountUpdateTime     interface{} `json:"vacantPlaceCountUpdateTime"`
					QuotaReservationHandlingParams *struct {
						Reserved            bool        `json:"reserved"`
						ReservationFixed    bool        `json:"reservationFixed"`
						ReservationCode     interface{} `json:"reservationCode"`
						AllowedOvertime     int         `json:"allowedOvertime"`
						LastReservationTime interface{} `json:"lastReservationTime"`
						PaymentDeadline     time.Time   `json:"paymentDeadline"`
						ReservationID       interface{} `json:"reservationID"`
					} `json:"quotaReservationHandlingParams"`
					Reserved            bool        `json:"reserved"`
					ReservationFixed    bool        `json:"reservationFixed"`
					ReservationCode     interface{} `json:"reservationCode"`
					AllowedOvertime     int         `json:"allowedOvertime"`
					LastReservationTime interface{} `json:"lastReservationTime"`
					PaymentDeadline     time.Time   `json:"paymentDeadline"`
					ReservationID       interface{} `json:"reservationID"`
				} `json:"placeTicket"`
				PlaceTicketSerialized *string `json:"placeTicketSerialized"`
				TicketReferenceCode   string  `json:"ticketReferenceCode"`
				StartStation          struct {
					Id                            int    `json:"id"`
					IsAlias                       bool   `json:"isAlias"`
					Name                          string `json:"name"`
					Code                          string `json:"code"`
					BaseCode                      string `json:"baseCode"`
					IsInternational               bool   `json:"isInternational"`
					CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
					CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
					Country                       string `json:"country"`
					CoutryIso                     string `json:"coutryIso"`
					IsIn1081                      bool   `json:"isIn108_1"`
				} `json:"startStation"`
				TouchedStations []struct {
					Id                            int         `json:"id"`
					IsAlias                       bool        `json:"isAlias"`
					Name                          string      `json:"name"`
					Code                          string      `json:"code"`
					BaseCode                      interface{} `json:"baseCode"`
					IsInternational               bool        `json:"isInternational"`
					CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
					CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
					Country                       interface{} `json:"country"`
					CoutryIso                     interface{} `json:"coutryIso"`
					IsIn1081                      bool        `json:"isIn108_1"`
				} `json:"touchedStations"`
				EndStation struct {
					Id                            int    `json:"id"`
					IsAlias                       bool   `json:"isAlias"`
					Name                          string `json:"name"`
					Code                          string `json:"code"`
					BaseCode                      string `json:"baseCode"`
					IsInternational               bool   `json:"isInternational"`
					CanUseForOfferRequest         bool   `json:"canUseForOfferRequest"`
					CanUseForPessengerInformation bool   `json:"canUseForPessengerInformation"`
					Country                       string `json:"country"`
					CoutryIso                     string `json:"coutryIso"`
					IsIn1081                      bool   `json:"isIn108_1"`
				} `json:"endStation"`
				Amount      int     `json:"amount"`
				CompanyCode *string `json:"companyCode"`
				Distance    float64 `json:"distance"`
			} `json:"tickets"`
			Routes []struct {
				ServiceIds   []int `json:"serviceIds"`
				TrainDetails struct {
					ViszonylatiJel struct {
						PiktogramFullName interface{} `json:"piktogramFullName"`
						FontSzinKod       *string     `json:"fontSzinKod"`
						HatterSzinKod     *string     `json:"hatterSzinKod"`
					} `json:"viszonylatiJel"`
					TrainKind struct {
						Name                string  `json:"name"`
						SortName            *string `json:"sortName"`
						Code                string  `json:"code"`
						Priority            int     `json:"priority"`
						BackgrouColorCode   string  `json:"backgrouColorCode"`
						ForegroundColorCode string  `json:"foregroundColorCode"`
						Sign                struct {
							FontName  *string `json:"fontName"`
							Character *string `json:"character"`
						} `json:"sign"`
						StartStation struct {
							Id                            int         `json:"id"`
							IsAlias                       bool        `json:"isAlias"`
							Name                          interface{} `json:"name"`
							Code                          interface{} `json:"code"`
							BaseCode                      interface{} `json:"baseCode"`
							IsInternational               bool        `json:"isInternational"`
							CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
							Country                       interface{} `json:"country"`
							CoutryIso                     interface{} `json:"coutryIso"`
							IsIn1081                      bool        `json:"isIn108_1"`
						} `json:"startStation"`
						EndStation struct {
							Id                            int         `json:"id"`
							IsAlias                       bool        `json:"isAlias"`
							Name                          interface{} `json:"name"`
							Code                          interface{} `json:"code"`
							BaseCode                      interface{} `json:"baseCode"`
							IsInternational               bool        `json:"isInternational"`
							CanUseForOfferRequest         bool        `json:"canUseForOfferRequest"`
							CanUseForPessengerInformation bool        `json:"canUseForPessengerInformation"`
							Country                       interface{} `json:"country"`
							CoutryIso                     interface{} `json:"coutryIso"`
							IsIn1081                      bool        `json:"isIn108_1"`
						} `json:"endStation"`
					} `json:"trainKind"`
					Type                   string  `json:"type"`
					Name                   *string `json:"name"`
					TrainNumber            string  `json:"trainNumber"`
					TrainId                string  `json:"trainId"`
					JeId                   string  `json:"jeId"`
					KozlekedesiTarsasagKod string  `json:"kozlekedesiTarsasagKod"`
				} `json:"trainDetails"`
				Departure struct {
					Time         time.Time `json:"time"`
					TimeExpected time.Time `json:"timeExpected"`
					TimeFact     time.Time `json:"timeFact"`
					DelayMin     int       `json:"delayMin"`
					TimeZone     string    `json:"timeZone"`
				} `json:"departure"`
				Arrival struct {
					Time         time.Time `json:"time"`
					TimeExpected time.Time `json:"timeExpected"`
					TimeFact     time.Time `json:"timeFact"`
					DelayMin     int       `json:"delayMin"`
					TimeZone     string    `json:"timeZone"`
				} `json:"arrival"`
				Services struct {
					Train []struct {
						ListOrder                   string `json:"listOrder"`
						Description                 string `json:"description"`
						RestrictiveStartStationCode string `json:"restrictiveStartStationCode"`
						RestrictiveEndStationCode   string `json:"restrictiveEndStationCode"`
						Sign                        struct {
							FontName  string `json:"fontName"`
							Character string `json:"character"`
						} `json:"sign"`
						TrainStopKind interface{} `json:"trainStopKind"`
					} `json:"train"`
					Station []struct {
						ListOrder                   interface{} `json:"listOrder"`
						Description                 string      `json:"description"`
						RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
						RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
						Sign                        struct {
							FontName  string `json:"fontName"`
							Character string `json:"character"`
						} `json:"sign"`
						TrainStopKind interface{} `json:"trainStopKind"`
					} `json:"station"`
				} `json:"services"`
				DepartureTrack struct {
					Name             interface{} `json:"name"`
					ChangedTrackName interface{} `json:"changedTrackName"`
				} `json:"departureTrack"`
				ArrivalTrack struct {
					Name             string `json:"name"`
					ChangedTrackName string `json:"changedTrackName"`
				} `json:"arrivalTrack"`
				SameCar       bool `json:"sameCar"`
				TravelClasses []struct {
					Name     string      `json:"name"`
					Fullness int         `json:"fullness"`
					Price    interface{} `json:"price"`
				} `json:"travelClasses"`
				StartStation struct {
					Name          string      `json:"name"`
					Code          string      `json:"code"`
					Coordinates   interface{} `json:"coordinates"`
					ArrivalTime   *time.Time  `json:"arrivalTime"`
					DepartureTime *time.Time  `json:"departureTime"`
				} `json:"startStation"`
				DestionationStation struct {
					Name          string      `json:"name"`
					Code          string      `json:"code"`
					Coordinates   interface{} `json:"coordinates"`
					ArrivalTime   string      `json:"arrivalTime"`
					DepartureTime *time.Time  `json:"departureTime"`
				} `json:"destionationStation"`
				TouchedStationsString          string        `json:"touchedStationsString"`
				TouchedStations                []interface{} `json:"touchedStations"`
				HavariaInfo                    interface{}   `json:"havariaInfo"`
				Distance                       float64       `json:"distance"`
				Description                    interface{}   `json:"description"`
				MasodlagosEszkozSzolgaltatasok []struct {
					Services               []string `json:"services"`
					EszkozSzam             string   `json:"eszkozSzam"`
					KozlekedesiTarsasagKod string   `json:"kozlekedesiTarsasagKod"`
				} `json:"masodlagosEszkozSzolgaltatasok"`
				TravelTime string `json:"travelTime"`
			} `json:"routes"`
			HasPlaceTicket              bool `json:"hasPlaceTicket"`
			PlaceTicketDutyDeviceNumber int  `json:"placeTicketDutyDeviceNumber"`
		} `json:"details"`
		SerializedOfferData   string      `json:"serializedOfferData"`
		SzabadHelyAllapot     int         `json:"szabadHelyAllapot"`
		OnlyForRegisteredUser bool        `json:"onlyForRegisteredUser"`
		AggregatedServiceIds  []int       `json:"aggregatedServiceIds"`
		OrderDisabled         bool        `json:"orderDisabled"`
		OrderDisabledReason   interface{} `json:"orderDisabledReason"`
	} `json:"route"`
	WarningMessages []string `json:"warningMessages"`
}
