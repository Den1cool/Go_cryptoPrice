package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type RAW struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      float64 `json:"LASTUPDATE"`
	SUPPLY          float64 `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}
type DISPLAY struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	LOW24HOUR       string `json:"LOW24HOUR"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`
	LASTUPDATE      string `json:"LASTUPDATE"`
	SUPPLY          string `json:"SUPPLY"`
	MKTCAP          string `json:"MKTCAP"`
}

var database *sql.DB

const APIKEY string = "3999f16f4879448ee4bf098e97accf8403792964ead5c97e680ba59dbc4a9620"

func cryptoPrice(w http.ResponseWriter, r *http.Request) {
	var fsymsArr, tsymsArr []string
	var RawArr []RAW
	var DisplayArr []DISPLAY
	fsyms := r.URL.Query().Get("fsyms")
	tsyms := r.URL.Query().Get("tsyms")
	fsymsArr = strings.Split(fsyms, ",")
	tsymsArr = strings.Split(tsyms, ",")
	i := 0
	for i < len(fsymsArr) {
		resp, err := http.Get("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=" + fsymsArr[i] + "&tsyms=" + tsymsArr[i] + "&api_key=" + APIKEY)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if json.Valid(body) {
			var Raw RAW
			var Display DISPLAY
			if err := json.Unmarshal(body, &Raw); err != nil {
				log.Fatalln(err)
			}
			type tsymsT struct {
				y map[string]map[string]map[string]interface{} `json:"-"`
			}
			type fsymsT struct {
				q tsymsT
			}
			type RawAndDisp struct {
				rw fsymsT
			}
			var ss RawAndDisp
			if err := json.Unmarshal(body, &ss.rw.q.y); err != nil {
				log.Fatalln(err)
			}
			b, err := json.Marshal(ss.rw.q.y["RAW"][fsymsArr[i]][tsymsArr[i]])
			if err != nil {
				log.Fatalln(err)
			}

			if err := json.Unmarshal(b, &Raw); err != nil {
				log.Fatalln(err)
			}
			RawArr = append(RawArr, Raw)
			b, err = json.Marshal(ss.rw.q.y["DISPLAY"][fsymsArr[i]][tsymsArr[i]])
			if err != nil {
				log.Fatalln(err)
			}

			if err := json.Unmarshal(b, &Display); err != nil {
				log.Fatalln(err)
			}

			DisplayArr = append(DisplayArr, Display)
			var exists bool
			_ = database.QueryRow("select exists(select 1 from public.pair where fsyms=$1 and tsyms=$2)", fsyms, tsyms).Scan(&exists)
			if !exists {
				_, err = database.Exec(`insert into public."RAW"(change24hour, changepct24hour, open24hour, volume24hour, volume24hourto, low24hour, high24hour, price, lastupdate, supply, mktcap) values($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11)`, Raw.CHANGE24HOUR, Raw.CHANGEPCT24HOUR, Raw.OPEN24HOUR, Raw.VOLUME24HOUR, Raw.VOLUME24HOURTO, Raw.LOW24HOUR, Raw.HIGH24HOUR, Raw.PRICE, Raw.LASTUPDATE, Raw.MKTCAP, Raw.SUPPLY)
				if err != nil {
					log.Fatalln(err)
				}
				var RawId, DisplayId string
				_ = database.QueryRow(`select raw_id from public."RAW" order by raw_id DESC limit 1`).Scan(&RawId)
				_, err = database.Exec(`insert into public."DISPLAY"(change24hour, changepct24hour, open24hour, volume24hour, volume24hourto, low24hour, high24hour, price, lastupdate, supply, mktcap) values($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11)`, Display.CHANGE24HOUR, Display.CHANGEPCT24HOUR, Display.OPEN24HOUR, Display.VOLUME24HOUR, Display.VOLUME24HOURTO, Display.LOW24HOUR, Display.HIGH24HOUR, Display.PRICE, Display.LASTUPDATE, Display.MKTCAP, Display.SUPPLY)
				if err != nil {
					log.Fatalln(err)
				}
				_ = database.QueryRow(`select display_id from public."DISPLAY" order by display_id DESC limit 1`).Scan(&DisplayId)
				dt := time.Now()
				_, err = database.Exec(`insert into public.pair(fsyms,tsyms,raw_id,display_id, updatetime) values($1,$2,$3,$4,$5)`, fsyms, tsyms, RawId, DisplayId, dt)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				var UpdateTime time.Time
				_ = database.QueryRow("select updatetime from public.pair where fsyms=$1 and tsyms=$2", fsyms, tsyms).Scan(&UpdateTime)
				dt := time.Now()
				TimeDiff := dt.Sub(UpdateTime)
				if TimeDiff.Minutes() > 3 {
					_, err = database.Exec(`update public."RAW" set change24hour=$1, changepct24hour=$2, open24hour=$3, volume24hour=$4, volume24hourto=$5, low24hour=$6, high24hour=$7, price=$8, lastupdate=$9, supply=$10, mktcap=$11`, Raw.CHANGE24HOUR, Raw.CHANGEPCT24HOUR, Raw.OPEN24HOUR, Raw.VOLUME24HOUR, Raw.VOLUME24HOURTO, Raw.LOW24HOUR, Raw.HIGH24HOUR, Raw.PRICE, Raw.LASTUPDATE, Raw.MKTCAP, Raw.SUPPLY)
					if err != nil {
						log.Fatalln(err)
					}
					_, err = database.Exec(`update public."DISPLAY" set change24hour=$1, changepct24hour=$2, open24hour=$3, volume24hour=$4, volume24hourto=$5, low24hour=$6, high24hour=$7, price=$8, lastupdate=$9, supply=$10, mktcap=$11`, Display.CHANGE24HOUR, Display.CHANGEPCT24HOUR, Display.OPEN24HOUR, Display.VOLUME24HOUR, Display.VOLUME24HOURTO, Display.LOW24HOUR, Display.HIGH24HOUR, Display.PRICE, Display.LASTUPDATE, Display.MKTCAP, Display.SUPPLY)
					if err != nil {
						log.Fatalln(err)
					}
					_, err = database.Exec(`update public.pair set updatetime=$1 where fsyms=$2 and tsyms=$3`, dt, fsyms, tsyms)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
			response := map[string]interface{}{
				"RAW": map[string]interface{}{
					fsymsArr[i]: map[string]interface{}{
						tsymsArr[i]: Raw,
					},
				},
				"DISPLAY": map[string]interface{}{
					fsymsArr[i]: map[string]interface{}{
						tsymsArr[i]: Display,
					},
				},
			}
			JsonData, err := json.MarshalIndent(response, "", "	")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(JsonData))
			w.Write(JsonData)
		} else {
			var exists bool
			_ = database.QueryRow("select exists(select 1 from public.pair where fsyms=$1 and tsyms=$2)", fsymsArr[i], tsymsArr[i]).Scan(&exists)
			if !exists {
				response := map[string]string{
					"error": fsymsArr[i] + " и " + tsymsArr[i] + " не найдены в базе данных",
				}
				JsonData, err := json.MarshalIndent(response, "", "	")
				if err != nil {
					log.Fatal(err)
				}
				w.Write(JsonData)
				fmt.Println("Пара не найдена")
			} else {
				var Raw RAW
				var Display DISPLAY
				var rawId, displayId string
				_ = database.QueryRow("select raw_id, display_id from public.pair where fsyms=$1 and tsyms=$2", fsymsArr[i], tsymsArr[i]).Scan(&rawId, &displayId)
				row := database.QueryRow(`select * from public."RAW" where raw_id=$1`, rawId)
				err := row.Scan(&rawId, &Raw.CHANGE24HOUR, &Raw.CHANGEPCT24HOUR, &Raw.OPEN24HOUR, &Raw.VOLUME24HOUR, &Raw.VOLUME24HOURTO, &Raw.LOW24HOUR, &Raw.HIGH24HOUR, &Raw.PRICE, &Raw.LASTUPDATE, &Raw.SUPPLY, &Raw.MKTCAP)
				if err != nil {
					log.Fatal(err)
				}
				row = database.QueryRow(`select * from public."DISPLAY" where display_id=$1`, displayId)
				err = row.Scan(&displayId, &Display.CHANGE24HOUR, &Display.CHANGEPCT24HOUR, &Display.OPEN24HOUR, &Display.VOLUME24HOUR, &Display.VOLUME24HOURTO, &Display.LOW24HOUR, &Display.HIGH24HOUR, &Display.PRICE, &Display.LASTUPDATE, &Display.SUPPLY, &Display.MKTCAP)
				if err != nil {
					log.Fatal(err)
				}
				response := map[string]interface{}{
					"RAW": map[string]interface{}{
						fsymsArr[i]: map[string]interface{}{
							tsymsArr[i]: Raw,
						},
					},
					"DISPLAY": map[string]interface{}{
						fsymsArr[i]: map[string]interface{}{
							tsymsArr[i]: Display,
						},
					},
				}
				JsonData, err := json.MarshalIndent(response, "", "	")
				if err != nil {
					log.Fatal(err)
				}
				w.Write(JsonData)
			}
		}
		i++
	}
}

func main() {
	connStr := "user=postgres password=root dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	database = db
	http.HandleFunc("/price", cryptoPrice)
	http.ListenAndServe(":9000", nil)
}
