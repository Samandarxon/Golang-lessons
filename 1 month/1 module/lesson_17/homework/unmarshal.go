package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	Metadata Metadata `json:"metadata"`
	Results  []Result `json:"results"`
}
type Result struct {
	Description string   `json:"description"`
	GuId        string   `json:"guid"`
	Id          string   `json:"id"`
	Photos      []string `json:"photos"`
	Terminal    Terminal `json:"terminal"`
	TimeZone    string   `json:"timezone"`
	Title       string   `json:"title"`
}
type Metadata struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Terminal struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Variant struct {
	Id           string        `json:"id"`
	AdultPrice   AdultPrice    `json:"adultPrice"`
	ChildPrice   ChildPrice    `json:"childPrice"`
	Features     []Feature     `json:"features"`
	Refund       Refund        `json:"refund"`
	TerminalType []string      `json:"terminalType"`
	Type         string        `json:"type"`
	Value        int           `json:"value"`
	WorkingTimes []WorkingTime `json:"workingTimes"`
}

type AdultPrice struct {
	GrossPrice int     `json:"grossPrice"`
	NetPrice   float64 `json:"netPrice"`
	PaxType    string  `json:"paxType"`
}

type ChildPrice struct {
	AdultPrice
}

type Feature struct {
	Terminal
}

type Refund struct {
	Type  int    `json:"type"`
	Value string `json:"value"`
}

type WorkingTime struct {
	Close string `json:"close"`
	Day   string `json:"day"`
	Open  string `json:"open"`
}

