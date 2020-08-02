### 入口

涉及包学习

- strings: Contains(s, substr string)、HasSuffix(s, suffix string)、Split(s, sep string) []string
- filepath: Abs(path string) (string, error)、Join(elem ...string) string、Walk(root string, walkFn WalkFunc) error
- ioutil: ReadFile(filename string) ([]byte, error)、ReadAll(r io.Reader) ([]byte, error)
- zip: OpenReader(name string) (*ReadCloser, error)