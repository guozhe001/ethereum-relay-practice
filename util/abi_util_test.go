package util

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetMethodId(t *testing.T) {
	methodId, err := GetMethodId("../abi/GZToken_metadata.json", "transfer")
	assert.NoError(t, err)
	assert.NotNil(t, methodId)
	log.Println(methodId)
}
