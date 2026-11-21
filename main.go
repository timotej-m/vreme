package main

import (
	"context"
	"log"
	"os"

	"github.com/timotej-m/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	vreme.Slovar = make(map[string]vreme.Postaja)

	cmd := &cli.Command{
		Name:  "Temperature Slovenskih Postaj",
		Usage: "[prikaz izmerjenih temperatur za različne postaje slovenije]",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "minTemp",
				Usage: "> nastavitev najnižje dovoljene temperature (privzeto = -50)",
				Value: -50,
			},
			&cli.IntFlag{
				Name:  "maxTemp",
				Usage: "> nastavitev najvišje dovoljene temperature (privzeto = 50)",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "> nastavitev najmanješe število meritev za izračun statistike (privzeto = 3)",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Int("minTemp")
			maxTemp := cmd.Int("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			return pozeni(minTemp, maxTemp, stMeritev)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	/*
		vreme.Slovar["LJ03"] = vreme.Postaja{"Ljubljana Šiška", "Ljubljana", []float64{23.3, 24.8, 39.4}}
		vreme.Slovar["LJ05"] = Postaja{"Ljubljana Bežigrad", "Ljubljana", []float64{19.4}}
		vreme.Slovar["LJ07"] = Postaja{"Ljubljana Črnuče", "Ljubljana", []float64{-20.4, 19.1, 18.9}}
		vreme.Slovar["KM01"] = Postaja{"Kamnik Graben", "Kamnik", []float64{24.4, 26.1, 23.6}}
		vreme.Slovar["MB05"] = Postaja{"Nekje v Mariboru", "Maribor", []float64{-25.5, -27.8, 20.3}}
	*/

	vreme.DodajPostajo("KM01", "Kamnik Graben", "Kamnik")
	vreme.DodajMeritev(vreme.Slovar, "KM01", 23.9)

	vreme.DodajPostajo("MB05", "Nekje v Mariboru", "Maribor")
	vreme.DodajMeritev(vreme.Slovar, "MB05", 34.8)
	vreme.DodajMeritev(vreme.Slovar, "MB05", 36.8)

	vreme.DodajPostajo("LJ03", "Ljubljana Šiška", "Ljubljana")
	vreme.DodajMeritev(vreme.Slovar, "LJ03", 32.2)
	vreme.DodajMeritev(vreme.Slovar, "LJ03", 24.8)
	vreme.DodajMeritev(vreme.Slovar, "LJ03", 39.4)

	vreme.DodajPostajo("LJ05", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajMeritev(vreme.Slovar, "LJ05", 19.2)
	vreme.DodajMeritev(vreme.Slovar, "LJ05", 25.6)
	vreme.DodajMeritev(vreme.Slovar, "LJ05", 18.9)
	vreme.DodajMeritev(vreme.Slovar, "LJ05", 15.9)

	vreme.DodajPostajo("LJ07", "Ljubljana Črnuče", "Ljubljana")
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 12.7)
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 23.2)
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 15.4)
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 11.3)
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 9.3)

	vreme.DodajMeritev(vreme.Slovar, "LJ03", -432.2)
	vreme.DodajMeritev(vreme.Slovar, "LJ07", 420.42)
	vreme.DodajMeritev(vreme.Slovar, "LJLJ", 32.2)

	vreme.IzpisMeritev(vreme.Slovar)

	vreme.IzpisPorocila(vreme.Slovar)

}

func pozeni(minTemp int, maxTemp int, stMeritev int) error {
	vreme.Constraints(minTemp, maxTemp, stMeritev)
	return nil
}
