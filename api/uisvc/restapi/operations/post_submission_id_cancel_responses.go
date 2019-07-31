package operations

import (
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/handlers"

	"github.com/gin-gonic/gin"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PostSubmissionIDCancelResponse responds for PostSubmissionIDCancel.
func PostSubmissionIDCancelResponse(h *handlers.H, ctx *gin.Context, resp interface{}, code int, err *errors.Error) *errors.Error {
	if err != nil {
		h.LogError(err, ctx, code)
		return err
	}

	ctx.JSON(code, resp)

	return nil
}
