package geoserver

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	gs "github.com/camptocamp/go-geoserver/client"
)

func resourceGeoserverLayerRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceGeoserverLayerRuleCreate,
		Read:   resourceGeoserverLayerRuleRead,
		Update: resourceGeoserverLayerRuleUpdate,
		Delete: resourceGeoserverLayerRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGeoserverLayerRuleImport,
		},

		Schema: map[string]*schema.Schema{
			"rule": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ACL Rule. Check https://docs.geoserver.org/latest/en/user/security/layer.html#rules for syntax. Use for ID",
			},
			"authorized_roles": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "List of roles associated to the rule. Use comma to separate roles",
			},
		},
	}
}

func resourceGeoserverLayerRuleCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating ACL Rule: %s", d.Id())

	client := meta.(*Config).GeoserverClient()

	acl_rule := d.Get("rule").(string)

	roles := d.Get("authorized_roles").(string)

	rule := &gs.LayerRule{
		Resource: acl_rule,
		Rule:     roles,
	}

	err := client.CreateLayerRule(rule)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", acl_rule))

	return resourceGeoserverLayerRuleRead(d, meta)
}

func resourceGeoserverLayerRuleRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Refreshing ACL Rule: %s", d.Id())

	acl_rule := d.Id()

	client := meta.(*Config).GeoserverClient()

	rule, err := client.GetLayerRule(acl_rule)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
		return err
	}

	if rule == nil {
		d.SetId("")
		return nil
	}
	d.Set("rule", rule.Resource)
	d.Set("authorized_roles", rule.Rule)

	return nil
}

func resourceGeoserverLayerRuleDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting ACL Rule: %s", d.Id())

	acl_rule := d.Id()

	client := meta.(*Config).GeoserverClient()

	err := client.DeleteLayerRule(acl_rule)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceGeoserverLayerRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating ACL Rule: %s", d.Id())

	client := meta.(*Config).GeoserverClient()

	rule := &gs.LayerRule{
		Resource: d.Get("rule").(string),
		Rule:     d.Get("authorized_roles").(string),
	}

	errUpdateUser := client.UpdateLayerRule(rule)
	if errUpdateUser != nil {
		return errUpdateUser
	}

	return nil
}

func resourceGeoserverLayerRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	d.SetId(d.Id())
	d.Set("rule", d.Id())

	log.Printf("[INFO] Importing ACL Rule `%s`", d.Id())

	err := resourceGeoserverLayerRuleRead(d, meta)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
