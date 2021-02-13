package engine

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
	}
}
