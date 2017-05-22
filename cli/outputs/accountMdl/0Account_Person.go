/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package accountMdl

type Account struct{
	Person personMdl.Person `json:"person"`
	}

