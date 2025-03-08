package views

import (
	"github.com/jinzhu/copier"
)

func Item[M any, V any](model M, view V) V {
	copier.Copy(&view, &model)
	return view
}

func List[M any, V any](models []M, view V) []V {
	views := make([]V, len(models))
	for i, model := range models {
		views[i] = Item(model, view)
	}
	return views
}
