package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"math"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shopspring/decimal"
)

func cartesianToPolar(x, y decimal.Decimal, precision int) (r, theta decimal.Decimal) {
	xFloat, yFloat := x.InexactFloat64(), y.InexactFloat64()
	rFloat := math.Sqrt(xFloat*xFloat + yFloat*yFloat)
	thetaFloat := math.Atan2(yFloat, xFloat)

	r = decimal.NewFromFloat(rFloat)
	theta = decimal.NewFromFloat(thetaFloat)

	// Округление результатов
	r = r.Round(int32(precision))
	theta = theta.Round(int32(precision))
	return
}

func polarToCartesian(r, theta decimal.Decimal, precision int) (x, y decimal.Decimal) {
	rFloat, thetaFloat := r.InexactFloat64(), theta.InexactFloat64()
	xFloat := rFloat * math.Cos(thetaFloat)
	yFloat := rFloat * math.Sin(thetaFloat)

	x = decimal.NewFromFloat(xFloat)
	y = decimal.NewFromFloat(yFloat)

	// Округление результатов
	x = x.Round(int32(precision))
	y = y.Round(int32(precision))
	return
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Координатный преобразователь")

	xEntry := widget.NewEntry()
	xEntry.SetPlaceHolder("Введите x (для декартовой)")

	yEntry := widget.NewEntry()
	yEntry.SetPlaceHolder("Введите y (для декартовой)")

	rEntry := widget.NewEntry()
	rEntry.SetPlaceHolder("Введите r (для полярной)")

	thetaEntry := widget.NewEntry()
	thetaEntry.SetPlaceHolder("Введите θ (для полярной)")

	precisionEntry := widget.NewEntry()
	precisionEntry.SetPlaceHolder("Введите точность")

	resultLabel := widget.NewLabel("Результат")

	conversionSelect := widget.NewSelect(
		[]string{"Декартовы -> Полярные", "Полярные -> Декартовы"},
		func(selected string) {
			// Очистка полей ввода в зависимости от выбора
			if selected == "Декартовы -> Полярные" {
				xEntry.SetPlaceHolder("Введите x")
				yEntry.SetPlaceHolder("Введите y")
				rEntry.SetPlaceHolder("")
				thetaEntry.SetPlaceHolder("")
			} else {
				rEntry.SetPlaceHolder("Введите r")
				thetaEntry.SetPlaceHolder("Введите θ в радианах")
				xEntry.SetPlaceHolder("")
				yEntry.SetPlaceHolder("")
			}
		},
	)
	conversionSelect.SetSelected("Декартовы -> Полярные")

	formulasLabel := widget.NewLabel("Формулы преобразования:\n" +
		"Декартовы -> Полярные:\n" +
		"  r = √(x² + y²)\n" +
		"  θ = atan2(y, x)\n\n" +
		"Полярные -> Декартовы:\n" +
		"  x = r * cos(θ)\n" +
		"  y = r * sin(θ)")

	calcButton := widget.NewButton("Преобразовать", func() {
		conversionType := conversionSelect.Selected

		var x, y, r, theta decimal.Decimal
		var err error
		var precision int

		precision, err = strconv.Atoi(precisionEntry.Text)
		if err != nil {
			resultLabel.SetText("Ошибка ввода точности.")
			return
		}

		switch conversionType {
		case "Декартовы -> Полярные":
			xStr := xEntry.Text
			yStr := yEntry.Text
			x, err = decimal.NewFromString(xStr)
			y, err = decimal.NewFromString(yStr)
			if err != nil {
				resultLabel.SetText("Ошибка ввода декартовых координат.")
				return
			}

			r, theta := cartesianToPolar(x, y, precision)
			resultLabel.SetText(fmt.Sprintf("В полярной системе координат:\n r = %s, θ = %s", r.String(), theta.String()))

		case "Полярные -> Декартовы":
			rStr := rEntry.Text
			thetaStr := thetaEntry.Text
			r, err = decimal.NewFromString(rStr)
			theta, err = decimal.NewFromString(thetaStr)
			if err != nil {
				resultLabel.SetText("Ошибка ввода полярных координат.")
				return
			}

			x, y := polarToCartesian(r, theta, precision)
			resultLabel.SetText(fmt.Sprintf("В декартовой системе координат:\n x = %s, y = %s", x.String(), y.String()))
		}
	})

	// Создание контейнера с элементами интерфейса
	content := container.NewVBox(
		conversionSelect,
		xEntry,
		yEntry,
		rEntry,
		thetaEntry,
		precisionEntry,
		calcButton,
		resultLabel,
		formulasLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 350))
	myWindow.ShowAndRun()
}
