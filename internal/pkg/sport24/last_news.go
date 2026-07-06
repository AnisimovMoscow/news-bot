package sport24

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	newsApi     = "/news/v1/materials/feed"
	commentsApi = "/community-service/v1/comment/count"

	newsURL = "https://sport24.ru/football/news-%s"
)

type UnixTimestamp struct {
	time.Time
}

func (ut *UnixTimestamp) UnmarshalJSON(b []byte) error {
	var msec int64
	if err := json.Unmarshal(b, &msec); err != nil {
		return err
	}
	ut.Time = time.UnixMilli(msec)
	return nil
}

type News struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	CommentsCount int
	PublishedAt   UnixTimestamp `json:"publishDate"`
	Urn           string        `json:"urn"`
}

func (n News) URL() string {
	return fmt.Sprintf(newsURL, n.Urn)
}

func LastNews(tagID, limit int) ([]News, error) {
	var resp struct {
		Items []News `json:"items"`
	}
	err := api(apiHost+newsApi, []Param{
		{
			Key:   "limit",
			Value: strconv.Itoa(limit),
		},
		{
			Key:   "tag",
			Value: strconv.Itoa(tagID),
		},
		{
			Key:   "type",
			Value: "NEWS",
		},
	}, &resp)
	if err != nil {
		return nil, err
	}

	// подгружаем комменты
	news, err := loadComments(resp.Items)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func loadComments(news []News) ([]News, error) {
	ids := make([]string, len(news))
	for i, n := range news {
		ids[i] = fmt.Sprintf("news:%d", n.ID)
	}
	var resp struct {
		Items []struct {
			ResourceID string `json:"resourceId"`
			Count      int    `json:"count"`
		} `json:"items"`
	}
	err := api(apiHost+commentsApi, []Param{
		{
			Key:   "id",
			Value: strings.Join(ids, ","),
		},
	}, &resp)
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int, len(resp.Items))
	for _, item := range resp.Items {
		counts[item.ResourceID] = item.Count
	}

	for i, n := range news {
		resource := fmt.Sprintf("news:%d", n.ID)
		count, ok := counts[resource]
		if ok {
			news[i].CommentsCount = count
		}
	}

	return news, nil
}
