// This script takes output of MUSCLE and generate consensus based on the frequencies of each nucleotide on each position

package main

import (
	"os"
	"log"
	"io"
	"bufio"
	"github.com/jessevdk/go-flags"
	"bytes"
	"fmt"
)

const (
	ConsensusHasA = 1  /* 0000001 */
	ConsensusHasC = 2  /* 0000010 */
	ConsensusHasG = 4  /* 0000100 */
	ConsensusHasT = 8  /* 0001000 */
	ConsensusHasD = 16 /* 0010000 */
	ConsensusHasN = 32 /* 0100000 */
	ConsCombinats = 64 /* 1000000 */
)

var (
	ConsensusTable = [ConsCombinats]byte{
		'N', /* 000000 None pass threshold */
		'A', /* 000001 A */
		'C', /* 000010 C */
		'M', /* 000011 A & C */
		'G', /* 000100 G */
		'R', /* 000101 G & A */
		'S', /* 000110 G & C */
		'V', /* 000111 G & A & C */
		'T', /* 001000 T */
		'W', /* 001001 T & A */
		'Y', /* 001010 T & C */
		'H', /* 001011 T & A & C */
		'K', /* 001100 T & G */
		'D', /* 001101 T & G & A */
		'B', /* 001110 T & G & C */
		'N', /* 001111 T & G & C & A */
		'-', /* 010000 Del */
		'a', /* 010001 Del & A */
		'c', /* 010010 Del & C */
		'm', /* 010011 Del & A & C */
		'g', /* 010100 Del & G */
		'r', /* 010101 Del & G & A */
		's', /* 010110 Del & G & C */
		'v', /* 010111 Del & G & A & C */
		't', /* 011000 Del & T */
		'w', /* 011001 Del & T & A */
		'y', /* 011010 Del & T & C */
		'h', /* 011011 Del & T & A & C */
		'k', /* 011100 Del & T & G */
		'd', /* 011101 Del & T & G & A */
		'b', /* 011110 Del & T & G & C */
		'n', /* 011111 Del & T & G & C & A */
		'N', /* 100000 Has N */
	}
)

func init() {
	// when it has N, this position is marked as N
	for i := ConsensusHasN; i < ConsCombinats; i++ {
		ConsensusTable[i] = 'N'
	}
}

var opts struct {
	InputFile     string  `short:"i" long:"input" description:"Input Filenames, each with its own flag" required:"true"`
	OutputFile    string  `short:"o" long:"output" description:"output file" default:"/dev/stdout"`
	Threshold     float32 `short:"t" long:"threshold" description:"threshold" default:"0.80"`
	ConcensusFile string  `short:"c" long:"concensus" description:"concensus file" required:"true"`
	NewName       string  `short:"n" long:"name" description:"new name for the consensus" default:"consensus"`
}

type FastaRecord struct {
	Name     string
	Sequence string
}

type NucCounter struct {
	NumOfA   int
	NumOfC   int
	NumOfG   int
	NumOfT   int
	NumOfN   int
	NumOfDel int
}

func (this *NucCounter) GetConsensus(t float32) byte {
	c := 0
	totalNuc := this.NumOfA + this.NumOfC + this.NumOfG + this.NumOfT + this.NumOfN + this.NumOfDel

	if float32(this.NumOfA)/float32(totalNuc) > t {
		c |= ConsensusHasA
	}

	if float32(this.NumOfC)/float32(totalNuc) > t {
		c |= ConsensusHasC
	}

	if float32(this.NumOfG)/float32(totalNuc) > t {
		c |= ConsensusHasG
	}

	if float32(this.NumOfT)/float32(totalNuc) > t {
		c |= ConsensusHasT
	}

	if float32(this.NumOfDel)/float32(totalNuc) > t {
		c |= ConsensusHasD
	}

	if float32(this.NumOfN)/float32(totalNuc) > t {
		c |= ConsensusHasN
	}

	return ConsensusTable[c]
}

type oneDimMatrix []NucCounter
type twoDimMatrix []oneDimMatrix

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}
	ofh, err := os.Create(opts.OutputFile)
	if err != nil {
		log.Fatal(err)
	}
	cfh, err := os.Create(opts.ConcensusFile)
	if err != nil {
		log.Fatal(err)
	}
	fh, err := os.Open(opts.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	reader := bufio.NewReader(fh)

	nRec := 0
	maxLen := 0
	seqs := make([]FastaRecord, 0, 1000)
	var fa FastaRecord

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		if line[0] == '>' { /* new record */
			if fa.Name != "" {
				seqs = append(seqs, fa)
				if len(fa.Sequence) > maxLen {
					maxLen = len(fa.Sequence)
				}
			}
			nRec++
			s := bytes.TrimFunc(line, func(r rune) bool {
				return r == '>' || r == '\n' || r == '\r'
			})
			fa.Name = string(s[0: func(bs []byte) int {
				i := 0
				for ; i < len(bs); i++ {
					if bs[i] == ' ' {
						break
					}
				}
				return i
			}(s)])
			fa.Sequence = ""
		} else { /* meet sequence */
			fa.Sequence += string(bytes.TrimRight(line, "\n"))
		}
	}
	seqs = append(seqs, fa)
	tmx := make(oneDimMatrix, maxLen)
	for _, fa := range seqs {
		for i, c := range fa.Sequence {
			switch c {
			case 'A', 'a':
				tmx[i].NumOfA ++
			case 'C', 'c':
				tmx[i].NumOfC ++
			case 'T', 't':
				tmx[i].NumOfT ++
			case 'G', 'g':
				tmx[i].NumOfG ++
			case '-', '.':
				tmx[i].NumOfDel ++
			default:
				tmx[i].NumOfN ++
			}
		}
	}
	fmt.Fprint(ofh, "pos\tnucleotide\tcount\tfraction\n")
	concensus := make([]byte, maxLen)
	for idx, rec := range tmx {
		concensus[idx] = rec.GetConsensus(opts.Threshold)
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "A", rec.NumOfA, float64(rec.NumOfA)/float64(nRec))
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "C", rec.NumOfC, float64(rec.NumOfC)/float64(nRec))
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "G", rec.NumOfG, float64(rec.NumOfG)/float64(nRec))
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "T", rec.NumOfT, float64(rec.NumOfT)/float64(nRec))
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "N", rec.NumOfN, float64(rec.NumOfN)/float64(nRec))
		fmt.Fprintf(ofh, "%d\t%s\t%d\t%.4f\n", idx, "-", rec.NumOfDel, float64(rec.NumOfDel)/float64(nRec))
	}
	fmt.Fprintf(cfh, ">%s\n%s\n", opts.NewName, string(GetConsensusWithoutDel(concensus)))
}

func GetConsensusWithoutDel(consensus []byte) (result []byte) {
	for _, b := range consensus {
		if b != '-' {
			result = append(result, b)
		}
	}
	return
}

