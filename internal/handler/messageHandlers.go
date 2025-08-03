package handler

import (
	"net/http"
	messageserver "pod/internal/messageServer"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type MessageHandler struct {
	service messageserver.MessageService
}

func NewMessageHandler(m messageserver.MessageService) *MessageHandler {
	return &MessageHandler{service: m}
}

func (h *MessageHandler) GetHandler(c echo.Context) error {
	messages, err := h.service.GetAllMessages()

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not get",
		})
	}

	return c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) PostHandler(c echo.Context) error {
	type request struct {
		Text string `json:"text"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid request body",
		})
	}

	msg, err := h.service.CreateMessage(req.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Failed to create message: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, msg)
}

func (h *MessageHandler) PutHandler(c echo.Context) error {
	idParam := c.Param("id")

	type request struct {
		Text string `json:"text"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid request body",
		})
	}

	err := h.service.UpdateMessage(idParam, req.Text)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "message not found" {
			status = http.StatusNotFound
		}

		return c.JSON(status, Response{
			Status:  "Error",
			Message: "Failed to update message: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, req.Text)
}

func (h *MessageHandler) DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")

	if err := h.service.DeleteMessage(idParam); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "message not found" {
			status = http.StatusNotFound
		}

		return c.JSON(status, Response{
			Status:  "Error",
			Message: "Failed to delete message: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message deleted successfully",
	})
}
