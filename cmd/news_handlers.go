package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListNews(c *fiber.Ctx) error {
	var news NewsListResponse
	var limit, offset int
	var err error

	log.Printf("ListNews: start %s", c.Request())

	limit, err = strconv.Atoi(c.Query("limit", "5"))
	if err != nil || limit < 1 {
		log.Printf("ListNews: err or wrong limit value %s %d %s", err, limit, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, "Limit must be gte 1 or not selected at all")
	}

	offset, err = strconv.Atoi(c.Query("offset", "0"))
	if err != nil || offset < 0 {
		log.Printf("ListNews: err or wrong offset value %s %d %s", err, limit, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, "Offset must be gte 0 or not selected at all")
	}

	query := `
		SELECT
			NEWS.ID,
			NEWS.TITLE,
			NEWS.CONTENT,
			COALESCE(array_agg(NEWS_CATEGORIES.ID), '{}'::bigint[])
		FROM
			NEWS
		LEFT JOIN
			NEWS_CATEGORIES
			ON NEWS.ID = NEWS_CATEGORIES.NEWS_ID
		GROUP BY
			NEWS.ID, NEWS.TITLE, NEWS.CONTENT
		LIMIT $1
		OFFSET $2
	`

	rows, err := MainDB.Query(query, limit, offset)
	if err != nil {
		log.Printf("ListNews: db err %s %s", err, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var nar NewsAggregateResponse
		if err := rows.Scan(&nar.ID, &nar.Title, &nar.Content, &nar.Categories); err != nil {
			log.Printf("ListNews: db err %s %s", err, c.Request())
			return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
		}

		news.News = append(news.News, nar)
	}
	log.Printf("ListNews: end %s", c.Request())

	return c.Status(fiber.StatusOK).JSON(news)
}

func EditNews(c *fiber.Ctx) error {
	log.Printf("ListNews: start %s", c.Request())

	paramNewsID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("EditNews: no struct passed from context %s", c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}

	news, ok := c.Locals(NewsKey).(NewsEditRequest)
	if !ok {
		log.Printf("EditNews: no struct passed from context %s", c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}

	tx, err := MainDB.Begin()
	if err != nil {
		log.Printf("EditNews: err starting transaction %s %s", err, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}

	// we either update in case of fields difference
	// or leave it as is in case of equality
	query := `
		UPDATE NEWS 
		SET 
			TITLE = CASE WHEN $1 = '' THEN TITLE ELSE $1 END,
			CONTENT = CASE WHEN $2 = '' THEN CONTENT ELSE $2 END
		WHERE ID = $3;
	`
	_, err = tx.Exec(query, news.Title, news.Content, paramNewsID)
	if err != nil {
		tx.Rollback()
		log.Printf("EditNews: err updating news table transaction %s %s", err, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}

	if news.Categories != nil {
		// guess there is not point in searching for keys and updating it
		// seems more efficent to re-write data based on input
		query = `DELETE FROM NEWS_CATEGORIES WHERE NEWS_ID = $1;`

		_, err = MainDB.Exec(query, paramNewsID)
		if err != nil {
			tx.Rollback()
			log.Printf("EditNews: err deleting news_categories table transaction %s %s", err, c.Request())
			return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
		}

		// insert new values instead of old ones
		// this might be a potential injection vulnerability
		// but since we added dive validation on "categories" field, we are fine
		query = fmt.Sprintf(`INSERT INTO NEWS_CATEGORIES (NEWS_ID, ID) VALUES %s`, news.CategoriesAsDBValues())
		_, err = MainDB.Exec(query, news.CategoriesAsDBArgs(paramNewsID)...)
		if err != nil {
			tx.Rollback()
			log.Printf("EditNews: err inserting news_categories table transaction %s %s %v", err, c.Request(), news)
			return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("EditNews: err commiting transaction %s %s", err, c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, ErrInternalDatabase.Error())
	}
	log.Printf("EditNews: end %s", c.Request())

	return c.SendStatus(fiber.StatusOK)
}
