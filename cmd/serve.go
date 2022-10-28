package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"github.com/vodeacloud/hr-api/concrete/company"
	"github.com/vodeacloud/hr-api/concrete/email"
	emailtemplate "github.com/vodeacloud/hr-api/concrete/email_template"
	"github.com/vodeacloud/hr-api/concrete/group"
	"github.com/vodeacloud/hr-api/concrete/notification"
	notificationconfig "github.com/vodeacloud/hr-api/concrete/notification_config"
	"github.com/vodeacloud/hr-api/concrete/oauth"
	"github.com/vodeacloud/hr-api/concrete/password"
	"github.com/vodeacloud/hr-api/concrete/permission"
	"github.com/vodeacloud/hr-api/concrete/role"
	"github.com/vodeacloud/hr-api/concrete/user"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/handler/api"
	"github.com/vodeacloud/hr-api/pkg/gracefully"
	"gorm.io/gorm"
	"strconv"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HTTP server",
	Run:   runGin,
}

var (
	RESTPort int
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&RESTPort, "port", "p", 80, "REST Port")
}

func runGin(_ *cobra.Command, _ []string) {
	r := gin.Default()
	r.Use(gin.Recovery())

	db := config.GetDatabase()
	defer config.CloseDatabase(db)

	redisCli := config.GetRedisCli()
	defer config.CloseRedis(redisCli)

	registerRouter(r, db, redisCli)

	gracefully.RunGinGracefully(r, strconv.Itoa(RESTPort))
}

func registerRouter(r *gin.Engine, db *gorm.DB, redisCli *redis.Client) {
	baseGroup := r.Group("/")
	apiGroup := baseGroup.Group("api")

	companyRepo := company.NewRepository(db)
	userRepo := user.NewRepository(db)
	permissionRepo := permission.NewRepository(db)
	roleRepo := role.NewRepository(db)
	groupRepo := group.NewRepository(db)
	emailTemplateRepo := emailtemplate.NewRepository(db)
	notificationRepo := notification.NewRepository(db)
	notificationConfigRepo := notificationconfig.NewRepository(db)
	oauthTokenRepo := oauth.NewTokenRepository(redisCli)
	oauthClientRepo := oauth.NewClientRepository(db)
	passwordRepo := password.NewRepository(redisCli)

	userUc := user.NewUsecase(userRepo, companyRepo)
	companyUc := company.NewUsecase(companyRepo, userUc)
	roleUc := role.NewUsecase(roleRepo, permissionRepo)
	groupUc := group.NewUsecase(groupRepo)
	emailUc := email.NewUsecase(emailTemplateRepo)
	notificationUc := notification.NewUsecase(notificationRepo, notificationConfigRepo)
	notificationConfigUc := notificationconfig.NewUsecase(notificationConfigRepo)
	oauthUc := oauth.NewUsecase(oauthTokenRepo, oauthClientRepo, userUc, roleUc)
	passwordUc := password.NewUsecase(passwordRepo, userRepo, emailUc)
	permissionUc := permission.NewUsecase(permissionRepo)

	api.NewRegisterHealthAPI(baseGroup)
	api.NewRegisterCompanyAPI(apiGroup, companyUc)
	api.NewRegisterUserAPI(apiGroup, userUc)
	api.NewRegisterRoleAPI(apiGroup, roleUc)
	api.NewRegisterGroupAPI(apiGroup, groupUc)
	api.NewRegisterNotificationAPI(apiGroup, notificationUc)
	api.NewRegisterNotificationConfigAPI(apiGroup, notificationConfigUc)
	api.NewRegisterOAuthAPI(baseGroup, oauthUc)
	api.NewRegisterPasswordAPI(apiGroup, passwordUc)
	api.NewRegisterPermissionAPI(apiGroup, permissionUc)
}
