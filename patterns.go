package main

import (
    "fmt"
)

var patternList = []string{
    "Off",
    "Forward Dreaming",
    "Backward Dreaming",
    "Forward 7 Colors",
    "Backward 7 Colors",
    "Forward RD/GN/BU",
    "Backward RD/GN/BU",
    "Forward YE/CN/VT",
    "Backward YE/CN/VT",
    "Forward 6 Colors RD",
    "Backward 6 Colors RD",
    "Forward 6 Colors GN",
    "Backward 6 Colors GN",
    "Forward 6 Colors BU",
    "Backward 6 Colors BU",
    "Forward 6 Colors CN",
    "Backward 6 Colors CN",
    "Forward 6 Colors YE",
    "Backward 6 Colors YE",
    "Forward 6 Colors VT",
    "Backward 6 Colors VT",
    "Forward 6 Colors WH",
    "Backward 6 Colors WH",
    "Forward Trailing 7 Colors",
    "Backward Trailing 7 Colors",
    "Forward Trailing RD",
    "Backward Trailing RD",
    "Forward Trailing GN",
    "Backward Trailing GN",
    "Forward Trailing BU",
    "Backward Trailing BU",
    "Forward Trailing YE",
    "Backward Trailing YE",
    "Forward Trailing CN",
    "Backward Trailing CN",
    "Forward Trailing VT",
    "Backward Trailing VT",
    "Forward Trailing WH",
    "Backward Trailing WH",
    "Forward Streaming 7 Colors",
    "Backward Streaming 7 Colors",
    "Forward Streaming RD/GN/BU",
    "Backward Streaming RD/GN/BU",
    "Forward Streaming YE/CN/VT",
    "Backward Streaming YE/CN/VT",
    "Forward Streaming RD/GN",
    "Backward Streaming RD/GN",
    "Forward Streaming GN/BU",
    "Backward Streaming GN/BU",
    "Forward Streaming YE/BU",
    "Backward Streaming YE/BU",
    "Forward Streaming YE/CN",
    "Backward Streaming YE/CN",
    "Forward Streaming CN/VT",
    "Backward Streaming CN/VT",
    "Forward Streaming BK/WH",
    "Backward Streaming BK/WH",
    "Open Curtain 7 Colors",
    "Close Curtain 7 Colors",
    "Open Curtain RD/GN/BU",
    "Close Curtain RD/GN/BU",
    "Open Curtain YE/CN/VT",
    "Close Curtain YE/CN/VT",
    "Forward Follow Spot 7 Colors",
    "Backward Follow Spot 7 Colors",
    "Forward Follow Spot RD/GN/BU",
    "Backward Follow Spot RD/GN/BU",
    "Forward Follow Spot YE/CN/VT",
    "Backward Follow Spot YE/CN/VT",
    "Forward Flutter 7 Colors",
    "Backward Flutter 7 Colors",
    "Forward Flutter RD/GN/BU",
    "Backward Flutter RD/GN/BU",
    "Forward Flutter YE/CN/VT",
    "Backward Flutter YE/CN/VT",
    "Hop 7 Colors",
    "Hop RD/GN/BU",
    "Hop RD/GN/BU",
    "Strobe 7 Colors",
    "Strobe RD/GN/BU",
    "Strobe YE/CN/VT",
    "Gradual 7 Colors",
    "Gradual RD/E",
    "Gradual RD/VT",
    "Gradual GN/CN",
    "Gradual GN/YE",
    "Gradual BU/VT",
    "Close Curtain RD",
    "Close Curtain GN",
    "Close Curtain BU",
    "Close Curtain YE ",
    "Close Curtain CN",
    "Close Curtain VT",
    "Close Curtain WH",
    "Open Curtain RD",
    "Open Curtain GN",
    "Open Curtain BU",
    "Open Curtain YE",
    "Open Curtain CN",
    "Open Curtain VT",
    "Open Curtain WH",
    "Horse Race RD",
    "Horse Race GN",
    "Horse Race BU",
    "Horse Race YE",
    "Horse Race CN",
    "Horse Race VT",
    "Horse Race WH",
    "Forward Run RD",
    "Backward Run RD",
    "Forward Run GN",
    "Backward Run GN",
    "Forward Run BU",
    "Backward Run BU",
    "Forward Run YE",
    "Backward Run YE",
    "Forward Run CN",
    "Backward Run CN",
    "Forward Run VT",
    "Backward Run VT",
    "Forward Run WH",
    "Backward Run WH",
    "Forward Run 7 Colors",
    "Backward Run 7 Colors",
    "Forward Run RD/GN/BU",
    "Backward Run RD/GN/BU",
    "Forward Run YE/CN/VT",
    "Backward Run YE/CN/VT",
    "Forward Run BU/YE/VT",
    "Forward Run GN/BU/YE",
    "Backward Run BU/YE/VT",
    "Forward Run RD/ON/WH",
    "Backward Run RD/ON/WH",
    "Forward Run GN ON RD",
    "Backward Run GN ON RD",
    "Forward Run BU ON GN",
    "Backward Run BU ON GN",
    "Forward Run YE ON BU",
    "Backward Run YE ON BU",
    "Forward Run CN ON YE",
    "Backward Run CN ON YE",
    "Forward Run VT ON CN",
    "Backward Run VT ON CN",
    "Forward Run WH ON VT",
    "Backward Run WH ON VT",
    "Forward Run WH ON RD",
    "Backward Run WH ON RD",
    "Forward Run 7 COLOR ON RD",
    "Backward Run 7 COLOR ON RD",
    "Forward Run 7 COLOR ON GN",
    "Backward Run 7 COLOR ON GN",
    "Forward Run 7 COLOR ON BU",
    "Backward Run 7 COLOR ON BU",
    "Forward Run 7 COLOR ON YE",
    "Backward Run 7 COLOR ON YE",
    "Forward Run 7 COLOR ON CN",
    "Backward Run 7 COLOR ON CN",
    "Forward Run 7 COLOR ON VT",
    "Backward Run 7 COLOR ON VT",
    "Forward Run 7 COLOR ON WH",
    "Backward Run 7 COLOR ON WH",
    "Forward Flow WH RD WH",
    "Backward Flow WH RD WH",
    "Forward Flow WH GN WH",
    "Backward Flow WH GN WH",
    "Forward Flow WH BU WH",
    "Backward Flow WH BU WH",
    "Forward Flow WH YE WH",
    "Backward Flow WH YE WH",
    "Forward Flow WH CN WH",
    "Backward Flow WH CN WH",
    "Forward Flow WH VT WH",
    "Backward Flow WH VT WH",
    "Forward Flow RD WH RD",
    "Backward Flow RD WH RD",
    "Forward Flow GN WH GN",
    "Backward Flow GN WH GN",
    "Forward Flow BU WH BU",
    "Backward Flow BU WH BU",
    "Forward Flow YE WH YE",
    "Backward Flow YE WH YE",
    "Forward Flow CN WH CN",
    "Backward Flow CN WH CN",
    "Forward Flow VT WH VT",
    "Backward Flow VT WH VT",
    "Forward Run GN ON BU",
    "Backward Run GN ON BU",
    "Forward Run GN ON RD",
    "Backward Run GN ON RD",
    "Forward Run RD ON BU",
    "Backward Run RD ON BU",
    "Forward Run CN ON YE",
    "Backward Run CN ON YE",
    "Forward Run YE ON VT",
    "Backward Run YE ON VT",
    "Forward Run WH ON YE",
    "Backward Run WH ON YE",
    "Forward Run YE ON WH",
    "Backward Run YE ON WH",
    "Forward Swab 7 COLORS",
    "Backward Swab 7 COLORS",
    "Forward Swab RD GN BU",
    "Backward Swab RD GN BU",
    "Forward Swab YE CN VT",
    "Backward Swab YE CN VT",
    "Open Curtain Swab 7 COLORS",
    "Close Curtain Swab 7 COLORS",
    "Open Curtain Swab R G B ",
    "Close Curtain Swab R G B ",
    "Open Curtain Swab Y C P",
    "Close Curtain Swab Y C P",
}

// return the index of a pattern by name
func GetPatternIndex(name string) int {
    for i, pattern := range patternList {
        if pattern == name {
            return i
        }
    }
    return -1
}

func ListPatterns() {
    for i, pattern := range patternList {
        fmt.Printf("%d: %s\n", i, pattern)
    }
}
