package search

// defaultMatcher 기본 검색기를 구현할 타입
type defaultMatcher struct{}

// init 함수에서는 기본 검색기를 프로그램에 등록한다.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 함수를 기본 검색기의 동작을 구현한다.
func (m defaultMatcher) Search(feed *Feed, serachTerm string) ([]*Result, error) {
	return nil, nil
}
