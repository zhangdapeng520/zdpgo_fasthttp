package bytebufferpool

import "io"

// ByteBuffer 提供字节缓冲区，可用于最小化内存分配。
//
// ByteBuffer 可以与向给定的[]byte追加数据的函数一起使用
// 使用Get获取空字节缓冲区。
type ByteBuffer struct {

	// B 是在类似于追加的工作负载中使用的字节缓冲区。
	B []byte
}

// Len 返回字节缓冲区的大小。
func (b *ByteBuffer) Len() int {
	return len(b.B)
}

// ReadFrom 实现 io.ReaderFrom.
//
// 返回字节缓冲区的大小。 该函数将从r读取的所有数据追加到b。
func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
	p := b.B
	nStart := int64(len(p))
	nMax := int64(cap(p))
	n := nStart
	if nMax == 0 {
		nMax = 64
		p = make([]byte, nMax)
	} else {
		p = p[:nMax]
	}
	for {
		if n == nMax {
			nMax *= 2
			bNew := make([]byte, nMax)
			copy(bNew, p)
			p = bNew
		}
		nn, err := r.Read(p[n:])
		n += int64(nn)
		if err != nil {
			b.B = p[:n]
			n -= nStart
			if err == io.EOF {
				return n, nil
			}
			return n, err
		}
	}
}

// WriteTo 实现 io.WriterTo.
func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.B)
	return int64(n), err
}

// Bytes 返回b.B，即缓冲区中累计的所有字节。
//
// The purpose of this function is bytes.Buffer compatibility.
func (b *ByteBuffer) Bytes() []byte {
	return b.B
}

// Write 实现 io.Writer - it appends p to ByteBuffer.B
func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

// WriteByte appends the byte c to the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
//
// The function always returns nil.
func (b *ByteBuffer) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

// WriteString appends s to ByteBuffer.B.
func (b *ByteBuffer) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

// Set sets ByteBuffer.B to p.
func (b *ByteBuffer) Set(p []byte) {
	b.B = append(b.B[:0], p...)
}

// SetString sets ByteBuffer.B to s.
func (b *ByteBuffer) SetString(s string) {
	b.B = append(b.B[:0], s...)
}

// String returns string representation of ByteBuffer.B.
func (b *ByteBuffer) String() string {
	return string(b.B)
}

// Reset makes ByteBuffer.B empty.
func (b *ByteBuffer) Reset() {
	b.B = b.B[:0]
}
