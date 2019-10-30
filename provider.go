package main

import (
	"github.com/ernesto-arm/go-coveo/sourceapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"organization_id":{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COVEO_ORG_ID", nil),
			},
			"api_key":{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COVEO_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"coveo_source": resourceSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData)(interface{}, error){
	log.Printf("[INFO] getting the configuration ")

	pushConfig := sourceapi.Config{
		OrganizationID: d.Get("organization_id").(string),
		APIKey:         d.Get("api_key").(string),
	}

	client, err := sourceapi.NewClient(pushConfig)

	if err != nil {
		log.Printf("[ERROR] Coveo client errord with %s", err.Error())
		return nil, err
	}

	log.Println("[INFO] Coveo client initialised")

	return client, nil
}