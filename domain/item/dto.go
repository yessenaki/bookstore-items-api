package item

const indexItems = "items"

func (i *Item) Save() *errors.RESTError {
	res, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.InternalServerError()
	}
	i.ID = res.ID

	return nil
}
