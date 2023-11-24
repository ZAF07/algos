V=""

go-two-sum:
	@if [ $V = "TRUE" ]; then\
		go test two-sum/two_sum_test.go -test.v;\
	else\
		go test two-sum/two_sum_test.go;\
  fi

go-two-sum-two:
	@if [ $V = "TRUE" ]; then\
		go test two-sum-two/two_sum_two_test.go -test.v;\
	else\
		go test two-sum-two/two_sum_two_test.go;\
	fi

go-palindrome-int:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/palindrome_integers_test.go -test.v;\
	else\
		go test palindrome/palindrome_integers_test.go;\
  fi


go-palindrome-string:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/palindrome_string_test.go -test.v;\
	else\
		go test palindrome/palindrome_string_test.go;\
  fi

go-longest-palindrome-manacher:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/longest_palindrome_manacher_test.go -test.v;\
	else\
		go test palindrome/longest_palindrome_manacher_test.go;\
  fi

go-longest-palindrome:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/longest_palindrome_test.go -test.v;\
	else\
		go test palindrome/longest_palindrome_test.go;\
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

go-first-and-last-position:
	@if [ $V = "TRUE" ]; then\
		go test binary-search/first_and_last_position.go -test.v;\
	else\
		go test binary-search/first_and_last_position_test.go;\
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

go-count-negative_two:
	@if [ $V = "TRUE" ]; then\
		go test arrays/count-negative-numbers-in-array/count_negative_two_test.go -test.v;\
	else\
		go test arrays/count-negative-numbers-in-array/count_negative_two_test.go;\
  fi
	
go-best-time-to-buy:
	@if [ $V = "TRUE" ]; then\
		go test arrays/best-time-to-buy-sell-stocks/best_time_to_buy_test.go -test.v;\
	else\
		go test arrays/best-time-to-buy-sell-stocks/best_time_to_buy_test.go;\
  fi
  
go-best-time-to-buy-sell-stock-2:
	@if [ $V = "TRUE" ]; then\
		go test dynamic-programming/best_time_buy_sell_stock_2_test.go -test.v;\
	else\
		go test dynamic-programming/best_time_buy_sell_stock_2_test.go;\
  fi


go-majority-element:
	@if [ $V == "TRUE" ]; then\
		go test arrays/majority-element/majority_element_test.go -test.v;\
	else\
		go test arrays/majority-element/majority_element_test.go;\
	fi

go-majority_element_boyer_moore:
	@if [ $V = "TRUE" ]; then\
		go test arrays/majority-element/boyer_moore_solution_test.go -test.v;\
	else\
		go test arrays/majority-element/boyer_moore_solution_test.go;\
	fi

go-search-in-rotated-array:
	@if [ $V == "TRUE" ]; then\
		go test arrays/search-in-rotated-array/search_in_rotated_array_test.go -test.v;\
	else\
		go test arrays/search-in-rotated-array/search_in_rotated_array_test.go;\
	fi

go-maximum-average-subarray-one:
	@if [ $V == "TRUE" ]; then\
		go test sliding-window/maximum_average_subarray_one_test.go -test.v;\
	else\
		go test sliding-window/maximum_average_subarray_one_test.go;\
	fi

go-minimum-size-subarray-sum:
	@if [ $V == "TRUE" ]; then\
		go test sliding-window/minimum_size_subarray_sum_test.go -test.v;\
	else\
		go test sliding-window/minimum_size_subarray_sum_test.go;\
	fi

go-merge-sort:
	@if [ $V = "TRUE" ]; then\
		go test merge-sort/merge_sort_test.go -test.v;\
	else\
		go test merge-sort/merge_sort_test.go;\
	fi

go-longest-substring-without-repeating-letters:
	@if [ $V = "TRUE" ]; then\
		go test sliding-window/longest_substring_without_repeating_letters_test.go -test.v;\
	else\
		go test sliding-window/longest_substring_without_repeating_letters_test.go;\
	fi


go-single-number:
	@if [ $V = "TRUE" ]; then\
		go test arrays/single-number/single_number_test.go -test.v;\
	else\
		go test arrays/single-number/single_number_test.go;\
	fi

go-valid-palindrome:
	@if [ $V = "TRUE" ]; then\
		go test palindrome/valid_palindrome_test.go -test.v;\
	else\
		go test palindrome/valid_palindrome_test.go;\
	fi

go-max-subarray-sum:
	@if [ $V = "TRUE" ]; then\
		go test max-sub-array-sum/max_subarray_sum_test.go -test.v;\
	else\
		go test max-sub-array-sum/max_subarray_sum_test.go;\
	fi

go-min-amt-time-collect-garbage:
	@if [ $V = "TRUE" ]; then\
		go test arrays/minimum-amount-of-time-to-collect-garbage/min_amt_time_collect_garbage_test.go -test.v;\
	else\
		go test arrays/minimum-amount-of-time-to-collect-garbage/min_amt_time_collect_garbage_test.go;\
	fi

go-balloon:
	@if [ $V = "TRUE" ]; then\
		go test balloon/balloon_test.go -test.v;\
	else\
		go test balloon/balloon_test.go;\
	fi

go-spiral-matrix:
	@if [ $V = "TRUE" ]; then\
		go test arrays/spiral-matrix/spiral_matrix_test.go -test.v;\
	else\
		go test arrays/spiral-matrix/spiral_matrix_test.go;\
	fi

go-three-sum:
	@if [ $V = "TRUE" ]; then\
		go test three-sum/three_sum_test.go -test.v;\
	else\
		go test three-sum/three_sum_test.go;\
	fi

			return []int{l + 1, r + 1}