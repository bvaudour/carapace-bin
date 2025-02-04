package os

import "github.com/rsteube/carapace"

// ActionHexColors completes hex color codes
//   #0000ff (Blue1)
//   #00d75f (SpringGreen3)
func ActionHexColors() carapace.Action {
	return carapace.ActionValuesDescribed(
		"#000000", "Black",
		"#800000", "Maroon",
		"#008000", "Green",
		"#808000", "Olive",
		"#000080", "Navy",
		"#800080", "Purple",
		"#008080", "Teal",
		"#c0c0c0", "Silver",
		"#808080", "Grey",
		"#ff0000", "Red",
		"#00ff00", "Lime",
		"#ffff00", "Yellow",
		"#0000ff", "Blue",
		"#ff00ff", "Fuchsia",
		"#00ffff", "Aqua",
		"#ffffff", "White",
		"#000000", "Grey0",
		"#00005f", "NavyBlue",
		"#000087", "DarkBlue",
		"#0000af", "Blue3",
		"#0000d7", "Blue3",
		"#0000ff", "Blue1",
		"#005f00", "DarkGreen",
		"#005f5f", "DeepSkyBlue4",
		"#005f87", "DeepSkyBlue4",
		"#005faf", "DeepSkyBlue4",
		"#005fd7", "DodgerBlue3",
		"#005fff", "DodgerBlue2",
		"#008700", "Green4",
		"#00875f", "SpringGreen4",
		"#008787", "Turquoise4",
		"#0087af", "DeepSkyBlue3",
		"#0087d7", "DeepSkyBlue3",
		"#0087ff", "DodgerBlue1",
		"#00af00", "Green3",
		"#00af5f", "SpringGreen3",
		"#00af87", "DarkCyan",
		"#00afaf", "LightSeaGreen",
		"#00afd7", "DeepSkyBlue2",
		"#00afff", "DeepSkyBlue1",
		"#00d700", "Green3",
		"#00d75f", "SpringGreen3",
		"#00d787", "SpringGreen2",
		"#00d7af", "Cyan3",
		"#00d7d7", "DarkTurquoise",
		"#00d7ff", "Turquoise2",
		"#00ff00", "Green1",
		"#00ff5f", "SpringGreen2",
		"#00ff87", "SpringGreen1",
		"#00ffaf", "MediumSpringGreen",
		"#00ffd7", "Cyan2",
		"#00ffff", "Cyan1",
		"#5f0000", "DarkRed",
		"#5f005f", "DeepPink4",
		"#5f0087", "Purple4",
		"#5f00af", "Purple4",
		"#5f00d7", "Purple3",
		"#5f00ff", "BlueViolet",
		"#5f5f00", "Orange4",
		"#5f5f5f", "Grey37",
		"#5f5f87", "MediumPurple4",
		"#5f5faf", "SlateBlue3",
		"#5f5fd7", "SlateBlue3",
		"#5f5fff", "RoyalBlue1",
		"#5f8700", "Chartreuse4",
		"#5f875f", "DarkSeaGreen4",
		"#5f8787", "PaleTurquoise4",
		"#5f87af", "SteelBlue",
		"#5f87d7", "SteelBlue3",
		"#5f87ff", "CornflowerBlue",
		"#5faf00", "Chartreuse3",
		"#5faf5f", "DarkSeaGreen4",
		"#5faf87", "CadetBlue",
		"#5fafaf", "CadetBlue",
		"#5fafd7", "SkyBlue3",
		"#5fafff", "SteelBlue1",
		"#5fd700", "Chartreuse3",
		"#5fd75f", "PaleGreen3",
		"#5fd787", "SeaGreen3",
		"#5fd7af", "Aquamarine3",
		"#5fd7d7", "MediumTurquoise",
		"#5fd7ff", "SteelBlue1",
		"#5fff00", "Chartreuse2",
		"#5fff5f", "SeaGreen2",
		"#5fff87", "SeaGreen1",
		"#5fffaf", "SeaGreen1",
		"#5fffd7", "Aquamarine1",
		"#5fffff", "DarkSlateGray2",
		"#870000", "DarkRed",
		"#87005f", "DeepPink4",
		"#870087", "DarkMagenta",
		"#8700af", "DarkMagenta",
		"#8700d7", "DarkViolet",
		"#8700ff", "Purple",
		"#875f00", "Orange4",
		"#875f5f", "LightPink4",
		"#875f87", "Plum4",
		"#875faf", "MediumPurple3",
		"#875fd7", "MediumPurple3",
		"#875fff", "SlateBlue1",
		"#878700", "Yellow4",
		"#87875f", "Wheat4",
		"#878787", "Grey53",
		"#8787af", "LightSlateGrey",
		"#8787d7", "MediumPurple",
		"#8787ff", "LightSlateBlue",
		"#87af00", "Yellow4",
		"#87af5f", "DarkOliveGreen3",
		"#87af87", "DarkSeaGreen",
		"#87afaf", "LightSkyBlue3",
		"#87afd7", "LightSkyBlue3",
		"#87afff", "SkyBlue2",
		"#87d700", "Chartreuse2",
		"#87d75f", "DarkOliveGreen3",
		"#87d787", "PaleGreen3",
		"#87d7af", "DarkSeaGreen3",
		"#87d7d7", "DarkSlateGray3",
		"#87d7ff", "SkyBlue1",
		"#87ff00", "Chartreuse1",
		"#87ff5f", "LightGreen",
		"#87ff87", "LightGreen",
		"#87ffaf", "PaleGreen1",
		"#87ffd7", "Aquamarine1",
		"#87ffff", "DarkSlateGray1",
		"#af0000", "Red3",
		"#af005f", "DeepPink4",
		"#af0087", "MediumVioletRed",
		"#af00af", "Magenta3",
		"#af00d7", "DarkViolet",
		"#af00ff", "Purple",
		"#af5f00", "DarkOrange3",
		"#af5f5f", "IndianRed",
		"#af5f87", "HotPink3",
		"#af5faf", "MediumOrchid3",
		"#af5fd7", "MediumOrchid",
		"#af5fff", "MediumPurple2",
		"#af8700", "DarkGoldenrod",
		"#af875f", "LightSalmon3",
		"#af8787", "RosyBrown",
		"#af87af", "Grey63",
		"#af87d7", "MediumPurple2",
		"#af87ff", "MediumPurple1",
		"#afaf00", "Gold3",
		"#afaf5f", "DarkKhaki",
		"#afaf87", "NavajoWhite3",
		"#afafaf", "Grey69",
		"#afafd7", "LightSteelBlue3",
		"#afafff", "LightSteelBlue",
		"#afd700", "Yellow3",
		"#afd75f", "DarkOliveGreen3",
		"#afd787", "DarkSeaGreen3",
		"#afd7af", "DarkSeaGreen2",
		"#afd7d7", "LightCyan3",
		"#afd7ff", "LightSkyBlue1",
		"#afff00", "GreenYellow",
		"#afff5f", "DarkOliveGreen2",
		"#afff87", "PaleGreen1",
		"#afffaf", "DarkSeaGreen2",
		"#afffd7", "DarkSeaGreen1",
		"#afffff", "PaleTurquoise1",
		"#d70000", "Red3",
		"#d7005f", "DeepPink3",
		"#d70087", "DeepPink3",
		"#d700af", "Magenta3",
		"#d700d7", "Magenta3",
		"#d700ff", "Magenta2",
		"#d75f00", "DarkOrange3",
		"#d75f5f", "IndianRed",
		"#d75f87", "HotPink3",
		"#d75faf", "HotPink2",
		"#d75fd7", "Orchid",
		"#d75fff", "MediumOrchid1",
		"#d78700", "Orange3",
		"#d7875f", "LightSalmon3",
		"#d78787", "LightPink3",
		"#d787af", "Pink3",
		"#d787d7", "Plum3",
		"#d787ff", "Violet",
		"#d7af00", "Gold3",
		"#d7af5f", "LightGoldenrod3",
		"#d7af87", "Tan",
		"#d7afaf", "MistyRose3",
		"#d7afd7", "Thistle3",
		"#d7afff", "Plum2",
		"#d7d700", "Yellow3",
		"#d7d75f", "Khaki3",
		"#d7d787", "LightGoldenrod2",
		"#d7d7af", "LightYellow3",
		"#d7d7d7", "Grey84",
		"#d7d7ff", "LightSteelBlue1",
		"#d7ff00", "Yellow2",
		"#d7ff5f", "DarkOliveGreen1",
		"#d7ff87", "DarkOliveGreen1",
		"#d7ffaf", "DarkSeaGreen1",
		"#d7ffd7", "Honeydew2",
		"#d7ffff", "LightCyan1",
		"#ff0000", "Red1",
		"#ff005f", "DeepPink2",
		"#ff0087", "DeepPink1",
		"#ff00af", "DeepPink1",
		"#ff00d7", "Magenta2",
		"#ff00ff", "Magenta1",
		"#ff5f00", "OrangeRed1",
		"#ff5f5f", "IndianRed1",
		"#ff5f87", "IndianRed1",
		"#ff5faf", "HotPink",
		"#ff5fd7", "HotPink",
		"#ff5fff", "MediumOrchid1",
		"#ff8700", "DarkOrange",
		"#ff875f", "Salmon1",
		"#ff8787", "LightCoral",
		"#ff87af", "PaleVioletRed1",
		"#ff87d7", "Orchid2",
		"#ff87ff", "Orchid1",
		"#ffaf00", "Orange1",
		"#ffaf5f", "SandyBrown",
		"#ffaf87", "LightSalmon1",
		"#ffafaf", "LightPink1",
		"#ffafd7", "Pink1",
		"#ffafff", "Plum1",
		"#ffd700", "Gold1",
		"#ffd75f", "LightGoldenrod2",
		"#ffd787", "LightGoldenrod2",
		"#ffd7af", "NavajoWhite1",
		"#ffd7d7", "MistyRose1",
		"#ffd7ff", "Thistle1",
		"#ffff00", "Yellow1",
		"#ffff5f", "LightGoldenrod1",
		"#ffff87", "Khaki1",
		"#ffffaf", "Wheat1",
		"#ffffd7", "Cornsilk1",
		"#ffffff", "Grey100",
		"#080808", "Grey3",
		"#121212", "Grey7",
		"#1c1c1c", "Grey11",
		"#262626", "Grey15",
		"#303030", "Grey19",
		"#3a3a3a", "Grey23",
		"#444444", "Grey27",
		"#4e4e4e", "Grey30",
		"#585858", "Grey35",
		"#626262", "Grey39",
		"#6c6c6c", "Grey42",
		"#767676", "Grey46",
		"#808080", "Grey50",
		"#8a8a8a", "Grey54",
		"#949494", "Grey58",
		"#9e9e9e", "Grey62",
		"#a8a8a8", "Grey66",
		"#b2b2b2", "Grey70",
		"#bcbcbc", "Grey74",
		"#c6c6c6", "Grey78",
		"#d0d0d0", "Grey82",
		"#dadada", "Grey85",
		"#e4e4e4", "Grey89",
		"#eeeeee", "Grey93",
	)
}

// ActionXtermColorNames completes xterm color names
//   Green
//   Olive
func ActionXtermColorNames() carapace.Action {
	return carapace.ActionValues(
		"Black",
		"Maroon",
		"Green",
		"Olive",
		"Navy",
		"Purple",
		"Teal",
		"Silver",
		"Grey",
		"Red",
		"Lime",
		"Yellow",
		"Blue",
		"Fuchsia",
		"Aqua",
		"White",
	)
}
