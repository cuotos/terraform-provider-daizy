package daizy

import (
	"context"
	"fmt"
	"github.com/cuotos/daizy-go/daizy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		UpdateContext: resourceProjectUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*daizy.API)

	project := daizy.Project{
		Name:   data.Get("name").(string),
		UserID: data.Get("user_id").(int),
	}

	p, err := client.CreateProject(project)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(fmt.Sprint(p.ID))

	return nil
}

func resourceProjectUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*daizy.API)

	id, err := strconv.Atoi(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	updateProject := daizy.Project{
		Name:   data.Get("name").(string),
		UserID: data.Get("user_id").(int),
	}

	respProject, err := client.UpdateProject(id, updateProject)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(fmt.Sprint(respProject.ID))
	data.Set("name", respProject.Name)
	data.Set("user_id", respProject.UserID)

	return nil
}

func resourceProjectRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*daizy.API)

	id, err := strconv.Atoi(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	project, err := client.GetProject(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(data.Id())
	data.Set("name", project.Name)
	data.Set("user_id", project.UserID)

	return nil
}

func resourceProjectDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*daizy.API)

	id, err := strconv.Atoi(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteProject(id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
