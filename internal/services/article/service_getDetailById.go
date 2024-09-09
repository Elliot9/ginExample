package article

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/elliot9/gin-example/internal/dtos"
)

func (s *service) GetDetailByID(id int, onlyActive bool) (*dtos.ArticleWithAuthor, error) {
	cache := s.cache.Tags([]string{"article", "article_" + strconv.Itoa(id)})
	result, err := cache.Remember(context.Background(), "article_"+strconv.Itoa(id), func() (any, error) {
		return s.repo.GetDetailByID(id, onlyActive)
	})

	if err != nil {
		return nil, err
	}

	var articleWithAuthor *dtos.ArticleWithAuthor
	err = json.Unmarshal(result, &articleWithAuthor)
	if err != nil {
		return nil, err
	}
	return articleWithAuthor, nil
}
