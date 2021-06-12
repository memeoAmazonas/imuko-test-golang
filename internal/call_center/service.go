package call_center

import (
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/imuko-test-golang/internal/entity"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var result entity.ResponseBuy

func Service(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date, _ := dateparse.ParseAny(vars["date"])
	days := r.URL.Query().Get("dias")
	err := call(date, days)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "A ocurrido un error calculando la estadistica")
	}
	json.NewEncoder(w).Encode(result)
}
func call(date time.Time, days string) error {
	number, err := strconv.Atoi(days)
	if err != nil {
		return err
	}
	for i := 0; i < number; i++ {
		callApi(date, i)
	}
	return nil
}
func callApi(date time.Time, i int) error {
	var data []entity.RequestBuy
	actual := convertDate(date, i)
	acDate := fmt.Sprintf("https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/%s", actual)
	res, err := http.Get(acDate)
	if err != nil {
		return err
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	calculate(data)
	return nil
}

func calculate(data []entity.RequestBuy) {

	tdc := make(map[string]int)
	for _, val := range data {
		result.Total = result.Total + val.Monto
		if val.Monto > result.BuyHigher {
			result.BuyHigher = val.Monto
		}
		if val.Compro != true {
			result.NotBuy = result.NotBuy + 1
		}
		td := strings.ReplaceAll(val.Tdc, " ", "")
		if td != "" {
			if tdc[td] == 0 {
				tdc[td] = 1
			} else {
				tdc[td] = tdc[td] + 1
			}
		}
		result.BuyByTdc = tdc
	}

}

func convertDate(date time.Time, day int) string {
	actual := date.AddDate(0, 0, day)
	month := strconv.Itoa(int(actual.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	da := strconv.Itoa(actual.Day())
	if len(da) == 1 {
		da = "0" + da
	}
	rs := fmt.Sprintf("%s-%s-%s", strconv.Itoa(actual.Year()), month, da)
	return rs
}
