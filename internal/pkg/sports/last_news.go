package sports

import (
	"strconv"
	"time"
)

const lastNewsQuery = `
	query ($id: ID!, $limit: Int!) {
		newsQueries {
			newsByTags(IDs: [$id], amount: $limit, filter: ALL, noAds: true, source: ALL) {
				news {
					title
					commentsCount
					publishedAt
				}
			}
		}
	}
`

type News struct {
	Title         string    `json:"title"`
	CommentsCount int       `json:"commentsCount"`
	PublishedAt   time.Time `json:"publishedAt"`
}

func LastNews(tagID, limit int) ([]News, error) {
	var resp struct {
		Data struct {
			NewsQueries struct {
				NewsByTags struct {
					News []News
				} `json:"newsByTags"`
			} `json:"newsQueries"`
		} `json:"data"`
	}
	err := request(lastNewsQuery, &resp, map[string]any{
		"id":    strconv.Itoa(tagID),
		"limit": limit,
	})
	if err != nil {
		return nil, err
	}

	return resp.Data.NewsQueries.NewsByTags.News, nil
}
