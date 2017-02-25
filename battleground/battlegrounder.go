package battleground

// Battlegrounder is the interface for a battleground.
type Battlegrounder interface {
	M() int
	Play() []string
}
