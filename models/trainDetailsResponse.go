package models

import "time"

type TrainDetailsResponse struct {
	TrainSchedulerDetails []struct {
		Train struct {
			AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
			Name                        string        `json:"name"`
			SeatReservationCode         string        `json:"seatReservationCode"`
			Code                        string        `json:"code"`
			CompanyCode                 interface{}   `json:"companyCode"`
			Route                       interface{}   `json:"route"`
			StartStationReservationCode interface{}   `json:"startStationReservationCode"`
			EndStationReservationCode   interface{}   `json:"endStationReservationCode"`
			StartStation                struct {
				ID                            int    `json:"id"`
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
				ID                            int    `json:"id"`
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
			StartDate        interface{} `json:"startDate"`
			OrigStartStation interface{} `json:"origStartStation"`
			OrigEndStation   interface{} `json:"origEndStation"`
			Start            time.Time   `json:"start"`
			VirtualStart     bool        `json:"virtualStart"`
			Arrive           time.Time   `json:"arrive"`
			VirtualArrive    bool        `json:"virtualArrive"`
			Distance         float64     `json:"distance"`
			ClosedTrackway   bool        `json:"closedTrackway"`
			FullName         string      `json:"fullName"`
			FullNameAndType  string      `json:"fullNameAndType"`
			Kinds            []struct {
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
					ID                            int    `json:"id"`
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
					ID                            int    `json:"id"`
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
			} `json:"kinds"`
			KindsToDisplay []struct {
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
					ID                            int    `json:"id"`
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
					ID                            int    `json:"id"`
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
			} `json:"kindsToDisplay"`
			Kind struct {
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
					ID                            int    `json:"id"`
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
					ID                            int    `json:"id"`
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
			} `json:"kind"`
			Services []struct {
				ListOrder                   string      `json:"listOrder"`
				Description                 string      `json:"description"`
				RestrictiveStartStationCode interface{} `json:"restrictiveStartStationCode"`
				RestrictiveEndStationCode   interface{} `json:"restrictiveEndStationCode"`
				Sign                        struct {
					FontName  string `json:"fontName"`
					Character string `json:"character"`
				} `json:"sign"`
				TrainStopKind interface{} `json:"trainStopKind"`
			} `json:"services"`
			ActualOrEstimatedStart  *time.Time `json:"actualOrEstimatedStart"`
			ActualOrEstimatedArrive *time.Time `json:"actualOrEstimatedArrive"`
			HavarianInfok           struct {
				AktualisKeses float64     `json:"aktualisKeses"`
				KesesiOk      interface{} `json:"kesesiOk"`
				HavariaInfo   interface{} `json:"havariaInfo"`
				UzletiInfo    interface{} `json:"uzletiInfo"`
				KesesInfo     string      `json:"kesesInfo"`
			} `json:"havarianInfok"`
			DirectTrains []struct {
				Train struct {
					AggregatedServiceIds        []interface{} `json:"aggregatedServiceIds"`
					Name                        interface{}   `json:"name"`
					SeatReservationCode         string        `json:"seatReservationCode"`
					Code                        string        `json:"code"`
					CompanyCode                 interface{}   `json:"companyCode"`
					Route                       interface{}   `json:"route"`
					StartStationReservationCode interface{}   `json:"startStationReservationCode"`
					EndStationReservationCode   interface{}   `json:"endStationReservationCode"`
					StartStation                struct {
						ID                            int    `json:"id"`
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
						ID                            int    `json:"id"`
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
					StartDate               interface{}   `json:"startDate"`
					OrigStartStation        interface{}   `json:"origStartStation"`
					OrigEndStation          interface{}   `json:"origEndStation"`
					Start                   interface{}   `json:"start"`
					VirtualStart            bool          `json:"virtualStart"`
					Arrive                  interface{}   `json:"arrive"`
					VirtualArrive           bool          `json:"virtualArrive"`
					Distance                float64       `json:"distance"`
					ClosedTrackway          bool          `json:"closedTrackway"`
					FullName                string        `json:"fullName"`
					FullNameAndType         string        `json:"fullNameAndType"`
					Kinds                   interface{}   `json:"kinds"`
					KindsToDisplay          interface{}   `json:"kindsToDisplay"`
					Kind                    interface{}   `json:"kind"`
					Services                []interface{} `json:"services"`
					ActualOrEstimatedStart  interface{}   `json:"actualOrEstimatedStart"`
					ActualOrEstimatedArrive interface{}   `json:"actualOrEstimatedArrive"`
					HavarianInfok           interface{}   `json:"havarianInfok"`
					DirectTrains            interface{}   `json:"directTrains"`
					CarrierTrains           interface{}   `json:"carrierTrains"`
					StartTrack              interface{}   `json:"startTrack"`
					EndTrack                interface{}   `json:"endTrack"`
					JeEszkozAlapID          float64       `json:"jeEszkozAlapId"`
					FullType                string        `json:"fullType"`
					FullShortType           string        `json:"fullShortType"`
					FullNameAndPiktogram    struct {
						Collection string `json:"(Collection)"`
					} `json:"fullNameAndPiktogram"`
					Footer           interface{} `json:"footer"`
					ViszonylatiJel   interface{} `json:"viszonylatiJel"`
					ViszonylatObject interface{} `json:"viszonylatObject"`
					Description      interface{} `json:"description"`
					SameCar          bool        `json:"sameCar"`
					StartTimeZone    interface{} `json:"startTimeZone"`
					ArriveTimeZone   interface{} `json:"arriveTimeZone"`
					TrainID          string      `json:"trainId"`
				} `json:"train"`
				StartStation struct {
					ID                            int    `json:"id"`
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
					ID                            int    `json:"id"`
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
				Footer interface{} `json:"footer"`
			} `json:"directTrains"`
			CarrierTrains        interface{} `json:"carrierTrains"`
			StartTrack           interface{} `json:"startTrack"`
			EndTrack             interface{} `json:"endTrack"`
			JeEszkozAlapID       float64     `json:"jeEszkozAlapId"`
			FullType             string      `json:"fullType"`
			FullShortType        string      `json:"fullShortType"`
			FullNameAndPiktogram struct {
				Collection string `json:"(Collection)"`
			} `json:"fullNameAndPiktogram"`
			Footer           string      `json:"footer"`
			ViszonylatiJel   interface{} `json:"viszonylatiJel"`
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
			TrainID        string      `json:"trainId"`
		} `json:"train"`
		Scheduler []TD_Scheduler `json:"scheduler"`
	} `json:"trainSchedulerDetails"`
	StationSchedulerDetails interface{} `json:"stationSchedulerDetails"`
	RouteSchedulerDetails   interface{} `json:"routeSchedulerDetails"`
}
