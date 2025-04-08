package utils

import (
	"cinema/models"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func SaveTicketAsJSON(ticket models.SavedTicket) error {
	data, err := json.MarshalIndent(ticket, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal ticket: %v", err)
	}

	filename := fmt.Sprintf("ticket-%s-%s.json", ticket.Movie, time.Now().Format("2006-01-02-15-04-05"))
	return os.WriteFile(filename, data, 0644)
}

func LoadTicketFromJSON(filepath string) (models.SavedTicket, error) {
	var ticket models.SavedTicket
	data, err := os.ReadFile(filepath)
	if err != nil {
		return ticket, fmt.Errorf("failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &ticket)
	if err != nil {
		return ticket, fmt.Errorf("failed to unmarshal ticket: %v", err)
	}

	return ticket, nil
}
