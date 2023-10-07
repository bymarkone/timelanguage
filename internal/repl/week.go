package repl

import (
	"github.com/bymarkone/timelanguage/internal/schedule"
	"github.com/bymarkone/timelanguage/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"strings"
	"time"
)

func init() {
	command := Command{
		Description: "Prints the plan for the week",
		Usage:       "week",
		Arguments:   []Argument{},
		Flags:       []Flag{},
		Function:    week,
	}
	RegisterCommands("week", command)
}

func week(out io.ReadWriter, _ []string) {
	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	times := []string{
		"04:00", "05:00", "06:00", "07:00", "08:00", "09:00",
		"10:00", "11:00", "12:00", "13:00", "14:00", "15:00",
		"16:00", "17:00", "18:00", "19:00", "20:00", "21:00",
	}
	//times := []string{
	//	"04:00", "04:30", "05:00", "05:30", "06:00", "06:30", "07:00", "07:30", "08:00", "08:30",
	//	"09:00", "09:30", "10:00", "10:30", "11:00", "11:30", "12:00", "12:30", "13:00", "13:30",
	//	"14:00", "14:30", "15:00", "15:30", "16:00", "16:30", "17:00", "17:30", "18:00", "18:30",
	//	"19:00", "19:30", "20:00", "20:30", "21:00", "21:30", "22:00", "22:30", "23:00", "23:30",
	//}

	t := table.NewWriter()
	t.SetOutputMirror(out)
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	var header []interface{}
	header = append(header, " ")
	for _, weekday := range weekdays {
		header = append(header, weekday.String())
	}
	t.AppendHeader(header)

	tracks := schedule.GetRepository().ListTracks()
	var rows []table.Row
	for _, timeSlot := range times {
		hour, _, _ := utils.Parse(timeSlot)
		row := table.Row{}
		row = append(row, timeSlot)
		for _, weekday := range weekdays {
			var name []string
			for _, track := range tracks {
				//if (track.Slot.Period.Start.Hour < hour || (track.Slot.Period.Start.Hour == hour && track.Slot.Period.Start.Minute <= minute)) &&
				//	(track.Slot.Period.End.Hour > hour || (track.Slot.Period.End.Hour == hour && track.Slot.Period.End.Minute > minute)) {
				if (track.Slot.Period.Start.Hour <= hour) && (track.Slot.Period.End.Hour > hour) {
					if ContainsWeekday(track.Slot.Period.Weekdays, weekday) {
						name = append(name, track.Name)
					}
				}
			}
			if len(name) == 0 {
				name = append(name, "x")
			}
			row = append(row, strings.Join(name, ","))
		}
		rows = append(rows, row)
	}
	t.AppendRows(rows, rowConfigAutoMerge)

	t.Style().Options.SeparateRows = true
	widthMax := 30
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true, WidthMax: widthMax},
		{Number: 2, AutoMerge: true, WidthMax: widthMax},
		{Number: 3, AutoMerge: true, WidthMax: widthMax},
		{Number: 4, AutoMerge: true, WidthMax: widthMax},
		{Number: 5, AutoMerge: true, WidthMax: widthMax},
		{Number: 6, AutoMerge: true, WidthMax: widthMax},
		{Number: 7, AutoMerge: true, WidthMax: widthMax},
		{Number: 8, AutoMerge: true, WidthMax: widthMax},
	})
	t.Render()
}
