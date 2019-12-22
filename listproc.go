package main

/*
   listproc v0.0.1  Dec 19 2019
   a list server for HNET and web
    
   listproc &
   (c) 2019 by moshix
   Program source is under Apache license             */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var version string = "v0.0.1"  // version of this program 

vary discussionsdir string = "/usr/spool/listproc"      // where discussion files are kept

var discussions_count int64         // total discussions managed, used /STATS command
var messages__received_count int64  // total notes received, used by /STATS command
var messages__sent_count int64      // total notes received, used by /STATS command

type discussions struct {
	subscribers  string //user@node
	lastmessage  int64  //when was the last message received for this discussion
	file         string //name of file holding this discussion in directory
	msgs         int64  //how many messages in this discussion
}

var discussiontable map[string]discussions // map of structs of all logged on users

func main() {
	discussions = make(map[string]discussions)
     
	fmt.Println("HNET LISTPROC server started....")
    // Open file with backup of all discussion files (backup of map discussions)
	

	// wait for commands to enter
	while {
	    when file with command came in catch it here somehow
		extract the first line 
		if not a listproc command line then ignore it!! 
		else 
	       readcommand(commandline)   //process the command after extracting it
		time.Sleep(400 * time.Millisecond) // wait a bit to avoid excessive CPU usage
	}
}

func openfile(discussion_file string) {

	file, err := os.OpenFile(discussionfile, os.O_CREATE, os.ModeNamedPipe) // how to open a normal text file ???XXXXXXXXXXX 
	if err != nil {
		log.Fatal("Open discussion file error:", err)
	} else {
		fmt.Print("Discussion file successfully opened and now listening\n")
	}

}



func readcommand(commandline string) {

	var fifouser string
	var fifomsg string
	var upperfifomsg string
	var upperfifouser string

	s := strings.Split(fifoline, "}") //split this message into sender and msg content

	fifouser = s[0]
	fifomsg = s[1]                            //this is the payload part of the incoming msg
	upperfifomsg = strings.ToUpper(fifomsg)   //make upper case for commands processing
	upperfifouser = strings.ToUpper(fifouser) //make user upper case
	fmt.Printf("'%s' '%s'\n", upperfifouser, fifomsg)

	//at this point we have the user at node and the payload in fifomsg/upperfifomsg
	//now we start some very simple processing
	//---------------------------------------------------------------------------------
	//   /HELP sends to the user a help menu with ossibilities
	//   /SEARCH
	//   /SUBSCRIBE 
	//   /UNBUSBSCRIBE 
	//   /LIST  
    //   /stats     	sends usage statistics
	//---------------------------------------------------------------------------------

	switch upperfifomsg {
	case "/HELP":
		//		fmt.Println("This is the help case")
		break
	case "/WHO":
		//		fmt.Println("This is the WHO case")
		senduserlist(upperfifouser)
		break
	case "/STATS":
		//		fmt.Println("This is the STATS case")
		sendstats(upperfifouser)
		break
	case "/LOGON":
		//		fmt.Println("This is the LOGON case")
		adduser(upperfifouser)
		break
	case "/LOGOFF":
		//		fmt.Println("This is the LOGOFF case")
		deluser(upperfifouser)
		break
	default:
		// must be a regular chat message

		if _, ok := table[upperfifouser]; ok {
			broacastmsg(upperfifouser, fifouser, fifomsg)
		} else {
			cmd := exec.Command("/usr/local/bin/send", upperfifouser, "You are not logged on currently to RELAY chat")
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			msgcount++

		}
	}
}

func senduserlist(upperfifouser string) {

	for user, _ := range table {
		cmd := exec.Command("/usr/local/bin/send", upperfifouser, "Online last 120 min: ", user)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		msgcount++

	}
}

func sendstats(user string) {
	s := strconv.FormatInt(msgcount, 10)
	t := strconv.FormatInt(totaluser, 10)
	cmd := exec.Command("/usr/local/bin/send", user, " Total messages: ", s, "     Total users:", t)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	msgcount++

}

func adduser(user string) {
	table[user] = users{
		lastactivity: time.Now().Unix(),
	}
	cmd := exec.Command("/usr/local/bin/send", user, " Welcome to RELAY CHAT v0.9.5")
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	msgcount++
	totaluser++

}

func deluser(user string) {
	delete(table, user)
	cmd := exec.Command("/usr/local/bin/send", user, " Goodbye from RELAY CHAT v0.9.5")
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	msgcount++

}

func broacastmsg(upperfifouser string, fifouser string, fifomsg string) {

	// remove users inactive for 120  minutes
	thirtyMinutesAgo := time.Now().Add(time.Duration(-120) * time.Minute).Unix()
	for username, userStruct := range table {
		if userStruct.lastactivity < thirtyMinutesAgo {
			log.Printf("Deleting inactive user '%s'", username)
			delete(table, username)
		}
	}

	loopmsg := fifomsg[0:3] //this is the begignning of a user who is not logged on anymore
	//Looping messages begin with DMT, filter those
	if loopmsg == "DMT" {
		delete(table, upperfifouser)

	}
	for upperfifouser, _ := range table {
		if _, ok := table[upperfifouser]; ok {
			cmd := exec.Command("/usr/local/bin/send", upperfifouser, "> ", fifouser, fifomsg)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			msgcount++

		}
	}
}
