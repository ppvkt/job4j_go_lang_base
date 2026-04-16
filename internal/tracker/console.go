package tracker

import "context"

type UI struct {
	In    Input
	Out   Output
	Store Store
}

func (u UI) Run(ctx context.Context) {
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

		err := action.Done(ctx, u.In, u.Out, u.Store)
		if err != nil {
			continue
		}
	}
}
