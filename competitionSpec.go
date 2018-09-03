package blckit_interface

type CompetitionSpec interface {
	GetBlockComparator() BlockComparator
}

// BlockComparator is a function that compares two blocks and the heights of
// their blockchains, and selects the most favored block. The function returns
// ComparatorResultKeepFirst to select the first block,
// ComparatorResultKeepSecond to select the second block,
// ComparatorResultKeepBoth to select both blocks, and
// ComparatorResultKeepNeither to select neither.
type BlockComparator func(Block, uint64, Block, uint64) BlockComparatorResult

type BlockComparatorResult int

const (
	ComparatorResultKeepFirst BlockComparatorResult = iota
	ComparatorResultKeepSecond
	ComparatorResultKeepBoth
	ComparatorResultKeepNeither
)
