package tracker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, uuid string, newName string) error
	Find(ctx context.Context, namePart string) []Item
}

type Usecase interface {
	Done(ctx context.Context, in Input, out Output, store Store) error
}

type AddUsecase struct{}

func (u AddUsecase) Done(
	ctx context.Context,
	in Input, out Output,
	store Store,
) error {
	out.Out("enter name: ")
	name := in.Get()
	id := uuid.New().String()

	if err := store.Create(ctx, Item{ID: id, Name: name}); err != nil {
		out.Out("failed to add item")
		return fmt.Errorf("failed to add item, %w", err)
	}
	msg := fmt.Sprintf("item added! id: %s, name: %s", id, name)
	out.Out(msg)

	return nil
}

type GetUsecase struct{}

func (u GetUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	items, err := store.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	}
	for _, item := range items {
		out.Out(item.toString())
	}
	return nil
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter target uuid: ")
	uuid := in.Get()

	out.Out("enter new name for this uuid: ")
	newName := in.Get()

	err := store.Update(ctx, uuid, newName)

	if err != nil {
		out.Out("unsuccess update")
		return fmt.Errorf("failed to upd item, %w", err)
	}

	out.Out("success update")
	return nil
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter uuid with you want delete: ")
	uuid := in.Get()

	err := store.Delete(ctx, uuid)

	if err != nil {
		out.Out("unsuccess delete")
		return fmt.Errorf("failed to delete item, %w", err)
	}

	out.Out("success delete")
	return nil
}

type FindUsecase struct{}

func (u FindUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter part of name witch you want find: ")
	part := in.Get()

	res := store.Find(ctx, part)
	if len(res) == 0 {
		out.Out("no results")
		return nil
	}
	for _, item := range res {
		out.Out(item.toString())
	}

	return nil
}
