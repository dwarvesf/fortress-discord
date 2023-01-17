package earn

type Earn struct{}

func New() EarnViewer {
	return &Earn{}
}
