package govalidator

import (
	"context"
	gov "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	val *gov.Validate
}

func (v *Validator) Validate(ctx context.Context, s any) error {
	return v.val.StructCtx(ctx, s)
}

func (v *Validator) ValidateFiberRequestBody(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.BodyParser(s); err != nil {
		return err
	}

	return v.Validate(ctx.Context(), s)
}

func (v *Validator) ValidateFiberQueryParams(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.QueryParser(s); err != nil {
		return err
	}

	return v.Validate(ctx.Context(), s)
}

func (v *Validator) ValidateFiberParams(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.ParamsParser(s); err != nil {
		return err
	}

	return v.Validate(ctx.Context(), s)
}

func New() *Validator {
	return &Validator{gov.New()}
}
