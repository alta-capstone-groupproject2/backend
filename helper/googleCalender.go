package helper

import (
	"fmt"
	"lami/app/config"
	"lami/app/features/events"
	"log"

	"google.golang.org/api/calendar/v3"
)

func InsertEvent(eventCore events.Core, email string) (eventId string, err error) {
	srv := config.CalendarService()
	event := &calendar.Event{
		Summary:     eventCore.Name,
		Location:    eventCore.City,
		Description: eventCore.Detail,
		Start: &calendar.EventDateTime{
			DateTime: eventCore.StartDate.GoString(),
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: eventCore.EndDate.GoString(),
			TimeZone: "Asia/Jakarta",
		},
		Attendees: []*calendar.EventAttendee{
			{
				Email: email,
			},
		},
	}
	calendarID := "primary"

	event, err = srv.Events.Insert(calendarID, event).SendNotifications(true).Do()
	if err != nil {
		return "", err
	}
	return event.Id, err
}

func UpdateEvent(srv *calendar.Service, eventID string) {
	event := &calendar.Event{
		Summary:     "Sample event 2",
		Location:    "Sample location",
		Description: "This is a sample event.",
		Start: &calendar.EventDateTime{
			DateTime: "2022-07-22T00:00:00+09:00",
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: "2022-07-22T01:00:00+09:00",
			TimeZone: "Asia/Jakarta",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: "alfin.7007@gmail.com"},
		},
		Attachments: []*calendar.EventAttachment{},
	}

	calendarID := "primary"

	event, err := srv.Events.Update(calendarID, eventID, event).Do()
	if err != nil {
		log.Fatalf("Unable to update event. %v\n", err)
	}
	fmt.Printf("Event updated: %s %s\n", event.HtmlLink, event.Id)

}

func eventList(srv *calendar.Service) {
	calendarID := "primary"
	events, err := srv.Events.List(calendarID).OrderBy("startTime").Do()
	fmt.Println(err)
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			eventID := item.Id
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%s %v (%v)\n", eventID, item.Summary, date)
		}
	}
}
