V=""

go-two-sum:
	@if [ $V = "TRUE" ]; then\
		go test two-sum/two_sum_test.go -test.v;\
	else\
		go test two-sum/two_sum_test.go;\
  fi

go-palindrome:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/palindrome_test.go -test.v;\
	else\
		go test palindrome/palindrome_test.go;\
  fi

go-bubble-sort:
	@if [ $V = "TRUE" ]; then\
		go test bubble_sort/bubble_test.go -test.v;\
	else\
		go test bubble_sort/bubble_test.go;\
  fi

go-binary-search:
	@if [ $V = "TRUE" ]; then\
		go test binary-search/binary_search_test.go -test.v;\
	else\
		go test binary-search/binary_search_test.go;\
  fi

go-top-k-frequent-nums:
	@if [ $V = "TRUE" ]; then\
		go test arrays/top-k-frequent-nums/top_k_frequent_nums_test.go -test.v;\
	else\
		go test arrays/top-k-frequent-nums/top_k_frequent_nums_test.go;\
  fi

go-rotate-array:
	@if [ $V = "TRUE" ]; then\
		go test arrays/rotate-array/rotate_array_test.go -test.v;\
	else\
		go test arrays/rotate-array/rotate_array_test.go;\
  fi

go-remove-duplicates:
	@if [ $V = "TRUE" ]; then\
		go test arrays/remove-duplicates/remove_duplicates_test.go -test.v;\
	else\
		go test arrays/remove-duplicates/remove_duplicates_test.go;\
  fi

go-plus-one:
	@if [ $V = "TRUE" ]; then\
		go test arrays/plus-one/plus_one_test.go -test.v;\
	else\
		go test arrays/plus-one/plus_one_test.go;\
  fi

go-count-negative:
	@if [ $V = "TRUE" ]; then\
		go test arrays/count-negative-numbers-in-array/count_negative_test.go -test.v;\
	else\
		go test arrays/count-negative-numbers-in-array/count_negative_test.go;\
  fi
	
go-best-time-to-buy:
	@if [ $V = "TRUE" ]; then\
		go test arrays/best-time-to-buy-sell-stocks/best_time_to_buy_test.go -test.v;\
	else\
		go test arrays/best-time-to-buy-sell-stocks/best_time_to_buy_test.go;\
  fi

go-majority-element:
	@if [ $V == "TRUE" ]; then\
		go test arrays/majority-element/majority_element_test.go -test.v;\
	else\
		go test arrays/majority-element/majority_element_test.go;\
	fi

go-search-in-rotated-array:
	@if [ $V == "TRUE" ]; then\
		go test arrays/search-in-rotated-array/search_in_rotated_array_test.go -test.v;\
	else\
		go test arrays/search-in-rotated-array/search_in_rotated_array_test.go;\
	fi