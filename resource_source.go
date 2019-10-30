package main

import (
	"encoding/json"
	"fmt"
	"github.com/ernesto-arm/go-coveo/sourceapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSourceCreate,
		Read:   resourceSourceRead,
		Update: resourceSourceUpdate,
		Delete: resourceSourceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"visibility": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"push_enabled": &schema.Schema {
				Type: schema.TypeBool,
				Optional: true,
				Default: false,
			},
		},
	}
}

func resourceSourceCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Creating resource %s", d.Get("name"))
	Source := sourceapi.Source{
		Name:       d.Get("name").(string),
		Type:       strings.ToUpper(d.Get("type").(string)),
		Visibility: strings.ToUpper(d.Get("visibility").(string)),
		Enabled:    d.Get("push_enabled").(bool),
	}
	resp, err := m.(sourceapi.Client).CreateSource(Source)
	if err != nil {
		return fmt.Errorf("error when trying to create a source %s with message %s", d.Get("name"), err.Error())
	}

	src := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &src)
	if err != nil {
		return fmt.Errorf("error unmarshalling the source response %s", resp)
	}

	log.Printf("Unmarshal resource %s", src)

	d.SetId(src["id"].(string))
	_ = d.Set("name", src["name"])
	_ = d.Set("type", src["sourceType"])
	_ = d.Set("visibility", src["sourceVisibility"])
	_ = d.Set("push_enabled", src["pushEnabled"])
	return nil
}

func resourceSourceRead(d *schema.ResourceData, m interface{}) error {
	// Attempt to read from an upstream API
	s, err := m.(sourceapi.Client).ReadSource(d.Id())

	if err != nil {
		d.SetId("")
		return fmt.Errorf("source id %s errored with message %s", d.Id(), err.Error())
	}

	src := make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &src)
	if err != nil {
		return fmt.Errorf("source id %s errored with message %s", d.Id(), err.Error())
	}

	_ = d.Set("name", src["name"])
	_ = d.Set("type", src["sourceType"])
	_ = d.Set("visibility", src["sourceVisibility"])
	_ = d.Set("push_enabled", src["pushEnabled"])
	return nil
}

func resourceSourceUpdate(d *schema.ResourceData, m interface{}) error {
	Source := sourceapi.Source{
		Id: 		d.Id(),
		Name:       d.Get("name").(string),
		Type:       strings.ToUpper(d.Get("type").(string)),
		Visibility: strings.ToUpper(d.Get("visibility").(string)),
		Enabled:    d.Get("push_enabled").(bool),
	}
	resp, err := m.(sourceapi.Client).UpdateSource(d.Id(), Source)
	if err != nil {
		return fmt.Errorf("error when trying to update a source %s with message %s", d.Id(), err.Error())
	}

	src := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &src)
	if err != nil {
		return fmt.Errorf("error unmarshalling the source %s with message %s", d.Id(), resp)
	}

	_ = d.Set("name", src["name"])
	_ = d.Set("type", src["sourceType"])
	_ = d.Set("visibility", src["sourceVisibility"])
	_ = d.Set("push_enabled", src["pushEnabled"])
	return nil
}


func resourceSourceDelete(d *schema.ResourceData, m interface{}) error {
	err := m.(sourceapi.Client).DeleteSource(d.Id())
	if err != nil {
		return fmt.Errorf("error deleting the source %s with message %s", d.Id(), err.Error())
	}
	d.SetId("")
	return nil
}