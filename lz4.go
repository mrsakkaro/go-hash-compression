package main


import (
	//"bytes"
	"io"
	"bufio"
	"os"
	"runtime"
	"github.com/pierrec/lz4"
)

//func compress(data []byte) int, string {
	// compress the uncompressed data
	// and compare with the original input
	/*
	buffer := bytes.NewBuffer(nil)
	compressor := lz4.NewWriter(buffer)
	n, err := compressor.Wrtie(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	if n != len(d) {
		panic("short write")
	}
	err = compressor.Close()
	if err != nil {
		panic(err)
	}
	*/
/*
	bufferWriter := bufio.NewWriter(nil)
	compressor := lz4.NewWriter(bufferWriter)
	compressor.Header = lz4.Header{
		BlockDependency:	false,
		BlockChecksum:		true,
		BlockMaxSize:		4 << 14,
		NoChecksum:			false,
		HighCompression:	true,
	}

	//compressor.Write([]byte("Hello World!"))

	//io.Copy(compressor, bufio.NewReaderSize(

	buffer := make([]byte, 1024)
	*/
//}

func compressFile2(srcdir, destdir string) {

	// Open Source File
	src, err := os.Open(srcdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			panic(err)
		}
	}()

	// Make a read buffer
	r := bufio.NewReader(src)

	// Open Output File
	dest, err := os.Create(destdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := dest.Close(); err != nil {
			panic(err)
		}
	}()

	// Make lz4 buffer writer
	lz4Writer := lz4.NewWriter(dest)

	// Make a buffer and do the magic
	buf := make([]byte, 1024)
	for {
		// Read a block
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// Write a block
		if _, err := lz4Writer.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = lz4Writer.Flush(); err != nil {
		panic(err)
	}

}

func compressFile3 (srcdir, destdir string) {

	// Open Source File
	src, err := os.Open(srcdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			panic(err)
		}
	}()

	// Open Output File
	dest, err := os.Create(destdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := dest.Close(); err != nil {
			panic(err)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	//bufferReader := bufio.NewReader(src)

	//bufferWriter := bufio.NewWriter(dest)
	compressor := lz4.NewWriter(dest)
	compressor.Header = lz4.Header{
		BlockChecksum:		false,
		BlockMaxSize:		4 << 20,
		NoChecksum:			false,
		CompressionLevel:	9,
	}
	/*
	compressorHeader = lz4.Header{
		BlockChecksum:		false,
		BlockMaxSize:		4 << 14,
		NoChecksum:			false,
		CompressionLevel:	1,
	}
	compressor.Header = compressorHeader
	*/

	//compressor.Write([]byte("Hello World!"))

	if _, err := io.Copy(compressor, src); err != nil {
		panic(err)
	}
	//bufferWriter.Flush()
	if err := compressor.Close(); err != nil {
		panic(err)
	}

}

func decompressFile(srcdir, destdir string) {

	// Open Source File
	src, err := os.Open(srcdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			panic(err)
		}
	}()

	// Open Output File
	dest, err := os.Create(destdir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := dest.Close(); err != nil {
			panic(err)
		}
	}()

	bufferWriter := bufio.NewWriter(dest)
	io.Copy(bufferWriter, lz4.NewReader(src))
	bufferWriter.Flush()
}
/*

func compressFile2( ) {
	var (
		blockMaxSizeDefault = 4 << 20
		flagStdout          = flag.Bool("c", false, "output to stdout")
		flagDecompress      = flag.Bool("d", false, "decompress flag")
		flagBlockMaxSize    = flag.Int("B", blockMaxSizeDefault, "block max size [64Kb,256Kb,1Mb,4Mb]")
		flagBlockDependency = flag.Bool("BD", false, "enable block dependency")
		flagBlockChecksum   = flag.Bool("BX", false, "enable block checksum")
		flagStreamChecksum  = flag.Bool("Sx", false, "disable stream checksum")
		flagHighCompression = flag.Bool("9", false, "enabled high compression")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [arg] [input]...\n\tNo input means [de]compress stdin to stdout\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println("output to stdout ", *flagStdout)
	fmt.Println("Decompress", *flagDecompress)
	// Use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	zr := lz4.NewReader(nil)
	zw := lz4.NewWriter(nil)
	zh := lz4.Header{
		BlockDependency: *flagBlockDependency,
		BlockChecksum:   *flagBlockChecksum,
		BlockMaxSize:    *flagBlockMaxSize,
		NoChecksum:      *flagStreamChecksum,
		HighCompression: *flagHighCompression,
	}

	worker := func(in io.Reader, out io.Writer) {
		if *flagDecompress {
			fmt.Println("\n Decompressing the data")
			zr.Reset(in)
			if _, err := io.Copy(out, zr); err != nil {
				log.Fatalf("Error while decompressing input: %v", err)
			}
		} else {
			zw.Reset(out)
			zw.Header = zh
			if _, err := io.Copy(zw, in); err != nil {
				log.Fatalf("Error while compressing input: %v", err)
			}
		}
	}
}

*/

func main() {
/*
	zr := lz4.NewReader(nil)
	zw := lz4.NewWriter(nil)
	zh := lz4.Header{
		BlockChecksum:		false,
		BlockMaxSize:		4 << 14,
		NoChecksum:			false,
		CompressionLevel:	1,
	}

	worker := func(in io.Reader, out io.Writer, compression bool) {
		if compression {
			fmt.Println("\n Compressing the data")
			zw.Reset(out)
			zw.Header = zh
			if _, err := io.Copy(zw, in); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("\n Decompressing the data")
			zr.Reset(in)
			if _, err := io.Copy(out, zr); err != nil {
				panic(err)
			}

		}
	}

*/




















	compressFile2("compress-me.txt", "result2.lz4")
	compressFile3("compress-me.txt", "result3.lz4")
	decompressFile3("result2.lz4", "result2-3.txt")
	decompressFile3("result3.lz4", "result3-3.txt")
}
