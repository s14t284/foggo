// Code generated by foggo; DO NOT EDIT.

package examples

type ImageOption func(*Image)

func NewImage(options ...ImageOption) *Image {
	s := &Image{}

	for _, option := range options {
		option(s)
	}

	return s
}

func WithWidth(Width int) ImageOption {
	return func(args *Image) {
		args.Width = Width
	}
}

func WithHeight(Height int) ImageOption {
	return func(args *Image) {
		args.Height = Height
	}
}

func WithAlt(Alt string) ImageOption {
	return func(args *Image) {
		args.Alt = Alt
	}
}
