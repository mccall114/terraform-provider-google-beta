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
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeRegionUrlMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionUrlMapCreate,
		Read:   resourceComputeRegionUrlMapRead,
		Update: resourceComputeRegionUrlMapUpdate,
		Delete: resourceComputeRegionUrlMapDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionUrlMapImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_service": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     computeRegionUrlMapHostRuleSchema(),
				// Default schema.HashSchema is used.
			},
			"path_matcher": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_service": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"path_rule": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"paths": {
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
									},
									"service": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: compareSelfLinkOrResourceName,
									},
								},
							},
						},
					},
				},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"test": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"map_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func computeRegionUrlMapHostRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"path_matcher": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeRegionUrlMapCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeRegionUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(defaultServiceProp)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeRegionUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hostRulesProp, err := expandComputeRegionUrlMapHostRule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(hostRulesProp)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	fingerprintProp, err := expandComputeRegionUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeRegionUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeRegionUrlMapPathMatcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(pathMatchersProp)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeRegionUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(testsProp)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}
	regionProp, err := expandComputeRegionUrlMapRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionUrlMap: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionUrlMap: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating RegionUrlMap",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionUrlMap: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating RegionUrlMap %q: %#v", d.Id(), res)

	return resourceComputeRegionUrlMapRead(d, meta)
}

func resourceComputeRegionUrlMapRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionUrlMap %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionUrlMapCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("default_service", flattenComputeRegionUrlMapDefaultService(res["defaultService"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionUrlMapDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("host_rule", flattenComputeRegionUrlMapHostRule(res["hostRules"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("map_id", flattenComputeRegionUrlMapMapId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeRegionUrlMapFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionUrlMapName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("path_matcher", flattenComputeRegionUrlMapPathMatcher(res["pathMatchers"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("test", flattenComputeRegionUrlMapTest(res["tests"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionUrlMapRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionUrlMap: %s", err)
	}

	return nil
}

func resourceComputeRegionUrlMapUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeRegionUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeRegionUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hostRulesProp, err := expandComputeRegionUrlMapHostRule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	fingerprintProp, err := expandComputeRegionUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeRegionUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeRegionUrlMapPathMatcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeRegionUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}
	regionProp, err := expandComputeRegionUrlMapRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionUrlMap %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionUrlMap %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating RegionUrlMap",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeRegionUrlMapRead(d, meta)
}

func resourceComputeRegionUrlMapDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionUrlMap %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionUrlMap")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting RegionUrlMap",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionUrlMap %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionUrlMapImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/urlMaps/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionUrlMapCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionUrlMapDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapHostRule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeRegionUrlMapHostRuleSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"description":  flattenComputeRegionUrlMapHostRuleDescription(original["description"], d),
			"hosts":        flattenComputeRegionUrlMapHostRuleHosts(original["hosts"], d),
			"path_matcher": flattenComputeRegionUrlMapHostRulePathMatcher(original["pathMatcher"], d),
		})
	}
	return transformed
}
func flattenComputeRegionUrlMapHostRuleDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapHostRuleHosts(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeRegionUrlMapHostRulePathMatcher(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapMapId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionUrlMapFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapPathMatcher(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"default_service": flattenComputeRegionUrlMapPathMatcherDefaultService(original["defaultService"], d),
			"description":     flattenComputeRegionUrlMapPathMatcherDescription(original["description"], d),
			"name":            flattenComputeRegionUrlMapPathMatcherName(original["name"], d),
			"path_rule":       flattenComputeRegionUrlMapPathMatcherPathRule(original["pathRules"], d),
		})
	}
	return transformed
}
func flattenComputeRegionUrlMapPathMatcherDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionUrlMapPathMatcherDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapPathMatcherName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapPathMatcherPathRule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"paths":   flattenComputeRegionUrlMapPathMatcherPathRulePaths(original["paths"], d),
			"service": flattenComputeRegionUrlMapPathMatcherPathRuleService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeRegionUrlMapPathMatcherPathRulePaths(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeRegionUrlMapPathMatcherPathRuleService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionUrlMapTest(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"description": flattenComputeRegionUrlMapTestDescription(original["description"], d),
			"host":        flattenComputeRegionUrlMapTestHost(original["host"], d),
			"path":        flattenComputeRegionUrlMapTestPath(original["path"], d),
			"service":     flattenComputeRegionUrlMapTestService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeRegionUrlMapTestDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapTestHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapTestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionUrlMapTestService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionUrlMapRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeRegionUrlMapDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("backendServices", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for default_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionUrlMapDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapHostRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeRegionUrlMapHostRuleDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHosts, err := expandComputeRegionUrlMapHostRuleHosts(original["hosts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHosts); val.IsValid() && !isEmptyValue(val) {
			transformed["hosts"] = transformedHosts
		}

		transformedPathMatcher, err := expandComputeRegionUrlMapHostRulePathMatcher(original["path_matcher"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPathMatcher); val.IsValid() && !isEmptyValue(val) {
			transformed["pathMatcher"] = transformedPathMatcher
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionUrlMapHostRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapHostRuleHosts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionUrlMapHostRulePathMatcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapPathMatcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDefaultService, err := expandComputeRegionUrlMapPathMatcherDefaultService(original["default_service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDefaultService); val.IsValid() && !isEmptyValue(val) {
			transformed["defaultService"] = transformedDefaultService
		}

		transformedDescription, err := expandComputeRegionUrlMapPathMatcherDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedName, err := expandComputeRegionUrlMapPathMatcherName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedPathRule, err := expandComputeRegionUrlMapPathMatcherPathRule(original["path_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPathRule); val.IsValid() && !isEmptyValue(val) {
			transformed["pathRules"] = transformedPathRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionUrlMapPathMatcherDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("backendServices", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for default_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionUrlMapPathMatcherDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapPathMatcherName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapPathMatcherPathRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPaths, err := expandComputeRegionUrlMapPathMatcherPathRulePaths(original["paths"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPaths); val.IsValid() && !isEmptyValue(val) {
			transformed["paths"] = transformedPaths
		}

		transformedService, err := expandComputeRegionUrlMapPathMatcherPathRuleService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionUrlMapPathMatcherPathRulePaths(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionUrlMapPathMatcherPathRuleService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("backendServices", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionUrlMapTest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeRegionUrlMapTestDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHost, err := expandComputeRegionUrlMapTestHost(original["host"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
			transformed["host"] = transformedHost
		}

		transformedPath, err := expandComputeRegionUrlMapTestPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandComputeRegionUrlMapTestService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionUrlMapTestDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapTestHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapTestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionUrlMapTestService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("backendServices", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionUrlMapRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
