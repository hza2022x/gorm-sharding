package log

import (
	internal "gopkg.inshopline.com/commons/sqlx/testing"
	"testing"
)

func TestSetErrorLevel(t *testing.T) {
	SetErrorLevel()
	assert := internal.NewAssert(t, "TestSetErrorLevel")
	assert.Equal(IsTraceEnabled, false)
	assert.Equal(IsDebugEnabled, false)
	assert.Equal(IsInfoEnabled, false)
	assert.Equal(IsWarnEnabled, false)
	assert.Equal(IsErrorEnabled, true)
}

func TestSetWarnLevel(t *testing.T) {
	SetWarnLevel()
	assert := internal.NewAssert(t, "TestSetWarnLevel")
	assert.Equal(IsTraceEnabled, false)
	assert.Equal(IsDebugEnabled, false)
	assert.Equal(IsInfoEnabled, false)
	assert.Equal(IsWarnEnabled, true)
	assert.Equal(IsErrorEnabled, true)
}

func TestSetInfoLevel(t *testing.T) {
	SetInfoLevel()
	assert := internal.NewAssert(t, "TestSetInfoLevel")
	assert.Equal(IsTraceEnabled, false)
	assert.Equal(IsDebugEnabled, false)
	assert.Equal(IsInfoEnabled, true)
	assert.Equal(IsWarnEnabled, true)
	assert.Equal(IsErrorEnabled, true)
}

func TestSetDebugLevel(t *testing.T) {
	SetDebugLevel()
	assert := internal.NewAssert(t, "TestSetDebugLevel")
	assert.Equal(IsTraceEnabled, false)
	assert.Equal(IsDebugEnabled, true)
	assert.Equal(IsInfoEnabled, true)
	assert.Equal(IsWarnEnabled, true)
	assert.Equal(IsErrorEnabled, true)
}

func TestSetTraceLevel(t *testing.T) {
	SetTraceLevel()
	assert := internal.NewAssert(t, "TestSetTraceLevel")
	assert.Equal(IsTraceEnabled, true)
	assert.Equal(IsDebugEnabled, true)
	assert.Equal(IsInfoEnabled, true)
	assert.Equal(IsWarnEnabled, true)
	assert.Equal(IsErrorEnabled, true)
}
