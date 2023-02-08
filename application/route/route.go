package route

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clienteId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}

	filePath := fmt.Sprintf("destinations/%v.txt", r.ID)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}
