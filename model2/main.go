package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
	"strconv"
)

const g = 9.81

func calculateTrajectory(h, v0, angle float64) ([]float64, []float64, []float64, []float64) {
	theta := angle * math.Pi / 180
	vx := v0 * math.Cos(theta)
	vy := v0 * math.Sin(theta)

	var xVals, yVals, timeVals, speedVals []float64

	timeOfFlight := (vy + math.Sqrt(vy*vy+2*g*h)) / g
	step := timeOfFlight / 100

	for t := 0.0; t <= timeOfFlight; t += step {
		x := vx * t
		y := h + vy*t - 0.5*g*t*t
		if y < 0 {
			break
		}
		v := math.Sqrt(vx*vx + (vy-g*t)*(vy-g*t))

		xVals = append(xVals, x)
		yVals = append(yVals, y)
		timeVals = append(timeVals, t)
		speedVals = append(speedVals, v)
	}
	return xVals, yVals, timeVals, speedVals
}

func createPlot(xVals, yVals []float64, title, xLabel, yLabel string) *plot.Plot {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	grid := plotter.NewGrid()
	p.Add(grid)

	pts := make(plotter.XYs, len(xVals))
	for i := range pts {
		pts[i].X = xVals[i]
		pts[i].Y = yVals[i]
	}

	line, _ := plotter.NewLine(pts)
	p.Add(line)
	return p
}

func savePlot(p *plot.Plot, filename string) error {
	return p.Save(10*vg.Inch, 5*vg.Inch, filename)
}

func main() {
	a := app.New()
	w := a.NewWindow("Визуализация баллистического движения")

	heightEntry := widget.NewEntry()
	heightEntry.SetPlaceHolder("Начальная высота")

	velocityEntry := widget.NewEntry()
	velocityEntry.SetPlaceHolder("Начальная скорость (м/с)")

	angleEntry := widget.NewEntry()
	angleEntry.SetPlaceHolder("Угол (в градусах)")

	output := widget.NewLabel("")

	calculateButton := widget.NewButton("Вычислить", func() {
		h, err1 := strconv.ParseFloat(heightEntry.Text, 64)
		v0, err2 := strconv.ParseFloat(velocityEntry.Text, 64)
		angle, err3 := strconv.ParseFloat(angleEntry.Text, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			output.SetText("Неверный формат ввода.")
			return
		}

		xVals, yVals, timeVals, speedVals := calculateTrajectory(h, v0, angle)

		// Create and save trajectory plot
		trajectoryPlot := createPlot(xVals, yVals, "Траектория", "X (м)", "Y (м)")
		savePlot(trajectoryPlot, "trajectory.png")

		// Create and save velocity plot
		velocityPlot := createPlot(timeVals, speedVals, "Скорость и Время", "Время (с)", "Скорость (м/с)")
		savePlot(velocityPlot, "velocity.png")

		output.SetText("Графики сохранены: trajectory.png, velocity.png")
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Визуализация баллистического движения"),
		heightEntry,
		velocityEntry,
		angleEntry,
		calculateButton,
		output,
	))

	w.ShowAndRun()
}
