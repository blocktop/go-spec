package spec

type Marshalled interface {
	Namespace() string
	ResourceType() string
	Name() string
	Version() string
	Hash() string
	Marshal() ([]byte, Links, error)
	Unmarshal([]byte, Links) error
}

func FullyQualifiedName(item Marshalled) string {
	return item.Namespace() + "." + item.Name()
}
