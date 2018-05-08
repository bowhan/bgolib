package utils

import "fmt"

const (
	PRIMARY   = iota
	SECONDARY
	SUCCESS
	DANGER
	WARNING
	INFO
)

const (
	KERNAL_RESET       = "\033[0m"
	KERNAL_BLACK       = "\033[30m"        /*  Black        */
	KERNAL_RED         = "\033[31m"        /*  Red          */
	KERNAL_GREEN       = "\033[32m"        /*  Green        */
	KERNAL_YELLOW      = "\033[33m"        /*  Yellow       */
	KERNAL_BLUE        = "\033[34m"        /*  Blue         */
	KERNAL_MAGENTA     = "\033[35m"        /*  Magenta      */
	KERNAL_CYAN        = "\033[36m"        /*  Cyan         */
	KERNAL_WHITE       = "\033[37m"        /*  White        */
	KERNAL_BOLDBLACK   = "\033[1m\033[30m" /*  Bold Black   */
	KERNAL_BOLDRED     = "\033[1m\033[31m" /*  Bold Red     */
	KERNAL_BOLDGREEN   = "\033[1m\033[32m" /*  Bold Green   */
	KERNAL_BOLDYELLOW  = "\033[1m\033[33m" /*  Bold Yellow  */
	KERNAL_BOLDBLUE    = "\033[1m\033[34m" /*  Bold Blue    */
	KERNAL_BOLDMAGENTA = "\033[1m\033[35m" /*  Bold Magenta */
	KERNAL_BOLDCYAN    = "\033[1m\033[36m" /*  Bold Cyan    */
	KERNAL_BOLDWHITE   = "\033[1m\033[37m" /*  Bold White   */
)

var Colors = [...]string{KERNAL_BOLDBLUE, /* PRIMARY */
	KERNAL_BOLDBLACK,                     /* SECONDARY */
	KERNAL_BOLDGREEN,                     /* SUCCESS */
	KERNAL_BOLDRED,                       /* DANGER */
	KERNAL_BOLDYELLOW,                    /* WARNING */
	KERNAL_BOLDCYAN,                      /* INFO */
}

func DecorateString(str interface{}, t int) string {
	if t >= len(Colors) {
		t = 0
	}
	return fmt.Sprintf("%s%v%s",
		Colors[t],
		str,
		KERNAL_RESET)
}
