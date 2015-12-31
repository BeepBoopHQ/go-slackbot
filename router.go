package slackbot

import "golang.org/x/net/context"

type Router interface {
	Match(context.Context, *RouteMatch) (bool, context.Context)
	NewRoute() *Route
	Hear(regex string) *Route
	Handler(handler Handler) *Route
	MessageHandler(handler MessageHandler) *Route
	Messages(types ...MessageType) *Route
	AddMatcher(m Matcher) *Route
	SetBotID(botID string)
}

type SimpleRouter struct {
	// Routes to be matched, in order.
	routes []*Route
	// Slack UserID of the bot UserID
	botUserID string
}

// Match matches registered routes against the request.
func (r *SimpleRouter) Match(ctx context.Context, match *RouteMatch) (bool, context.Context) {
	for _, route := range r.routes {
		var matched bool
		if matched, ctx = route.Match(ctx, match); matched {
			return true, ctx
		}
	}

	return false, ctx
}

// NewRoute registers an empty route.
func (r *SimpleRouter) NewRoute() *Route {
	route := &Route{}
	r.routes = append(r.routes, route)
	return route
}

func (r *SimpleRouter) Hear(regex string) *Route {
	return r.NewRoute().Hear(regex)
}

func (r *SimpleRouter) Handler(handler Handler) *Route {
	return r.NewRoute().Handler(handler)
}

func (r *SimpleRouter) MessageHandler(handler MessageHandler) *Route {
	return r.NewRoute().MessageHandler(handler)
}

func (r *SimpleRouter) Messages(types ...MessageType) *Route {
	return r.NewRoute().Messages(types...)
}

func (r *SimpleRouter) AddMatcher(m Matcher) *Route {
	return r.NewRoute().AddMatcher(m)
}

func (r *SimpleRouter) SetBotID(botID string) {
	r.botUserID = botID
	for _, route := range r.routes {
		route.setBotID(botID)
	}
}