func main() {
	var data Data

	err := json.Unmarshal([]byte(DataJSON), &data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", data.Metadata)
	for _, data := range data.Results {
		fmt.Printf("%+v\n", data)
	}

}

var DataJSON = `{
    "metadata": {
        "count": 3,
        "limit": 0,
        "offset": 0,
        "total": 0
    },
    "results": [
        {
            "description": "<p>A green hideaway awaits you near Gate 35 at Hong Kong International Airport to unwind and relax. Savour your Airport Moment overlooking a comforting tarmac view. From freshly-made meals, shower and rest cabin to relaxation zone, we are here to make your travel effortless with Plaza Premium Lounge, the perfect sanctuary to complement with your travel journey.</p><p><br></p><p>While we are open to serve you, your health and wellbeing is of paramount importance. Guided by our “We Care of Your Wellbeing” campaign, in-lounge meal service is modified to pre-portioned light refreshments and a range of beverages including coffee, tea and soft drinks. Additionally, all areas and dining tables are thoroughly cleaned and disinfected immediately after each use.</p>",
            "guid": "44898eae-4fd6-49af-87ad-35fceac6d8bb",
            "id": "P-0000000003",
            "photos": [
                "https://cdn.u-code.io/ucode/e1937f7f-3b81-4625-ac92-7583cd579a78_SignatureWall&DiningArea02copy.jpg",
                "https://cdn.u-code.io/ucode/c379aa30-095c-42b5-9ae5-51cd0078f8ea_Self-ServedFnBStation02copy.jpg",
                "https://cdn.u-code.io/ucode/8a59b15f-0288-49bb-845b-ec89b99ae980_Self-ServeFnBStation01copy.jpg",
                "https://cdn.u-code.io/ucode/ddd8771e-d534-48a5-ae14-4beeabb82daa_LoungeBay02.jpg",
                "https://cdn.u-code.io/ucode/7e70fcc2-b3ab-4f44-a164-786369acba6a_LoungeBayCocoon01copy.jpg",
                "https://cdn.u-code.io/ucode/62f44251-ad12-4869-a5dc-ee5793161eb8_LoungeBay01copy.jpg",
                "https://cdn.u-code.io/ucode/de3c606d-57e2-4a17-bac5-53836af7f2a2_DiningArea06copy.jpg",
                "https://cdn.u-code.io/ucode/269bd342-e7bd-4576-b0e9-3ca12bf95c1b_DiningArea05copy.jpg",
                "https://cdn.u-code.io/ucode/8b810e80-ca8c-47bf-a9c0-6a392b9bb082_DiningArea04copy.jpg",
                "https://cdn.u-code.io/ucode/5f6edbb9-ef05-493f-b056-a4d0ac493726_DiningArea03copy.jpg",
                "https://cdn.u-code.io/ucode/b3bf3950-f90d-4f38-acb5-584141e8daf7_DiningArea01copy.jpg",
                "https://cdn.u-code.io/ucode/e99d8f92-cc02-4f92-a0f6-ac1764140139_Concierge02copy.jpg"
            ],
            "terminal": {
                "id": "T-001",
                "title": "Terminal 1"
            },
            "timezone": "Asia/Hong_Kong",
            "title": "Plaza Premium Lounge (Near Gate 35, Departures)",
            "variants": [
                {
                    "adultPrice": {
                        "grossPrice": 82,
                        "netPrice": 65.60000000000001,
                        "paxType": "adult"
                    },
                    "childPrice": {
                        "grossPrice": 41,
                        "netPrice": 32.800000000000004,
                        "paxType": "child"
                    },
                    "features": [
                        {
                            "id": "3f2b3718-5b07-4b27-99fa-7e2bb71ef2ff",
                            "title": "Private Resting Suite"
                        },
                        {
                            "id": "602c057f-991f-4923-97dc-589dc5a134d9",
                            "title": "Massage"
                        },
                        {
                            "id": "76ae3bad-88d5-484e-95f5-6de8c81ba645",
                            "title": "Shower"
                        },
                        {
                            "id": "f23620fd-bb3f-4e28-a60b-0b5c95a92124",
                            "title": "Flight Information"
                        },
                        {
                            "id": "ab932e7a-a3a8-4f2a-a10b-72bf0863b05a",
                            "title": "Wi-Fi"
                        },
                        {
                            "id": "acce43df-3314-4b6c-b56e-e66c02dccd72",
                            "title": "Newspaper & Magazines"
                        },
                        {
                            "id": "d10370d9-ef4c-4c18-b21a-3e7aea2cad3f",
                            "title": "Washroom"
                        },
                        {
                            "id": "82107a29-8ee7-4631-b80f-e1fa6bf4e934",
                            "title": "Disabled Shower"
                        }
                    ],
                    "id": "PV-00021",
                    "refund": {
                        "type": 1,
                        "value": "day"
                    },
                    "terminalType": [
                        "departure"
                    ],
                    "type": "Hour",
                    "value": 3,
                    "workingTime": [
                        {
                            "close": "08:00",
                            "day": "saturday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "friday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "thursday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "wednesday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "tuesday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "monday",
                            "open": "23:00"
                        },
                        {
                            "close": "08:00",
                            "day": "sunday",
                            "open": "23:00"
                        }
                    ]
                }
            ]
        },
        {
            "description": "<p>This lounge is conveniently located near Gate 60 at the North Departures Hall in Hong Kong International Airport. The lounge is well-furnished with comfortable seating, shower rooms, dining area, Wi-Fi and stunning views of the tarmac.</p><p>&nbsp;</p><p>While we are open to serve you, your health and wellbeing is of paramount importance. Guided by our “We Care for Your Wellbeing” initiative, we have put in place heightened hygienic measures and practice social distancing, whilst in-lounge meal service have modified to pre-portioned light refreshment with a range of beverages including coffee, tea and soft drinks.</p>",
            "guid": "2e9ccc42-efc2-400f-813a-d4d7fd0c6c0b",
            "id": "P-0000000002",
            "photos": [
                "https://cdn.u-code.io/ucode/355aec9e-9114-4dbb-8f7f-ec982205b8f7_AeroBar(1).jpg",
                "https://cdn.u-code.io/ucode/accfe6ba-ca5b-441f-9e99-c2eb67b1bc50_NearCafe.jpg",
                "https://cdn.u-code.io/ucode/08b744dc-01c1-44fc-86b3-974ae2ea72e3_NearAeroBar(1).jpg",
                "https://cdn.u-code.io/ucode/76f848e3-4908-4635-8586-9b378350468a_LoungeBay(2).jpg",
                "https://cdn.u-code.io/ucode/b1cae032-d44d-4a0e-afef-cc1a1156854b_LoungeBay(1).jpg",
                "https://cdn.u-code.io/ucode/915addca-41ad-4b1f-aa51-daf76e270040_DiningArea(3).jpg",
                "https://cdn.u-code.io/ucode/d8d48873-446c-421a-83c0-b0686234c1e5_Cafe(3).jpg",
                "https://cdn.u-code.io/ucode/754e49b2-3ae6-4d42-bfdb-4ae2f9cd1606_Cafe(2).jpg",
                "https://cdn.u-code.io/ucode/63c4105f-8f9e-4bae-bac5-841d05aea676_Cafe(1)-.jpg",
                "https://cdn.u-code.io/ucode/ee14e187-741c-4cae-a41c-6161d023f621_AeroBar(3).jpg",
                "https://cdn.u-code.io/ucode/420e2d74-86b5-4d88-b32d-211f0ad442cc_AeroBar(2).jpg"
            ],
            "terminal": {
                "id": null,
                "title": null
            },
            "timezone": "Asia/Hong_Kong",
            "title": "Plaza Premium Lounge (Near Gate 60, Departures, Terminal 1)",
            "variants": [
                {
                    "adultPrice": {
                        "grossPrice": 82,
                        "netPrice": 65.60000000000001,
                        "paxType": "adult"
                    },
                    "childPrice": {
                        "grossPrice": 41,
                        "netPrice": 32.800000000000004,
                        "paxType": "child"
                    },
                    "features": [
                        {
                            "id": "602c057f-991f-4923-97dc-589dc5a134d9",
                            "title": "Massage"
                        },
                        {
                            "id": "6df1b5ad-7ca6-43ec-a0ff-42cd676d7423",
                            "title": "Alcoholic drinks"
                        },
                        {
                            "id": "f23620fd-bb3f-4e28-a60b-0b5c95a92124",
                            "title": "Flight Information"
                        },
                        {
                            "id": "e7cf79c7-1e42-43d3-a6bb-1482072be521",
                            "title": "Charging Station"
                        },
                        {
                            "id": "0e63caf9-0aed-45e5-a82b-c22f6508b053",
                            "title": "TV Channels"
                        },
                        {
                            "id": "ab932e7a-a3a8-4f2a-a10b-72bf0863b05a",
                            "title": "Wi-Fi"
                        },
                        {
                            "id": "acce43df-3314-4b6c-b56e-e66c02dccd72",
                            "title": "Newspaper & Magazines"
                        },
                        {
                            "id": "d10370d9-ef4c-4c18-b21a-3e7aea2cad3f",
                            "title": "Washroom"
                        }
                    ],
                    "id": "PV-00019",
                    "refund": {
                        "type": 1,
                        "value": "day"
                    },
                    "terminalType": [
                        "departure"
                    ],
                    "type": "Hour",
                    "value": 3,
                    "workingTime": [
                        {
                            "close": "01:00",
                            "day": "saturday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "friday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "thursday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "wednesday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "tuesday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "monday",
                            "open": "08:30"
                        },
                        {
                            "close": "01:00",
                            "day": "sunday",
                            "open": "08:30"
                        }
                    ]
                }
            ]
        },
        {
            "description": "<p>This lounge is conveniently located near Gate 1 at the South Departures Hall in Hong Kong International Airport. The lounge is well-furnished with comfortable seating, shower rooms, dining area, Wi-Fi and international TV channels.</p><p><br></p><p>While heightened hygienic measures and social distancing are implemented throughout the lounge, in-lounge dining can be safe and seamless as never before. With the state-of-the-art Smart Order, the contactless food ordering service, you can self-order main dishes from your smart phone, including our Signature Hong Kong-style Fish ball Noodles, by scanning a QR code from where you are seated - no queue, no hassle! Then sit back and relax. Your dishes will be served to your seat in a timely manner.</p>",
            "guid": "3679fbe5-0732-46f8-bfaf-51d204c9a123",
            "id": "P-0000000001",
            "photos": [
                "https://cdn.u-code.io/ucode/0dcfb70e-d4d1-42b4-b83d-5dc6f50ca117_IMG_0504.JPG",
                "https://cdn.u-code.io/ucode/060bf1fa-7bde-4c94-b332-f50c480bd86e_IMG_0681.JPG",
                "https://cdn.u-code.io/ucode/a26211c5-41dd-495f-8d58-32725f2b879e_IMG_0503.jpg",
                "https://cdn.u-code.io/ucode/c9b9a522-c74e-4e8d-97c9-5da463e58b4a_IMG_0495.jpg"
            ],
            "terminal": {
                "id": null,
                "title": null
            },
            "timezone": "Asia/Hong_Kong",
            "title": "Plaza Premium Lounge (Near Gate 1, Departures, Terminal 1)",
            "variants": [
                {
                    "adultPrice": {
                        "grossPrice": 82,
                        "netPrice": 65.60000000000001,
                        "paxType": "adult"
                    },
                    "childPrice": {
                        "grossPrice": 41,
                        "netPrice": 32.800000000000004,
                        "paxType": "child"
                    },
                    "features": [
                        {
                            "id": "76ae3bad-88d5-484e-95f5-6de8c81ba645",
                            "title": "Shower"
                        },
                        {
                            "id": "6df1b5ad-7ca6-43ec-a0ff-42cd676d7423",
                            "title": "Alcoholic drinks"
                        },
                        {
                            "id": "f23620fd-bb3f-4e28-a60b-0b5c95a92124",
                            "title": "Flight Information"
                        },
                        {
                            "id": "e7cf79c7-1e42-43d3-a6bb-1482072be521",
                            "title": "Charging Station"
                        },
                        {
                            "id": "0e63caf9-0aed-45e5-a82b-c22f6508b053",
                            "title": "TV Channels"
                        },
                        {
                            "id": "ab932e7a-a3a8-4f2a-a10b-72bf0863b05a",
                            "title": "Wi-Fi"
                        },
                        {
                            "id": "acce43df-3314-4b6c-b56e-e66c02dccd72",
                            "title": "Newspaper & Magazines"
                        },
                        {
                            "id": "d10370d9-ef4c-4c18-b21a-3e7aea2cad3f",
                            "title": "Washroom"
                        },
                        {
                            "id": "82107a29-8ee7-4631-b80f-e1fa6bf4e934",
                            "title": "Disabled Shower"
                        }
                    ],
                    "id": "PV-00017",
                    "refund": {
                        "type": 1,
                        "value": "day"
                    },
                    "terminalType": [
                        "departure"
                    ],
                    "type": "Hour",
                    "value": 3,
                    "workingTime": [
                        {
                            "close": "01:00",
                            "day": "saturday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "friday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "thursday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "wednesday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "tuesday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "monday",
                            "open": "06:30"
                        },
                        {
                            "close": "01:00",
                            "day": "sunday",
                            "open": "06:30"
                        }
                    ]
                }
            ]
        }
    ]
}`
