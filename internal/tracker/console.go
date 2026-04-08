package tracker

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
}

type ConsoleInput struct{}

func (c ConsoleInput) Get() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name: ")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter target uuid: ")
	uuid := in.Get()

	out.Out("enter new name for this uuid: ")
	newName := in.Get()

	ok := tracker.UpdateItem(uuid, newName)

	if !ok {
		out.Out("unsuccess update")
		return
	}

	out.Out("success update")
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid with you want delete: ")
	uuid := in.Get()

	ok := tracker.DeleteItem(uuid)

	if !ok {
		out.Out("unsuccess delete")
		return
	}

	out.Out("success delete")
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter part of name witch you want find: ")
	part := in.Get()

	res := tracker.FindItem(part)
	if len(res) == 0 {
		out.Out("no results")
		return
	}
	for _, item := range res {
		out.Out(item.toString())
	}
}

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) Run() {
	actions := map[string]Usecase{
		"add":    AddUsecase{},
		"get":    GetUsecase{},
		"update": UpdateUsecase{},
		"delete": DeleteUsecase{},
		"find":   FindUsecase{},
	}
	for {
		u.Out.Out("select action (add/get/update/delete/find/exit):")
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("not found action")
			continue
		}
		action.Done(u.In, u.Out, u.Tracker)
	}
}
