package duo

import (
	"encoding/json"
	"fmt"
	"github.com/duosecurity/duo_api_golang"
	admin "github.com/duosecurity/duo_api_golang/admin"
	"github.com/hashicorp/terraform/helper/schema"
	"net/url"
)

type UserEnrollResult struct {
	Stat string
}

func resourceUserEnrollCreate(d *schema.ResourceData, meta interface{}) error {
	duoclient := meta.(*duoapi.DuoApi)
	duoAdminClient := admin.New(*duoclient)

	params := url.Values{}
	params.Set("username", d.Get("username").(string))
	params.Set("email", d.Get("email").(string))
	_, enrollmentBody, err := duoAdminClient.SignedCall("POST", "/admin/v1/users/enroll", params, duoapi.UseTimeout)
	if err != nil {
		return err
	}
	result := &UserEnrollResult{}
	err = json.Unmarshal(enrollmentBody, result)
	if err != nil {
		return err
	}
	if result.Stat != "OK" {
		return fmt.Errorf("could not enroll user %s", result.Stat)
	}

	return nil
}
