package game

import (
	"context"
	"fmt"
)

// contextKey is a type used to store values in the context.
type contextKey string

// viewerKey is the contextKey used to set & get the viewer from the context.
const viewerKey contextKey = "viewer"

// A Viewer is used for restricted viewing of nodes.
type Viewer interface {
	View(Node) Node
}

// WithViewer associates a viewer with a context.
func WithViewer(ctx context.Context, viewer Viewer) context.Context {
	if viewer == nil {
		panic(fmt.Errorf("called WithViewer with nil viewer"))
	}
	return context.WithValue(ctx, viewerKey, viewer)
}

// GetViewer returns the current viewer associated with the context.
func GetViewer(ctx context.Context) Viewer {
	viewer, ok := ctx.Value(viewerKey).(Viewer)
	if !ok {
		panic(fmt.Errorf("received invalid viewer type from context"))
	}
	if viewer == nil {
		panic(fmt.Errorf("no viewer stored in provided context"))
	}
	return viewer
}

// example for jon (requires WithViewer being called on the provided context first)
func DoThing(ctx context.Context, node Node) {
	viewer := GetViewer(ctx)

	// You probably won't need this, I just needed to print the team name
	team, ok := viewer.(Team) //wtf is this?
	if ok {
		panic(fmt.Errorf("viewer is not a team"))
	}

	visibleNode := team.View(node)
	fmt.Printf("Here's what team %s can see: %+v", team.Name, visibleNode)
}
