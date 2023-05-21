package log

import (
	"io"
	"os"
	"testing"

	api "github.com/bayamasa/proglog/api/v1"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, err := os.MkdirTemp("", "segment_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	
	want := &api.Record{Value: []byte("hello world")}
	
	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3
	
	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset)
	require.False(t, s.IsMaxed())
	
	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		// 相対オフセットなので、ベースオフセットとの差分を取る
		require.Equal(t, 16 + i, off)
		
		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}
	
	// MaxIndexBytesがentWidth * 3なので、4つ目のAppendでインデックスがMaxIndexBytesを超える
	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)
	require.True(t, s.IsMaxed())
	require.NoError(t, s.Close())
	
	p, _ := proto.Marshal(want)
	c.Segment.MaxStoreBytes = uint64(len(p) + lenWidth) * 4
}
