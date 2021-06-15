package cond

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	a byte = 1 << iota
	b
	c
)

type (
	fn      func(d float64, e, f int) float64
	condMap map[byte]map[string]fn

	Input struct {
		A bool    `json:"A"`
		B bool    `json:"B"`
		C bool    `json:"C"`
		D float64 `json:"D"`
		E int     `json:"E"`
		F int     `json:"F"`
	}

	Cond struct {
		Map condMap
	}

	Output struct {
		H string
		K float64
	}
)

func calcM(d float64, e, f int) float64 {
	return d + (d * float64(e) / 10)
}
func overrideM(d float64, e, f int) float64 {
	return float64(f) + d + (d * float64(e) / 100)
}
func calcP(d float64, e, f int) float64 {
	return d + (d * float64(e-f) / 25.5)
}
func calcT(d float64, e, f int) float64 {
	return d - (d * float64(f) / 30)
}
func convertBoolToMask(i *Input) byte {
	var mask byte
	if !i.A && i.B && i.C {
		mask = b | c
		return mask
	} else if i.A && !i.B && i.C {
		mask = a | c
		return mask
	} else if i.A && i.B && !i.C {
		mask = a | b
		return mask
	} else if i.A && i.B && i.C {
		mask = a | b | c
		return mask
	}
	return mask
}
func dataProceeding(i *Input) error {
	precond := &Cond{
		Map: condMap{
			3: {"M": calcM},
			5: {"M": overrideM},
			6: {"T": calcT},
			7: {"P": calcP},
		},
	}
	out := &Output{}

	mask := convertBoolToMask(i)
	if res, ok := precond.Map[mask]; ok {
		for key, val := range res {
			(*out).H, (*out).K = key, val(i.D, i.E, i.F)
			return nil
		}
	}
	return errors.New("There is no a scheme for a given input")
}

func GetInputData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		i := &Input{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		err = json.Unmarshal(body, &i)
		if err != nil {
			http.Error(w, "Error unmarshaling body", http.StatusInternalServerError)
		}

		if err = dataProceeding(i); err != nil {
			log.Fatal(err)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
