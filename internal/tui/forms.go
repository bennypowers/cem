package tui

import (
	"errors"

	"charm.land/huh/v2"
)

var ErrCancelled = errors.New("cancelled by user")

func WrapAbort(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, huh.ErrUserAborted) {
		return ErrCancelled
	}
	return err
}

func Confirm(title string, defaultValue bool) (bool, error) {
	result := defaultValue
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Affirmative("Yes").
				Negative("No").
				Value(&result),
		),
	).Run()
	if err != nil {
		return false, WrapAbort(err)
	}
	return result, nil
}

func Select[T comparable](title string, options []huh.Option[T]) (T, error) {
	var result T
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[T]().
				Title(title).
				Options(options...).
				Value(&result),
		),
	).Run()
	if err != nil {
		return result, WrapAbort(err)
	}
	return result, nil
}

func MultiSelect[T comparable](title string, options []huh.Option[T]) ([]T, error) {
	var result []T
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[T]().
				Title(title).
				Options(options...).
				Value(&result),
		),
	).Run()
	if err != nil {
		return nil, WrapAbort(err)
	}
	return result, nil
}

func TextInput(title string, placeholder string) (string, error) {
	var result string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(title).
				Placeholder(placeholder).
				Value(&result),
		),
	).Run()
	if err != nil {
		return "", WrapAbort(err)
	}
	return result, nil
}

func StringOptions(values ...string) []huh.Option[string] {
	opts := make([]huh.Option[string], len(values))
	for i, v := range values {
		opts[i] = huh.NewOption(v, v)
	}
	return opts
}