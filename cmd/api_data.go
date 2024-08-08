package main

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

var validate = validator.New()

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GenericValidator struct {
	// validator *validator.Validate
}

func (gv *GenericValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

type NewsCategories struct {
	ID     int64 `json:"id"`
	NewsID int64 `json:"news_id"`
}

type News struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsEditRequest struct {
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	Categories pq.Int64Array `json:"categories" validate:"dive,gte=0"`
}

func (ner *NewsEditRequest) CategoriesAsDBValues() string {
	// INSERT INTO NEWS_CATEGORIES (NEWS_ID, ID) VALUES ($1, 1), ($1, 2)
	var values []string

	for i := 1; i <= len(ner.Categories); i++ {
		values = append(values, fmt.Sprintf("($1, $%d)", i+1))
	}

	return strings.Join(values, ", ")
}

func (ner *NewsEditRequest) CategoriesAsDBArgs(newsID int) []interface{} {
	var args []interface{}

	args = append(args, newsID)
	for _, category := range ner.Categories {
		args = append(args, category)
	}
	return args
}

type NewsAggregateResponse struct {
	ID         int64         `json:"id"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	Categories pq.Int64Array `json:"categories"`
}

type NewsListResponse struct {
	Success bool                    `json:"success"`
	News    []NewsAggregateResponse `json:"news"`
}
