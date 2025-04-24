package apis

import (
	"github.com/mss-boot-io/mss-boot/pkg/response/actions"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"github.com/mss-boot-io/mss-boot/pkg/response/controller"

	"nsqauthd/config"
	"nsqauthd/dto"
	"nsqauthd/models"
)

func init() {
	e := &Auth{
		Simple: controller.NewSimple(
			controller.WithAuth(false),
			controller.WithModelProvider(actions.ModelProviderGorm),
		),
	}
	response.AppendController(e)
}

type Auth struct {
	*controller.Simple
}

func (e *Auth) GetAction(_ string) response.Action {
	return nil
}

func (e *Auth) Other(r *gin.RouterGroup) {
	r.GET("/auth", e.Auth)
}

// Auth authenticate
// @Summary authenticate
// @Description authenticate
// @Tags auth
// @Accept application/json
// @Product application/json
// @Param secret query string true "secret"
// @Param remote_ip query string true "remote_ip"
// @Param tls query bool false "tls"
// @Success 200 {object} dto.AuthResp
// @Router /auth [get]
func (e *Auth) Auth(ctx *gin.Context) {
	api := response.Make(ctx)
	req := &dto.AuthReq{}
	if api.Bind(req).Error != nil {
		api.Err(http.StatusForbidden, "NOT_AUTHORIZED")
		return
	}
	resp := &dto.AuthResp{
		Identity:       req.Secret,
		TTL:            config.Cfg.TTL,
		Authorizations: make([]*dto.AuthAccount, 0),
	}
	auth := models.Auth{
		Login: req.Secret,
	}
	err := auth.GetByLogin(ctx, gormdb.DB)
	if err != nil {
		api.AddError(err).Log.Error("get auth failed", "err", err)
		api.Err(http.StatusForbidden, "NOT_AUTHORIZED")
		return
	}
	if req.TLS != auth.TlsRequired {
		api.Err(http.StatusForbidden, "NOT_AUTHORIZED")
		return
	}
	authAccount := &dto.AuthAccount{}
	authAccount.Channels, authAccount.Topic, authAccount.Permissions = auth.ToAuthPermissions()
	resp.Authorizations = append(resp.Authorizations, authAccount)
	api.OK(resp)
}
