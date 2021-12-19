// Code generated by schema-generate. DO NOT EDIT.

package schema

import (
    "bytes"
    "encoding/json"
    "errors"
)

// Product A product from Acme's catalog
type Product struct {

  // The unique identifier for a product
  ProductId int `json:"productId"`
}

func (strct *Product) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ProductId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "productId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"productId\": ")
	if tmp, err := json.Marshal(strct.ProductId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Product) UnmarshalJSON(b []byte) error {
    productIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "productId":
            if err := json.Unmarshal([]byte(v), &strct.ProductId); err != nil {
                return err
             }
            productIdReceived = true
        }
    }
    // check if productId (a required property) was received
    if !productIdReceived {
        return errors.New("\"productId\" is required but was not present")
    }
    return nil
}
