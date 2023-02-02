package errs

import (
	"github.com/rendau/dop/dopErrs"
)

const (
	ConfFileNameRequired    = dopErrs.Err("conf_file_name_required")
	FailToSaveFile          = dopErrs.Err("fail_to_save_file")
	FailToSendDeployWebhook = dopErrs.Err("fail_to_send_deploy_webhook")
)
