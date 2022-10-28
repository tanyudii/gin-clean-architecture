package cmd

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"github.com/vodeacloud/hr-api/cmd/prompt"
	"github.com/vodeacloud/hr-api/concrete/company"
	"github.com/vodeacloud/hr-api/concrete/user"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/pkg/logger"
)

// companyRegisterCmd represents the companyRegister command
var companyRegisterCmd = &cobra.Command{
	Use:   "create",
	Short: "Register new Company",
	Run:   runCompanyRegister,
}

func init() {
	companyCmd.AddCommand(companyRegisterCmd)
}

func runCompanyRegister(_ *cobra.Command, _ []string) {
	namePc := prompt.PromptContent{Label: "Name: ", ErrMsg: "name field is required"}
	adminEmailPc := prompt.PromptContent{
		Label: "Admin Email: ",
		Validate: func(val string) error {
			if !govalidator.IsEmail(val) {
				return errors.New("email format is invalid")
			}
			return nil
		},
	}
	adminPasswordPc := prompt.PromptContent{
		Label:  "Admin Password: ",
		ErrMsg: "admin password field is required",
		Mask:   '*',
	}

	registerCompanyReq := &entities.RegisterCompanyRequest{
		Name:          namePc.GetString(),
		AdminEmail:    adminEmailPc.GetString(),
		AdminPassword: adminPasswordPc.GetString(),
	}

	executeCompanyRegister(registerCompanyReq)
}

func executeCompanyRegister(req *entities.RegisterCompanyRequest) {
	db := config.GetDatabase()
	defer config.CloseDatabase(db)

	companyRepo := company.NewRepository(db)
	userRepo := user.NewRepository(db)

	userUc := user.NewUsecase(userRepo, companyRepo)
	companyUc := company.NewUsecase(companyRepo, userUc)

	err := companyUc.RegisterCompany(context.Background(), req)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
