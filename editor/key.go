package editor

import . "github.com/itchyny/bed/common"

func defaultKeyManagers() map[Mode]*KeyManager {
	kms := make(map[Mode]*KeyManager)
	km := NewKeyManager(true)
	km.Register(EventQuit, "Z", "Q")
	km.Register(EventQuit, "c-w", "q")
	km.Register(EventQuit, "c-w", "c-q")
	km.Register(EventQuit, "c-w", "c")
	km.Register(EventCursorUp, "up")
	km.Register(EventCursorDown, "down")
	km.Register(EventCursorLeft, "left")
	km.Register(EventCursorRight, "right")
	km.Register(EventPageUp, "pgup")
	km.Register(EventPageDown, "pgdn")
	km.Register(EventPageTop, "home")
	km.Register(EventPageEnd, "end")
	km.Register(EventCursorUp, "k")
	km.Register(EventCursorDown, "j")
	km.Register(EventCursorLeft, "h")
	km.Register(EventCursorRight, "l")
	km.Register(EventCursorPrev, "b")
	km.Register(EventCursorNext, "w")
	km.Register(EventCursorHead, "0")
	km.Register(EventCursorHead, "^")
	km.Register(EventCursorEnd, "$")
	km.Register(EventScrollUp, "c-y")
	km.Register(EventScrollDown, "c-e")
	km.Register(EventPageUp, "c-b")
	km.Register(EventPageDown, "c-f")
	km.Register(EventPageUpHalf, "c-u")
	km.Register(EventPageDownHalf, "c-d")
	km.Register(EventPageTop, "g", "g")
	km.Register(EventPageEnd, "G")
	km.Register(EventJumpTo, "\x1d")
	km.Register(EventJumpBack, "c-t")
	km.Register(EventDeleteByte, "x")
	km.Register(EventDeletePrevByte, "X")
	km.Register(EventIncrement, "c-a")
	km.Register(EventIncrement, "+")
	km.Register(EventDecrement, "c-x")
	km.Register(EventDecrement, "-")

	km.Register(EventStartInsert, "i")
	km.Register(EventStartInsertHead, "I")
	km.Register(EventStartAppend, "a")
	km.Register(EventStartAppendEnd, "A")
	km.Register(EventStartReplaceByte, "r")
	km.Register(EventStartReplace, "R")

	km.Register(EventSwitchFocus, "tab")
	km.Register(EventStartCmdline, ":")

	km.Register(EventNew, "c-w", "n")
	km.Register(EventNew, "c-w", "c-n")
	km.Register(EventFocusWindowDown, "c-w", "down")
	km.Register(EventFocusWindowDown, "c-w", "c-j")
	km.Register(EventFocusWindowDown, "c-w", "j")
	km.Register(EventFocusWindowUp, "c-w", "up")
	km.Register(EventFocusWindowUp, "c-w", "c-k")
	km.Register(EventFocusWindowUp, "c-w", "k")
	km.Register(EventFocusWindowLeft, "c-w", "left")
	km.Register(EventFocusWindowLeft, "c-w", "c-h")
	km.Register(EventFocusWindowLeft, "c-w", "backspace")
	km.Register(EventFocusWindowLeft, "c-w", "h")
	km.Register(EventFocusWindowRight, "c-w", "right")
	km.Register(EventFocusWindowRight, "c-w", "c-l")
	km.Register(EventFocusWindowRight, "c-w", "l")
	km.Register(EventFocusWindowTopLeft, "c-w", "t")
	km.Register(EventFocusWindowTopLeft, "c-w", "c-t")
	km.Register(EventFocusWindowBottomRight, "c-w", "b")
	km.Register(EventFocusWindowBottomRight, "c-w", "c-b")
	km.Register(EventMoveWindowTop, "c-w", "K")
	km.Register(EventMoveWindowBottom, "c-w", "J")
	km.Register(EventMoveWindowLeft, "c-w", "H")
	km.Register(EventMoveWindowRight, "c-w", "L")
	kms[ModeNormal] = km

	km = NewKeyManager(false)
	km.Register(EventExitInsert, "escape")
	km.Register(EventExitInsert, "c-c")
	km.Register(EventCursorUp, "up")
	km.Register(EventCursorDown, "down")
	km.Register(EventCursorLeft, "left")
	km.Register(EventCursorRight, "right")
	km.Register(EventCursorUp, "c-p")
	km.Register(EventCursorDown, "c-n")
	km.Register(EventCursorPrev, "c-b")
	km.Register(EventCursorNext, "c-f")
	km.Register(EventPageUp, "pgup")
	km.Register(EventPageDown, "pgdn")
	km.Register(EventPageTop, "home")
	km.Register(EventPageEnd, "end")
	km.Register(EventBackspace, "backspace")
	km.Register(EventBackspace, "backspace2")
	km.Register(EventDelete, "delete")
	km.Register(EventSwitchFocus, "tab")
	kms[ModeInsert] = km
	kms[ModeReplace] = km

	km = NewKeyManager(false)
	km.Register(EventCursorLeft, "left")
	km.Register(EventCursorLeft, "c-b")
	km.Register(EventCursorRight, "right")
	km.Register(EventCursorRight, "c-f")
	km.Register(EventCursorHead, "home")
	km.Register(EventCursorHead, "c-a")
	km.Register(EventCursorEnd, "end")
	km.Register(EventCursorEnd, "c-e")
	km.Register(EventBackspaceCmdline, "c-h")
	km.Register(EventBackspaceCmdline, "backspace")
	km.Register(EventBackspaceCmdline, "backspace2")
	km.Register(EventDeleteCmdline, "delete")
	km.Register(EventDeleteWordCmdline, "c-w")
	km.Register(EventClearToHeadCmdline, "c-u")
	km.Register(EventClearCmdline, "c-k")
	km.Register(EventExitCmdline, "escape")
	km.Register(EventExitCmdline, "c-c")
	km.Register(EventExecuteCmdline, "enter")
	km.Register(EventExecuteCmdline, "c-j")
	km.Register(EventExecuteCmdline, "c-m")
	kms[ModeCmdline] = km
	return kms
}
