package xdaemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const ENV_NAME = "XW_DAEMON_IDX"

// Number of times Background() is called during runtime
var runIdx int = 0

// Daemon process
type Daemon struct {
	LogFile     string // Log file for recording daemon and child process stdout/stderr. If empty, no logging
	MaxCount    int    // Maximum number of restart cycles. If 0, unlimited restarts
	MaxError    int    // Maximum number of consecutive start failures or abnormal exits before daemon quits
	MinExitTime int64  // Minimum runtime (seconds) for normal exit. Less than this is considered abnormal
}

// Convert the program itself to run in the background (start a child process, then exit)
// logFile: if not empty, the child process's stdout and stderr will be written to this file
// isExit: whether to exit the main program after starting the child process
//
//	if false, main process returns *os.Process, child process returns nil. Need to handle accordingly
func Background(logFile string, isExit bool) (*exec.Cmd, error) {
	// Check if child or parent process
	runIdx++
	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	if runIdx <= envIdx { // Child process, exit
		return nil, nil
	}

	// Set child process environment variables
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))

	// Start child process
	cmd, err := startProc(os.Args, env, logFile)
	if err != nil {
		log.Println(os.Getpid(), "Failed to start child process:", err)
		return nil, err
	} else {
		// Success
		log.Println(os.Getpid(), ":", "Successfully started child process:", "->", cmd.Process.Pid, "\n ")
	}

	if isExit {
		os.Exit(0)
	}

	return cmd, nil
}

func NewDaemon(logFile string) *Daemon {
	return &Daemon{
		LogFile:     logFile,
		MaxCount:    0,
		MaxError:    3,
		MinExitTime: 10,
	}
}

// Start the background daemon process
func (d *Daemon) Run() {
	// Start a daemon process and exit
	Background(d.LogFile, true)

	// Daemon process starts a child process and monitors it in a loop
	var t int64
	count := 1
	errNum := 0
	for {
		// Daemon info description
		dInfo := fmt.Sprintf("Daemon process(pid:%d; count:%d/%d; errNum:%d/%d):",
			os.Getpid(), count, d.MaxCount, errNum, d.MaxError)
		if errNum > d.MaxError {
			log.Println(dInfo, "Too many child process start failures, exiting")
			os.Exit(1)
		}
		if d.MaxCount > 0 && count > d.MaxCount {
			log.Println(dInfo, "Too many restarts, exiting")
			os.Exit(0)
		}
		count++

		t = time.Now().Unix() // Start timestamp
		cmd, err := Background(d.LogFile, false)
		if err != nil { // Start failed
			log.Println(dInfo, "Child process start failed;", "err:", err)
			errNum++
			continue
		}

		// Child process
		if cmd == nil {
			log.Printf("Child process pid=%d: Starting...", os.Getpid())
			break
		}

		// Parent process: wait for child process to exit
		err = cmd.Wait()
		dat := time.Now().Unix() - t // Child process runtime in seconds
		if dat < d.MinExitTime {     // Abnormal exit
			errNum++
		} else { // Normal exit
			errNum = 0
		}
		log.Printf("%s Detected child process(%d) exit, ran for %d seconds: %v\n", dInfo, cmd.ProcessState.Pid(), dat, err)
	}
}

func startProc(args, env []string, logFile string) (*exec.Cmd, error) {
	cmd := &exec.Cmd{
		Path:        args[0],
		Args:        args,
		Env:         env,
		SysProcAttr: NewSysProcAttr(),
	}

	if logFile != "" {
		stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Println(os.Getpid(), ": Error opening log file:", err)
			return nil, err
		}
		cmd.Stderr = stdout
		cmd.Stdout = stdout
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
