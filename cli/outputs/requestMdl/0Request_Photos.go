/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package requestMdl

type Request struct{
	Photos photoMdl.Photos `json:"photoss"`
	}

