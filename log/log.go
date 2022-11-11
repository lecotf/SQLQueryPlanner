package logsqp

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sqlqp/sqp"
	"strconv"
	"strings"
	"time"
)

const LOG = 1
const DEBUG = 2
const ERROR = 3
const FATALERROR = 4
const SQL = 5
const SQLDEBUG = 6
const SQLERROR = 7

type LogWriterError struct{}

func Print(msg string, typeLog int) {
	// Si les modes DEBUG ne sont pas actif, ne rien faire
	if typeLog == DEBUG && sqp.Config.Debug == 0 {
		return
	}
	if typeLog == SQLDEBUG && sqp.Config.SQLDebug == 0 {
		return
	}

	f := initLog(typeLog)
	switch typeLog {
	case 2, 6:
		log.Println("{DEBUG} " + msg)
	case 3:
		logger := log.New(LogWriterError{}, "Error message: ", 0)
		logger.Print(msg)
	case 4:
		logger := log.New(LogWriterError{}, "Error message: ", 0)
		logger.Print(msg)
		log.Println("--- Exit due to error above ---")
		os.Exit(1)
	default:
		log.Println(msg)
	}
	f.Close()
}

func initLog(logType int) (f *os.File) {
	// Définition du nom du fichier de log
	var fileName string
	if logType >= 5 && logType <= 7 {
		fileName = sqp.Config.ProjectShortName + "_" + "SQLQueries"
	} else {
		fileName = sqp.Config.ProjectName
	}
	currentTime := time.Now()
	fileName = sqp.Config.DataFolder + "/" + currentTime.Format("2006-01-02") + "_" + fileName + ".log"

	// Création et ouverture en écriture du fichier de log
	err := os.MkdirAll(sqp.Config.DataFolder, 0755)
	if err != nil {
		panic(err)
	}
	f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// Défini l'output de la libraire log sur le fichier ouvert et indique que chaque log sera précédé par l'heure
	log.SetOutput(f)
	log.SetFlags(log.Ltime)

	return f
}

func (f LogWriterError) Write(p []byte) (n int, err error) {
	var file, callStack, fnName, pkgName string
	var pc uintptr
	var line int

	for i, ok := 4, true; ok == true; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if ok == true {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				tmp := fn.Name()
				fnName = strings.TrimLeft(filepath.Ext(tmp), ".")
				if pos := strings.Index(tmp, "."); pos != -1 {
					pkgName = tmp[0:pos]
				}
			}
			callStack = "[Pkg: " + pkgName + "]\t[File: " + filepath.Base(file) + "]\t[Func: " + fnName + "]\t[line " + strconv.Itoa(line) + "]\r\n" + callStack
		}
		file, fnName, pkgName = "?", "?", "?"
	}
	log.Printf("[ERROR]\r\n***\r\n%s\r\n%s\r\n***\r\n", p, callStack)
	return len(p), nil
}

/*
func (f LogWriterError2) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(5)
	if !ok {
		file = "?"
		line = 0
	}
	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}
	log.Printf("[ERROR]\r\n***\r\n%s:%d - %s\r\n%s", filepath.Base(file), line, fnName, p)
	return len(p), nil
}

func writeLogFile(msg string, nameFile string, mode int, logger *log.Logger) {
	_, f := initLog(nameFile)
	switch mode {
	case 1:
		log.Println(msg)
	case 2:
		logger.Println("Error message: \"" + msg + "\"")
	case 3:
		log.Println(msg)
	}
	f.Close()
}

func Log(msg string) {
	writeLogFile(msg, sqp.Config.ProjectName, 1, nil)
}

func LogError(msg string, exit bool) {
	logger := log.New(LogWriterError{}, "", 0)
	writeLogFile(msg, sqp.Config.ProjectName, 2, logger)
	if exit {
		writeLogFile("--- Exit due to error above ---", sqp.Config.ProjectName, 1, nil)
		os.Exit(1)
	}
}

func LogSQLError(msg string) {
	writeLogFile(msg, sqp.Config.ProjectShortName+"_FailedSQL", 3, nil)
}*/
