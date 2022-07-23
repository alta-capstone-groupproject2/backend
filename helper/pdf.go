package helper

import (
	"bytes"
	"fmt"
	"lami/app/config"
	"lami/app/features/events"
	"os"
	"path/filepath"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func ExportPDF(data events.Core) string {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 0, 20)

	buildHeading(m, data)
	buildFruitList(m, data.AttendeesData)
	buf, err := m.Output()

	if err != nil {
		fmt.Println("⚠️  Could not get PDF:", err)
		os.Exit(1)
	}
	reader := bytes.NewReader(buf.Bytes())

	url, errUpPDF := UploadPDFToS3(config.AttendeesDocuments, time.Now().GoString(), config.ContentDocuments, reader)
	if errUpPDF != nil {
		fmt.Println("⚠️  Could not get PDF:", errUpPDF)
		os.Exit(1)
	}
	return url
}

func buildHeading(m pdf.Maroto, data events.Core) {
	imagePath, errPath := filepath.Abs("./images/logo.png")
	fmt.Print(errPath)
	m.RegisterHeader(func() {
		m.Row(30, func() {
			m.Col(12, func() {
				err := m.FileImage(imagePath, props.Rect{
					Center:  true,
					Percent: 40,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("%s", data.Name), props.Text{
				Size:  20,
				Style: consts.Bold,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("- City : %s", data.City), props.Text{
				Style: consts.Normal,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("- Location : %s", data.Location), props.Text{
				Style: consts.Normal,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("- Start Date : %s", data.StartDate.String()), props.Text{
				Style: consts.Normal,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("- End Date : %s", data.EndDate.String()), props.Text{
				Style: consts.Normal,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("- Attendees : %d", len(data.AttendeesData)), props.Text{
				Style: consts.Normal,
				Align: consts.Left,
				Color: color.Color{Red: 20, Green: 20, Blue: 20},
			})
		})
	})
}

func buildFruitList(m pdf.Maroto, data []events.AttendeesData) {
	tableHeadings := []string{"Num", "Name", "Email", "City", "Present"}
	contents := dataList(data)
	lightColor := getLightColor()

	m.SetBackgroundColor(getPrimaryColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Attendees", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{1, 4, 4, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 4, 4, 2, 1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}

func getLightColor() color.Color {
	return color.Color{
		Red:   240,
		Green: 216,
		Blue:  222,
	}
}

func getPrimaryColor() color.Color {
	return color.Color{
		Red:   240,
		Green: 26,
		Blue:  76,
	}
}

func generateData(datum events.AttendeesData, num int) []string {

	attendee := []string{}
	attendee = append(attendee, fmt.Sprintf("%d", num))
	attendee = append(attendee, datum.Name)
	attendee = append(attendee, datum.Email)
	attendee = append(attendee, datum.City)
	attendee = append(attendee, "")

	return attendee
}

func dataList(data []events.AttendeesData) [][]string {
	var attendees [][]string
	for i, valData := range data {
		attendee := generateData(valData, i+1)
		attendees = append(attendees, attendee)
	}

	return attendees
}
