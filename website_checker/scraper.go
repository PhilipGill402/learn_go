package main

import (
	"fmt"
	"bufio"
	"os"	
	"net/http"
	_ "sync"
)

type Job struct {
	ID	int
	URL	string
};

type JobResult struct {
	ID			int
	URL			string
	StatusCode	int
	Err			error
};

func worker(jobsChan <-chan Job, resultsChan chan<- JobResult) {
	for job := range jobsChan {
		res, err := http.Get(job.URL);
		if (err != nil) {
			result := JobResult {
				ID: job.ID,
				URL: job.URL,
				StatusCode: 0,
				Err: err,
			};

			resultsChan <- result;
			continue;
		}

		result := JobResult {
			ID: job.ID,
			URL: job.URL,
			StatusCode: res.StatusCode,
			Err: nil,
		};

		resultsChan <- result;
	}
}

func main() {
	file, err := os.Open("websites.txt");
	if (err != nil) {
		fmt.Println("Error:", err);
		return;
	}
	defer file.Close();
	
	var jobs []Job;
	reader := bufio.NewReader(file);
	counter := 0;	

	for {
		line, err := reader.ReadString('\n');
		if (err != nil) {
			break;
		}

		job := Job {
			ID: counter,
			URL: line[:len(line) - 1],
		};

		jobs = append(jobs, job);
		counter++;
	}
	
	jobsChan := make(chan Job);
	resultsChan := make(chan JobResult);
	
	workerCount := 5;

	for i := 0; i < workerCount; i++ {
		go worker(jobsChan, resultsChan);
	}

	for _, job := range jobs {
		jobsChan <- job;
	}

	close(jobsChan);

	for i := 0; i < counter; i++ {
		result := <- resultsChan;
		
		if (result.StatusCode == 200) {
			fmt.Printf("[OK]: %s %d\n", result.URL, result.StatusCode);
		} else if (result.Err != nil) {
			fmt.Printf("[Fail]: %s %s\n", result.URL, result.Err);
		} else {
			fmt.Printf("[Fail]: %s %d\n", result.URL, result.StatusCode);
		}
	}
}


