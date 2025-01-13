package mainMenuOpts

import (
	"fmt"
	"strings"
)

const (
	CreateNote  = "1. Create Note"
	ReadNotes   = "2. List Notes"
	ReadOneNote = "3. Read One Note"
	UpdateNote  = "4. Update Note"
	DeleteNote  = "5. Delete Note"
)

func Show() {
	options := []string{
		CreateNote,
		ReadNotes,
		ReadOneNote,
		UpdateNote,
		DeleteNote,
	}

	fmt.Println("List of available actions:")
	fmt.Println(strings.Join(options, "\n"))
}
