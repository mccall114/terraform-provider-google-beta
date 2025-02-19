// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var BinaryAuthorizationAttestorIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"attestor": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type BinaryAuthorizationAttestorIamUpdater struct {
	project  string
	attestor string
	d        *schema.ResourceData
	Config   *Config
}

func BinaryAuthorizationAttestorIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project
	if v, ok := d.GetOk("attestor"); ok {
		values["attestor"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/attestors/(?P<attestor>[^/]+)", "(?P<project>[^/]+)/(?P<attestor>[^/]+)", "(?P<attestor>[^/]+)"}, d, config, d.Get("attestor").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BinaryAuthorizationAttestorIamUpdater{
		project:  values["project"],
		attestor: values["attestor"],
		d:        d,
		Config:   config,
	}

	d.Set("project", u.project)
	d.Set("attestor", u.GetResourceId())

	d.SetId(u.GetResourceId())

	return u, nil
}

func BinaryAuthorizationAttestorIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/attestors/(?P<attestor>[^/]+)", "(?P<project>[^/]+)/(?P<attestor>[^/]+)", "(?P<attestor>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BinaryAuthorizationAttestorIamUpdater{
		project:  values["project"],
		attestor: values["attestor"],
		d:        d,
		Config:   config,
	}
	d.Set("attestor", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url := u.qualifyAttestorUrl("getIamPolicy")

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}

	policy, err := sendRequest(u.Config, "GET", project, url, nil)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *BinaryAuthorizationAttestorIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url := u.qualifyAttestorUrl("setIamPolicy")

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *BinaryAuthorizationAttestorIamUpdater) qualifyAttestorUrl(methodIdentifier string) string {
	return fmt.Sprintf("https://binaryauthorization.googleapis.com/v1beta1/%s:%s", fmt.Sprintf("projects/%s/attestors/%s", u.project, u.attestor), methodIdentifier)
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s/%s", u.project, u.attestor)
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-binaryauthorization-attestor-%s", u.GetResourceId())
}

func (u *BinaryAuthorizationAttestorIamUpdater) DescribeResource() string {
	return fmt.Sprintf("binaryauthorization attestor %q", u.GetResourceId())
}
