// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

// GetIsMaxTotalReposSizeReached check is MaxTotalREposSize reached
func GetIsMaxTotalReposSizeReached(username string) (bool, error) {
	// Ask for running deliver hook and test pull request tasks.
	reqURL := setting.LocalURL + fmt.Sprintf("api/internal/user/%s/is-total-repos-size-reached", username)
	log.GitLogger.Trace("GetIsMaxTotalReposSizeReached: %s", reqURL)

	resp, err := newRequest(reqURL, "GET").SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	}).Response()
	if err != nil {
		return true, err
	}

	var isReached bool
	if err := json.NewDecoder(resp.Body).Decode(&isReached); err != nil {
		return true, err
	}

	defer resp.Body.Close()

	// All 2XX status codes are accepted and others will return an error
	if resp.StatusCode/100 != 2 {
		return true, fmt.Errorf("Failed fetch internal information: %s", decodeJSONError(resp).Err)
	}

	return isReached, nil
}
