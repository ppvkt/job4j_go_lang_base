package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_validate(t *testing.T) {
	t.Parallel()

	t.Run("ptr validation - ReqValidateMessage", func(t *testing.T) {
		t.Parallel()

		var req *base.ValidateRequest

		expected := []string{base.ReqValidateMessage}
		result := base.Validate(req)

		assert.Equal(t, expected, result)
	})

	t.Run("success validation - empty slice", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "1",
			Title:       "1",
			Description: "1",
		}

		expected := []string{}
		result := base.Validate(req)

		assert.Equal(t, expected, result)
	})

	t.Run("UserId validation - UserIdValidateMessage", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			Title:       "1",
			Description: "1",
		}

		expected := []string{base.UserIdValidateMessage}
		result := base.Validate(req)

		assert.Equal(t, expected, result)
	})

	t.Run("Title validation - TitleValidateMessage", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "1",
			Description: "1",
		}

		expected := []string{base.TitleValidateMessage}
		result := base.Validate(req)

		assert.Equal(t, expected, result)
	})

	t.Run("Description validation - DescriptionValidateMessage", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId: "1",
			Title:  "1",
		}

		expected := []string{base.DescriptionValidateMessag}
		result := base.Validate(req)

		assert.Equal(t, expected, result)
	})

	t.Run(
		"Empty struct validation - UserIdValidateMessage, DescriptionValidateMessage, TitleValidateMessage",
		func(t *testing.T) {
			t.Parallel()

			req := &base.ValidateRequest{}

			expected := []string{base.UserIdValidateMessage, base.TitleValidateMessage, base.DescriptionValidateMessag}
			result := base.Validate(req)

			assert.Equal(t, expected, result)
		})

}
