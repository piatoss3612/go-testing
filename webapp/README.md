### Test Build Tags

```go
//go:build integration
```

테스트 파일 최상단에 ```//go:build {tagname}``` 형식으로 테스트 빌드 태그를 지정

빌드 태그가 설정된 테스트 파일은 해당 태그를 명시해주지 않으면 테스트가 실행되지 않는다

시간이 오래걸리는 일부 테스트를 제외하는데 유용하다

##### 빌드 태그 미지정: integration 테스트 미실행

```bash
go test -v ./...
```

##### 빌드 태그 지정: integration 테스트를 포함하여 실행

```bash
go test -v -tags=integration ./...
```