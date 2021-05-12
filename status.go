package dht

type Status struct {
	BlackList *blackList
	Queue     *syncedMap
	Requests  int
	Responses int
	Tokens    int
}
