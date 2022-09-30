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

2. 커버리지 프로파일을 파일로 저장

```bash
$ go test -coverprofile=coverage.out
PASS
coverage: 72.7% of statements
ok      primeapp        0.438s
```

3. 커버리지 프로파일을 html로 시각화

```bash
$ go tool cover -html=coverage.out
```

### Section04 Running individual tests & Test Suites

#### Running a single test

```bash
$ go test -run Test_isPrime
PASS
ok      primeapp        0.580s
```

```bash
$ go test -v -run Test_isPrime
=== RUN   Test_isPrime
--- PASS: Test_isPrime (0.00s)
PASS
ok      primeapp        0.610s
```

#### Running groups of tests (test suites)

테스트 함수를 테스트 그룹명으로 묶어줄 수 있으며 테스트 그룹명을 이용해 그룹 테스트를 실행할 수 있다

ex) [Test_alpha_Xxx, Test_alpha_Yyy], [Test_beta_Xxx, Test_beta_Yyy]

```bash
$ go test -v -run Test_alpha
=== RUN   Test_alpha_isPrime
--- PASS: Test_alpha_isPrime (0.00s)
=== RUN   Test_alpha_prompt
--- PASS: Test_alpha_prompt (0.00s)
PASS
ok      primeapp        0.715s
```