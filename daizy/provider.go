package daizy

import (
	"context"
	"github.com/cuotos/daizy-go/daizy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"organisation": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authtoken": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"daizy_project": resourceProject(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: func(context context.Context, data *schema.ResourceData) (i interface{}, diagnostics diag.Diagnostics) {
			client, err := daizy.New(data.Get("organisation").(string), data.Get("authtoken").(string))
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return client, nil
		},
	}
}
