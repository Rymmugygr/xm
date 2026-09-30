package v1

import (
	"errors"
	"net/http"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/controller/restapi/v1/request"
	"github.com/Rymmugygr/xm/internal/app/controller/restapi/v1/response"
	"github.com/Rymmugygr/xm/internal/app/repository"
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/gofiber/fiber/v2"
)

func (r *V1) getCompany(ctx *fiber.Ctx) error {
	companyId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, " restapi / v1 / getCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid Company ID")
	}

	out, err := r.uc.FindOneCompanyUseCase.Execute(ctx.UserContext(), usecase.FindOneCompanyInput{
		ID: companyId,
	})
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return errorResponse(ctx, http.StatusNotFound, "Invalid Company ID")
		}

		r.l.Error(err, " restapi / v1 / getCompany")

		return errorResponse(ctx, http.StatusInternalServerError, "Something went wrong")
	}

	res := (&response.Company{}).FromEntity(out.Company)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (r *V1) createCompany(ctx *fiber.Ctx) error {
	var body request.CreateCompany

	if err := ctx.BodyParser(&body); err != nil {
		r.l.Error(err, " restapi / v1 / createCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, " restapi / v1 / createCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	out, err := r.uc.CreateCompanyUseCase.Execute(ctx.UserContext(), usecase.CreateCompanyInput{
		Name:              body.Name,
		Description:       body.Description,
		AmountOfEmployees: body.AmountOfEmployees,
		Registered:        body.Registered,
		Type:              body.Type,
	})
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return errorResponse(ctx, http.StatusBadRequest, "Company already existed")
		}

		r.l.Error(err, " restapi / v1 / createCompany")

		return errorResponse(ctx, http.StatusInternalServerError, "Something went wrong")
	}

	res := (&response.Company{}).FromEntity(out.Company)
	return ctx.Status(http.StatusCreated).JSON(res)
}

func (r *V1) pathCompany(ctx *fiber.Ctx) error {
	companyId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, " restapi / v1 / deleteCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid Company ID")
	}

	var body request.PatchCompany

	if err := ctx.BodyParser(&body); err != nil {
		r.l.Error(err, " restapi / v1 / pathCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, " restapi / v1 / pathCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	out, err := r.uc.ModifyCompanyUseCase.Execute(ctx.UserContext(), usecase.ModifyCompanyInput{
		ID:                companyId,
		Name:              body.Name,
		Description:       body.Description,
		AmountOfEmployees: body.AmountOfEmployees,
		Registered:        body.Registered,
		Type:              body.Type,
	})
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return errorResponse(ctx, http.StatusBadRequest, "Company already existed")
		}

		r.l.Error(err, " restapi / v1 / pathCompany")

		return errorResponse(ctx, http.StatusInternalServerError, "Something went wrong")
	}

	res := (&response.Company{}).FromEntity(out.Company)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (r *V1) deleteCompany(ctx *fiber.Ctx) error {
	companyId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, " restapi / v1 / deleteCompany")

		return errorResponse(ctx, http.StatusBadRequest, "Invalid Company ID")
	}

	_, err = r.uc.DeleteCompanyUseCase.Execute(ctx.UserContext(), usecase.DeleteCompanyInput{
		ID: companyId,
	})
	if err != nil {
		r.l.Error(err, " restapi / v1 / deleteCompany")

		return errorResponse(ctx, http.StatusInternalServerError, "Something went wrong")
	}

	return ctx.Status(http.StatusNoContent).JSON(nil)
}
