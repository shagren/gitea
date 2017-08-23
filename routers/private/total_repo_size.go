// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"code.gitea.io/gitea/models"

	macaron "gopkg.in/macaron.v1"
	"code.gitea.io/gitea/modules/log"
)

// GetIsMaxTotalReposSizeReached return if user reach MaxTotalReposSize limit
func GetIsMaxTotalReposSizeReached(ctx *macaron.Context) {
	username := ctx.Params(":username")
	user, err := models.GetUserByName(username)
	if err != nil {
		ctx.JSON(500, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}
	isReposTotalSizeLimitReached, err := user.IsReposTotalSizeLimitReached()
	log.Info("Result is %b", isReposTotalSizeLimitReached)
	if err != nil {
		ctx.JSON(500, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, isReposTotalSizeLimitReached)
	return
}
