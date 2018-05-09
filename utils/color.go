package utils

import "fmt"

const (
	// PRIMARY flag to represent primary message
	PRIMARY = iota
	// SECONDARY flag to represent secondary message
	SECONDARY
	// SUCCESS flag to represent succesful message
	SUCCESS
	// DANGER flag to represent critical message
	DANGER
	// WARNING flag to represent warning message
	WARNING
	// INFO flag to represent information
	INFO
)

const (
	// KERNAL_RESET      Reset Linux Kernel Color
	KERNAL_RESET = "\033[0m"
	// KERNAL_BLACK      Linux Kernel Color     Black
	KERNAL_BLACK = "\033[30m"
	// KERNAL_RED        Linux Kernel Color     Red
	KERNAL_RED = "\033[31m"
	// KERNAL_GREEN      Linux Kernel Color     Green
	KERNAL_GREEN = "\033[32m"
	// KERNAL_YELLOW     Linux Kernel Color     Yellow
	KERNAL_YELLOW = "\033[33m"
	// KERNAL_BLUE       Linux Kernel Color     Blue
	KERNAL_BLUE = "\033[34m"
	// KERNAL_MAGENTA    Linux Kernel Color     Magenta
	KERNAL_MAGENTA = "\033[35m"
	// KERNAL_CYAN       Linux Kernel Color     Cyan
	KERNAL_CYAN = "\033[36m"
	// KERNAL_WHITE      Linux Kernel Color     White
	KERNAL_WHITE = "\033[37m"
	// KERNAL_BOLDBLACK  Linux Kernel Color     Bold Black
	KERNAL_BOLDBLACK = "\033[1m\033[30m"
	// KERNAL_BOLDRED    Linux Kernel Color     Bold Red
	KERNAL_BOLDRED = "\033[1m\033[31m"
	// KERNAL_BOLDGREEN  Linux Kernel Color     Bold Green
	KERNAL_BOLDGREEN = "\033[1m\033[32m"
	// KERNAL_BOLDYELLOW Linux Kernel Color     Bold Yellow
	KERNAL_BOLDYELLOW = "\033[1m\033[33m"
	// KERNAL_BOLDBLUE   Linux Kernel Color     Bold Blue
	KERNAL_BOLDBLUE = "\033[1m\033[34m"
	// KERNAL_BOLDMAGENTALinux Kernel Color     Bold Magenta
	KERNAL_BOLDMAGENTA = "\033[1m\033[35m"
	// KERNAL_BOLDCYAN   Linux Kernel Color     Bold Cyan
	KERNAL_BOLDCYAN = "\033[1m\033[36m"
	// KERNAL_BOLDWHITE  Linux Kernel Color     Bold White
	KERNAL_BOLDWHITE = "\033[1m\033[37m"
)

// Colors Kernel Colors
var Colors = [...]string{
	// KERNAL_BOLDBLUE    PRIMARY
	KERNAL_BOLDBLUE,
	// KERNAL_BOLDBLACK   SECONDARY
	KERNAL_BOLDBLACK,
	// KERNAL_BOLDGREEN   SUCCESS
	KERNAL_BOLDGREEN,
	// KERNAL_BOLDRED     DANGER
	KERNAL_BOLDRED,
	// KERNAL_BOLDYELLOW  WARNING
	KERNAL_BOLDYELLOW,
	// KERNAL_BOLDCYAN    INFO
	KERNAL_BOLDCYAN,
}

// DecorateString Take a generic type and a color (represented such as utils.PRIMARY) and return a decorated string
func DecorateString(str interface{}, t int) string {
	if t >= len(Colors) {
		t = 0
	}
	return fmt.Sprintf("%s%v%s",
		Colors[t],
		str,
		KERNAL_RESET)
}
