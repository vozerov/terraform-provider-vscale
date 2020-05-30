package main

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pkg/errors"
	vscale "github.com/vozerov/go-vscale"
)

func dataSourceVScaleSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVScaleSSHKeyRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "name of the ssh key",
				ValidateFunc: validation.NoZeroValues,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "public key part of the ssh key",
				Computed:    true,
			},
		},
	}
}

func dataSourceVScaleSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	name, ok := d.Get("name").(string)
	if !ok {
		return errors.New("can't find name ssh key")
	}

	keys, _, err := client.SSHKey.List()
	if err != nil {
		return errors.Wrap(err, "listing ssh keys failed")
	}

	var sshKey vscale.SSHKey

	if keys != nil {
		for _, key := range *keys {
			if key.Name == name {
				sshKey = key
			}
		}
	}

	if sshKey.ID == 0 {
		d.SetId("")
		return nil
	}

	d.SetId(strconv.Itoa(int(sshKey.ID)))
	d.Set("key", sshKey.Key)
	d.Set("name", sshKey.Name)

	return nil
}
