package zkillboard_redisq

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type ZKillboardRedisQClient struct{}

func NewZKillboardRedisQClient() *ZKillboardRedisQClient {
	return &ZKillboardRedisQClient{}
}

func (z ZKillboardRedisQClient) ListenOnce() (*Package, []error) {
	resp, body, errs := gorequest.New().Get("https://redisq.zkillboard.com/listen.php").End()

	if resp.StatusCode == 429 {
		// Very occasionally we asked to back off. This will wait 30 sec then try again.
		time.Sleep(30 * time.Second)
		resp, body, errs = gorequest.New().Get("https://redisq.zkillboard.com/listen.php").End()
	}

	if len(errs) > 0 {
		return nil, errs
	}

	if resp.StatusCode != 200 {
		return nil, []error{fmt.Errorf("Unexpected status code %d %s", resp.StatusCode, resp.Status)}
	}

	var apiResponse ApiResponse

	err := json.Unmarshal([]byte(body), &apiResponse)
	if err != nil {

		return nil, []error{err}
	}

	if apiResponse.Package != nil {
		return apiResponse.Package, nil
	}

	return nil, []error{}
}

func (z ZKillboardRedisQClient) Listen(killPackages chan *Package, errs chan error) {
	for {
		killPackage, apiErrs := z.ListenOnce()

		if len(apiErrs) > 0 {
			for errI := range apiErrs {
				errs <- apiErrs[errI]
			}
		}

		if killPackage != nil {
			killPackages <- killPackage
		}
	}
}
