package main

import ( 
    "fmt"
    "os"
    "encoding/csv"
    "flag"
    "time"
)


func main() {
    //set up the flag to check
    filenamePtr := flag.String("csv", "./problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
    timePtr := flag.Int("limit", 30, "The time limit, in seconds")
    //read the flags
    flag.Parse()

    //need to deref pointer for it to work 
    problems := csvReader(*filenamePtr)
    
    fmt.Printf("Hit enter to begin")
    fmt.Scanf("\n")
    timer := time.NewTimer(time.Duration(*timePtr) * time.Second)

    //print the final part early
    count := 0
    length := len(problems)

    //loop through
    for index, problem := range problems {
        fmt.Printf("Problem #%d: %s = \n", index+1, problem.q)

        //ok, so much going on here
        //go: defines a go routine. Enables concurrency
        //func () {} defines an anonymous func
        //answerCh: makes a new channel to pass answers between threads
        answerCh := make(chan string)
        go func() {
            var answer string
            //take user input 
            fmt.Scanf("%s\n", &answer)
            answerCh <- answer
        }()

        select {
        case <-timer.C:
            fmt.Printf("\nYou scored %d out of %d. \n", count, length)
            return
        case answer := <-answerCh:
            if answer == problem.a {
                count += 1
            }
        }
    }
    fmt.Printf("You scored %d out of %d. \n", count, length)
    
}

type problem struct {
    q string
    a string
}

func csvReader(filename string) ([]problem) {
    // 1. Open the file
    //:= declares a var and sets it. Otherwise, you'd use var name type or var name = 
    f, err := os.Open(filename)
    //error check
    if err != nil {
        fmt.Println("Failed to open: %s", filename)
        os.Exit(1)
    }

    //defer waits until the current function is done to run
    defer f.Close()

    reader := csv.NewReader(f)

    //records is type [][]string (line, index)
    records, _ := reader.ReadAll()

    //makes an empty array
    ret := make([]problem, len(records))

    for i, line := range records {
        ret[i] = problem{
            q : line[0],
            a : line[1],
        }
    }
    return(ret)
}