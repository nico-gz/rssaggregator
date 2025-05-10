package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"rssgator/internal/database"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	var rssFeed RSSFeed
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(data, &rssFeed); err != nil {
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for idx, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[idx].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[idx].Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)

	feed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		log.Fatal(err)
	}
	timestampFormat := "Sun, 22 Sep 2019 00:00:00 +0000"

	for _, item := range feed.Channel.Item {
		publishedDate, err := time.Parse(item.PubDate, timestampFormat)
		if err != nil {
			fmt.Println("Error parsing publish time")
			fmt.Println(item.PubDate)
		}
		postDescription := sql.NullString{
			String: item.Description,
			Valid:  true,
		}

		if item.Description == "" {
			postDescription.Valid = false
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: postDescription,
			PublishedAt: publishedDate,
			FeedID:      feedToFetch.ID,
		})

		// Scuffed validation for duplicate urls
		errString := fmt.Sprintf("%v", err)
		if err != nil && !strings.Contains(errString, "pq: duplicate key value violates unique constraint") {
			fmt.Println("ERROR while creating post")
			log.Fatal(err)
		}
	}
}
