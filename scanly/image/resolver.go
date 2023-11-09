package image

type resolver interface {
	fetch(id string) (*Image, error)
	build(options []string) (Image, error)
}
