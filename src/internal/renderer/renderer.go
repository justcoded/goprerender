package renderer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"net/http"
	"os"
	"prerender/internal/archive"
	"prerender/internal/cachers"
	"prerender/internal/urlparser"
	"prerender/pkg/log"
	"strconv"
	"time"
)

func DoRender(ctx context.Context, queryString string, pc cachers.Сacher, force bool, logger log.Logger) (string, error) {
	waitSecondsStr, exists := os.LookupEnv("PAGE_WAIT_TIME")
	waitSeconds := 5

	if exists && waitSecondsStr != "" {
		var err error
		waitSeconds, err = strconv.Atoi(waitSecondsStr)

		if err != nil {
			logger.Error(err)
		}
	}

	requestURL, hostPath, err := urlparser.ParseUrl(queryString)
	if err != nil {
		logger.Error(err)
	}

	key, err := urlparser.GetHashKey(queryString)
	if err != nil {
		logger.Error(err)
	}

	value, err := pc.Get(key)
	var res string

	if force || err != nil {
		err := chromedp.Run(ctx,
			chromedp.Navigate(requestURL),
			chromedp.ActionFunc(func(ctx context.Context) error {
				logger.Infof("Waiting %v sec for rendering...", time.Second*time.Duration(waitSeconds))

				return nil
			}),
			chromedp.Sleep(time.Second*time.Duration(waitSeconds)),
			chromedp.ActionFunc(func(ctx context.Context) error {
				node, err := dom.GetDocument().Do(ctx)
				if err != nil {
					return err
				}
				res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
				return err
			}),
		)

		if err != nil {
			fmt.Println(err)
			return "", err
		}
		htmlGzip := archive.GzipHtml(res, hostPath, "", logger)
		err = pc.Put(key, htmlGzip)
		if err != nil {
			return "", err
		}
	} else {
		res = archive.UnzipHtml(value, logger)
	}

	return res, nil
}

func GetDebugURL(logger log.Logger) (string, error) {
	resp, err := http.Get("http://localhost:9222/json/version")
	if err != nil {
		logger.Warn(err)
		return "", err
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Error(err)
		return "", err
	}
	return result["webSocketDebuggerUrl"].(string), nil
}
