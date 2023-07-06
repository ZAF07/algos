V=""

two-sum-go:
	@if [ $V = "TRUE" ]; then\
		go test two-sum/twoSum_test.go -test.v;\
	else\
		go test two-sum/twoSum_test.go;\
  fi
