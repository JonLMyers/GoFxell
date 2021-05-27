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
	View() // TODO: Implement this
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
