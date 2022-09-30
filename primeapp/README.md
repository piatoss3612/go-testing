### Section03 Simple Testing

#### Test

1. 테스트 파일 형식

```
xxx_test.go
```

2. 테스트 함수 형식

```go
func TestMain(m *testing.M)
func TestXxx(t *testing.T)
```

3. 테스트 실행 명령어

```bash
$ go test -v
=== RUN   Test_isPrime
--- PASS: Test_isPrime (0.00s)
PASS
ok      primeapp        0.610s
```

#### Test Coverage

1. 커버리지 확인 명령어

```bash
$ go test -cover .
ok      primeapp        0.447s  coverage: 72.7% of statements
```

1. 커버리지 프로파일을 파일로 저장

```bash
$ go test -coverprofile=coverage.out
PASS
coverage: 72.7% of statements
ok      primeapp        0.438s
```

1. 커버리지 프로파일을 html로 시각화

```bash
$ go tool cover -html=coverage.out
```