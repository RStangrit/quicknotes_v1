package mainMenuNavigation

import (
	"fmt"
	"main/errHandler"
	"main/fileOps"
	"main/note"
)

func createNote() {
	fileOps.CheckFileExistence()
	title, content, err := note.GetNoteData()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	fileOps.WriteNoteToFile(title, content)
	fmt.Println("Note has been successfully created:", title)
}

func readNotes() {
	fileOps.CheckFileExistence()
	fileOps.ReadNotesFromFile()
}

func readOneNote() {
	fileOps.CheckFileExistence()
	noteIndex, err := note.GetNoteIndex()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	fileOps.ReadNoteFromFile(noteIndex)
}

func updateNote() {
	fileOps.CheckFileExistence()
	noteIndex, err := note.GetNoteIndex()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	title, content, err := note.GetNoteData()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	fileOps.UpdateNoteInFile(noteIndex, title, content)
}

func deleteNote() {
	fileOps.CheckFileExistence()
	noteIndex, err := note.GetNoteIndex()
	if err != nil {
		errHandler.HandleError(err)
		return
	}
	fileOps.DeleteNoteInFile(noteIndex)
}

func Selector(choice string) {
	switch choice {
	case "1":
		createNote()
	case "2":
		readNotes()
	case "3":
		readOneNote()
	case "4":
		updateNote()
	case "5":
		deleteNote()
	default:
		fmt.Println("Wrong input, please try again!")
	}
}
