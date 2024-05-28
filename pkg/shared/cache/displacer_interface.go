package cache

type Displacer interface {
	Run(storage Storage)
	Stop()
}
