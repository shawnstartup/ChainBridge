package chains

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProposalStatus(t *testing.T) {

	testTable := []struct {
		name   string
		status int
		expect string
	}{
		{
			"Invalid case 1",
			5,
			"Unknown",
		},
		{
			"Invalid case 2",
			-1,
			"Unknown",
		},
		{
			"Valid case 1",
			0,
			"Inactive",
		},
		{
			"Valid case 2",
			1,
			"Active",
		},
		{
			"Valid case 3",
			2,
			"Passed",
		},
		{
			"Valid case 4",
			3,
			"Executed",
		},
		{
			"Valid case 5",
			4,
			"Cancelled",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.expect,
				ProposalStatus(testCase.status).String(),
			)
		})
	}

}
