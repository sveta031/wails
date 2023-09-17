//go:build windows

package application

var VirtualKeyCodes = map[uint]string{
	0x01: "lbutton",
	0x02: "rbutton",
	0x03: "cancel",
	0x04: "mbutton",
	0x05: "xbutton1",
	0x06: "xbutton2",
	0x08: "back",
	0x09: "tab",
	0x0C: "clear",
	0x0D: "return",
	0x10: "shift",
	0x11: "control",
	0x12: "menu",
	0x13: "pause",
	0x14: "capital",
	0x15: "kana",
	0x17: "junja",
	0x18: "final",
	0x19: "hanja",
	0x1B: "escape",
	0x1C: "convert",
	0x1D: "nonconvert",
	0x1E: "accept",
	0x1F: "modechange",
	0x20: "space",
	0x21: "prior",
	0x22: "next",
	0x23: "end",
	0x24: "home",
	0x25: "left",
	0x26: "up",
	0x27: "right",
	0x28: "down",
	0x29: "select",
	0x2A: "print",
	0x2B: "execute",
	0x2C: "snapshot",
	0x2D: "insert",
	0x2E: "delete",
	0x2F: "help",
	0x30: "0",
	0x31: "1",
	0x32: "2",
	0x33: "3",
	0x34: "4",
	0x35: "5",
	0x36: "6",
	0x37: "7",
	0x38: "8",
	0x39: "9",
	0x41: "a",
	0x42: "b",
	0x43: "c",
	0x44: "d",
	0x45: "e",
	0x46: "f",
	0x47: "g",
	0x48: "h",
	0x49: "i",
	0x4A: "j",
	0x4B: "k",
	0x4C: "l",
	0x4D: "m",
	0x4E: "n",
	0x4F: "o",
	0x50: "p",
	0x51: "q",
	0x52: "r",
	0x53: "s",
	0x54: "t",
	0x55: "u",
	0x56: "v",
	0x57: "w",
	0x58: "x",
	0x59: "y",
	0x5A: "z",
	0x5B: "lwin",
	0x5C: "rwin",
	0x5D: "apps",
	0x5F: "sleep",
	0x60: "numpad0",
	0x61: "numpad1",
	0x62: "numpad2",
	0x63: "numpad3",
	0x64: "numpad4",
	0x65: "numpad5",
	0x66: "numpad6",
	0x67: "numpad7",
	0x68: "numpad8",
	0x69: "numpad9",
	0x6A: "multiply",
	0x6B: "add",
	0x6C: "separator",
	0x6D: "subtract",
	0x6E: "decimal",
	0x6F: "divide",
	0x70: "f1",
	0x71: "f2",
	0x72: "f3",
	0x73: "f4",
	0x74: "f5",
	0x75: "f6",
	0x76: "f7",
	0x77: "f8",
	0x78: "f9",
	0x79: "f10",
	0x7A: "f11",
	0x7B: "f12",
	0x7C: "f13",
	0x7D: "f14",
	0x7E: "f15",
	0x7F: "f16",
	0x80: "f17",
	0x81: "f18",
	0x82: "f19",
	0x83: "f20",
	0x84: "f21",
	0x85: "f22",
	0x86: "f23",
	0x87: "f24",
	0x88: "navigation_view",
	0x89: "navigation_menu",
	0x8A: "navigation_up",
	0x8B: "navigation_down",
	0x8C: "navigation_left",
	0x8D: "navigation_right",
	0x8E: "navigation_accept",
	0x8F: "navigation_cancel",
	0x90: "numlock",
	0x91: "scroll",
	0x92: "oem_nec_equal",
	0x93: "oem_fj_masshou",
	0x94: "oem_fj_touroku",
	0x95: "oem_fj_loya",
	0x96: "oem_fj_roya",
	0xA0: "lshift",
	0xA1: "rshift",
	0xA2: "lcontrol",
	0xA3: "rcontrol",
	0xA4: "lmenu",
	0xA5: "rmenu",
	0xA6: "browser_back",
	0xA7: "browser_forward",
	0xA8: "browser_refresh",
	0xA9: "browser_stop",
	0xAA: "browser_search",
	0xAB: "browser_favorites",
	0xAC: "browser_home",
	0xAD: "volume_mute",
	0xAE: "volume_down",
	0xAF: "volume_up",
	0xB0: "media_next_track",
	0xB1: "media_prev_track",
	0xB2: "media_stop",
	0xB3: "media_play_pause",
	0xB4: "launch_mail",
	0xB5: "launch_media_select",
	0xB6: "launch_app1",
	0xB7: "launch_app2",
	0xBA: "oem_1",
	0xBB: "oem_plus",
	0xBC: "oem_comma",
	0xBD: "oem_minus",
	0xBE: "oem_period",
	0xBF: "oem_2",
	0xC0: "oem_3",
	0xC3: "gamepad_a",
	0xC4: "gamepad_b",
	0xC5: "gamepad_x",
	0xC6: "gamepad_y",
	0xC7: "gamepad_right_shoulder",
	0xC8: "gamepad_left_shoulder",
	0xC9: "gamepad_left_trigger",
	0xCA: "gamepad_right_trigger",
	0xCB: "gamepad_dpad_up",
	0xCC: "gamepad_dpad_down",
	0xCD: "gamepad_dpad_left",
	0xCE: "gamepad_dpad_right",
	0xCF: "gamepad_menu",
	0xD0: "gamepad_view",
	0xD1: "gamepad_left_thumbstick_button",
	0xD2: "gamepad_right_thumbstick_button",
	0xD3: "gamepad_left_thumbstick_up",
	0xD4: "gamepad_left_thumbstick_down",
	0xD5: "gamepad_left_thumbstick_right",
	0xD6: "gamepad_left_thumbstick_left",
	0xD7: "gamepad_right_thumbstick_up",
	0xD8: "gamepad_right_thumbstick_down",
	0xD9: "gamepad_right_thumbstick_right",
	0xDA: "gamepad_right_thumbstick_left",
	0xDB: "oem_4",
	0xDC: "oem_5",
	0xDD: "oem_6",
	0xDE: "oem_7",
	0xDF: "oem_8",
	0xE1: "oem_ax",
	0xE2: "oem_102",
	0xE3: "ico_help",
	0xE4: "ico_00",
	0xE5: "processkey",
	0xE6: "ico_clear",
	0xE7: "packet",
	0xE9: "oem_reset",
	0xEA: "oem_jump",
	0xEB: "oem_pa1",
	0xEC: "oem_pa2",
	0xED: "oem_pa3",
	0xEE: "oem_wsctrl",
	0xEF: "oem_cusel",
	0xF0: "oem_attn",
	0xF1: "oem_finish",
	0xF2: "oem_copy",
	0xF3: "oem_auto",
	0xF4: "oem_enlw",
	0xF5: "oem_backtab",
	0xF6: "attn",
	0xF7: "crsel",
	0xF8: "exsel",
	0xF9: "ereof",
	0xFA: "play",
	0xFB: "zoom",
	0xFC: "noname",
	0xFD: "pa1",
	0xFE: "oem_clear",
}
