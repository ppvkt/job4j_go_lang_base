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
	Done(in Input, out Output, tracker *Tracker) error
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter name: ")
	name := in.Get()
	id := uuid.New().String()
	item, err := tracker.AddItem(Item{Name: name, ID: id})
	if err != nil {
		out.Out("failed to add item")
		return fmt.Errorf("failed to add item, %w", err)
	}
	msg := fmt.Sprintf("item added: %s", item.toString())
	out.Out(msg)

	return nil
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) error {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}

	return nil
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter target uuid: ")
	uuid := in.Get()

	out.Out("enter new name for this uuid: ")
	newName := in.Get()

	err := tracker.UpdateItem(uuid, newName)

	if err != nil {
		out.Out("unsuccess update")
		return fmt.Errorf("failed to upd item, %w", err)
	}

	out.Out("success update")
	return nil
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter uuid with you want delete: ")
	uuid := in.Get()

	err := tracker.DeleteItem(uuid)

	if err != nil {
		out.Out("unsuccess delete")
		return fmt.Errorf("failed to delete item, %w", err)
	}

	out.Out("success delete")
	return nil
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter part of name witch you want find: ")
	part := in.Get()

	res := tracker.FindItem(part)
	if len(res) == 0 {
		out.Out("no results")
		return nil
	}
	for _, item := range res {
		out.Out(item.toString())
	}

	return nil
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
		err := action.Done(u.In, u.Out, u.Tracker)
		if err != nil {
			continue
		}
	}
}
