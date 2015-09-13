package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LOWER = "abcdefghijklmnopqrstuvwxyz"

func runProg(cmd *exec.Cmd) (io.WriteCloser, io.ReadCloser, error) {
	w, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}
	return w, stdout, nil
}

var inputLog bytes.Buffer

func runIt(r io.Reader, prog1 *exec.Cmd, prog2 *exec.Cmd) (io.ReadCloser, io.ReadCloser) {

	iw := bufio.NewWriter(&inputLog)
	w1, r1, err := runProg(prog1)
	if err != nil {
		log.Fatal(err)
	}
	w2, r2, err := runProg(prog2)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer w1.Close()
		defer w2.Close()
		defer iw.Flush()
		mw := io.MultiWriter(w1, w2, iw)
		io.Copy(mw, r)
	}()

	return r1, r2
}

func diff(r1, r2 io.Reader) {
	for {
		b1 := make([]byte, 1)
		b2 := make([]byte, 1)
		n1, err1 := r1.Read(b1)
		n2, err2 := r2.Read(b2)

		if err1 == io.EOF || err2 == io.EOF {
			break
		}

		var n int
		if n1 < n2 {
			n = n1
			r2 = io.MultiReader(bytes.NewReader(b2[n:n2]), r2)
		} else {
			n = n2
			r1 = io.MultiReader(bytes.NewReader(b1[n:n1]), r1)
		}
		fmt.Printf("So far %d: %s, %s\n", n, b1[:n], b2[:n])
		if string(b1[:n]) != string(b2[:n]) {
			log.Fatal("Mimatch found!")
		}

		//fmt.Printf("Read %d bytes: %s. Err: %v\n", n1, b1, err1)
		//fmt.Printf("Read %d bytes: %s. Err: %v\n", n2, b2, err2)
	}

}

func diff2(r1, r2 io.Reader) {
	scanner1 := bufio.NewScanner(r1)
	scanner2 := bufio.NewScanner(r2)
	for {
		n1 := scanner1.Scan()
		n2 := scanner2.Scan()
		err1 := scanner1.Err()
		err2 := scanner1.Err()
		fmt.Println(n1, n2, err1, err2)
		if n1 != n2 || n1 == false || n2 == false {
			break
		}
		if err1 != nil || err2 != nil {
			break
		}
		line1 := scanner1.Text()
		line2 := scanner2.Text()
		fmt.Printf("So far: %s, %s\n", line1, line2)
		if line1 != line2 {
			fmt.Printf("Mismatch:\n->%s\n=>%s\n", line1, line2)
		}
	}
}

func main() {
	genBinary, prog1Binary, prog2Binary := os.Args[1], os.Args[2], os.Args[3]

	generator := exec.Command(genBinary)
	r, err := generator.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	prog1 := exec.Command(prog1Binary)
	prog2 := exec.Command(prog2Binary)

	r1, r2 := runIt(r, prog1, prog2)

	generator.Run()

	diff2(r1, r2)

	err = prog1.Wait()
	err1 := prog2.Wait()
	if err != nil || err1 != nil {
		fmt.Println(err, err1)
	}

	fmt.Printf("Here: %s", &inputLog)
}
