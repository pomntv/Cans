package main

import (
	"testing"
)

func TestFormatPhoneNumberWithoutLeadingZero(t *testing.T) {
	testCases := []struct {
		input          string
		zerolead       bool
		expectedResult string
	}{
		{
			input:          "123456789",
			zerolead:       false,
			expectedResult: "12-345-6789",
		},
		{
			input:          "11223344",
			zerolead:       false,
			expectedResult: "11-223-344",
		},
		{
			input:          "44556741",
			zerolead:       false,
			expectedResult: "44-556-741",
		},
	}

	for _, tc := range testCases {
		result := formatPhoneNumber(tc.input, tc.zerolead)
		if result != tc.expectedResult {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expectedResult, result)
		}
	}
}

func TestFormatPhoneNumberWithLeadingZero(t *testing.T) {
	testCases := []struct {
		input          string
		zerolead       bool
		expectedResult string
	}{
		{
			input:          "0397087700",
			zerolead:       true,
			expectedResult: "039-708-7700",
		},
		{
			input:          "012469260",
			zerolead:       true,
			expectedResult: "012-469-260",
		},
		{
			input:          "000123",
			zerolead:       true,
			expectedResult: "000-123",
		},
	}

	for _, tc := range testCases {
		result := formatPhoneNumber(tc.input, tc.zerolead)
		if result != tc.expectedResult {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expectedResult, result)
		}
	}
}

func TestFormatPhoneNumberLongNumber(t *testing.T) {
	testCases := []struct {
		input          string
		zerolead       bool
		expectedResult string
	}{
		{
			input:          "0011223344556677",
			zerolead:       true,
			expectedResult: "0011223344556677",
		},
		{
			input:          "12345678912345",
			zerolead:       false,
			expectedResult: "12345678912345",
		},
	}

	for _, tc := range testCases {
		result := formatPhoneNumber(tc.input, tc.zerolead)
		if result != tc.expectedResult {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expectedResult, result)
		}
	}
}

func TestFormatPhoneNumberShortNumber(t *testing.T) {
	testCases := []struct {
		input          string
		zerolead       bool
		expectedResult string
	}{
		{
			input:          "039708",
			zerolead:       true,
			expectedResult: "039-708",
		},
		{
			input:          "03970",
			zerolead:       true,
			expectedResult: "039-70",
		},
		{
			input:          "0397",
			zerolead:       true,
			expectedResult: "039-7",
		},
		{
			input:          "44123",
			zerolead:       false,
			expectedResult: "44-123",
		},
		{
			input:          "4412",
			zerolead:       false,
			expectedResult: "44-12",
		},
		{
			input:          "441",
			zerolead:       false,
			expectedResult: "44-1",
		},
		{
			input:          "22",
			zerolead:       false,
			expectedResult: "22",
		},
		{
			input:          "039",
			zerolead:       true,
			expectedResult: "039",
		},
		{
			input:          "01",
			zerolead:       true,
			expectedResult: "01",
		},
		{
			input:          "1",
			zerolead:       false,
			expectedResult: "1",
		},
		{
			input:          "0",
			zerolead:       true,
			expectedResult: "0",
		},
	}

	for _, tc := range testCases {
		result := formatPhoneNumber(tc.input, tc.zerolead)
		if result != tc.expectedResult {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expectedResult, result)
		}
	}
}
