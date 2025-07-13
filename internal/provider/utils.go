package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func convertInterfaceMapToStringMap(input *map[string]interface{}) (types.Map, error) {
	if input != nil {
		output := make(map[string]types.String)
		for k, v := range *input {
			str := types.StringValue(fmt.Sprintf("%v", v))
			output[k] = str
		}

		mapReturn, diags := types.MapValueFrom(context.Background(), types.StringType, output)
		if diags.HasError() {
			return types.MapNull(types.StringType), fmt.Errorf("error creating new map value: %v", diags.Errors())
		}
		return mapReturn, nil
	} else {
		return types.MapNull(types.StringType), nil
	}
}
