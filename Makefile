V=""

go-two-sum:
	@if [ $V = "TRUE" ]; then\
		go test two-sum/two_sum_test.go -test.v;\
	else\
		go test two-sum/two_sum_test.go;\
  fi
