package expand

type ExpanderHandler struct {
	expanders []Expander
}

func NewExpanderHandler(expanders ...Expander) *ExpanderHandler {
	return &ExpanderHandler{expanders: expanders}
}

func (eh *ExpanderHandler) Expand(input string) string {
	for _, expander := range eh.expanders {
		input = expander.Expand(input)
	}
	return input
}
