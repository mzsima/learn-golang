package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

type Point struct {
	X, Y float32
}

type Rectangle struct {
	Min, Max Point
}

var progressIncrementer chan float32
var progress float32

func main() {
	progressIncrementer = make(chan float32)

	go func() {
		for {
			time.Sleep(time.Second / 25)
			progressIncrementer <- 0.004
		}
	}()

	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Egg timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	var startButton widget.Clickable

	var boilDurationInput widget.Editor
	var boiling bool
	var boilDuration float32
	th := material.NewTheme(gofont.Collection())

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {

			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				if startButton.Clicked() {
					boiling = !boiling

					if progress >= 1 {
						progress = 0
					}

					inputString := boilDurationInput.Text()
					inputString = strings.TrimSpace(inputString)
					inputFloat, _ := strconv.ParseFloat(inputString, 32)
					boilDuration = float32(inputFloat)
					boilDuration = boilDuration / (1 - progress)
				}

				if startButton.Clicked() {
					boiling = !boiling
				}
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx C) D {
							var eggPath clip.Path
							op.Offset(f32.Pt(400, 150)).Add(gtx.Ops)
							eggPath.Begin(gtx.Ops)
							for deg := 0.0; deg <= 360; deg++ {
								rad := deg / 360 * 2 * math.Pi
								cosT := math.Cos(rad)
								sinT := math.Sin(rad)
								a := 110.0
								b := 150.0
								d := 20.0
								x := a * cosT
								y := -(math.Sqrt(b*b-d*d*cosT*cosT) + d*sinT) * sinT
								p := f32.Pt(float32(x), float32(y))
								eggPath.LineTo(p)
							}
							eggPath.Close()

							eggArea := clip.Outline{Path: eggPath.End()}.Op()
							color := color.NRGBA{R: 255, G: uint8(239 * (1 - progress)), B: uint8(174 * (1 - progress)), A: 255}
							paint.FillShape(gtx.Ops, color, eggArea)

							d := image.Point{Y: 375}
							return layout.Dimensions{Size: d}
						},
					),
					layout.Rigid(
						func(gtx C) D {
							ed := material.Editor(th, &boilDurationInput, "sec")
							boilDurationInput.SingleLine = true
							boilDurationInput.Alignment = text.Middle

							if boiling && progress < 1 {
								boilRemain := (1 - progress) * boilDuration
								inputStr := fmt.Sprintf("%.1f", math.Round(float64(boilRemain)*10)/10)
								boilDurationInput.SetText(inputStr)
							}

							margins := layout.Inset{
								Top:    unit.Dp(0),
								Right:  unit.Dp(170),
								Bottom: unit.Dp(170),
								Left:   unit.Dp(170),
							}
							border := widget.Border{
								Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
								CornerRadius: unit.Dp(3),
								Width:        unit.Dp(2),
							}
							return margins.Layout(gtx,
								func(gtx C) D {
									return border.Layout(gtx, ed.Layout)
								},
							)
						},
					),
					layout.Rigid(
						func(gtx C) D {
							bar := material.ProgressBar(th, progress)
							return bar.Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx C) D {
							margins := layout.Inset{
								Top:    unit.Dp(25),
								Bottom: unit.Dp(25),
								Right:  unit.Dp(35),
								Left:   unit.Dp(35),
							}
							return margins.Layout(gtx,
								func(gtx C) D {
									var text string
									if !boiling {
										text = "Start"
									}
									if boiling && progress < 1 {
										text = "Stop"
									}
									if boiling && progress >= 1 {
										text = "Finished"
									}
									btn := material.Button(th, &startButton, text)
									return btn.Layout(gtx)
								},
							)
						},
					),
					layout.Rigid(
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)

			case system.DestroyEvent:
				return e.Err
			}

		case <-progressIncrementer:
			if boiling && progress < 1 {
				progress += 1.0 / 25.0
				if progress > 1 {
					progress = 1
				}
				w.Invalidate()
			}
		}
	}
}
